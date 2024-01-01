package man

import (
    "fmt" 
    "math/rand"
)

func Main() {
    // creating an array
    var arr []int = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
    fmt.Printf("type: %T\n", arr)
    fmt.Printf("Array: %+v\n", arr)
    fmt.Printf("Length: %d\n", len(arr))

    // slice
    // slice is a reference to an array segment 
    fmt.Printf("Slice: %+v\n", arr[3:len(arr)-1])
}

