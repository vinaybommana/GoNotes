## types in go:

``` go
func add(x float64, y float64) float64 {
    return x + y
}
```

can also be written as
``` go
func add(x, y float64) float64 {
    return x + y
}

var num1, num2 float64 = 5.6, 9.5
// short circuit declaration
```

inorder to return multiple elements from
a function

``` go
func return_multiple_items(a, b string) (string, string) {
    return a, b
}
```

# pointers in go
```go
func main() {
    x := 15
    a := &x // memory address

    fmt.Println(a)

    // print the value of x
    fmt.Println(*a)

    *a = 5
    fmt.Println(x)
    fmt.Println(a)
}
```
