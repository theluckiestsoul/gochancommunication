package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var sg sync.WaitGroup
	sg.Add(1)
	ch := make(chan employee)
	go printEmployee(ch, &sg)
	go func() {
		defer close(ch)
		e := newEmployee("Kiran", 30, "IT")
		ch <- e
	}()

	sg.Wait()
}

type employee struct {
	name       string
	age        int
	department string
}

func newEmployee(name string, age int, department string) employee {
	return employee{name: name, age: age, department: department}
}

func (e employee) String() string {
	fmt.Println("To String")
	return e.name + strconv.Itoa(e.age) + e.department
}

func printEmployee(ch chan employee, sg *sync.WaitGroup) {
	defer sg.Done()
	for c := range ch {
		fmt.Println(c)
	}

}
