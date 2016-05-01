package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"./ConsoleManager"
	"./Util"
	"./Tour"
)

func main() {
	buffer := ""

	console := &ConsoleManager.ConsoleManager{}
	util := &Util.Util{}
	tour := &Tour.Tour{}

	console.Reader = bufio.NewReader(os.Stdin)

	tour.Pathspliter = util.SupportMultiPlatform()

	if tour.Pathspliter == "UnKnown" {
		log.Fatal("Unsupported Operating System")
	}

	// Scanln is very slow~~
	for {
		fmt.Print("재귀탐색할 경로를 입력하세요 : ")
		console.ReadFromConsole(&buffer)

		if util.CheckIsExist(buffer) == true {
			tour.Explore(buffer)
		} else {
			fmt.Println("올바르지 않은 경로입니다.")
		}
	}

}