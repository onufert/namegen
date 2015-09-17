package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	otherWord = "*"
	pre       = 0
	none      = 1
	post      = 2
)

var positions = []int{
	pre, none, post,
}

var words = []string{
	"app",
	"site",
	"time",
	"get",
	"go",
	"lets",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		p := positions[rand.Intn(len(positions))]
		w := words[rand.Intn(len(words))]
		t := otherWord
		switch p {
		case pre:
			t = w + otherWord
		case post:
			t = otherWord + w
		}
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
