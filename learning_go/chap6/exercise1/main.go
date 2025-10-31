package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

// go build -gcflags="-m" main.go
func main() {
	MakePerson("John", "Doe", 30)
	MakePersonPointer("John", "Doe", 30)
}
