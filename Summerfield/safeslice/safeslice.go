package main

import "fmt"

type SafeSlice interface {
	Append(interface{})
	At(int) interface{}
	Close() []interface{}
	Delete(int)
	Len() int
	Update(int, UpdateFunc)
}

type UpdateFunc func(interface{}) interface{}

type safeSlice chan commandData
type commandData struct {
	action  commandAction
	key     int
	value   interface{}
	result  chan<- interface{}
	data    chan<- []interface{}
	updater UpdateFunc
}
type commandAction int

const (
	remove commandAction = iota
	end
	find
	insert
	length
	update
)

func (ss safeSlice) Append(value interface{}) {
	ss <- commandData{action: insert, value: value}
}

func (ss safeSlice) Delete(key int) {
	ss <- commandData{action: remove, key: key}
}

type atValue struct {
	value interface{}
}

func (ss safeSlice) At(key int) (value interface{}) {
	reply := make(chan interface{})
	ss <- commandData{action: find, key: key, result: reply}
	result := (<-reply).(atValue)
	return result.value
}

func (ss safeSlice) Len() int {
	reply := make(chan interface{})
	ss <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (ss safeSlice) Update(key int, updater UpdateFunc) {
	ss <- commandData{action: update, key: key, updater: updater}
}

func (ss safeSlice) Close() []interface{} {
	reply := make(chan []interface{})
	ss <- commandData{action: end, data: reply}
	return <-reply
}

func New() SafeSlice {
	ss := make(safeSlice)
	go ss.run()
	return ss
}

func (ss safeSlice) run() {
	store := make([]interface{}, 0)
	for command := range ss {
		switch command.action {
		case insert:
			store[command.key] = command.value
		case remove:
			store = append(store[:command.key], store[command.key+1:]...)
		case find:
			value := store[command.key]
			command.result <- atValue{value}
		case length:
			command.result <- len(store)
		case update:
			value := store[command.key]
			store[command.key] = command.updater(value)
		case end:
			close(ss)
			command.data <- store
		}
	}
}

func main() {
	fmt.Println("Hi")
}
