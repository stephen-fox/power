package main

import (
	"log"
	"os"

	"github.com/stephen-fox/power"
)

func main() {
	power, err := power.Get()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-s":
			handleResult(power.Sleep())
		case "-r":
			handleResult(power.Restart())
		case "-o":
			handleResult(power.Shutdown())
		case "-h":
			printHelp()
		default:
			log.Println("Please supply valid arguments")
			os.Exit(1)
		}
	} else {
		printHelp()
	}

	os.Exit(0)
}

func handleResult(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func printHelp() {
	log.Println("power - Manage a machine's power state.\n\n",
		"Specify action with second argument: <application> [-s, -r, -o]\n",
		"    -s     Sleep the machine\n",
		"    -r     Restart the machine\n",
		"    -o     Power the machine off\n")
}
