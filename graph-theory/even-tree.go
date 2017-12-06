package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"
)

func main() {
    tree := make(map[int][]int) // map[node][]childNodes
    childNodeCount := make(map[int]int) // map[node]numberOfRecursedChildNodes

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    scanner.Scan() // skip first entry

    // read all values into tree map
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
                tree[parentNode] = append(tree[parentNode], childNode)
            }
        }
    }

    childNodeCount = countChildNodes(1, tree, childNodeCount)

    /*
     * Find the maximum number of edges you can remove from the tree to get
     * a forest such that each connected component of the forest contains
     * an even number of nodes.
     */
    removalsPossible := 0
    for _, val := range childNodeCount {
        if (isEven(val)) {
            removalsPossible++
        }
    }
    fmt.Println(removalsPossible - 1)
}

/*
 * Recurse once through the tree and store count of all child nodes
 */
func countChildNodes(node int, tree map[int][]int, childNodeCount map[int]int) map[int]int {
    for _, val := range tree[node] {
        childNodeCount = countChildNodes(val, tree, childNodeCount)
        childNodeCount[node] += childNodeCount[val]
    }
    childNodeCount[node] += len(tree[node])
    return childNodeCount
}

func isEven(childNodes int) bool {
    if ((childNodes - 1) % 2 == 0) {
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
