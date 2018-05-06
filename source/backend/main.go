package main

import (
	"sports-event-timer/source/backend/router"
)

// The entry point of the project, where router.Server loads
// the project routes and httprouter.
func main() {
	router.Serve()
}
