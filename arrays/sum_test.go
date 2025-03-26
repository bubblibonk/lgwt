package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T){
    mynums:=[] int{1,2,3,4,1}
    got:=Sum(mynums);
    want:=12;
    if got!=want{
        t.Errorf("got %d want %d the array is %v",got,want,mynums)
    }
}
func TestSumAll(t *testing.T){
    got:=SumAll([]int{1,2},[]int {3,4})
    want:= [] int{3,7}
    if slices.Equal(got,want){
        t.Errorf("got %v want %v",got,want)
    }
}
