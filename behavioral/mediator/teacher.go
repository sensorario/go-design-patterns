package main

type Teacher struct {
	attendees []Student
}

func (m *Teacher) TeachesTo(c Student) {
	m.attendees = append(m.attendees, c)
}

func (m *Teacher) Spread(message string) {
	for _, a := range m.attendees {
		a.Learn(message)
	}
}
