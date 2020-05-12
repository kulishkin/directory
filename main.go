package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func scanDir(startDir string, wg *sync.WaitGroup, ch chan<- []int) {
	defer wg.Done()
	collectionCounts := []int{}
	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if info.Name() == "count" {
			count, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			intCount, err := strconv.Atoi(string(count))
			if err != nil {
				panic(err)
			}
			collectionCounts = append(collectionCounts, intCount)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	ch <- collectionCounts
}

func waitDirScan(wg *sync.WaitGroup, ch chan []int) {
	wg.Wait()
	close(ch)
}

func main() {
	dirs := []string{"dir1", "dir2", "dir3"}
	dirsLength := len(dirs)
	wg := &sync.WaitGroup{}
	wg.Add(dirsLength)
	ch := make(chan []int)
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			panic(err)
			continue
		}
		go scanDir(dir, wg, ch)
	}
	go waitDirScan(wg, ch)

	collectionCounts := []int{}
	for count := range ch {
		collectionCounts = append(collectionCounts, count...)
	}
	fmt.Println(sum(collectionCounts))
}
