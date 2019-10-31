package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

var xx = "Hello, world, xx!"

type Point struct {
	x, y float64
}

type Circle struct {
	origin Point
	r      float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	Person Person
	Model  string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func f() {
	fmt.Println(xx)
}

func average(xs []float64) float64 {
	// panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f2returnWithName() (r int) {
	r = 1
	return
}

func f2returnMultipleValue() (int, int) {
	return 5, 6
}

func f2VariadicParams(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Factorial
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func sFirstFunc() {
	fmt.Println("1st")
}

func sSecondFunc() {
	fmt.Println("2nd")
}

func zeroIt(x int) {
	x = 0
}

func zeroItWithPointer(xPtr *int) {
	*xPtr = 0
}

func oneItWithPointer(xPtr *int) {
	*xPtr = 1
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleAreaOverload(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func totalArea(shape ...Shape) float64 {
	var area float64
	for _, v := range shape {
		area += v.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello, world!"))
	fmt.Println("Hello, world!"[1])
	fmt.Println("Hello, " + "world!")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x = "HeLlO, wOrld!"
	fmt.Println(x)

	var y string
	y = "Hello, wOrld!"
	fmt.Println(y)

	z := "Hello, wORld!"
	fmt.Println(z)

	var a = "hello"
	var b = "world"
	fmt.Println(a == b)

	var c = "hello"
	var d = "hello"
	fmt.Println(c == d)

	f()

	// const testing
	const cy = "Hello, world from const!"
	fmt.Println(cy)

	const (
		name  = "evian"
		stars = "muguruza"
	)

	//
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)

	//
	xxx := 5
	xxx++

	fmt.Println(xxx)

	//
	fmt.Println("--------------------------")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}

		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Unknown number!")
		}

	}

	fmt.Println("--------------------------")
	var xyz [5]int

	xyz[3] = 100
	fmt.Println(xyz)

	var xyy [5]float64
	xyy[0] = 98
	xyy[1] = 93
	xyy[2] = 77
	xyy[3] = 82
	xyy[4] = 83

	var total float64
	for i := 0; i < 5; i++ {
		total += xyy[i]
	}

	fmt.Println(total / 5)
	fmt.Println(total / float64(len(xyy)))

	var totalx float64
	for _, value := range xyy {
		totalx += value
	}

	fmt.Println(totalx / 5)

	xxy := [4]float64{
		98,
		// 93,
		77,
		82,
		83,
	}

	for i, value := range xxy {
		fmt.Printf("%d = %f \n", i, value)
	}

	xxyx := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(xxyx[2:5])

	fmt.Println("--------------------------")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice11 := []int{1, 2, 3}
	slice22 := make([]int, 2)

	copy(slice22, slice11)
	fmt.Println(slice11)
	fmt.Println(slice22)

	fmt.Println("--------------------------")
	var xzmap map[string]int

	xzmap = make(map[string]int)
	xzmap["keya"] = 10

	fmt.Println(xzmap["keya"])

	xymap := make(map[int]int)
	xymap[1] = 12
	fmt.Println(xymap[1])

	delete(xymap, 1)
	fmt.Println(xymap[1])

	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["O"] = "Oxygen"

	fmt.Println(elements["O"])
	fmt.Println(elements["A"])

	namex, ok := elements["A"]
	fmt.Println(namex, ok)

	if namey, ok := elements["A"]; ok {
		fmt.Println(namey, ok)
	} else {
		fmt.Println("Not in Ok")
	}

	elementz := map[string]string{
		"C": "Carbon",
		"F": "Fluorine",
	}

	fmt.Println(elementz["H"])

	elementy := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
	}

	if el, ok := elementy["N"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println("--------------------------")
	xss := []float64{98, 93, 77, 83, 92, 300, 23}
	fmt.Println(average(xss))

	xm, ym := f2returnMultipleValue()
	fmt.Println(xm, ym)

	fmt.Println(f2VariadicParams(1, 2, 3, 4))

	xsx := []int{93, 4, 98, 23, 53, 5}
	fmt.Println(f2VariadicParams(xsx...))

	sAdd := func(x, y int) int {
		return x + y
	}

	fmt.Println(sAdd(5, 6))

	// closure
	xinc := 0
	sIncrement := func() int {
		xinc++
		return xinc
	}

	fmt.Println(sIncrement())
	fmt.Println(sIncrement())

	nextEvent := makeEvenGenerator()
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())

	fmt.Println(factorial(10))

	defer sSecondFunc()
	sFirstFunc()

	// defer
	// f, _ := os.Open(fileName)
	// defer f.Close()
	fmt.Println("--------------------------")
	// defer func() {
	//	 sstr := recover()
	//	 fmt.Println(sstr)
	// }()
	// panic("PANIC here")

	fmt.Println("--------------------------")
	sXx := 5
	fmt.Println(sXx)
	zeroIt(sXx)
	fmt.Println(sXx)

	zeroItWithPointer(&sXx)
	fmt.Println(sXx)

	sXs := new(int)
	fmt.Println(*sXs)

	oneItWithPointer(sXs)
	fmt.Println(*sXs)

	fmt.Println("--------------------------")
	fmt.Println(distance(0, 0, 3, 4))
	fmt.Println(circleArea(0, 0, 3))

	fmt.Println("--------------------------")
	var xsP Point
	var xsC Circle
	var xsI float32

	fmt.Println(xsP)
	fmt.Println(xsC)
	fmt.Println(xsI)

	xxP := new(Point)
	xxC := new(Circle)

	fmt.Println(xxP) // Pointer
	fmt.Println(xxC) // Pointer

	xyP := Point{x: 0, y: 0}
	xyC := Circle{origin: xyP, r: 5}

	fmt.Println(xyC)
	fmt.Println(&xyC) // Pointer

	xyP.x = 5
	xyP.y = 5
	fmt.Println(xyP)

	xyC.r = 5
	fmt.Println("The origin is:", xyC.origin)
	fmt.Println(circleAreaOverload(xyC))

	r := Rectangle{0, 0, 10, 20}
	fmt.Println(r.area())

	fmt.Println("--------------------------")
	xsAP := new(Android)
	xsAP.Person.Talk()

	xifRect1 := Rectangle{0, 0, 4, 3}
	xifRect2 := Rectangle{1, 1, 8, 9}

	xifRects := [2]Rectangle{xifRect1, xifRect2}
	fmt.Println(totalArea(&xifRects[0], &xifRects[1]))

	multiShapeXs := MultiShape{
		shapes: []Shape{
			&Rectangle{0, 0, 4, 3},
			&Rectangle{1, 1, 8, 9},
		},
	}

	fmt.Println(multiShapeXs.area())

	fmt.Println("--------------------------")
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test8jsdfsf", "d"))
	fmt.Println(strings.Join([]string{"138", "8743", "3117"}, "-"))
	fmt.Println(strings.Repeat("8", 3))
	fmt.Println(strings.Replace("aaaaaaa", "a", "b", 5))
	fmt.Println(strings.Split("138-0983-3275", "-"))
	fmt.Println(strings.ToLower("EjsdfTJAJDFAENdkei"))

	// byte[] <-> string
	sArr := []byte("testdfs")
	sStr := string([]byte{'t', 'e', 's', 't'})

	fmt.Println(sArr)
	fmt.Println(sStr)

	fmt.Println("--------------------------")
	var buf bytes.Buffer
	buf.Write([]byte("test"))

}
