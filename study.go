package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// default to Pomodoro times
var cycles = 1
var focus = 25
var rest = 5

func init() {
	var err error
	args := os.Args
	for n, arg := range args {
		switch n {
		case 1:
			if focus, err = strconv.Atoi(arg); err != nil {
				panic(err)
			}
		case 2:
			if rest, err = strconv.Atoi(arg); err != nil {
				panic(err)
			}
		case 3:
			if cycles, err = strconv.Atoi(arg); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	fmt.Printf("Running %v cycles of %v : %v intervals of study, rest\n", cycles, focus, rest)

	for i := 0; i < cycles; i++ {
		remaining := focus
		for range time.Tick(1 * time.Minute) {
			remaining--
			if remaining == 0 {
				break
			}
			fmt.Printf("%v : %v\n", remaining, rest)
		}

		remaining = rest
		fmt.Printf("Break for %v minutes\n", rest)
		for range time.Tick(1 * time.Minute) {
			remaining--
			if remaining == 0 {
				break
			}
			fmt.Println("Break remaining:")
			fmt.Printf("%v : %v\n", focus, remaining)
		}
		fmt.Printf("%v cycles remaining\n", cycles-1)
	}

	total := cycles * (rest + focus)
	fmt.Printf("Done! You studied for %v minutes\n", total)

}
