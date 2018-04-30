import * as actionTypes from "./actionTypes";

export const updateRaceData = athletes => ({
  type: actionTypes.internalNonSideEffect.UPDATE_RACE_DATA,
  athletes
});
