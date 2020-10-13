package main

import ("fmt"
	"math"
)

type phone struct {
	credit uint16
	color string
	brand string
	weight float64
}

func foo() {
	fmt.Println("Hello from Golang", math.Sqrt(16))
}

func add(x float64, y float64) float64 {
	return x + y
}

func main() {
	foo()
	var num1,num2 float64 = 8.5, 3.2	
	fmt.Println(add(num1, num2))
	xiaomi := phone{credit: 1000, 
		color: "black", 
		brand: "Xiaomi", 
		weight: 50.5}
	fmt.Println("La couleur de mon téléphone est", xiaomi.color)
}