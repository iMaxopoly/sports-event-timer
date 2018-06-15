package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// handleError is a generic helper to handle mux errors.
func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

// writeJSON is a generic helper to write JSON to client.
func writeJSON(w http.ResponseWriter, payload *raceJSONPayload) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(payload)
	handleError(w, err)
}

// payloadHasCommand checks if the client payload is empty or not.
func payloadHasCommand(payload *raceJSONPayload) bool {
	if strings.TrimSpace(string(payload.RequestReceived)) == "" {
		return false
	}
	return true
}

// loadEntitiesToShow fetches entities currently in the database and prepares it to be
// sent as a JSON payload to the client.
func loadEntitiesToShow(raceResp *raceJSONPayload) {
	entities := *event.Entities()

	if len(entities) <= 0 {
		return
	}

	for _, entity := range entities {
		raceResp.Athletes = append(raceResp.Athletes, athleteJSONPayload{
			StartNumber:                    entity.StartNumber(),
			FullName:                       entity.EntityName(),
			Identifier:                     entity.Chip().Identifier(),
			InFinishCorridor:               entity.InFinishCorridor(),
			HasFinished:                    entity.HasFinished(),
			TimeTakenToReachFinishCorridor: float64(entity.TimeTakenToReachFinishCorridor()),
			TimeTakenToFinish:              float64(entity.TimeTakenToFinish()),
			Location:                       entity.Location(),
		})
	}
}

// loadEntitiesToShow fetches timepoints currently in the database and prepares it to be
// sent as a JSON payload to the client.
func loadTimePointsToShow(raceResp *raceJSONPayload) {
	timePoints := *event.TimePoints()

	if len(timePoints) <= 0 {
		return
	}

	for _, timePoint := range timePoints {
		raceResp.TimePoints = append(raceResp.TimePoints, timePointJSONPayload{
			Name:       string(timePoint.Name()),
			Location:   int(timePoint.Location()),
			Identifier: timePoint.Chip().Identifier(),
		})
	}
}
