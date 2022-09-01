package main

// This only works partially at the moment.
// todo: Add functionality read a file's contents and print them out with funny colors
import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Color struct {
	a, b, c int
}

// Streamline error checking proccess
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	info, _ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes")
		fmt.Println("Usage: Print funny colors")
	}

	file_name := os.Args[1:]
	if len(os.Args) > 2 {
		fmt.Println("Only enter one text file, please.")
		return
	}
	file_join := strings.Join(file_name, "")
	file, err := os.Open(file_join)
	check(err)

	reader := bufio.NewReader(file)
	for {
		// Main loop. Each iteration a new 'random' is created and will generate a random color for each rune.
		s1 := rand.NewSource(time.Now().UnixNano())
		s2 := rand.New(s1)
		rgb := Color{s2.Intn(250), s2.Intn(250), s2.Intn(250)}

		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", rgb.a, rgb.b, rgb.c, input)
	}
}

//	var phrases []string
//	for i := 1; i < 3; i++ {
//		phrases = append(phrases, faker.Hacker().Phrases()...)
//	}
