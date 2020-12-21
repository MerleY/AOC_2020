package main

import (
	"aoc/arrays"
	"aoc/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type picture struct {
	pixels [][]byte
	id     int
	orient int
}

// type BiGPicture struct {
// 	content map[int]map[int]*picture
// }

func (p *picture) reverse() {
	p.orient *= -1
}

func (p *picture) show() {
	fmt.Printf("\nid:%v\n", p.id)
	if p.orient == 1 {
		for _, l := range p.pixels {
			for _, c := range l {
				fmt.Printf("%v", string(c))
			}
			fmt.Printf("\n")
		}
	} else {
		for i := len(p.pixels) - 1; i > 0; i-- {
			line := ""
			for _, c := range p.pixels[i] {
				line = string(c) + line
			}
			fmt.Println(line)
		}
	}
}

func (p *picture) getTop() []byte {
	top := []byte{}
	if p.orient == 1 {
		top = p.pixels[0]
	} else {
		top = arrays.Reverse(p.pixels[len(p.pixels)-1])
	}
	return top
}

func (p *picture) getBottom() []byte {
	bottom := []byte{}
	if p.orient == 1 {
		bottom = p.pixels[len(p.pixels)-1]
	} else {
		bottom = arrays.Reverse(p.pixels[0])
	}
	return bottom
}

func (p *picture) getLeft() []byte {
	left := []byte{}
	if p.orient == 1 {
		for _, pix := range p.pixels {
			left = append(left, pix[0])
		}
	} else {
		for i := len(p.pixels) - 1; i > 0; i-- {
			left = append(left, p.pixels[i][len(p.pixels[i])-1])
		}
	}
	return left
}

func (p *picture) getRight() []byte {
	right := []byte{}
	if p.orient == 1 {
		for _, pix := range p.pixels {
			right = append(right, pix[len(pix)-1])
		}
	} else {
		for i := len(p.pixels) - 1; i > 0; i-- {
			right = append(right, p.pixels[i][0])
		}
	}
	return right
}

func (p *picture) fit(p2 picture) string {
	if arrays.Equal(p.getLeft(), p2.getRight()) {
		return "left"
	}
	if arrays.Equal(p.getRight(), p2.getLeft()) {
		return "right"
	}
	if arrays.Equal(p.getTop(), p2.getBottom()) {
		return "top"
	}
	if arrays.Equal(p.getBottom(), p2.getTop()) {
		return "bottom"
	}
	return ""
}

func (p *picture) fitBoth(p2 picture) (string, bool) {
	if p.id == 1171 && (p2.id == 2473 || p2.id == 1489) {
		p.show()
		p2.show()
	}
	if side := p.fit(p2); side != "" {
		return side, false
	}
	p2.reverse()
	if p.id == 1171 && (p2.id == 2473 || p2.id == 1489) {
		p.show()
		p2.show()
	}
	if side := p.fit(p2); side != "" {
		return side, true
	}
	return "", false
}

// func (bp *BigPicture) touch(x int, y int, picture *picture, dir int) int {

// }

// func (bp *BigPicture) add(x int, y int, picture *picture) []int {

// }

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
