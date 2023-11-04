package main

import (
	"fmt"
	"html/template"
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
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("layout")
	tmpl := template.Must(t.Parse(layout))
	upto, err := parseURLPath(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	var exercises []Exercise
	exercises = append(exercises, plus(upto))
	exercises = append(exercises, minus(upto))
	exercises = append(exercises, div(upto))
	exercises = append(exercises, mult(upto))
	data := Page{Exercises: exercises}
	tmpl.Execute(w, data)
}

// parseURLPath parses path into an integer number. The number must be greater
// than 0 and less then thousand. It's meant to sanitize user input.
func parseURLPath(path string) (upto int, err error) {
	s := strings.TrimPrefix(path, "/")
	upto, err = strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("URL path %q is not an integer", s)
	}
	if upto <= 0 || upto > 1000 {
		return 0, fmt.Errorf("URL path %q is not an integer greater than 0 and less than 1000", s)
	}
	return upto, nil
}

func plus(upto int) Exercise {
	if upto <= 0 {
		return Exercise{}
	}
	x := rand.Intn(upto)
	y := rand.Intn(upto)
	result := x + y
	for result > upto {
		x = rand.Intn(upto)
		y = rand.Intn(upto)
		result = x + y
	}
	return Exercise{
		Sign:   "+",
		X:      x,
		Y:      y,
		Result: result,
	}
}

func minus(upto int) Exercise {
	if upto <= 0 {
		return Exercise{}
	}
	x := rand.Intn(upto)
	y := rand.Intn(upto)
	result := x - y
	for result < 0 || result > upto {
		x = rand.Intn(upto)
		y = rand.Intn(upto)
		result = x - y
	}
	return Exercise{
		Sign:   "-",
		X:      x,
		Y:      y,
		Result: result,
	}
}

func div(upto int) Exercise {
	if upto <= 0 {
		return Exercise{}
	}
	x := rand.Intn(upto)
	y := rand.Intn(upto)
	for y == 0 {
		y = rand.Intn(upto)
	}
	result := x / y
	for result > upto || x%y != 0 {
		x = rand.Intn(upto)
		y = rand.Intn(upto)
		for y == 0 {
			y = rand.Intn(upto)
		}
		result = x / y
	}
	return Exercise{
		Sign:   ":",
		X:      x,
		Y:      y,
		Result: result,
	}
}

func mult(upto int) Exercise {
	if upto <= 0 {
		return Exercise{}
	}
	x := rand.Intn(upto)
	y := rand.Intn(upto)
	result := x * y
	for result > upto {
		x = rand.Intn(upto)
		y = rand.Intn(upto)
		result = x * y
	}
	return Exercise{
		Sign:   "x",
		X:      x,
		Y:      y,
		Result: result,
	}
}
