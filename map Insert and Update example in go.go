// map Insert and Update example in go

package main
import "fmt"

func main() {
    m := make(map[string]int)
    fmt.Println(m)
    m["key1"] =10
    m["key2"] =20
    m["key3"] =30
    
    fmt.Println(m)
    
    m["key2"] = 555
    fmt.Println(m)
}
