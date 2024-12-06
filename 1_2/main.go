package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
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

    // Organize the split text into ints and slices
    for index, line := range lines {
        vals := strings.Split(line, " ")
        if len(vals) > 1 {
            var e1, e2 error
            first[index], e1 = strconv.Atoi(vals[0])
            second[index], e2 = strconv.Atoi(vals[1])

            if e1 != nil {
                panic(e1)
            }

            if e2 != nil {
                panic(e2)
            }
        }
    }
    
    // Use a hashmap to store what is going on
    m := make(map[int]int)
    result := 0

    for _, n := range first {
        // Check if it exists, and skip if it does
        if _, ok := m[n]; !ok {
            amount := 0
            for _, val := range second {
                if val == n {
                    amount += 1
                }
            }
            m[n] = amount
        }

        result += n * m[n]
    }

    fmt.Println(result);
}
