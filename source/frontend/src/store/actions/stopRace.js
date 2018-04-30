import * as actionTypes from "./actionTypes";

export const stopRace = simulationType => ({
  type: actionTypes.networkSideEffect.STOP_RACE,
  simulationType
});

export const stopRaceStart = () => ({
  type: actionTypes.networkNonSideEffect.STOP_RACE_START
});

export const stopRaceSuccess = () => ({
  type: actionTypes.networkNonSideEffect.STOP_RACE_SUCCESS
});

export const stopRaceFail = lastKnownError => ({
  type: actionTypes.networkNonSideEffect.STOP_RACE_FAIL,
  lastKnownError
});
