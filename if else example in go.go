
// if else example in go
package main
import "fmt"

func main() {
    
    var a int = 10
    if(a%2==0){
        fmt.Println("a is even number")
    }else{
       fmt.Println("a is odd number")  
    }
    fmt.Println("Value of a is: \n ",a)
}