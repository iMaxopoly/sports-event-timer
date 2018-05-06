package router

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"sports-event-timer/source/backend/racesimulator"

	"github.com/julienschmidt/httprouter"
)

var templates = template.Must(template.ParseFiles("index.html"))

// notFoundHandler is a custom handler that handles routes that aren't found
func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Not Found"))
	handleError(w, err)
}

// indexHandler handles the index page request at /
func indexHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	handleError(w, err)
}

// faviconHandler handles request to fetch the favicon
func faviconHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./favicon.ico")
}

// ServerRaceStartHandler initates a server simulated race and sets athlete goroutines into motion
// the results are thereby fetchable via the ServerRaceFetchLiveHandler
func ServerRaceStartHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var raceResp raceJSONPayload

	switch event.EventState() {
	case racesimulator.RaceNotRunning:
		event.ResetPlatform()
		loadEntitiesToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageStartingRace
		go event.StartServerSimulatedRace()
		break
	case racesimulator.RaceRunning:
		raceResp.ResponseToSend = responseMessageOnGoingTryLater
		break
	}

	writeJSON(w, &raceResp)
}

// ServerRaceFetchLiveHandler allows user to fetch the data of an on-going server-simulated
// race. The data is fetched from the database where the records are stored.
func ServerRaceFetchLiveHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var raceResp raceJSONPayload

	switch event.EventState() {
	case racesimulator.RaceNotRunning:
		loadEntitiesToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageNotInProcessOldData
		break
	case racesimulator.RaceRunning:
		loadEntitiesToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageInProcessLiveData
		break
	}

	writeJSON(w, &raceResp)
}

// ServerRaceFetchLastHandler loads the last race results from the database.
func ServerRaceFetchLastHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var raceResp raceJSONPayload

	switch event.EventState() {
	case racesimulator.RaceNotRunning:
		loadEntitiesToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageNotInProcessOldData
		break
	case racesimulator.RaceRunning:
		loadEntitiesToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageInProcessLiveData
		break
	}

	writeJSON(w, &raceResp)
}

// ClientRaceStartHandler boots up a new event environment for the race to be processed.
func ClientRaceStartHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var raceResp raceJSONPayload

	switch event.EventState() {
	case racesimulator.RaceNotRunning:
		event.ResetPlatform()
		event.StartClientSimulatedRace()
		loadEntitiesToShow(&raceResp)
		loadTimePointsToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageStartingRace
		break
	case racesimulator.RaceRunning:
		raceResp.ResponseToSend = responseMessageOnGoingTryLater
		break
	}

	writeJSON(w, &raceResp)
}

// RaceStopHandler stops the current race and all underlying goroutines.
func RaceStopHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var raceResp raceJSONPayload

	switch event.EventState() {
	case racesimulator.RaceNotRunning:
		raceResp.ResponseToSend = responseMessageNotInProcess
		break
	case racesimulator.RaceRunning:
		event.StopSimulatedRace()
		raceResp.ResponseToSend = responseMessageStoppingRace
		break
	}

	writeJSON(w, &raceResp)
}

// ClientRaceFetchHandler fetches live data from the database in regards to the status of
// all athletes participating in the race.
func ClientRaceFetchHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var raceResp raceJSONPayload

	switch event.EventState() {
	case racesimulator.RaceNotRunning:
		loadEntitiesToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageNotInProcessOldData
		break
	case racesimulator.RaceRunning:
		loadEntitiesToShow(&raceResp)
		raceResp.ResponseToSend = responseMessageInProcessLiveData
		break
	}

	writeJSON(w, &raceResp)
}

// ClientRacePushHandler handles the data pushed by the client side, updating the race status
// and setting up the data into the database as per requirements.
// Due to constraints of the project, only athlete chip information, timepoint chip information
// and time the athlete reached given timepoint is sent to this handler. The server
// then decides who entered the finish corridor or has crossed the finish line.
func ClientRacePushHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var payload raceJSONPayload
	err := decoder.Decode(&payload)
	handleError(w, err)

	defer r.Body.Close()

	if !payloadHasCommand(&payload) {
		writeJSON(w, &raceJSONPayload{ResponseToSend: responseMessageUnidentifiedPayload})
		return
	}

	var raceResp raceJSONPayload

	switch event.EventState() {
	case racesimulator.RaceNotRunning:
		raceResp.ResponseToSend = responseMessageNotInProcess
		break
	case racesimulator.RaceRunning:
		if len(payload.Athletes) <= 0 {
			loadEntitiesToShow(&raceResp)
			raceResp.ResponseToSend = responseMessageNonCommittedData
		} else {
			for _, ath := range payload.Athletes {
				var athChip racesimulator.Chip
				athChip.SetIdentifier(ath.Identifier)
				var timePointChip racesimulator.Chip
				timePointChip.SetIdentifier(ath.TimePointIdentifier)
				event.UpdateEntitiesFromTimePoint(&athChip, &timePointChip, time.Duration(ath.TimeElapsed))
			}
			loadEntitiesToShow(&raceResp)
			raceResp.ResponseToSend = responseMessageCommittedData
			break
		}
	}

	writeJSON(w, &raceResp)
}
