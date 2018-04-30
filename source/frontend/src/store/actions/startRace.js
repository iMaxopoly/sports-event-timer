import * as actionTypes from "./actionTypes";

export const startRace = (simulationType, timePoints) => ({
  type: actionTypes.networkSideEffect.START_RACE,
  simulationType,
  timePoints
});

export const startRaceStart = () => ({
  type: actionTypes.networkNonSideEffect.START_RACE_START
});

export const startRaceSuccess = (athletes, timePoints) => ({
  type: actionTypes.networkNonSideEffect.START_RACE_SUCCESS,
  athletes,
  timePoints
});

export const startRaceFail = lastKnownError => ({
  type: actionTypes.networkNonSideEffect.START_RACE_FAIL,
  lastKnownError
});
