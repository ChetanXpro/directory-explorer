package main

import (
	"fmt"
	"os"
)

func main() {

	stop := false

	for stop == false {

		step := 0

		files, _ := os.Open(".")

		defer files.Close()

		filesInfo, _ := files.Readdir(-1)

		i := 0
		for _, file := range filesInfo {

			if file.IsDir() {
				i++
				fmt.Println("Dir", i, ":", file.Name())

			} else {
				fmt.Println("File:", file.Name())
			}
		}

		fmt.Println("Choose a Step")
		fmt.Scan(&step)

		fmt.Println("Choosen step", step)

		stop = true

	}

}
