//https://tour.golang.org/moretypes/18
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    x := make([][]uint8, dx)
    for i := range(x) {
        y := make([]uint8, dy)
        x[i] = y
        for j := range(x[i]) {
            //x[i][j] = uint8((i+j)/2)
            x[i][j] = uint8(i*j)
            //x[i][j] = uint8(i^j)
        }
    }
    return x
}

func main() {
    pic.Show(Pic)
}

