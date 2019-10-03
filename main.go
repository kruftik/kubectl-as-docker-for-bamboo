package main

import (
	"fmt"
	"os"
	"strings"
)

const KUBECTL_BIN  = "kubectl"

var (
	//wellKnownDirs = map[string]bool{
	//	"/usr/bin": true,
	//	"/usr/local/bin": true,
	//}
	wellKnownDirs = map[string]bool{
		"/tmp": true,
	}
)

func main() {
	var kubectl_path string

	for d := range wellKnownDirs {
		if IsFileExists(d + string(os.PathSeparator) +  KUBECTL_BIN){
			kubectl_path = d + string(os.PathSeparator) +  KUBECTL_BIN
		}
	}

	if kubectl_path == "" {
		//fmt.Printf("%s found at %s", KUBECTL_BIN, kubectl_path)
		path := os.Getenv("PATH")
		if path != "" {
			for _, d := range strings.Split(path, string(os.PathListSeparator)){
				_, ok := wellKnownDirs[d]
				if ok {
					fmt.Printf("%s skipped as already checked\n", d)
					continue
				}

				if IsFileExists(d + string(os.PathSeparator) +  KUBECTL_BIN){
					kubectl_path = d + string(os.PathSeparator) +  KUBECTL_BIN
				}
			}
		}
	}

	if kubectl_path == "" {
		fmt.Printf("%s binary was not found.", KUBECTL_BIN)
		os.Exit(1)
	}

	ParseAndRunCommand()
}

