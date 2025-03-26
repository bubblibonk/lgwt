package main

import "testing"

func TestHello(t *testing.T){
    t.Run("Test for names", func(t *testing.T){
        got:=Hello("");
        want:= "Hello Waorld";

        if got!=want{
            t.Errorf("got %q want %q",got,want)
        }
 
    })
   t.Run("if funnction returned with an empty string function should return string as world", func(t *testing.T){
        got:=Hello("")
        want:="Hello World"
        if got!=want{
            t.Errorf("got %q want %q",got,want);
        }
    })
}
