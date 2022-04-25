package interfaces

import "fmt"

// Stringer

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s string
	for i := 0; i < len(ip); i++ {
		s += fmt.Sprint(".", ip[i])
	}
	return s[1:]
}

func runIPAddrExampl() {
	var ip = IPAddr{1, 2, 3, 4}
	fmt.Println(ip)
}

type Car struct {
	model   string
	maker   string
	carType string
}

func (c Car) String() string {
	return "Model " + c.model + " is a new addition of " + c.maker
}

func runCarExample() {
	c := Car{model: "S", maker: "Tesla"}
	fmt.Println(c)
}

func runTypeSwitch() {
	typeSwitch("hello")
	typeSwitch(2222)
	typeSwitch(3i + 2)

	typeSwitch(&Vertex{})
}
func typeSwitch(i interface{}) {

	switch t := i.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)
	default:
		fmt.Println("default", t)
	}

}

func typeAssertion() {
	var i interface{} = "Hello"

	var s string
	s, ok := i.(string)
	fmt.Println(ok, s)
}

func interfaceExample() {
	var f Abser

	var v *Vertex
	f = v
	describe(f)
	f.Abs()

	var any Any
	any = f
	fmt.Println(any.Abs())
	// f = MyFloat64(0)

	// f = &Vertex{}

	// f.Abs()
	// fmt.Println(f.Abs())
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Abser interface {
	Abs() float64
}

type Any interface {
	Abs() float64
}

type Vertex struct {
	x, y float64
}

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (p *Vertex) Abs() float64 {
	if p == nil {
		return 0
	}
	return p.x*p.x + p.y*p.y
}
