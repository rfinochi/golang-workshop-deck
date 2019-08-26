func (e *Employee) updateAge(newAge int) {
    e.age = newAge
}
emp := Employee{
    age: 33,
}

fmt.Println("Before:", emp.age)
emp.updateAge(34)
fmt.Println("After:", emp.age)