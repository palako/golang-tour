//https://tour.golang.org/flowcontrol/8
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i++ {
        aux := z
        z -= (z*z - x) / (2 * z)
        fmt.Printf("%v: %g\n", i, z)
        if math.Abs(z-aux) < 0.000001 {
            return z
        }
    }
    return z
}

func main() {
    x := 1024.0
    y := Sqrt(x)
    z := math.Sqrt(x)
    fmt.Printf("x=%g  Sqrt=%g  math.sqrt=%g  diff=%v", x, y, z, math.Abs(z - y))
}
