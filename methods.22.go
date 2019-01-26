//https://tour.golang.org/methods/22
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
    for i:= 0; i<len(b); i++ {    
        b[i] = 0x41
    }
    return len(b), nil
}

func main() {
    reader.Validate(MyReader{})
}
