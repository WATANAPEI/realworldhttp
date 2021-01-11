package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
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

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Println("z is", z, "(", i+1, "times)")
// 	}
// 	return z
// }

type Vertex struct {
	X, Y float64
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

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

type Abser interface {
	Abs() float64
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	str := fmt.Sprint(float64(e))
	return fmt.Sprint("cannot Sqrt negative number: ", str)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return x, e
	}
	return x, nil
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b = append(b, 'A')
	return 1, nil
}

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		if rune(b[i]) >= 'a' && rune(b[i]) <= 'z' {
			b[i] = (b[i]-97+13)%26 + 97
		} else if rune(b[i]) >= 'A' && rune(b[i]) <= 'Z' {
			b[i] = (b[i]-65+13)%26 + 65
		}
	}
	if err == io.EOF {
		return n, io.EOF
	}
	return n, nil

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
