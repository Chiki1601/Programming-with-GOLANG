// if else if chain example in go
package main
import "fmt"

func main() {
    
   
    fmt.Println("Enter text:")
     var input int
     fmt.Scanln(&input)
     
    if(input<0 || input>100){
        fmt.Println("Please enter valid number")
    }else if(input>=0 || input<50){
       fmt.Println(" Fail")  
    }else if(input>=50 || input<60){
       fmt.Println(" D grade")  
    }else if(input>=60 || input<70){
       fmt.Println(" C grade")  
    }else if(input>=70 || input<80){
       fmt.Println(" B grade")  
    }else if(input>=80 || input<90){
       fmt.Println(" A grade")  
    }else if(input>=90 || input<=100){
       fmt.Println(" A+ grade")  
    }
  
}