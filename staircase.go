package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var size int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    size = bytesToInt(scanner.Bytes()) - 1

    for row := 0; row <= size; row++ {
        for column := 0; column <= size; column++ {
            if (row + column < size) {
                fmt.Print(" ")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Print("\n")
    }
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}
