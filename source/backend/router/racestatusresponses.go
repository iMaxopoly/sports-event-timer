package router

const (
	responseMessageStartingRace        responseMessage = "starting race"
	responseMessageStoppingRace        responseMessage = "stopping race"
	responseMessageOnGoingTryLater     responseMessage = "race ongoing, please try again later"
	responseMessageNotInProcess        responseMessage = "race is currently not in process"
	responseMessageNotInProcessOldData responseMessage = "race is currently not in process, showing last race data"
	responseMessageInProcessLiveData   responseMessage = "race is in process, showing live data"
	responseMessageCommittedData       responseMessage = "race is in process, data was committed, showing committed data"
	responseMessageNonCommittedData    responseMessage = "race is in process, no data was committed, showing existing data"
	responseMessageUnidentifiedPayload responseMessage = "unidentified payload"
)
