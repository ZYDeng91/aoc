package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
	z int
}

type Brick struct {
	start Coord
	end   Coord
}

func main() {
	content, err := os.Open("22.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	lines := make([]Brick, 0)

	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), "~")
		left := strings.Split(splits[0], ",")
		right := strings.Split(splits[1], ",")
		var line Brick
		temp1, _ := strconv.Atoi(left[0])
		temp2, _ := strconv.Atoi(left[1])
		temp3, _ := strconv.Atoi(left[2])
		line.start = Coord{temp1, temp2, temp3}
		temp1, _ = strconv.Atoi(right[0])
		temp2, _ = strconv.Atoi(right[1])
		temp3, _ = strconv.Atoi(right[2])
		line.end = Coord{temp1, temp2, temp3}
		lines = append(lines, line)
	}

	sort.Slice(lines, func(i, j int) bool { return lines[i].start.z < lines[j].start.z })
	fmt.Println(day22(lines))
}

func day22(bricks []Brick) int {
	total := 0
	dependency := make(map[Brick][]Brick)
	dependencyOf := make(map[Brick][]Brick)
	newBricks := make([]Brick, 0)
	top := make([][]Brick, 10)
	visited := make(map[Brick]bool)
	for x := 0; x < 10; x++ {
		top1 := make([]Brick, 10)
		for y := 0; y < 10; y++ {
			top1[y] = Brick{Coord{0, 0, 0}, Coord{0, 0, 0}}
		}
		top[x] = top1
	}
	for _, brick := range bricks {
		//fmt.Println(brick.end.x-brick.start.x+brick.end.y-brick.start.y+brick.end.z-brick.start.z)
		candidates := make([]Brick, 0)
		seenCandidate := make(map[Brick]bool)
		for x := min(brick.start.x, brick.end.x); x <= max(brick.end.x, brick.start.x); x++ {
			for y := min(brick.start.y, brick.end.y); y <= max(brick.end.y, brick.end.y); y++ {
				if seenCandidate[top[x][y]] {
					continue
				}
				candidates = append(candidates, top[x][y])
				seenCandidate[top[x][y]] = true

			}
		}
		sort.Slice(candidates, func(i, j int) bool { return candidates[i].end.z > candidates[j].end.z })

		newBrick := Brick{Coord{brick.start.x, brick.start.y, candidates[0].end.z + 1}, Coord{brick.end.x, brick.end.y, brick.end.z - brick.start.z + candidates[0].end.z + 1}}
		for x := brick.start.x; x <= brick.end.x; x++ {
			for y := brick.start.y; y <= brick.end.y; y++ {
				top[x][y] = newBrick
			}
		}
		newBricks = append(newBricks, newBrick)
		visited[newBrick] = true
		for _, candidate := range candidates {
			if candidate.end.z == candidates[0].end.z {
				dependency[newBrick] = append(dependency[newBrick], candidate)
				dependencyOf[candidate] = append(dependencyOf[candidate], newBrick)
			} else {
				break
			}
		}
	}

	for _, val := range dependency {
		if len(val) == 1 {
			visited[val[0]] = false
		}
	}
	for _, val := range visited {
		if val {
			total += 1
		}
	}
	total2 := 0
	for _, brick := range newBricks {
		total2 += day22_2(brick, dependency)
	}
	fmt.Println(total2)
	return total
}

func day22_2(brick Brick, dependency map[Brick][]Brick) int {
	total := 0
	remove := []Brick{brick}
	dependency_copy := make(map[Brick][]Brick)
	for key, val := range dependency {
		dependency_copy[key] = val
	}

	for len(remove) > 0 {
		remove_new := make([]Brick, 0)
		for key, val := range dependency_copy {
			if len(val) == 0 {
				continue
			}
			removed := make([]Brick, 0)
			for _, item := range val {
				for _, fallen := range remove {
					if fallen == item {
						goto removal
					}
				}
				removed = append(removed, item)

			removal:
				continue
			}
			dependency_copy[key] = removed
			if len(removed) == 0 {
				remove_new = append(remove_new, key)
				total += 1
			}
		}
		remove = remove_new
	}

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
