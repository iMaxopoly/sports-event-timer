package router

type athleteJSONPayload struct {
	StartNumber                    int     `json:"startNumber,omitempty"`
	FullName                       string  `json:"fullName,omitempty"`
	Identifier                     string  `json:"identifier,omitempty"`
	TimePointIdentifier            string  `json:"timePointIdentifier,omitempty"`
	Location                       int     `json:"location,omitempty"`
	InFinishCorridor               bool    `json:"inFinishCorridor,omitempty"`
	HasFinished                    bool    `json:"hasFinished,omitempty"`
	TimeTakenToReachFinishCorridor float64 `json:"timeTakenToReachFinishCorridor,omitempty"`
	TimeTakenToFinish              float64 `json:"timeTakenToFinish,omitempty"`
	TimeElapsed                    float64 `json:"timeElapsed,omitempty"`
}

type timePointJSONPayload struct {
	Name       string `json:"name,omitempty"`
	Location   int    `json:"location,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type raceJSONPayload struct {
	ResponseToSend  responseMessage        `json:"responseMessage,omitempty"`
	RequestReceived requestCommand         `json:"requestCommand,omitempty"`
	Athletes        []athleteJSONPayload   `json:"athletes,omitempty"`
	TimePoints      []timePointJSONPayload `json:"timePoints,omitempty"`
}
