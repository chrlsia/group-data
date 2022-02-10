package main

type employee struct {
	id   int
	name string
	dept string
}

func main() {
	var coder employee
	coder.id = 1
	coder.name = "Chris"
	coder.dept = "I.T"

}
