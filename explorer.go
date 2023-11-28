package main

import (
	"fmt"
	"os"
)

func main() {

	stop := false
	totalsteps := 0
	lastCommand := ""

	move_back := "../"
	var Currfiles *os.File
	if totalsteps == 0 {

		files, _ := os.Open(".")
		Currfiles = files
	}
	defer Currfiles.Close()

	for stop == false {

		step := 0

		if totalsteps != 0 {

			os.Chdir(lastCommand)

		}

		files, _ := os.Open(".")
		Currfiles = files

		filesInfo, _ := Currfiles.Readdir(-1)

		i := 1

		filesMap := make(map[int]string)

		filesMap[1] = move_back

		fmt.Println("Back:", 1, "=>", move_back)

		for _, file := range filesInfo {
			if file.IsDir() {
				i++

				filesMap[i] = file.Name()

				folderName := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, file.Name())

				fmt.Println("Dir :", i, "=>", folderName)

			} else {

				fileName := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, file.Name())

				fmt.Println("File:", "=>", fileName)
			}
		}

		fmt.Println("Enter Dir number to Enter")
		fmt.Scan(&step)

		lastCommand = filesMap[step]
		totalsteps++

		continue

	}

}
