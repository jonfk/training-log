package main

import (
	"../common"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Statistics struct {
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("usage: statistics <directory>\n")
		os.Exit(0)
	}

	traininglogs := parseYaml(args[0])

	calculateStatistics(traininglogs)
}

func parseYaml(directory string) []common.TrainingLog {
	var result []common.TrainingLog
	toProcess, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for i := range toProcess {
		file := filepath.Join(directory, toProcess[i].Name())
		t := common.TrainingLog{}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("%s\n", err)
		}

		err = yaml.Unmarshal(data, &t)
		if err != nil {
			log.Fatalf("Error parsing yaml file %s\n%v", file, err)
		}
		result = append(result, t)
	}

	// fmt.Printf("--- TrainingLogs:\n%#v\n", result)
	return result
}

func calculateStatistics(logs []common.TrainingLog) Statistics {
	return Statistics{}
}
