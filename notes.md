# Notes

## Slices

To delete items from a slice in Go, it is idiomatic (although expensive) to use a technique called re-slicing.

```go
func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

func main() {
    all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
    all = RemoveIndex(all, 5)
    fmt.Println(all) //[0 1 2 3 4 6 7 8 9]
}
```

Everything in Go is passed by values, including slices. A slice is header that contains the length of the slice and a pointer to the underlying array. When we `append` an element to a slice, Go will:
  1. Either update the underlying array if it has enough capacity, or allocate a new array.
  2. Additional it will update the length of the slice.
Because of 1 and 2, we cannot reliably depend on `append` to update the slice passed in as its first argument. Instead, `append` will return the slice.
