package main

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type rectangle struct{
	width, height float64
}

type circle struct{
	radius float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

func (c circle) area() float64{
	return math.Pi * c.radius*c.radius
}

func getArea(shape shape) float64{
	return shape.area()
}

func demo1(){
	R:=rectangle{3,4}
	C:=circle{5}
	fmt.Println(getArea(R))
	fmt.Println(getArea(C))
}