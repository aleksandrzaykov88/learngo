package main

import (
	"fmt"
)

type employee struct {
	name string
	age  int
}

type department struct {
	id        int
	name      string
	employees []employee
}

type storageDepartment struct {
	counter int
	data    map[int]*department
}

func (s *storageDepartment) insertEmployee(id int, e employee) {
	s.data[id].employees = append(s.data[id].employees, e)
}

func main() {
	var alex = employee{"alex", 34}
	var jacob = employee{"jacob", 22}

	var it = department{1, "IT", []employee{}}

	data := make(map[int]*department, 2)
	data[1] = &it

	var sd = storageDepartment{1, data}

	fmt.Println(alex)

	sd.data[1].employees = append(sd.data[1].employees, alex)
	sd.insertEmployee(1, jacob)

	fmt.Println(sd.data[1].employees)
}
