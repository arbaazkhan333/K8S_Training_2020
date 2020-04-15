// Client program
package main

import (
	"crypto/sha1"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	BUFFERSIZE := int64(1024)
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	fileName := make([]byte, 64)
	fileSize := make([]byte, 64)
	checkSum := make([]byte, 20)
	buffer := make([]byte, BUFFERSIZE)
	// Read the file name.
	connection.Read(fileName)
	// Create a file with the same file name
	newFile, err := os.Create("clientdata/" + strings.Trim(string(fileName), ":"))
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	// Read the file size
	connection.Read(fileSize)
	fs, _ := strconv.Atoi(strings.Trim(string(fileSize), ":"))
	// chunks := int(math.Ceil(float64(fs) / float64(BUFFERSIZE)))
	// for i := 0; i < chunks; i++ {
	// 	io.CopyN(newFile, connection, BUFFERSIZE)
	// 	//connection.Read(bufferFileSize)
	// }
	log.Println("File size to be received", fs)
	// Read the checksum
	connection.Read(checkSum)
	// Read the contents of the file.
	connection.Read(buffer)
	// Calculate the checksum
	hash := sha1.Sum(buffer)
	// Match the checksum
	if string(hash[:]) != string(checkSum) {
		panic("Checksum mismatch")
	}
	log.Println("Checksum matched.")
	// Write the contents to the file
	_, err = newFile.Write(buffer)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully received the file.")
	}
	// Hang client to verify whether it has done it's job
	log.Println("Hi I'm client! Can you please verify whether I have done my job correctly or not.")
	log.Println("Please exec into my container and check the files, till then I won't exit :)")
	for {
	}
}
