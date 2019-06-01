# Behavioral » Mediator

## Elements

 - Mediator - interface to mediate between colleagues
 - Concrete Mediator - transfer messages between colleague
 - Colleague Class - communicate to the mediator to communicate with other object

## Description

It defined an object that encapsulate the manner of how a set of objects
interact each others. This patterns also solve the problem of coupling
delegating to the mediator the responsibility to manipulate objects.

## Implementation

In the following implementation we consider the mediator as a sort of teacher.
The teacher will teaches to students and all students will learn.

```go
type Student interface {
	Class() string
	Learn(message string)
	Learned() string
}

type Mediator interface {
	TeachesTo(c *Student)
	Spread(messqger string)
}
```

A class mate implements the interface Student. It learn something and tell what
  he or she have learned. In this mediator implementation students dont speak
  each others.

```go
type ClassMate struct {
	name        string
	lastMessage string
	forum       string
}

func (a *ClassMate) Class() string {
	return a.forum
}

func (a *ClassMate) Learn(message string) {
	a.lastMessage = message
}

func (a *ClassMate) Learned() string {
	return a.lastMessage
}
```

Finally the teacher. Teach spread knowledge and decide the list of students to
teach to.

```go
type Teacher struct {
	learners []Student
}

func (m *Teacher) TeachesTo(c Student) {
	m.learners = append(m.learners, c)
}

func (m *Teacher) Spread(message string) {
	for _, a := range m.learners {
		a.Learn(message)
	}
}
```

All works inside the main function. When the teacher say something, a student
learn.

```go
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

	eteacher.Spread("Message sent to everyone")

	fmt.Println(student.Learned()) // Message sent to everyone

}
```

Last but not least a test to prove that student really learned lesson!


```go
func TestStudentLearn(t *testing.T) {
	teacher := Teacher{}
	student := NewClassMate("Mario")
	teacher.TeachesTo(student)
	teacher.Spread("Message sent to everyone")
	if student.Learned() != "Message sent to everyone" {
		t.Error("Student should learn")
	}
}
```
