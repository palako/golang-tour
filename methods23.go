//https://tour.golang.org/methods/23
package main

import (
    "io"
    "os"
    "strings"
)

const ascii_a = 97
const ascii_z = 123
const ascii_A = 65
const ascii_Z = 90
const numChar = ascii_z - ascii_a
const shift = 13

type rot13Reader struct {
    r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (l int, e error) {
    l, e = r13.r.Read(b)
    for i := 0; i < l; i++ {
        if b[i] > ascii_a && b[i] <= ascii_z {
            b[i] = (((b[i] - ascii_a) + shift) % numChar) + ascii_a
        } else if b[i] > ascii_A && b[i] <= ascii_Z {
            b[i] = (((b[i] - ascii_A) + shift) % numChar) + ascii_A
        } else {
            b[i] = b[i]
        }
    }
    return l, e
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
