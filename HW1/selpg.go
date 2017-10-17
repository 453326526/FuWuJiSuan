package main

import (
	"fmt"
	"io"
	"flag"
	"bufio"
	"os"
)

type Selpg_args struct {
	start_page int
	end_page int
	input_filename string
	page_len int
	page_type int
	form_feed_delimited bool
}
var progname string

func main() {
	progname = os.Args[0]
	var sa Selpg_args

	FlagInit(&sa)
	process_args(&sa)
	process_input(&sa)
}

func FlagInit(sa *Selpg_args) {
	flag.Usage = usage
	flag.IntVar(&sa.start_page, "s", -1, "Start page.")
	flag.IntVar(&sa.end_page, "e", -1, "End page.")
	flag.IntVar(&sa.page_len, "l", 72, "Line number per page.")
	flag.BoolVar(&sa.form_feed_delimited, "f", false, "Determine form-feed-delimited")
	flag.Parse()
}

func process_args(sa *Selpg_args) {
	if sa.start_page == -1 || sa.end_page == -1 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}
	if sa.start_page > sa.end_page || sa.start_page < 0 ||
		sa.end_page < 0 {
		fmt.Fprintln(os.Stderr, "Invalid input")
		flag.Usage()
		os.Exit(1)
	}
}

func process_input(sa *Selpg_args) {
	if flag.NArg() > 0 {
		sa.input_filename = flag.Arg(0)
		fname, err := os.Open(sa.input_filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		reader := bufio.NewReader(fname)
		counter := 0
		for {
			line, _, err := reader.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err == io.EOF {
				break
			}
			positions := counter/sa.page_len
			if positions <= sa.end_page && positions >= sa.start_page {
				fmt.Println(string(line))
			}
			counter++
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		counter := 0
		response := ""
		for scanner.Scan() {
			line := scanner.Text()
			line += "\n"
			positions := counter/sa.page_len
			if positions <= sa.end_page && positions >= sa.start_page {
				response += line
			}
			counter++
		}
		fmt.Printf("%s", response)
	}
}

func usage() {
	fmt.Println("This is the selpg.")
	fmt.Println("\tselpg -s=Number -e=Number [options] [filename]")
	fmt.Printf("\t-s=Number\tStart from Page <number>.\n")
	fmt.Printf("\t-e=Number\tEnd to Page <number>.\n")
	fmt.Printf("\t-f\t\t[options]Specify that the pages are sperated by \\f.\n")
	fmt.Println("\t[filename] ---------- Read input from this file.")
	fmt.Println("\tIf filename is not given, it will read input from stdin. Ctrl+D to cutout.")
}

