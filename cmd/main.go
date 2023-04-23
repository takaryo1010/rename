package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/takaryo1010/rename/internal/rename"
)

func main() {

	var (
		a = flag.String("a", "", "Renames all files in the specified directory to the argument + number.\nUsage:'-a NewName")
		p = flag.String("p", "", "Changes the specified part of a file name in the specified directory to the argument.\nUsage:'-p FromName,ToName")
	)
	flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s -p or -a [ARGUMENTS] [PATH] \n", os.Args[0])
        fmt.Fprintln(os.Stderr, "This is my command.")
        fmt.Fprintln(os.Stderr, "Options:")
        flag.PrintDefaults()
    }

	flag.Parse()


	if *a !="" && *p !=""{
		fmt.Println("-a and -p cannot be used together")
		os.Exit(0)
	} else if flag.NArg() != 1 ||flag.NFlag()==0 {
		fmt.Println("rename: missing operand")
		fmt.Println("Try 'rename -h' for more information")
		os.Exit(0)
	} else if *a !=""{
		// func
		rename.RenameAll(flag.Arg(0),*a)

	}else if *p !=""{
		// func
		rename.RenamePart(flag.Arg(0),*p)
		os.Exit(0)
	}



	
}