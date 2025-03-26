package main

import "math"

type Rectangle struct {
    Width float64
    Length float64
}
type Circle struct {
    Radius float64
}
type Triangle struct{
    Base float64
    Height float64
}


func(r Rectangle)Perimeter() float64{
    return 2*(r.Length+r.Width)
}
func (c Circle) Perimeter()float64{
    return math.Pi*c.Radius*c.Radius
}
func(t Triangle) Perimeter()float64{
    return 0.5*t.Base*t.Height;
}

type Shape interface{
    Perimeter() float64;
}


