package json

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

type ValueType int

const (
	NULL_VALUE ValueType = iota
	BOOLEAN_VALUE
	NUMBER_VALUE
	STRING_VALUE
	ARRAY_VALUE
	MAP_VALUE
	OBJECT_VALUE
)

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
