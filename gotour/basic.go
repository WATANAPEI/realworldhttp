package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("z is", z, "(", i+1, "times)")
	}
	return z
}

type Vertex struct {
	X int
	Y int
}

//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Pic(dx, dy int) [][]uint8 {
	var row []uint8
	var matrix [][]uint8
	for i := 0; i < dx; i++ {
		row = row[:0]
		for j := 0; j < dy; j++ {
			row = append(row, uint8(i*j))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func main() {
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
