package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
)

// FileDetails stores the last modified time and if the file is to be reaped
type FileDetails struct {
	LastMod time.Time
	Reap    bool
}

var (
	sigc = make(chan os.Signal, 1)
	// FileList is a map of filenames -> *FileDetails
	FileList = make(map[string]*FileDetails)
)

func checkfile(filenow string, timenow time.Time) {
	fdets, ok := FileList[filenow]

	// If the file exists and it's changed,
	if ok && fdets.LastMod != timenow {
		FileList[filenow].Reap = false
		fmt.Printf("Touched %s\n", filenow)

		currentTime := time.Now().Local()
		err := os.Chtimes(filenow, currentTime, currentTime)
		if err != nil {
			fmt.Println(err)
		}

		FileList[filenow].LastMod = currentTime
		return
	}

	// If the file still exists, ensure it dosen't get reaped.
	if ok {
		FileList[filenow].Reap = false
		return
	}

	// The file must not exist, so let's add it (unreapable)
	fmt.Printf("Added " + filenow + "\n")
	FileList[filenow] = &FileDetails{timenow, false}
}

func dirlist(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if ignorelist(f.Name()) {
			continue
		}

		if f.IsDir() {
			subdir := directory + f.Name() + "/"
			dirlist(subdir)
		} else {
			k := directory + f.Name()
			v := f.ModTime()
			checkfile(k, v)
		}
	}
}

func ignorelist(name string) bool {
	ilist := []string{
		"node_modules",
		"tmp",
		"log",
		".git",
	}

	for _, i := range ilist {
		if i == name {
			return true
		}
	}

	return false
}

func help() string {
	return "Usage: `touchngo` from inside directory you care about."
}

func main() {
	signal.Notify(sigc, os.Interrupt)

	go func() {
		for {

			// Set everything as reapable
			for k := range FileList {
				FileList[k].Reap = true
			}

			dirlist("./")

			// Reap: remove anything that's no longer there.
			for k, v := range FileList {
				if v.Reap {
					delete(FileList, k)
					fmt.Printf("Reaped " + k + "\n")
				}
			}

			time.Sleep(3 * time.Second)

		}
	}()

	<-sigc
}
