package main

import (
	"os"
	"time"
	"fmt"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: filescan <path>\n")
		os.Exit(1)
	}
	scan(os.Args[1])
}

func scan(keyDir string) error {
	fmt.Printf("1. Scanning directory %v...\n", keyDir)
	// List all the files from the keystore folder
	t0 := time.Now()
	files, err := ioutil.ReadDir(keyDir)
	//files, err := os.ReadDir(keyDir) <-- not available pre 1.16
	if err != nil {
		fmt.Printf("ERROR: Scan error after %v : %v\n", time.Since(t0), err)
		return err
	}
	fmt.Printf("   Scan ok, %d files, elapsed %v\n", len(files), time.Since(t0))
	fmt.Printf("2. Iterating over files...\n")
	for _, fi := range files {
		fi.ModTime()
		//info, err := fi.Info()
		//if err != nil {
		//	fmt.Printf("ERROR: Err while getting file info for %v: %v\n", fi.Name(), err)
		//	return err
		//}
		//info.ModTime()
	}
	fmt.Printf("   Iteration done, elapsed %v\n", time.Since(t0))
	return nil
}
