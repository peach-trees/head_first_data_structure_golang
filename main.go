package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
)

const exercisesDir = "./exercises"

var logger *log.Logger

func LoopExercisesDir(dir string) []string {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.Println(err)
		return nil
	}
	ret := make([]string, 0)
	for _, v := range fileInfos {
		ret = append(ret, v.Name())
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return ret
}

func init() {
	logger = log.New(os.Stdout, "head first:",
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetOutput(os.Stderr)
}

func main() {

}
