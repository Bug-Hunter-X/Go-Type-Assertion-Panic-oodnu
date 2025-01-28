```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    switch j := i.(type) {
    case int:
        fmt.Println(j * 2)
    case string:
        fmt.Println(j)
    default:
        fmt.Println("Unexpected type")
    }

    //Alternative solution
    if i, ok := i.(int); ok {
        fmt.Println(i * 2)
    } else {
        fmt.Println("i is not an int")
    }

    if i, ok := i.(string); ok {
        fmt.Println(i)
    } else {
        fmt.Println("i is not a string")
    }
}
```