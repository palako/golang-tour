//https://tour.golang.org/methods/20
package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }
    z := 1.0
    for i := 0; i < 10; i++ {
        aux := z
        z -= (z*z - x) / (2 * z)
        fmt.Printf("%v: %g\n", i, z)
        if math.Abs(z-aux) < 0.000001 {
            return z, nil
        }
    }
    return z, nil
}

func main() {
    sol, err := Sqrt(2)
    if err == nil {
        fmt.Println(sol)
    } else {
        fmt.Println(err)
    }
    sol, err = Sqrt(-2)
    if err == nil {
        fmt.Println(sol)
    } else {
        fmt.Println(err)
    }
}
