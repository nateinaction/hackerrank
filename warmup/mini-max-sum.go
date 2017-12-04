package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var sum, number, max, min int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    // setup initial values
    scanner.Scan()
    max = bytesToInt(scanner.Bytes())
    min = bytesToInt(scanner.Bytes())
    sum = bytesToInt(scanner.Bytes())
    
    for scanner.Scan() {
        number = bytesToInt(scanner.Bytes())
        sum += number
        if (number > max) {
            max = number
        } else if (number < min) {
            min = number
        }
    }

    fmt.Println(sum - max, sum - min)   
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}
