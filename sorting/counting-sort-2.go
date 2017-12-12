package main

import (
	"strconv"
	"bytes"
	"bufio"
	"fmt"
	"os"
)

func main() {
	store := make(map[int][]int)

    scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
    scanner.Scan() // skip first line

    for scanner.Scan() {
        item := bytesToInt(scanner.Bytes())
        store[item] = append(store[item], item)
    }

    var buffer bytes.Buffer
    for i := 0; i <= 99; i++ {
        for _, v := range store[i] {
            buffer.WriteString(strconv.Itoa(v))
            buffer.WriteString(" ")
        }
    }
    fmt.Println(buffer.String())
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}