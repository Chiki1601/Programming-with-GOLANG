// map Delete example in go

package main
import "fmt"

func main() {
    m := make(map[string]int)
    fmt.Println(m)
    m["key1"] =10
    m["key2"] =20
    m["key3"] =30
    
    fmt.Println(m)
    delete(m,"key3")
    fmt.Println(m)
    delete(m,"key2")
       fmt.Println(m)
  
}
