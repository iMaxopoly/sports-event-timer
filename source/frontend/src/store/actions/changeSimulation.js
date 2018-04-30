import * as actionTypes from "./actionTypes";

export const changeSimulation = simulationType => ({
  type: actionTypes.internalNonSideEffect.CHANGE_SIMULATION_TYPE,
  simulationType
});
