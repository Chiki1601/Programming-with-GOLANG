// if else example with input from user in go
package main
import "fmt"

func main() {
    
   
    fmt.Println("Enter number")
     var input int
     fmt.Scanln(&input)
     fmt.Print(input)
    if(input%2==0){
        fmt.Println(" is even number")
    }else{
       fmt.Println(" is odd number")  
    }
  
}