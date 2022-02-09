package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var tasks []func() error

	f := func() error {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		num, err := strconv.Atoi(strconv.Itoa(rand.Intn(3)))
		fmt.Println(num)
		if err != nil || num == 2 {
			return errors.New("Ошибка")
		}
		return nil
	}
	for i := 0; i < 6; i++ {
		tasks = append(tasks, f)
	}

	parallelExecuter(tasks, 3, 2)
}

func parallelExecuter(tasks []func() error, tasksAmount int, errorMax int) {
	chTasks := make(chan func() error, len(tasks))
	listenError := make(chan error, tasksAmount)

	go addTasks(tasks, chTasks)

	for i := 0; i < tasksAmount; i++ {
		go doTask(chTasks, listenError)
	}

	errorListener(listenError, tasksAmount, errorMax)
}

func addTasks(tasks []func() error, chTasks chan func() error) {
	for _, task := range tasks {
		chTasks <- task
	}
	close(chTasks)
}

func doTask(chTasks chan func() error, errorChan chan error) {
	var err error
	for task := range chTasks {
		err = task()
		if err != nil {
			errorChan <- err
		}
	}
	if err == nil {
		errorChan <- err
	}
}

func errorListener(errChan <-chan error, tasksAmount int, maxError int) {
	count := 0
	for i := 0; i < tasksAmount; i++ {
		err := <-errChan
		if err != nil {
			count++
		}
		if count >= maxError {
			fmt.Println("Cчётчик ошибок:", count, err)
			log.Fatal("exit")
		}
	}
}
