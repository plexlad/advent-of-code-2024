package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
    "sort"
    "math"
)

func main() {
    data, err := os.ReadFile("input")
    if err != nil {
        fmt.Println("File read issue")
        return
    }

    str := string(data)
    lines := strings.Split(str, "\n")

    // Initialize
    size := len(lines)
    first := make([]int, size)
    second := make([]int, size)
    var result int

    // Organize the split text into ints and slices
    for i, line := range lines {
        vals := strings.Split(line, " ")
        if len(vals) > 1 {
            var e1, e2 error
            first[i], e1 = strconv.Atoi(vals[0])
            second[i], e2 = strconv.Atoi(vals[1])

            if e1 != nil {
                panic(e1)
            }

            if e2 != nil {
                panic(e2)
            }
        }
    }
    
    // Sort the numbers
    sort.Ints(first)
    sort.Ints(second)

    for i, n := range first {
        result += int(math.Abs(float64(n - second[i])))
    }

    fmt.Println(result)
}
