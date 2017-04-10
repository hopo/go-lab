package main

import "math"

// Shape method 원형들의 정의(선언)
type Shape interface {
	area() float64
	perimeter() float64
}

//Rect struct 정의
type Rect struct {
	width, height float64
}

//Circle struct 정의
type Circle struct {
	radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 사용
func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() //인터페이스 메서드 호출
		println(a)
	}
}
