package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

func main() {
    var denominator, n int
    var positive, negative, zero []int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    denominator = stringToInt(scanner.Text())

    for scanner.Scan() {
        n = stringToInt(scanner.Text())
        if (n > 0) {
            positive = append(positive, n)
        } else if (n < 0) {
            negative = append(negative, n)
        } else {
            zero = append(zero, n)
        }
    }

    fmt.Printf("%0.6f\n", float64(len(positive)) / float64(denominator))
    fmt.Printf("%0.6f\n", float64(len(negative)) / float64(denominator))
    fmt.Printf("%0.6f\n", float64(len(zero)) / float64(denominator))
}

func stringToInt(s string) (n int) {
    n, err := strconv.Atoi(string(s))
    _ = err
    return
}
