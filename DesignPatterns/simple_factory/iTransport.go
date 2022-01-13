package main

type iTransport interface {
	setName(name string)
	getName() string
	deliver()
}
