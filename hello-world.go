package main

import "fmt"

const FIELD_SIZE = 3

func main() {

	var field [FIELD_SIZE][FIELD_SIZE]int

	for y, row := range field {
		for x := range row {
			field[y][x] = (x+y)%2 + 1
		}
	}

	printField(field)
	fmt.Println(getFieldSum(field))

}

func printField(field [FIELD_SIZE][FIELD_SIZE]int) {
	for _, row := range field {
		for _, value := range row {
			fmt.Print(value)
		}
		fmt.Println()
	}
}

func getFieldSum(field [FIELD_SIZE][FIELD_SIZE]int) int {
	sum := 0
	for _, row := range field {
		for _, value := range row {
			sum += value
		}
	}
	return sum
}
