package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var find, number int
    var lookup map[int]int //map[number] = index
    lookup = make(map[int]int)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    find = bytesToInt(scanner.Bytes())

    scanner.Scan() // skip input "number of items in array"

    for index := 0; scanner.Scan(); index++ {
        number = bytesToInt(scanner.Bytes())
        lookup[number] = index
    }

    fmt.Println(lookup[find])
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}
