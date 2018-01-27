package main

import (
	"os"
	"fmt"
)

func main()  {
	var ss string
	for i := 1; i< len(os.Args); i++ {
		ss += os.Args[i] + "\n"
	}
	fmt.Printf("%s", ss)

}
