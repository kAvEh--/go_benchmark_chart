package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println("Invalid input files...")
		return
	}
	// get chart header from CLI
	chartNamePtr := flag.String("name", "Benchmark Chart", "char name")
	//TODO get more parameters from CLI here ....
	flag.Parse()
	for _, arg := range argsWithoutProg {
		if !strings.Contains(arg, "-") {
			fmt.Println(arg)
			readFile(arg)
		}
	}

	fmt.Println(">>>>", *chartNamePtr)
}

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}
	defer f.Close()

	in := f

	buf := bufio.NewScanner(in)

	for i := 0; i < 5; i++ {
		if !buf.Scan() {
			break
		}
		fmt.Println(buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}
}
