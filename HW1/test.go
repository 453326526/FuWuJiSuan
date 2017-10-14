package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"flag"
	"os/exec"
)



func main() {
	progname = os.Args[0]

	var args selpgArgs

	initArgs(&args)
	processArgs(&args)
	processInput(&args)
}
