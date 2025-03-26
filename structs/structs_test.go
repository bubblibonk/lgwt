package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T){

    checkPerimeter:= func(t testing.TB,shape Shape,want float64){

        t.Helper()

        got:=shape.Perimeter()

        if got!=want{
            t.Errorf("got %.2f want %.2f",got,want);
        }
    }

    t.Run("rectangles",func(t *testing.T){
        rectangle:= Rectangle{4,5}
        checkPerimeter(t,rectangle,18.0)
    })

    t.Run("circles",func(t *testing.T){
        circle:= Circle{10}
        checkPerimeter(t,circle,math.Pi*100)
    })

    t.Run("triangles",func(t *testing.T){
        triangle:= Triangle{10,20}
        checkPerimeter(t,triangle,100)
    })
}
