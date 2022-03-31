package main

// set GOOS=linux
// set GOARCH=amd64 go build
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ossDir := "oss"
	ossFile := "oss.txt"
	propDir := "prop"
	propFile := "prop.txt"
	var folderPath string
	if len(os.Args) > 1 {
		folderPath = os.Args[1]

		// do something with command
	} else {
		fmt.Println("Please provide folder to process, as an argument")
		os.Exit(3)
	}
	fmt.Println(folderPath)

	// file handle for proprietary file
	f, err := os.Open(ossFile)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// file handle for proprietary file
	p, err := os.Open(propFile)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer p.Close()

	// present working directory
	// mydir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fullpropDir := folderPath + string(os.PathSeparator) + propDir
	fullossDir := folderPath + string(os.PathSeparator) + ossDir
	if _, err := os.Stat(fullpropDir); os.IsNotExist(err) {
		fmt.Println("Creating proprietary directory")
		err = os.Mkdir(fullpropDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(fullossDir); os.IsNotExist(err) {
		fmt.Println("Creating oss directory")
		err = os.Mkdir(fullossDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// read the file line by line using scanner
	scanner1 := bufio.NewScanner(f)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for scanner1.Scan() {
			if scanner1.Text() != "" {
				fmt.Printf("%s\n", scanner1.Text())
				oldLocation := folderPath + string(os.PathSeparator) + scanner1.Text()
				newLocation := fullossDir + string(os.PathSeparator) + scanner1.Text()
				err := os.Rename(oldLocation, newLocation)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if err := scanner1.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(p)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for scanner.Scan() {
			if scanner.Text() != "" {
				fmt.Printf("%s\n", scanner.Text())
				oldLocation := folderPath + string(os.PathSeparator) + scanner.Text()
				newLocation := fullpropDir + string(os.PathSeparator) + scanner.Text()
				err := os.Rename(oldLocation, newLocation)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if err := scanner1.Err(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
