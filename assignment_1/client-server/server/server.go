// Server program
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

func main() {
	// fmt.Println("Server running at port 8080...")
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	fileName := "tempnew.txt"
	createFile(fileName)
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err, "Error listening on 8080")
	}
	defer server.Close()
	log.Println("Waiting for connection...")
	//New goroutine is created whenever a client connects
	for {
		// Accept connection from client
		connection, err := server.Accept()
		if err != nil {
			fmt.Println(err, "Error accepting connection")
		}
		log.Println("Client connected")
		go sendFileToClient(connection, "serverdata/"+fileName)
	}
}

// Creates a file, writes some text on it.
func createFile(name string) {
	//os.Mkdir("serverdata", os.FileMode(os.ModePerm))
	f, err := os.Create("serverdata/" + name)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello Kubernetes")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "Message written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Sends the file details and contents to the client.
func sendFileToClient(connection net.Conn, fName string) {
	file, err := os.Open(fName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	sendBuffer := make([]byte, 1024)
	_, err = file.Read(sendBuffer)
	if err == io.EOF {
		//End of file reached, break out of for loop
		log.Println("EOF")
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fileName := fileInfo.Name()
	fileSize := int(fileInfo.Size())
	strconv.Itoa(fileSize)
	connection.Write([]byte(fillString(fileInfo.Name(), 64)))
	connection.Write([]byte(fillString(strconv.Itoa(fileSize), 64)))
	// Calculate checksum
	var checksum []byte
	hash := sha1.Sum(sendBuffer)
	checksum = hash[:]
	connection.Write(checksum)
	if err != nil {
		//End of file reached, break out of for loop
		log.Println("Error")
	} else {
		connection.Write(sendBuffer)
	}
}

// Helper function to fill the empty portion of the string with :
func fillString(retunString string, toLength int) string {
	for {
		lengthString := len(retunString)
		if lengthString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

/*
// Not used.
func checkSum(msg []byte) int {
	sum := 0
	for i := 0; i < len(msg); i++ {
		if msg[i] == 0 {
			break
		}
		sum += int(msg[i])
	}
	return sum
}*/
