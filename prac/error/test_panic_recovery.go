package main

import(
    "fmt"
)

func recoveryFullName() {
    if r := recover(); r != nil {
        fmt.Println("recovered from ", r)
    }
}

func fullName(firstName, lastName *string) {
    
    defer recoveryFullName()

    if firstName == nil {
        panic("runtime error: firstname is nil")
    }

    if lastName == nil {
        panic("runtime error: lastname  is nil")
    }

    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullname function")

}

func main() {

    defer fmt.Println("defer from main")
    firstName := "Elon"
    fullName(&firstName, nil)

    fmt.Println("returned normally from the main function")

}
