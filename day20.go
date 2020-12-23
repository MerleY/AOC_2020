package main

import (
	"aoc/arrays"
	"aoc/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func day20() {
	input := input.Load("20-test").ToDoubleStringGroupedArray()
	re := regexp.MustCompile(`^Tile\s(\d+):`)
	allPictures := []picture{}
	for _, lines := range input {
		matches := re.FindStringSubmatch(lines[0])
		if len(matches) != 2 {
			log.Fatal("Parsing error")
		}
		id, _ := strconv.Atoi(matches[1])
		p := picture{id: id, orient: 1}
		p.pixels = make([][]byte, len(lines[1:]))
		for i, subline := range lines[1:] {
			p.pixels[i] = make([]byte, len(subline))
			for j, c := range subline {
				p.pixels[i][j] = byte(c)
			}
		}
		allPictures = append(allPictures, p)
	}
	//bigPic := BigPicture{content: make(map[int][int]*picture)}
	//picturesPlaced := []int{allPictures[0].id}
	fmt.Println(len(allPictures))
	ids := []int{}
	for i, p := range allPictures {
		sides := []string{}
		for j, p2 := range allPictures {
			if i == j {
				continue
			}
			if side, _ := p.fitBoth(p2); side != "" {
				// if !arrays.StringIn(side, sides) {
				sides = append(sides, side)
				// }
			}
			if len(sides) >= 2 {
				break
			}
		}
		if len(sides) < 2 {
			ids = append(ids, p.id)
		}
	}
	fmt.Println(len(ids))
	fmt.Println(ids)
	//bigPic.add(0, 0, allPicture[0])
}

type picture struct {
	pixels [][]byte
	id     int
	orient int
}

// type BiGPicture struct {
// 	content map[int]map[int]*picture
// }

func (p *picture) flip() {
	newPixels := [][]byte{}
	if p.orient == 0 || p.orient == 2 {
		for _, line := range p.pixels {
			newLine := arrays.Reverse(line)
			newPixels = append(newPixels, newLine)
		}
	} else if p.orient == 1 || p.orient == 3 {
		for i := len(p.pixels) - 1; i >= 0; i-- {
			newPixels = append(newPixels, p.pixels[i])
		}
	}
	p.pixels = newPixels
	p.orient += 1
	p.orient %= 4
}

func (p *picture) show() {
	fmt.Printf("\nid:%v\n", p.id)
	for _, l := range p.pixels {
		for _, c := range l {
			fmt.Printf("%v", string(c))
		}
		fmt.Printf("\n")
	}
}

func (p *picture) getTop() []byte {
	return p.pixels[0]
}

func (p *picture) getBottom() []byte {
	return p.pixels[len(p.pixels)-1]
}

func (p *picture) getLeft() []byte {
	line := []int{}
	for _, pix := range p.pixels {
		line = append(pix[0])
	}
	return line
}

func (p *picture) getLeft() []byte {
	line := []int{}
	for _, pix := range p.pixels {
		line = append(pix[len(pix)-1])
	}
	return line
}

// func (bp *BigPicture) touch(x int, y int, picture *picture, dir int) int {

// }

// func (bp *BigPicture) add(x int, y int, picture *picture) []int {

// }
