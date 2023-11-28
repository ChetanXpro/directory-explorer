package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func main() {

	stop := false
	totalsteps := 0
	lastCommand := ""

	move_back := "../"
	exit := "Exit"
	var Currfiles *os.File
	if totalsteps == 0 {

		files, err := os.Open(".")
		Currfiles = files

		if err != nil {
			fmt.Println("error opening directory:", err)

		}
	}
	defer Currfiles.Close()

	for stop == false {

		step := 0

		if totalsteps != 0 {

			os.Chdir(lastCommand)

		}

		files, _ := os.Open(".")

		Currfiles = files

		defer Currfiles.Close()

		filesInfo, err := Currfiles.Readdir(-1)

		if err != nil {
			fmt.Println("error reading directory:", err)
			continue
		}

		i := 2

		filesMap := make(map[int]string)

		filesMap[1] = exit
		filesMap[2] = move_back

		fmt.Println("Exit:", 1, "=>", "Exit")
		fmt.Println("Back:", 2, "=>", move_back)

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

		// fmt.Println("Enter Dir number to Enter")
		// ask := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "Enter Dir number to Enter")

		// fmt.Println(ask)

		prompt := &survey.Input{
			Message: "Enter Dir number to Enter",
		}
		survey.AskOne(prompt, &step)

		lastCommand = filesMap[step]

		if lastCommand == exit {
			stop = true
			println("Exiting ...")
		}
		totalsteps++

		continue

	}

}
