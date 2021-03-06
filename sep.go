package main

// set GOOS=linux
// set GOOS=windows
// set GOARCH=amd64 go build
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	jsonDir := "jsonFiles"
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
		os.Exit(1)
	}
	fmt.Println(folderPath)

	// file handle for proprietary file

	// present working directory
	// mydir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	wg.Add(3)
	go func() {
		defer wg.Done()
		fullossDir := folderPath + string(os.PathSeparator) + ossDir
		if _, err := os.Stat(fullossDir); os.IsNotExist(err) {
			fmt.Println("Creating oss directory")
			err = os.Mkdir(fullossDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
		f, err := os.Open(ossFile)
		if err != nil {
			log.Fatal(err)
		}
		// remember to close the file at the end of the program
		defer f.Close()
		// read the file line by line using scanner
		scanner1 := bufio.NewScanner(f)
		for scanner1.Scan() {
			if fileName1 := scanner1.Text(); fileName1 != "" {
				// fmt.Printf("%s\n", scanner1.Text())
				oldLocation := folderPath + string(os.PathSeparator) + fileName1
				newLocation := fullossDir + string(os.PathSeparator) + fileName1
				_, err := os.Stat(newLocation)
				_, err1 := os.Stat(oldLocation)
				if os.IsNotExist(err) && !os.IsNotExist(err1) {
					err := os.Rename(oldLocation, newLocation)
					if err != nil {
						log.Fatal(err)
					}
					if err1 != nil {
						log.Fatal(err1)
					}
				} else {
					// fmt.Println(strings.Repeat("=", 10), "File", fileName1, "exist in", fullossDir, strings.Repeat("=", 10))
				}
			}
		}
		if err := scanner1.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		// file handle for proprietary file
		p, err := os.Open(propFile)
		if err != nil {
			log.Fatal(err)
		}
		// remember to close the file at the end of the program
		defer p.Close()
		fullpropDir := folderPath + string(os.PathSeparator) + propDir
		if _, err := os.Stat(fullpropDir); os.IsNotExist(err) {
			fmt.Println("Creating proprietary directory")
			err = os.Mkdir(fullpropDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
		scanner := bufio.NewScanner(p)
		for scanner.Scan() {
			if fileName2 := scanner.Text(); fileName2 != "" {
				// fmt.Printf("%s\n", scanner.Text())
				oldLocation := folderPath + string(os.PathSeparator) + fileName2
				newLocation := fullpropDir + string(os.PathSeparator) + fileName2
				_, err := os.Stat(newLocation)
				_, err1 := os.Stat(oldLocation)
				fmt.Println(fileName2,"======",os.IsNotExist(err),os.IsNotExist(err1))
				if os.IsNotExist(err) && !os.IsNotExist(err1) {
					err := os.Rename(oldLocation, newLocation)
					if err != nil {
						log.Fatal(err)
					}
					if err1 != nil {
						log.Fatal(err1)
					}
				} else {
					// fmt.Println(strings.Repeat("=", 10), "File", fileName2, "exist in", fullpropDir, strings.Repeat("=", 10))
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		fulljsonDir := folderPath + string(os.PathSeparator) + jsonDir
		if _, err := os.Stat(fulljsonDir); os.IsNotExist(err) {
			fmt.Println("Creating json directory")
			err = os.Mkdir(fulljsonDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
		dir, err := ioutil.ReadDir(folderPath)
		if err != nil {
			msg := fmt.Sprintf("An error occured reading the %v directory.\n%s", folderPath, err)
			fmt.Println(msg)
			os.Exit(1)
		}
		for _, file := range dir {
			if !file.IsDir() {
				jsonFileName := file.Name()
				oldLocation := folderPath + string(os.PathSeparator) + jsonFileName
				newLocation := fulljsonDir + string(os.PathSeparator) + jsonFileName
				isJsonFile := strings.Contains(jsonFileName, ".json")
				_, err := os.Stat(newLocation)
				_, err1 := os.Stat(oldLocation)
				if os.IsNotExist(err) && !os.IsNotExist(err1) {
					if isJsonFile {
						err := os.Rename(oldLocation, newLocation)
						if err != nil {
							log.Fatal(err)
						}
						if err1 != nil {
							log.Fatal(err1)
						}
					} 
				}else {
					// fmt.Println(strings.Repeat("=", 10), "File", jsonFileName, "exist in", fulljsonDir, strings.Repeat("=", 10))
				}
			}
		}
	}()
	wg.Wait()
}
