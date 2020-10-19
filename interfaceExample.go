package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow string

func (animal Cow) Eat() {
	fmt.Println("grass")
}
func (animal Cow) Move() {
	fmt.Println("walk")
}
func (animal Cow) Speak() {
	fmt.Println("moo")
}

type Bird string

func (animal Bird) Eat() {
	fmt.Println("worms")
}
func (animal Bird) Move() {
	fmt.Println("fly")
}
func (animal Bird) Speak() {
	fmt.Println("peep")
}

type Snake string

func (animal Snake) Eat() {
	fmt.Println("mice")
}
func (animal Snake) Move() {
	fmt.Println("slither")
}
func (animal Snake) Speak() {
	fmt.Println("hsss")
}

func queryRespond(animal Animal, i string) {
	switch s := animal.(type) {
	case Cow:
		switch i {
		case "eat", "Eat":
			s.Eat()
		case "speak", "Speak":
			s.Speak()
		case "move", "Move":
			s.Move()
		}
	case Bird:
		switch i {
		case "eat", "Eat":
			s.Eat()
		case "speak", "Speak":
			s.Speak()
		case "move", "Move":
			s.Move()
		}
	case Snake:
		switch i {
		case "eat", "Eat":
			s.Eat()
		case "speak", "Speak":
			s.Speak()
		case "move", "Move":
			s.Move()
		}
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string]Animal)
	var animal Animal
	fmt.Printf("> ")
	for scanner.Scan() {
		data := scanner.Text()
		d := strings.Split(data, " ")
		switch d[0] {
		// In case newanimal, d :=["newanimal", name, type]
		case "newanimal":
			switch d[2] {
			case "cow", "Cow":
				animal = Cow(d[1])
			case "bird", "Bird":
				animal = Bird(d[1])
			case "snake", "Snake":
				animal = Snake(d[1])
			}
			m[d[1]] = animal
			fmt.Println("Created it!")
		// In case query, d :=["query", name, information requested, i.e eat, move, speak]
		case "query":
			if _, ok := m[d[1]]; ok {
				queryRespond(m[d[1]], d[2])
			} else {
				fmt.Println("Given animal doesn't exist")
			}
		}
		fmt.Printf("> ")
	}
}
