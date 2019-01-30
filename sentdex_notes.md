## types in go:

`uint16` // min 0 max 65535
`int16` // -32k --> 32k

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

# Basics of empty web tutorial

* http.HandleFunc
* http.ListenAndServe
  http.ListenAndServe --> (takes_port_number)
* http.ResponseWriter
* *http.Request

# Structs

`structs` are the basic building blocks in `go`

```go
type <struct_name> struct {
    // attributes
}

// usage
myStruct := <struct_name> {
    // attribute_values
}

// accessing
myStruct.attribute_value
```

## methods in go
* value receivers
    receive values and perform calculations

```go
func (<struct_attribute_name> <struct_name>) <method_name> <return_value> {
    return ...
}
```
* pointer receivers
    if we want to modify values in the struct
    we need pointer receivers

``` go
func (<struct_attribute_name> *<struct_name>) <method_name> <return_value> {
    return ...
}
```

## Should I define methods on values or pointers ?

``` go
func (s *MyStruct) pointerMethod() { } // method on pointer

func (s MyStruc) valueMethod() { } // method on value
```

When defining a method on a type, the receiver (s in the above example) behaves exactly as if it were an argument to the method.

Whether to define the receiver as a value or as
a pointer is the same question, then, as whether
a function argument should be a value or a pointer.

First, and most important, does the method need to modify the receiver? if it does, the receiver *must* be a pointer.

(Slices and maps acts as references, so their story is a little more subtle, but for instance to change the length of a slice in a method the receiver must still be a pointer.)

In the examples above, if pointerMethod modifies the fields of s, the caller will see those changes, but valueMethod is called with a copy of the caller's argument (that's the definition of passing a value), so changes it makes will be invisible to the caller.

Second is the consideration of efficiency.
if the receiver is large, a big `struct` for instance, it will be much cheaper to use a pointer receiver.

Next is consistency. If some of the methods of the type must have pointer receivers, the rest should too, so the method set is consistent regardless of how the type is used.

For types such as basic types, slices, and small structs, a value receiver is very cheap so unless the semantics of the method requires a pointer, a value receiver is efficient and clear


## arrays and slices
[5 5]int == array
[]int == slice


# looping structures
```go
for i:=0;i<10;i++ {
    //body of for
}
```