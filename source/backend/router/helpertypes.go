package router

type (
	// responseMessage is a helper type that masks an string type to help with better segregation of constant variables.
	// This is currently used to declare the messages that are sent from server to the client side.
	responseMessage string

	// requestCommand is a helper type that masks an string type to help with better segregation of constant variables.
	// This is currently used to declare the messages that are received from the client.
	requestCommand string
)
