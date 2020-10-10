package main

import (
	"log"
	"os"
	"path"
)

const tusFolder = "tus_fileserver"

// GetFileDir return the folder path tus file stored
func GetFileDir() (string, error) {

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to get user home directory.", err)
		return "", err
	}

	tusPath := path.Join(userHome, tusFolder)

	if _, err := os.Stat(tusPath); os.IsNotExist(err) {
		log.Println("Tus file folder not exist.", err)
		return "", err
	}

	return tusPath, nil

}

// MakeFileDir create file folder for tus upload
func MakeFileDir() (string, error) {

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to get user home directory.", err)
		return "", err
	}

	tusPath := path.Join(userHome, tusFolder)

	err = os.Mkdir(tusPath, 0744)
	if err != nil {
		log.Println("Unable to create file directory.", err)
		return "", err
	}

	return tusPath, nil

}
