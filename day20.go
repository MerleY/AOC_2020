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
	input := input.Load("20").ToDoubleStringGroupedArray()
	re := regexp.MustCompile(`^Tile\s(\d+):`)
	allPictures := []picture{}
	for _, lines := range input {
		matches := re.FindStringSubmatch(lines[0])
		if len(matches) != 2 {
			log.Fatal("Parsing error")
		}
		id, _ := strconv.Atoi(matches[1])
		p := picture{id: id, orient: 0}
		p.pixels = make([][]byte, len(lines[1:]))
		for i, subline := range lines[1:] {
			p.pixels[i] = make([]byte, len(subline))
			for j, c := range subline {
				p.pixels[i][j] = byte(c)
			}
		}
		allPictures = append(allPictures, p)
	}

	// part 1
	productAngle := 1
	for i, p := range allPictures {
		fitCount := 0
		for j, p2 := range allPictures {
			if i == j {
				continue
			}
			if p.fitAll(p2) {
				fitCount++
			}
		}
		fmt.Printf("Image %v, fits: %v\n", p.id, fitCount)
		if fitCount == 2 {
			productAngle *= p.id
		}
	}

	fmt.Printf("Star 1: %v\n", productAngle)

	bigPic := BigPicture{content: make(map[int][int]*picture)}
	for {
		for _, p := range allPictures {
			bigPic.place(p)
		}
	}

}

type BiGPicture struct {
	content map[int]map[int]*picture
}

func (bp *BigPicture) place(pic picture) []int {
	if len(bp.content) == 0 {
		bp[0][0] = &p
	} else {
		for x, my := range bp.content {
			for y, p := range my {
			}
		}
	}
}

type picture struct {
	pixels [][]byte
	id     int
	orient int
}

func (p *picture) fitAll(p2 picture) bool {
	for i := 0; i < 8; i++ {
		if arrays.Equal(p.getBottom(), p2.getTop()) || arrays.Equal(p.getTop(), p2.getBottom()) || arrays.Equal(p.getLeft(), p2.getRight()) || arrays.Equal(p.getRight(), p2.getLeft()) {
			return true
		}
		p2.flip()
	}
	return false
}

func (p *picture) fitRight(p2 *picture) bool {
	for i := 0; i < 8; i++ {
		if arrays.Equal(p.getRight(), p2.getLeft()) {
			return true
		}
		p2.flip()
	}
	return false
}

func (p *picture) fitLeft(p2 *picture) bool {
	for i := 0; i < 8; i++ {
		if arrays.Equal(p.getLeft(), p2.getRight()) {
			return true, i
		}
		p2.flip()
	}
	return false, 0
}

func (p *picture) fitTop(p2 *picture) bool {
	for i := 0; i < 8; i++ {
		if arrays.Equal(p.getTop(), p2.getBottom()) {
			return true
		}
		p2.flip()
	}
	return false
}

func (p *picture) fitBottom(p2 *picture) bool {
	for i := 0; i < 8; i++ {
		if arrays.Equal(p.getBottom(), p2.getTop()) {
			return true
		}
		p2.flip()
	}
	return false
}

func (p *picture) flip() {
	newPixels := [][]byte{}
	if p.orient%2 == 0 {
		for _, line := range p.pixels {
			newLine := arrays.Reverse(line)
			newPixels = append(newPixels, newLine)
		}
	} else if p.orient%4 == 1 {
		for i := len(p.pixels) - 1; i >= 0; i-- {
			newPixels = append(newPixels, p.pixels[i])
		}
	} else if p.orient%4 == 3 {
		for i := 0; i < len(p.pixels); i++ {
			newPixelLine := []byte{}
			for _, line := range p.pixels {
				newPixelLine = append(newPixelLine, line[i])
			}
			newPixels = append(newPixels, newPixelLine)
		}
	}

	p.pixels = newPixels
	p.orient += 1
	p.orient %= 8
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
	line := []byte{}
	for _, pix := range p.pixels {
		line = append(line, pix[0])
	}
	return line
}

func (p *picture) getRight() []byte {
	line := []byte{}
	for _, pix := range p.pixels {
		line = append(line, pix[len(pix)-1])
	}
	return line
}
