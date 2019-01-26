//https://tour.golang.org/moretypes/23
package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    count := make(map[string]int)
    
    sArray := strings.Fields(s)
    for _,w := range(sArray) {
        count[w] = count[w]+1
    }
    return count//map[string]int{"x": 1}
}

func main() {
    wc.Test(WordCount)
}

