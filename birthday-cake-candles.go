package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var height, max int
    var occurances map[int]int
    occurances = make(map[int]int) // initialize map

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan() // skip first entry

    for scanner.Scan() {
        height = bytesToInt(scanner.Bytes())
        occurances[height]++
        if (height > max) {
            max = height
        }
    }

    fmt.Println(occurances[max])
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}
