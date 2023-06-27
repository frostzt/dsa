package main

import (
	"fmt"
	"reflect"
)

func reverse(input interface{}) any {
	switch input := input.(type) {
	case string:
		{
			bytes := []byte(input)

			var i int = 0
			var j int = len(bytes) - 1

			for j > i {
				bytes[i], bytes[j] = bytes[j], bytes[i]

				i = i + 1
				j = j - 1
			}

			return string(bytes)
		}

	case []int:
		{
			reversed := make([]int, len(input))
			copy(reversed, input)

			var i int = 0
			var j int = len(reversed) - 1

			for j > i {
				reversed[i], reversed[j] = reversed[j], reversed[i]

				i = i + 1
				j = j - 1
			}

			return reversed
		}

	default:
		{
			fmt.Println("Type ", reflect.TypeOf(input), " is not supported!")
			return -1
		}
	}
}

func main() {
	hello := "Hello World"
	reversedHello := reverse(hello)

	fmt.Println(reversedHello)

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reversedNumber := reverse(number)

	fmt.Println(reversedNumber)
}
