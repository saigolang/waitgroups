package main

import (
	"sync"
	"waitgroups/service"
)

func main() {

	// initil
	var wg sync.WaitGroup

	wg.Add(1)
	go service.GetEmployees(&wg)

	wg.Add(1)
	go service.GetActivities(&wg)

	wg.Wait()
}
