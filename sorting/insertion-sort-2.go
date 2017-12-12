package main

import (
    "strconv"
    "bufio"
    "bytes"
    "fmt"
    "os"
)

func main() {
    arr := []int{}
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()

    for scanner.Scan() {
        arr = append(arr, bytesToInt(scanner.Bytes()))
    }

    for i := 1; i < len(arr); i++ { // starting at second index and comparing left
        if (arr[i - 1] > arr[i]) {
            insertionSort(arr, i, arr[i])
        }
        printArray(arr)
    }
}

func insertionSort(arr []int, index int, num int) {
    if (index != 0 && arr[index - 1] > num) {
        arr[index] = arr[index - 1]
        insertionSort(arr, index - 1, num)
    } else {
        arr[index] = num
    }
}

func printArray(arr []int) {
    var buffer bytes.Buffer
    for i, n := range arr {
        buffer.WriteString(strconv.Itoa(n))
        if (i + 1 < len(arr)) {
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