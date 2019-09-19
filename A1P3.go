package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

var wgTracker sync.WaitGroup

var lowChan = make(chan int, 1)
var highChan = make(chan int, 1)

type Point struct {
	x float64
	y float64
}
type Triangle struct {
	A Point
	B Point
	C Point
}

func triangles10000() (result [10000]Triangle) {
	rand.Seed(2120)
	for i := 0; i < 10000; i++ {
		result[i].A = Point{rand.Float64() * 100., rand.Float64() * 100.}
		result[i].B = Point{rand.Float64() * 100., rand.Float64() * 100.}
		result[i].C = Point{rand.Float64() * 100., rand.Float64() * 100.}
	}
	return
}
func (t Triangle) Perimeter() float64 {
	var side1, side2, side3 float64
	side1 = math.Sqrt(((t.A.x - t.B.x) * (t.A.x - t.B.x)) + ((t.A.y - t.B.y) * (t.A.y - t.B.y)))
	side2 = math.Sqrt(((t.B.x - t.C.x) * (t.B.x - t.C.x)) + ((t.B.y - t.C.y) * (t.B.y - t.C.y)))
	side3 = math.Sqrt(((t.A.x - t.C.x) * (t.A.x - t.C.x)) + ((t.A.y - t.C.y) * (t.A.y - t.C.y)))
	peri := side1 + side2 + side3
	return peri

}
func (t Triangle) Area() float64 {
	var side1, side2, side3 float64
	
	side1 = (t.A.x )* (t.B.y - t.C.y) 
	side2 = (t.B.x )* (t.C.y - t.A.y)
	side3 = (t.C.x )* (t.A.y - t.B.y)
	
	value1:= side1 +side2 + side3
	value2:= (0.5)* value1
	value3:= math.Abs(value2)
	
	return value3
}

type Stack struct {
	myTriangles []Triangle
	
}

func (s *Stack) push(t *Triangle) {
	s.myTriangles = append(s.myTriangles, *t)
}
func (s *Stack) pop() (t *Triangle) {
	var temp *Triangle
	if len(s.myTriangles) > 0 {
		temp = &s.myTriangles[len(s.myTriangles)-1]
		s.myTriangles = s.myTriangles[:len(s.myTriangles)-1]
	}
	return temp
}

func (s *Stack) peek() (t Triangle) {
	temp := s.myTriangles[len(s.myTriangles)-1]
	return temp
}

func classifyTriangles(highRatio *Stack, lowRatio *Stack, ratioThreshold float64, triangles []Triangle) {
	var ratio float64
	
	for _, i := range triangles {
		ratio = (i.Perimeter()) / (i.Area())
		
		if ratio > ratioThreshold {
			highChan <- 1
			highRatio.push(&i)
			
			<-highChan

		} else {
			lowChan <- 1
			lowRatio.push(&i)
			
			<-lowChan
		}

	}
	
	wgTracker.Done()
}
func main() {
	myTriangles := triangles10000()
	var highStack, lowStack Stack
	for i := 0; i < 10; i++ {
		triangleSlices := myTriangles[i*1000 : (i+1)*1000]

		wgTracker.Add(1)
		go classifyTriangles(&highStack, &lowStack, 1.0, triangleSlices) 

	}

	wgTracker.Wait()
	fmt.Println("number of triangles in lower Stack:  ", len(lowStack.myTriangles))
	fmt.Println("top item in lower Stack is :  ", lowStack.peek())

	fmt.Println("number of triangles in higher Stack:  ", len(highStack.myTriangles))
	fmt.Println("top item in higher Stack is:  ", highStack.peek())
	
}
