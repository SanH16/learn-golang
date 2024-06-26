package main

import (
	"errors"
	"fmt"
)

func main()  {
	value := 40
	result := CheckNilai(value)
	fmt.Println("Hello konoha!")
	fmt.Println(result)

	num1 := 2
	num2 := 10
	res, err := Increment(num1, num2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func CheckNilai(val int) (string) {
	var hasilNilai string
		if val <= 50 {
			hasilNilai = "Minor Value"
		} else if val <= 60 {
			hasilNilai = "Medium Value"
		} else if val >= 61 {
			hasilNilai = "High Value"
		}
	return hasilNilai
}

func Increment(num1, num2 int) (int, error) {
	if num1 == 0 || num2 == 0 {
		return 0, errors.New("Invalid Number 0")
	}
	
	result := num1 + num2
	return result, nil
}

//
// TODO: 
// buat function nerima 2 param type int, 1 string (operator)
// jika operator tidak di kenal return error "operator tidak di kenal"
//

// func Todo(v int, v2 string) int {
// 	if v == 0 || v2 == "" {
// 		return 0
// 	}
// 	results := v  v2
// 	return results
// }
