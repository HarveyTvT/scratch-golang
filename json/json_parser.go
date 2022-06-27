package json

import (
	"errors"
	"strconv"
)

type TokenType string

const (
	BEGIN_OBJECT TokenType = "BEGIN_OBJECT"
	END_OBJECT   TokenType = "END_OBJECT"
	BEGIN_ARRAY  TokenType = "BEGIN_ARRAY"
	END_ARRAY    TokenType = "END_ARRAY"
	NULL         TokenType = "NULL"
	NUMBER       TokenType = "NUMBER"
	STRING       TokenType = "STRING"
	BOOLEAN      TokenType = "BOOLEAN"
	SEP_COLON    TokenType = "SEP_COLON"
	SEP_COMMA    TokenType = "SEP_COMMA"
	END_DOCUMENT TokenType = "END_DOCUMENT"
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(json string) []Token {
	tokens := make([]Token, 0)
	for i := 0; i < len(json); i++ {
		switch json[i] {
		case '{':
			tokens = append(tokens, Token{BEGIN_OBJECT, ""})
		case '}':
			tokens = append(tokens, Token{END_OBJECT, ""})
		case '[':
			tokens = append(tokens, Token{BEGIN_ARRAY, ""})
		case ']':
			tokens = append(tokens, Token{END_ARRAY, ""})
		case '"':
			j := i + 1
			for ; j < len(json); j++ {
				if json[j] == '"' && json[j-1] != '\\' {
					break
				}
			}
			tokens = append(tokens, Token{STRING, json[i : j+1]})
			i = j
		case 'n':
			j := i + 3
			if json[i+1] == 'u' && json[i+2] == 'l' && json[i+3] == 'l' {
				tokens = append(tokens, Token{NULL, json[i : j+1]})
				i = j
			}
		case 't':
			j := i + 3
			if json[i+1] == 'r' && json[i+2] == 'u' && json[i+3] == 'e' {
				tokens = append(tokens, Token{BOOLEAN, json[i : j+1]})
				i = j
			}
		case 'f':
			j := i + 4
			if json[i+1] == 'a' && json[i+2] == 'l' && json[i+3] == 's' && json[i+4] == 'e' {
				tokens = append(tokens, Token{BOOLEAN, json[i : j+1]})
				i = j
			}
		case ':':
			tokens = append(tokens, Token{SEP_COLON, ""})
		case ',':
			tokens = append(tokens, Token{SEP_COMMA, ""})
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j := i + 1
			for ; j < len(json); j++ {
				if json[j] != '-' && json[j] != '0' && json[j] != '1' && json[j] != '2' && json[j] != '3' && json[j] != '4' && json[j] != '5' && json[j] != '6' && json[j] != '7' && json[j] != '8' && json[j] != '9' && json[j] != '.' {
					break
				}
			}
			tokens = append(tokens, Token{NUMBER, json[i:j]})
			i = j - 1
		default:
			j := i + 1
			for ; j < len(json); j++ {
				if json[j] == ' ' || json[j] == '\n' || json[j] == '\r' || json[j] == '\t' {
					break
				}
			}
			tokens = append(tokens, Token{STRING, json[i : j+1]})
			i = j
		}
	}
	return tokens
} // tokenize

type JsonValue interface{}

type JsonNull struct{}

type JsonNumber float64

type JsonString string

type JsonBoolean bool

type JsonArray []JsonValue

type JsonObject map[string]JsonValue

func ParseValue(tokens []Token) (JsonValue, error) {
	switch tokens[0].Type {
	case BEGIN_OBJECT:
		return ParseObject(tokens)
	case BEGIN_ARRAY:
		return ParseArray(tokens)
	case NULL:
		return JsonNull{}, nil
	case NUMBER:
		number, err := strconv.ParseFloat(tokens[0].Value, 64)
		if err != nil {
			return nil, err
		}
		return JsonNumber(number), nil
	case STRING:
		return JsonString(tokens[0].Value), nil
	case BOOLEAN:
		return JsonBoolean(tokens[0].Value == "true"), nil
	default:
		panic(errors.New("ParseValue: unknown token type"))
	}
}

func ParseArray(tokens []Token) (JsonArray, error) {
	array := make(JsonArray, 0)
	for i := 1; i < len(tokens); i++ {
		if tokens[i].Type == END_ARRAY {
			return array, nil
		}
		value, err := ParseValue(tokens[i:])
		if err != nil {
			return nil, err
		}
		array = append(array, value)
	}
	return nil, errors.New("ParseArray: no end token")
}

func ParseObject(tokens []Token) (JsonObject, error) {
	object := make(JsonObject)
	for i := 1; i < len(tokens); i++ {
		if tokens[i].Type == END_OBJECT {
			return object, nil
		}
		if tokens[i].Type != STRING {
			return nil, errors.New("ParseObject: expected string")
		}
		key := tokens[i].Value
		i++
		if tokens[i].Type != SEP_COLON {
			return nil, errors.New("ParseObject: expected colon")
		}
		i++
		value, err := ParseValue(tokens[i:])
		if err != nil {
			return nil, err
		}
		object[key] = value
		j := i + 1
		if j < len(tokens) && tokens[j].Type == SEP_COMMA {
			i = j
		}
	}
	return nil, errors.New("ParseObject: no end token")
}

func ParseJson(json string) (JsonValue, error) {
	tokens := Tokenize(json)
	return ParseValue(tokens)
}
