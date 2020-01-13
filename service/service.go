package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var url = "http://dummy.restapiexample.com/api/v1/employees"

var url1 = "http://fakerestapi.azurewebsites.net/api/Activities/"

/*
type Response struct {
	status       string
	employeeData []EmployeeData
}

type EmployeeData struct {
	id string
	employee_name string
	employee_salary string
	employee_age  string
	profile_image  string
}*/

func GetEmployees(wg *sync.WaitGroup) {
	// Make a get request
	rs, err := http.Get(url)
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	fmt.Println("bodyString is ", bodyString)

	defer wg.Done()
}

func GetActivities(wg *sync.WaitGroup) {
	// Make a get request
	rs, err := http.Get(url1)
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	fmt.Println("bodyString is ", bodyString)

	defer wg.Done()

}
