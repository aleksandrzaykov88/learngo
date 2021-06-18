package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
	address
}

type employee struct {
	name   string
	salary float64
	address
}

type address struct {
	street     string
	city       string
	state      string
	postalCode string
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)

	applyDiscount(subscriber2)
	printInfo(subscriber2)

	var myAddress address
	myAddress.street = "123 Oak St"
	myAddress.city = "Omaha"
	myAddress.state = "NE"
	myAddress.postalCode = "68111"

	var worker employee
	worker.name = "Joy Carr"
	worker.salary = 60000
	fmt.Println(worker.name)
	fmt.Println(worker.salary)
	worker.city = "Moscow"
	worker.postalCode = "612352"
	fmt.Printf("%#v\n", worker.address)
}
