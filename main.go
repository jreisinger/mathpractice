package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Exercise struct {
	Sign         string
	X, Y, Result int
}

type Page struct {
	Exercises []Exercise
}

func main() {
	http.HandleFunc("/", handler)
	addr := "localhost:8000"
	log.Printf("starting to listen at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

const minXY = 2

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("layout")
	tmpl := template.Must(t.Parse(layout))
	maxXY, err := parseInput(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	var exercises []Exercise
	exercises = append(exercises, plus(minXY, maxXY))
	exercises = append(exercises, minus(minXY, maxXY))
	exercises = append(exercises, div(minXY, maxXY))
	exercises = append(exercises, mult(minXY, maxXY/2))
	data := Page{Exercises: exercises}
	tmpl.Execute(w, data)
}

// parseInput parses input string into an integer number. If input cannot be
// converted to a number greater than 0 and less then thousand, an error is
// returned. The function is meant to sanitize the user input.
func parseInput(input string) (upto int, err error) {
	s := strings.TrimPrefix(input, "/")
	upto, err = strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("URL path %q can't be converted into an integer", s)
	}
	if upto <= 0 || upto > 1000 {
		return 0, fmt.Errorf("URL path %q can't be converted into an integer greater than 0 and less than 1000", s)
	}
	return upto, nil
}

// randInt returns a random integer in the half-open interval [min, max). If min
// > max, 0 is returned.
func randInt(min, max int) int {
	if min > max {
		return 0
	}
	if min == max { // exact number
		return min
	}
	return rand.Intn(max-min) + min
}

func plus(minXY, maxXY int) Exercise {
	x := randInt(minXY, maxXY)
	y := randInt(minXY, maxXY)
	result := x + y
	return Exercise{
		Sign:   "+",
		X:      x,
		Y:      y,
		Result: result,
	}
}

func minus(minXY, maxXY int) Exercise {
	y := randInt(minXY, maxXY)
	x := randInt(y, maxXY)
	result := x - y
	return Exercise{
		Sign:   "-",
		X:      x,
		Y:      y,
		Result: result,
	}
}

func div(minXY, maxXY int) Exercise {
	x := randInt(minXY, maxXY)
	y := randInt(minXY, maxXY)
	for y == 0 || x%y != 0 || x == y {
		if x == 0 && y == 0 { // to avoid infinite loop
			return Exercise{Sign: ":"}
		}
		x = randInt(minXY, maxXY)
		y = randInt(minXY, maxXY)
	}
	result := x / y
	return Exercise{
		Sign:   ":",
		X:      x,
		Y:      y,
		Result: result,
	}
}

func mult(minXY, maxXY int) Exercise {
	x := randInt(minXY, maxXY)
	y := randInt(minXY, maxXY)
	result := x * y
	return Exercise{
		Sign:   "x",
		X:      x,
		Y:      y,
		Result: result,
	}
}
