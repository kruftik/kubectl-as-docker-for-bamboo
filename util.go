package main

import "os"

func IsPathExists(path string) bool{
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func IsDirExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func IsFileExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
