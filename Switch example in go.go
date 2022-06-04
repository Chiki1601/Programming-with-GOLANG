
//Switch example in go
package main
import "fmt"

func main() {
    
    fmt.Print("Enter number:")
    var input int 
    fmt.Scanln(&input)
    
switch(input){
        case 10:
        {
        fmt.Printf("The value is 10")
        }
        break
        case 20:
        fmt.Printf("The value is 20")
         break
        case 30:
        fmt.Printf("The value is 30")
         break
        case 40:
        fmt.Printf("The value is 40")
    
        case 100:
        fmt.Printf("The value is 100")
         break
        default:
        fmt.Printf("The value is not 10,20,30,40,50,60,70,80,90,100")
        
    }
    
    
  
}
