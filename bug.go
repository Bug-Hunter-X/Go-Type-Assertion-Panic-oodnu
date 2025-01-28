```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    j := i.(int)
    fmt.Println(j * 2) //This will panic if i is not an int
    fmt.Println(i.(string)) // This will panic if i is not a string
}