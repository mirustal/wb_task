package main

import (
	"fmt"
	"math/big"
	"wildberries-task/01_inheritance"
	"wildberries-task/02_square"
	sumsquare "wildberries-task/03_sum_square"
	bitredactor "wildberries-task/08_bit_redactor"
	stepsort "wildberries-task/10_step_sort"
	set "wildberries-task/12_set"
	swapnum "wildberries-task/13_swap_num"
	reversechar "wildberries-task/19_reverseChar"
	reverseword "wildberries-task/20_reverseWord"
	bignum "wildberries-task/22_bigNum"
	pointer "wildberries-task/24_pointer"
	contains "wildberries-task/26_contains"

	removeslice "wildberries-task/23_remove_slice"
)


func main() {
	test_01()
	test_02()
	test_03()
	test_08()
	test_10()
	test_12()
	test_13()
	test_19()
	test_20()
	test_22()
	test_23()
	test_24()
	test_26()
}

func test_01() {
	fmt.Println("Test_01")
	inheritance.Inheritance("Miroslav")
	fmt.Println()
}

func test_02() {
	fmt.Println("Test_02")
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(square.Square(nums), "\n")
}

func test_03() {
	fmt.Println("Test_03")
	nums := []int{2, 4, 6, 8, 10}
	sumsquare.SumSquare(nums)
	fmt.Println(square.Square(nums), "\n")
}

func test_08() {
	fmt.Println("Test_08")
	num := int64(2)
	bitPos := uint8(1)
	bit := true
	fmt.Println("Num before: ", num)
	num = bitredactor.BitRedactor(num, bitPos, bit)
	fmt.Println("Num after: num: ", num, "\n")
}

func test_10() {
	fmt.Println("Test_10")
	var arr = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := float32(10)
	num := stepsort.StepSort(arr, step)
	fmt.Println("num: ", num, "\n")
}

func test_12() {
	fmt.Println("Test_12")
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := set.SetValue(strings)
	for element := range set {
		fmt.Println(element)
}
	fmt.Println()
}

func test_13() {
	fmt.Println("Test_13")
	fNum := 1
	sNum := 5
	fmt.Printf("Before swap:\nFirst num: %d Second num: %d\n", fNum, sNum)
	swapnum.SwapNum(&fNum, &sNum)
	fmt.Printf("After swap:\nFirst num: %d Second num: %d\n", fNum, sNum)
}

func test_14() {
	fmt.Println("Test_14")
	fNum := 1
	sNum := 5
	fmt.Printf("Before swap:\nFirst num: %d Second num: %d\n", fNum, sNum)
	swapnum.SwapNum(&fNum, &sNum)
	fmt.Printf("After swap:\nFirst num: %d Second num: %d\n", fNum, sNum)
}


func test_19() {
	fmt.Println("\nTest_19")
	str := "главрыба"
	fmt.Println("Before str: ", str)
	str = reversechar.ReverseString(str)
	fmt.Println("After str: ", str, "\n")
}



func test_20() {
	fmt.Println("\nTest_20")
	str := "now dog sun"
	fmt.Println("Before str: ", str)
	str = reverseword.ReverseWords(str)
	fmt.Println("After str: ", str, "\n")
}

func test_22() {
	fmt.Println("\nTest_22")
	a, b := big.NewInt(0), big.NewInt(0)
	a.SetString("100000000000000000000", 10)
	b.SetString("100000000000000000000", 10) 

	bignum.BigSum(a, b)
	bignum.BigSub(a, b)
	bignum.BigMul(a, b)
	bignum.BigDiv(a, b)
	fmt.Println()
}


func test_23() {
	fmt.Println("\nTest_23")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Before slice:", nums)
	nums = removeslice.Remove(nums, 2)
	fmt.Println("After slice:", nums, "\n")
}


func test_24() {
	fmt.Println("\nTest_24")
	point1 := pointer.NewPoint(1, 1)
    point2 := pointer.NewPoint(2, 2)
    fmt.Printf("Distance: %.2f\n\n", pointer.Distance(point1, point2))
}

func test_26() {
	fmt.Println("\nTest_26")
	strs := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, str := range strs { 
		fmt.Printf("strs: %s, uniq: %v\n", str, contains.ContainsSym(str))
	}
	fmt.Println()
}