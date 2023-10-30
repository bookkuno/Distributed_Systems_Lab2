package stubs

//Variables: define the name of the nethods called on server by client
var ReverseHandler = "SecretStringOperation.Reverse"
var PremiumReverseHandler = "SecretStringOperation.FastReverse"

//Structs
type Response struct {
	Message string
}

type Request struct {
	Message string
}
