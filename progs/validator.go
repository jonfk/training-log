package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type TrainingLog struct {
	Date       string `date`
	Time       string `time,omitempty`
	Length     string `length,omitempty`
	Bodyweight string `bodyweight,omitempty`
	Event      string `event,omitempty`
	Wilks      string `wilks,omitempty`
	Total      string `total,omitempty`
	Workout    []struct {
		Name   string `name`
		Weight string `weight,omitempty`
		Sets   string `sets`
		Reps   string `reps`
	} `workout`
	Notes []string `notes,omitempty`
}

var Verbose bool = false
var Flexible bool = false
var Output bool = false

func main() {
	// Parse Command Line arguments
	args := os.Args[1:] // Ignore program name

	for i := range args {
		switch args[i] {
		case "-h":
			printUsage()
			os.Exit(0)
		case "--help":
			printUsage()
			os.Exit(0)
		case "-v":
			Verbose = true
		case "--verbose":
			Verbose = true
		case "-f":
			Flexible = true
		case "--flexible":
			Flexible = true
		case "-o":
			Output = true
		case "--output":
			Output = true
		default:
			process(args[i])
		}
	}

}

func process(arg string) {

	// fmt.Println("Processing " + arg)

	isDir, err := IsDirectory(arg)

	if err != nil {
		log.Fatalf("Invalid path to file or directory\n")
	}
	if isDir {
		fmt.Println("isdirectory")

	} else {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatalf("Error reading file %s\n", arg)
			return
		}

		if Flexible {
			m := make(map[interface{}]interface{})
			err := yaml.Unmarshal(data, &m)
			if err != nil {
				log.Fatalf("Error parsing yaml file %s\n%v", arg, err)
			}
			if Verbose {
				fmt.Printf("--- m:\n%#v\n\n", m)
			}
			return
		}

		t := TrainingLog{}

		err = yaml.Unmarshal(data, &t)
		if err != nil {
			log.Fatalf("Error parsing yaml file %s\n%v", arg, err)
		}
		if Verbose {
			fmt.Printf("--- t:\n%#v\n\n", t)
		}

		// Output flag set
		if Output {
			d, err := yaml.Marshal(&t)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			fmt.Printf("%s\n", string(d))
		}
	}

}

func printUsage() {
	fmt.Println("Usage TODO")
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}
