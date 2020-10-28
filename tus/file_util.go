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
		log.Println("Unable to get user home dir.", err)
		return "", err
	}

	tusPath := path.Join(userHome, tusFolder)
	_, err = os.Stat(tusPath)

	if os.IsNotExist(err) {
		tusPath, err = MakeFileDir()
		if err != nil {
			return "", err
		}
	}

	return tusPath, nil

}

// MakeFileDir create file folder for tus upload
func MakeFileDir() (string, error) {

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to get user home dir.", err)
		return "", err
	}

	tusPath := path.Join(userHome, tusFolder)

	err = os.Mkdir(tusPath, 0744)
	if err != nil {
		log.Println("Unable to create file dir.", err)
		return "", err
	}

	return tusPath, nil

}
