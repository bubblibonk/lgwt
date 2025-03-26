package main

func Sum(nums []int) int{
    result:=0;
    for _,number:= range nums{
        result+=nums[number]
    }
    return result;
}
func SumAll(numbersToSum ...[]int) []int{

    var sums [] int
    for _,numbers:=range numbersToSum{
        sums = append(sums,Sum(numbers));
    }
    return sums;
}
