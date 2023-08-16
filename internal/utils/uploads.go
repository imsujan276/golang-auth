package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadImage uploads an image to the server and returns the file path
func UploadImage(file multipart.File, header *multipart.FileHeader, destination ...string) (string, error) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Set the default destination if not provided
	dest := filepath.Join(currentDir, "uploads/images")
	if len(destination) > 0 {
		dest = destination[0]
	}

	// Create the destination folder if it doesn't exist
	err = os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Generate a unique filename
	filename := generateUniqueFilename()
	// Get the file extension from the original filename
	extension := filepath.Ext(header.Filename)

	// Create a new file on the server inside the destination folder
	filePath := filepath.Join(dest, filename+extension)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the server
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// UploadImage2 uploads an uploadedFormDataImg image to the server and returns the file path
func UploadedFormDataImg(uploadedFormImg *multipart.FileHeader, destination ...string) (string, error) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Set the default destination if not provided
	dest := "uploads/images"
	if len(destination) > 0 {
		dest = destination[0]
	}
	fullDestination := filepath.Join(currentDir, dest)

	// Create the destination folder if it doesn't exist
	err = os.MkdirAll(fullDestination, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Generate a unique filename
	filename := generateUniqueFilename()
	// Get the file extension from the original filename
	extension := filepath.Ext(uploadedFormImg.Filename)

	// Create a new file on the server inside the destination folder
	filePath := filepath.Join(dest, filename+extension)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Retrieve the uploaded file
	file, err := uploadedFormImg.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the uploaded file to the server
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// Delete file from path
func DeleteFile(pathToFile string) error {
	err := os.Remove(pathToFile)
	if err != nil {
		return err
	}
	return nil
}

func generateUniqueFilename() string {
	// Generate a unique filename based on the current timestamp
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	filename := fmt.Sprintf("file_%d", timestamp)
	return filename
}
