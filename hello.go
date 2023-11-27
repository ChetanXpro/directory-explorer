package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	stop := false
	totalsteps := 0
	lastCommand := ""

	move_back := "../"

	for stop == false {

		step := 0

		var Currfiles *os.File

		if totalsteps == 0 {
			files, _ := os.Open(".")
			Currfiles = files
		} else {

			files, _ := os.Open(lastCommand)
			Currfiles = files
		}

		// println(Currfiles.Chdir())

		defer Currfiles.Close()

		filesInfo, _ := Currfiles.Readdir(-1)

		i := 1

		filesMap := make(map[int]string)

		filesMap[1] = move_back

		for _, file := range filesInfo {
			if file.IsDir() {
				i++
				// fmt.Println("Dir", i, ":", file.Name(), file.Size(), "Bytes")

				filesMap[i] = file.Name()

			} else {
				// fmt.Println("File:", file.Name())
			}
		}

		// fmt.Println("map:", filesMap)

		keys := make([]int, 0, len(filesMap))
		for k := range filesMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for key, element := range filesMap {
			fmt.Println("Step:", key, "=>", element)
		}

		fmt.Println("Choose a Step")
		fmt.Scan(&step)

		lastCommand = filesMap[step]
		totalsteps++

		continue

	}

}
