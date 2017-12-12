package main

import (
	"strconv"
	"bytes"
	"bufio"
	"fmt"
	"os"
)

func main() {
	store := make(map[int]int)

    scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
    scanner.Scan() // skip first line

    for scanner.Scan() {
        count := bytesToInt(scanner.Bytes())
        store[count]++
    }

    var buffer bytes.Buffer
    for i := 0; i <= 99; i++ {
    	buffer.WriteString(strconv.Itoa(store[i]))
        if (i != 99) {
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