package main

import "fmt"

func NewClassMate(name string) Student {
	return &ClassMate{name, "meeting 1", ""}
}

func main() {

	teacher := Teacher{}

	student := NewClassMate("Mario")

	teacher.TeachesTo(student)
	teacher.TeachesTo(NewClassMate("Demo"))
	teacher.TeachesTo(NewClassMate("Mattia"))
	teacher.TeachesTo(NewClassMate("Simone"))

	fmt.Println(len(teacher.attendees))
	fmt.Println(student.Class())

	teacher.Spread("Message sent to everyone")

	fmt.Println(student.Learned())

}
