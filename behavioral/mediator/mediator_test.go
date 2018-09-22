package main

import (
	"testing"
)

func TestStudentLearn(t *testing.T) {
	teacher := Teacher{}
	student := NewClassMate("Mario")
	teacher.TeachesTo(student)
	teacher.Spread("Message sent to everyone")
	if student.Learned() != "Message sent to everyone" {
		t.Error("Student should learn")
	}
}
