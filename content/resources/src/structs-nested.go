package main

type Address struct {
    city, state string
}

type Employee struct {
    firstName, lastName string
    age int
    address Address
}

var emp Employee

emp.firstName = "SomeThing"
emp.address = Address {
    city  : "AA",
    state : "CO",
}