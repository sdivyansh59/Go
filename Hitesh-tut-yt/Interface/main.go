package main

// video link
//https://www.youtube.com/watch?v=lbW-KVdIXaY&ab_channel=TheNetNinja

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square method
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle method
func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) circumf()float64 {
	return 2 * math.Pi * c.radius
}

// interface ( it can take both square and circle)
type Shape interface {
	area() float64
	circumf() float64
}


// we can pass circle or square struct 
func printShapeInfo (s Shape) {
	fmt.Printf("area of %T is %0.2f\n",s, s.area())
	fmt.Printf("circumference of %T is : %0.2f\n", s, s.circumf())
}



func main () {
	shapes := []Shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("....")
	}

}
