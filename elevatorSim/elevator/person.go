package elevator

type Person struct {
	weight       int
	workingFloor int
}

func NewPerson(weight int, workingFloor int) *Person {
	return &Person{
		weight:       weight,
		workingFloor: workingFloor,
	}
}
