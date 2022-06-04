
//Nested if else  chain example in go
package main
import "fmt"

func main() {
    
   
    var x int = 10
    var y  int = 20
    if(x>=10){
        if(y>=10){
            fmt.Println("inside nested if statement")
        }
    }
    
    fmt.Println("Value of x is: ",x)
    fmt.Println("Value of y is",y)
  
}