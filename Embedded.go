// Embedded struct example on go

package main
import "fmt"
type person struct{
    fname string
    lname string
}
type employee struct{
    person
    empId int
}

func(p person)details(){
    fmt.Println(p,""+"I am a person")
}

func(e employee)details(){
    fmt.Println(e,""+"I am an Employee")
}

func main() {
    p1 := person{"Pooja","Patel"}
    p1.details()
    e1:= employee{person:person{"Aditya","Patel"},empId :11}
    e1.details
}