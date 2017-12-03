package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    a, b := 0, 0
    var length, ascore, bscore int
    var scores []int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        length = len(scores)
        scores = append(scores, toInt(scanner.Bytes()))
        if (length > 2) {
            ascore = scores[length - 3]
            bscore = scores[length]
            if (ascore > bscore) {
                a += 1
            } else if (ascore < bscore) {
                b += 1
            }
        }
    }

    fmt.Println(a, b)
}

func toInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}
