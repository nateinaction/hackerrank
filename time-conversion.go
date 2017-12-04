package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

func main() {
    var timestring, hourstring, period, nochange string
    var hour int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    timestring = scanner.Text()

    period = string(timestring[8:])
    hourstring = string(timestring[:2])
    hour = stringToInt(hourstring)
    nochange = string(timestring[2:8])
    
    if (period == "PM" && hour != 12) {
        hour += 12
        hourstring = strconv.Itoa(hour)
    } else if (period == "AM" && hour == 12) {
        hourstring = "00"
    }

    fmt.Println(hourstring + nochange)
}

func stringToInt(s string) (n int) {
    n, err := strconv.Atoi(string(s))
    _ = err
    return
}
