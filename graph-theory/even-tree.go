package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"
)

func main() {
    var nodes map[int][]int
    nodes = make(map[int][]int) // initialize map

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    scanner.Scan() // skip first entry

    for scanner.Scan() {
        edge := strings.NewReader(scanner.Text())
        edgeScanner := bufio.NewScanner(edge)
        edgeScanner.Split(bufio.ScanWords)

        childNode := 0 //set defautl value to 0
        var parentNode int
        for edgeScanner.Scan() {
            if (childNode == 0) {
                childNode = bytesToInt(edgeScanner.Bytes())
            } else {
                parentNode = bytesToInt(edgeScanner.Bytes())
                nodes[parentNode] = append(nodes[parentNode], childNode)
            }
        }
    }

    fmt.Println(nodes)
    fmt.Println(run(1, nodes, 0) - 1)
    //run(1, nodes, 0)
}

func run(node int, nodes map[int][]int, evenTreesFound int) int {
    for _, val := range nodes[node] {
        evenTreesFound = run(val, nodes, evenTreesFound)
    }
    if (isEvenTree(nodes[node])) {
        evenTreesFound++
        fmt.Println(node)
    }
    return evenTreesFound
}

func isEvenTree(childNodes []int) bool {
    if ((len(childNodes) - 1) % 2 == 0) {
        return true
    }
    return false
}

func bytesToInt(buffer []byte) (n int) {
    for _, v := range buffer {
        n = n*10 + int(v-'0')
    }
    return
}
