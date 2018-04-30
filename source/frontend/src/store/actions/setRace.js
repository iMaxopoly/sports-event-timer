import * as actionTypes from "./actionTypes";

export const setRaceData = athletes => ({
  type: actionTypes.internalNonSideEffect.SET_RACE_DATA,
  athletes
});
