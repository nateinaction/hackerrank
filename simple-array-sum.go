package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    sum := 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    
    scanner.Scan() // skip first token
    
    for scanner.Scan() {
        sum += toInt(scanner.Bytes())
    }
    
    fmt.Println(sum)
}

/*
 * Borrowed this byte conversion from
 * https://stackoverflow.com/questions/31333353/faster-input-scanning
 * I'm still not quite sure how it works but I'm looking into it.
 */
func toInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}
