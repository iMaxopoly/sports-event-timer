package router

import (
	"log"
	"net/http"

	"sports-event-timing/source/backend/racesimulator"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var event racesimulator.IEvent

func init() {
	// Initialize RaceEvent and flag that it is not running
	var rE racesimulator.RaceEvent
	event = &rE
	event.SetEventState(racesimulator.RaceNotRunning)
}

// Serve loads up the routes and serves at given point.
func Serve() {
	mux := httprouter.New()

	// Port for mux
	muxPort := ":8082"

	// Notfound redirect to index for react to handle
	mux.NotFound = notFoundHandler

	// Index page handler
	mux.GET("/", indexHandler)

	// // Server simulated REST API
	mux.GET("/server/race-simulation/start", ServerRaceStartHandler)

	mux.GET("/server/race-simulation/stop", RaceStopHandler)

	mux.GET("/server/race-simulation/fetch/live-standings", ServerRaceFetchLiveHandler)

	mux.GET("/server/race-simulation/fetch/last-standings", ServerRaceFetchLastHandler)

	// Client simulated REST API
	mux.GET("/client/race-simulation/start", ClientRaceStartHandler)

	mux.GET("/client/race-simulation/stop", RaceStopHandler)

	mux.GET("/client/race-simulation/fetch/current-standings", ClientRaceFetchHandler)

	mux.POST("/client/race-simulation/push/current-standings", ClientRacePushHandler)

	// Favicon
	mux.GET("/favicon.ico", faviconHandler)

	// Static files
	mux.ServeFiles("/static/*filepath", http.Dir("./static"))

	log.Println("Server running at", muxPort)

	log.Fatal(http.ListenAndServe(muxPort, cors.Default().Handler(mux)))
}
