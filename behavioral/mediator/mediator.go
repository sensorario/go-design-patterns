package main

import "fmt"

func NewClassMate(name string) Student {
	return &ClassMate{name, "meeting 1", ""}
}

/**
@todo: add a course as mediator
object that interacts with other
objects:
 - teacher
 - student
*/

func main() {

	teacher := Teacher{}

	student := NewClassMate("Mario")

	/** @todo: add variadic arguments
	and add student in one step */
	teacher.TeachesTo(student)
	teacher.TeachesTo(NewClassMate("Demo"))
	teacher.TeachesTo(NewClassMate("Mattia"))
	teacher.TeachesTo(NewClassMate("Simone"))

	fmt.Println(len(teacher.attendees))
	fmt.Println(student.Class())

	teacher.Spread("Message sent to everyone")

	fmt.Println(student.Learned())

}
