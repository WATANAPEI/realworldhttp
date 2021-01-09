package main

import (
	"fmt"
	"math"
	"strings"
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
	Lat, Long float64
}

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

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	splited := strings.Fields(s)
	for i := 0; i < len(splited); i++ {
		result[splited[i]] += 1
	}
	return result
}

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	initials := make([]int, 0)
	initials = append(initials, 0)
	prev := 0
	current := 0
	tmp := 0

	return func() int {
		if prev == 0 && current == 0 {
			current = 1
			return 0
		} else if prev == 0 && current == 1 {
			prev = 1
			current = 1
			return 1
		} else {
			tmp = prev
			prev = current
			current += tmp
			return current
		}
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
