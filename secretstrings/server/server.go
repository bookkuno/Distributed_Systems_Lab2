package main

import (
	"flag"
	"math/rand"
	"net"
	"net/rpc"
	"time"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

//Server Procedure: Super-Secret `reversing a string' method we can't allow clients to see.
func ReverseString(s string, i int) string {
	time.Sleep(time.Duration(rand.Intn(i)) * time.Second)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//Server Stub Procedure = these methods invoked by clients through RPC (Remote Procedure Call): Reverse() & FastReverse()
type SecretStringOperation struct{} //a receiver type for methods associated with this struct

func (s *SecretStringOperation) Reverse(req stubs.Request, res *stubs.Response) (err error) {
	//(s *SecretStringOperation) = this method is associated with instances of the SecretStringOperation struct
	//s, is a pointer to a SecretStringOperation struct.
	res.Message = ReverseString(req.Message, 10)
	return
}
func (s *SecretStringOperation) FastReverse(req stubs.Request, res *stubs.Response) (err error) {
	res.Message = ReverseString(req.Message, 2) //decrease the delay that get passed to ReverseString()
	return
}

func main() { //Listen for indications from the client on calling the func (s *SecretStringOperation), and will handle the communications by the listener.
	pAddr := flag.String("port", "8030", "Port to listen on")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	rpc.Register(&SecretStringOperation{}) //'&' creates a new instance of SecretStringOperation and obtain its address
	listener, _ := net.Listen("tcp", ":"+*pAddr)
	defer listener.Close()
	rpc.Accept(listener)
}
