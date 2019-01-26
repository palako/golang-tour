//https://tour.golang.org/methods/25
package main

import {
    "golang.org/x/tour/pic"
    "image/color"
    "image"
}

type Image struct{
    Height int
    Width int
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x*y), uint8(x*y), 255, 255}    
}

func main() {
    m := Image{Height:250, Width: 250}
    pic.ShowImage(m)
}
