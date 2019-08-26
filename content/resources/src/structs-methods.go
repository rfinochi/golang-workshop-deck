type Employee struct {
    firstName string
    lastName string
    age int
}

func (e Employee) Print() {
    fmt.Println("Employee Record:")
    fmt.Println("Name:", e.firstName, e.lastName)
    fmt.Println("Address:", e.address)
}

var emp Employee

emp.Print()