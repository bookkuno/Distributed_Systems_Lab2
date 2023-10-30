package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

/*Client Stub Procedure = marshalling the client's request, sending it to the server, receiving the response,
and unmarshalling the response (but not explicitly present in the code)
	- rpc.Dial = establish a connection to the server
	- rpc.Call = send an RPC request and receive the response.
*/

func main() { //Client Program = the main function serves as the client program
	//Option1: Reversing single input text
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)

	client, err := rpc.Dial("tcp", *server)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer client.Close()

	// Ask the user for input
	fmt.Print("Enter a string to reverse: ")
	var input string
	_, err = fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	request := stubs.Request{Message: input}
	response := new(stubs.Response)
	err = client.Call(stubs.PremiumReverseHandler, request, response) //stubs.___ tell choose between reverse()/fastreverse()
	if err != nil {
		fmt.Println("Error calling RPC:", err)
		return
	}

	fmt.Println("Reversed string:", response.Message)

	//Option2: Reversing textfile.txt
	//server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	//wordlistFile := flag.String("wordlist", "/Users/book_kuno/Desktop/Golang/Distributed_Systems_Lab2/secretstrings/GoTextFile.txt", "File containing a list of words")
	//flag.Parse()
	//fmt.Println("Server:", *server)
	//
	//client, err := rpc.Dial("tcp", *server)
	//if err != nil {
	//	fmt.Println("Error connecting to the server:", err)
	//	return
	//}
	//defer client.Close()
	//
	//// Open the wordlist file
	//file, err := os.Open(*wordlistFile)
	//if err != nil {
	//	fmt.Println("Error opening wordlist file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//scanner := bufio.NewScanner(file)
	//
	//for scanner.Scan() {
	//	word := scanner.Text()
	//	request := stubs.Request{Message: word}
	//	response := new(stubs.Response)
	//	err = client.Call(stubs.PremiumReverseHandler, request, response)
	//	if err != nil {
	//		fmt.Printf("Error reversing word '%s': %v\n", word, err)
	//	} else {
	//		fmt.Printf("Original: %s, Reversed: %s\n", word, response.Message)
	//	}
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	fmt.Println("Error reading wordlist file:", err)
	//}
}

/*
	1)Parses command-line arguments to determine the server's IP and port.
	2)Establishes a connection to the server using rpc.Dial.
	3)Prompts the user to input a string to be reversed.
	4)Reads the user's input using fmt.Scanln.
	5)Creates an RPC request using the user's input and sends it to the server using client.Call.
	6)Receives and displays the reversed string from the server.
*/
