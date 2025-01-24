package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var a [9]int

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type Point struct {
	X, Y int
}

type Movie struct {
	Title  string
	Year   int  `json:"realised"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	qk := [...]int{1, 2, 2, 3}
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for i, v := range qk {
		fmt.Printf("%d %d\n", i, v)
	}

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	var movies = []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Huphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"},
		},
		{
			Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}

	data, err := json.MarshalIndent(movies, "", "  ")

	if err != nil {
		log.Fatalf("Сбой! в JSON: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println(1990 % 100)

}

func CountPositivesSumNegatives(numbers []int) []int {
	var res [2]int = [2]int{0, 0}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			res[0] += numbers[i]
		} else {
			res[1] += numbers[i]
		}
	}

	return res // your code here
}
