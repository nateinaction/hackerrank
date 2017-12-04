package main

import (
    "math"
    "bufio"
    "strings"
    "strconv"
    "fmt"
    "os"
)

func main() {
    sumA, sumB := 0, 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)

    scanner.Scan() // skip initial empty line
    size := bytesToInt(scanner.Bytes()) - 1

    for row := 0; scanner.Scan(); row++ {
        rowReader := strings.NewReader(scanner.Text())
        rowScanner := bufio.NewScanner(rowReader)
        rowScanner.Split(bufio.ScanWords)
        for column := 0; rowScanner.Scan(); column++ {
            if (row == column) {
                sumA += stringToInt(rowScanner.Text())
            }
            if (row + column == size) {
                sumB += stringToInt(rowScanner.Text())
            }
        }
    }
    fmt.Println(math.Abs(float64(sumA - sumB)))
}

func stringToInt(s string) (n int) {
    n, err := strconv.Atoi(string(s))
    _ = err
    return
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}