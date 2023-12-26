package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type lens struct {
	label string
	inst  string
}

func main() {
	content, err := os.ReadFile("15.in")
	if err != nil {
		return
	}

	lines := strings.Split(string(content), ",")

	total := 0
	items := make([]lens, len(lines))
	re := regexp.MustCompile(`(\w+)([-=]\d*)`)
	for i, line := range lines {
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		total += hash(line)
		split := re.FindStringSubmatch(line)
		items[i] = lens{split[1], split[2]}
	}

	fmt.Println(total)
	fmt.Println(day15(items))
}

func hash(input string) int {
	current := 0
	for _, j := range input {
		current += int(j)
		current *= 17
		current = current % 256
	}
	return current
}

func day15(items []lens) int {
	foc := make(map[string]int)
	boxes := make(map[int][]string)
	for _, item := range items {
		if item.inst[0] == '-' {
			foc[item.label] = 0
			boxes[hash(item.label)] = pop(boxes[hash(item.label)], item.label)
		}
		if item.inst[0] == '=' {
			if foc[item.label] == 0 {
				boxes[hash(item.label)] = append(boxes[hash(item.label)], item.label)
			}
			foc[item.label] = int(item.inst[1] - '0')
		}
	}

	total := 0
	for key, content := range boxes {
		for slot, label := range content {
			fmt.Println(label, key+1, slot+1, foc[label])
			total += (key + 1) * (slot + 1) * foc[label]
		}
	}
	return total
}

func pop(box []string, label string) []string {
	newBox := make([]string, 0)
	for _, item := range box {
		if item != label {
			newBox = append(newBox, item)
		}
	}
	return newBox
}
