package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x float64
	y float64
	z float64
}

type Hail struct {
	pos Coord
	v   Coord
}

func main() {
	content, err := os.Open("24.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	lines := make([]Hail, 0)
	re := regexp.MustCompile(`\s+`)

	for scanner.Scan() {
		text := scanner.Text()
		text = re.ReplaceAllString(text, " ")
		splits := strings.Split(text, " @ ")
		left := strings.Split(splits[0], ", ")
		right := strings.Split(splits[1], ", ")
		var line Hail
		temp1, _ := strconv.Atoi(left[0])
		temp2, _ := strconv.Atoi(left[1])
		temp3, _ := strconv.Atoi(left[2])
		line.pos = Coord{float64(temp1), float64(temp2), float64(temp3)}
		temp1, _ = strconv.Atoi(right[0])
		temp2, _ = strconv.Atoi(right[1])
		temp3, _ = strconv.Atoi(right[2])
		line.v = Coord{float64(temp1), float64(temp2), float64(temp3)}
		lines = append(lines, line)
	}

	fmt.Println(day24(lines))
}

func day24(hails []Hail) int {
	total := 0
	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			total += testarea(intersect(hails[i], hails[j]))
		}
	}
	return total
}

func intersect(l1, l2 Hail) Coord {
	if l1.v.x*l2.v.y == l1.v.y*l2.v.x {
		return Coord{-1, -1, -1}
	}
	t1 := ((l1.pos.y-l2.pos.y)*(l2.v.x) - (l1.pos.x-l2.pos.x)*(l2.v.y)) / ((l1.v.x * l2.v.y) - (l1.v.y * l2.v.x))
	if t1 < 0 {
		return Coord{-1, -1, -1}
	}
	//t2 := ((l1.pos.y-l2.pos.y)*(l1.v.x)-(l1.pos.x-l2.pos.x)*(l2.v.y))/((l1.v.x*l2.v.y)-(l1.v.y*l2.v.x))
	Px := l1.pos.x + t1*l1.v.x
	Py := l1.pos.y + t1*l1.v.y
	t2 := (Px - l2.pos.x) / l2.v.x
	if t2 < 0 {
		return Coord{-1, -1, -1}
	}
	//fmt.Println(l1, l2, Px, Py, l2.pos.x + t2*l2.v.x, l2.pos.y + t2*l2.v.y)
	return Coord{Px, Py, 0}
}

func testarea(intersection Coord) int {
	high := 400000000000000.0
	low := 200000000000000.0
	//high := 20.0
	//low := 10.0
	if intersection.x < low || intersection.x > high || intersection.y < low || intersection.y > high {
		return 0
	}
	return 1
}
