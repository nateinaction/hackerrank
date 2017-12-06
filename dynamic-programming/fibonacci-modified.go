package main

import (
    "math/big"
    "bufio"
    "fmt"
    "os"
)

func main() {
    memoize := make(map[int]*big.Int)

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    memoize[1] = big.NewInt(bytesToInt(scanner.Bytes()))

    scanner.Scan()
    memoize[2] = big.NewInt(bytesToInt(scanner.Bytes()))

    scanner.Scan()
    n := int(bytesToInt(scanner.Bytes()))

    for i := 3; i <= n; i++ {
        b := big.NewInt(0).Mul(memoize[i - 1], memoize[i - 1])
        memoize[i] = big.NewInt(0).Add(memoize[i - 2], b)
        delete(memoize, i - 2) // keeping only the values I need
    }

    fmt.Println(memoize[n])
}

func bytesToInt(buffer []byte) (n int64) {
    for _, v := range buffer {
        n = n*10 + int64(v-'0')
    }
    return
}