package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FloatFormat(bit int, num float64) (str string) {
	str = strconv.FormatFloat(num, 'f', bit, 64)
	if len(str) < 4 {
		return str
	}

	arr := strings.Split(str, ".")
	length := len(arr[0])
	if length < 4 {
		return str
	}

	count := (length - 1) / 3

	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length-(i+1)*3] + "," + arr[0][length-(i+1)*3:]
	}
	fmt.Println(arr)
	return strings.Join(arr, ".")
}

func main() {
	num := 123489564545.789
	str := FloatFormat(6, num)
	fmt.Println(str)
}
