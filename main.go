package main

import (
	"fmt"
	"os"

	"github.com/xiechuanj/blog/cmd"
)

func init() {
	//
}

func main() {
	fmt.Println("22222222222222")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println("1111111111111")
		fmt.Println(err)
		os.Exit(-1)
	}
}
