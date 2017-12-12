package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    arr := []int{}
    count := 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()

    for scanner.Scan() {
        arr = append(arr, bytesToInt(scanner.Bytes()))
    }

    for i := 1; i < len(arr); i++ { // starting at second index and comparing left
        if (arr[i - 1] > arr[i]) {
            count = insertionSort(arr, i, arr[i], count)
        }
    }

    fmt.Println(count)
}

func insertionSort(arr []int, index, num, count int) int {
    if (index != 0 && arr[index - 1] > num) {
        count += 1
        arr[index] = arr[index - 1]
        count = insertionSort(arr, index - 1, num, count)
    } else {
        arr[index] = num
    }
    return count
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}