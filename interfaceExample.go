package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

type rectangle struct {
	x, y float64
}

func (r rectangle) area() float64 {
	return r.x * r.y
}
func (r rectangle) perim() float64 {
	return 2 * (r.x + r.y)
}

type triangle struct {
	a, b, c float64
}

func (t triangle) area() float64 {
	return t.a + t.b + t.c
}
func (t triangle) perim() float64 {
	s := t.area()
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func measure(g geometry, i string) {
	switch shape := g.(type) {
	case circle:
		shape.area()
		shape.perim()
	case rectangle:
		shape.area()
		shape.perim()
	case triangle:
		shape.area()
		shape.perim()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var g geometry
	fmt.Printf("> ")
	for scanner.Scan() {
		data := scanner.Text()
		switch data {
		case "circle", "Circle":
			fmt.Printf("Radius = ")
			scanner.Scan()
			d := scanner.Text()
			r, _ := strconv.ParseFloat(d, 64)
			g = circle{r}
		case "rectangle", "Rectangle":
			fmt.Printf("length and width = ")
			scanner.Scan()
			d := scanner.Text()
			slice := strings.Split(d, " ")
			l, _ := strconv.ParseFloat(slice[0], 64)
			w, _ := strconv.ParseFloat(slice[1], 64)
			g = rectangle{l, w}
		case "triangle", "Triangle":
			fmt.Printf("sides = ")
			scanner.Scan()
			d := scanner.Text()
			slice := strings.Split(d, " ")
			a, _ := strconv.ParseFloat(slice[0], 64)
			b, _ := strconv.ParseFloat(slice[1], 64)
			c, _ := strconv.ParseFloat(slice[2], 64)
			g = triangle{a, b, c}
		}
		fmt.Printf("Area      : %0.3f\n", g.area())
		fmt.Printf("Perimeter : %0.3f\n", g.perim())
		fmt.Printf("> ")
	}
}
