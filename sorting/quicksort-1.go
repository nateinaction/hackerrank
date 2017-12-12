package main

import (
    "strconv"
    "bytes"
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()
    scanner.Scan()

    left := []int{}
    right := []int{}
    equal := []int{stringToInt(scanner.Text())}

    var n int
    for scanner.Scan() {
        n = stringToInt(scanner.Text())
        if (n > equal[0]) {
            right = append(right, n)
        } else if (n < equal[0]) {
            left = append(left, n)
        } else {
            equal = append(equal, n)
        }
    }

    fmt.Println(printArray(left), printArray(equal), printArray(right))
}

func printArray(arr []int) string {
    var buffer bytes.Buffer
    for i, n := range arr {
        buffer.WriteString(strconv.Itoa(n))
        if (i + 1 < len(arr)) {
            buffer.WriteString(" ")
        }
    }
    return buffer.String()
}

func stringToInt(s string) (n int) {
    n, err := strconv.Atoi(string(s))
    _ = err
    return
}