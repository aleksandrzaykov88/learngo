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

//Append adds element to slice
func (ss safeSlice) Append(value interface{}) {
	ss <- commandData{action: insert, value: value}
}

//Delete removes element from slice
func (ss safeSlice) Delete(key int) {
	ss <- commandData{action: remove, key: key}
}

type atKey struct {
	value interface{}
}

//Ar returns value by inout key
func (ss safeSlice) At(key int) (value interface{}) {
	reply := make(chan interface{})
	ss <- commandData{action: find, key: key, result: reply}
	result := (<-reply).(atKey)
	return result.value
}

//Len returns length of slice
func (ss safeSlice) Len() int {
	reply := make(chan interface{})
	ss <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

//Update updates slice value by key using UpdateFunc
func (ss safeSlice) Update(key int, updater UpdateFunc) {
	ss <- commandData{action: update, key: key, updater: updater}
}

//Close closes commandData channel
func (ss safeSlice) Close() []interface{} {
	reply := make(chan []interface{})
	ss <- commandData{action: end, data: reply}
	return <-reply
}

//New returns new safeSlice object
func New() SafeSlice {
	ss := make(safeSlice)
	go ss.run()
	return ss
}

//run process input commands and starts some action
func (ss safeSlice) run() {
	store := make([]interface{}, 0)
	for command := range ss {
		switch command.action {
		case insert:
			store = append(store, command.value)
		case remove:
			store = append(store[:command.key], store[command.key+1:]...)
		case find:
			key := store[command.key]
			command.result <- atKey{key}
		case length:
			command.result <- len(store)
		case update:
			key := store[command.key]
			store[command.key] = command.updater(key)
		case end:
			close(ss)
			command.data <- store
		}
	}
}

func main() {
	ss := New()
	ss.Append(111123)
	ss.Append(1342342)
	fmt.Println(ss.At(1))
	ss.Delete(1)
	fmt.Println(ss.At(0))
	fmt.Println(ss.Len())
	updater := func(value interface{}) interface{} {
		return 5
	}
	ss.Update(0, updater)
	fmt.Println(ss.At(0))
	ss.Close()
}
