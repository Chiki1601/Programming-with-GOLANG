// for loop example in while fashion

package main 
import "fmt" 
func main(){
sum:= 1 
for sum<100{
    sum += sum 
    fmt.Println(sum)
    }
}
