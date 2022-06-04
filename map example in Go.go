// map example in go

package main
import "fmt"

func main() {
    x := map[string]int{"Pooja": 16,"John":76,"Rohit":90}
    fmt.Print(x)
    fmt.Print("\n",x["Pooja"])
}
