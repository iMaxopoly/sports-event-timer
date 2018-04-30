import {
  internalNonSideEffect,
  networkNonSideEffect
} from "../actions/actionTypes";
import { startRaceFail, startRaceStart, startRaceSuccess } from "./startRace";
import {
  commitFeedFail,
  commitFeedStart,
  commitFeedSuccess
} from "./commitFeed";
import { fetchFeedFail } from "./fetchFeed";
import { stopRaceFail, stopRaceStart, stopRaceSuccess } from "./stopRace";
import updateRaceData from "./updateRace";
import changeSimulation from "./changeSimulation";
import setRaceData from "./setRace";
import { simulationType } from "../../simulationTypes";

const initialState = {
  simulationType: simulationType.CLIENT,
  athletes: [],
  timePoints: [],
  loadingData: false,
  manualStop: false,
  raceInProgress: false,
  lastKnownError: null
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case internalNonSideEffect.UPDATE_RACE_DATA:
      return updateRaceData(state, action);
    case internalNonSideEffect.SET_RACE_DATA:
      return setRaceData(state, action);
    case internalNonSideEffect.CHANGE_SIMULATION_TYPE:
      return changeSimulation(state, action);

    case networkNonSideEffect.START_RACE_START:
      return startRaceStart(state);
    case networkNonSideEffect.START_RACE_SUCCESS:
      return startRaceSuccess(state, action);
    case networkNonSideEffect.START_RACE_FAIL:
      return startRaceFail(state, action);
    case networkNonSideEffect.COMMIT_FEED_START:
      return commitFeedStart(state);
    case networkNonSideEffect.COMMIT_FEED_SUCCESS:
      return commitFeedSuccess(state);
    case networkNonSideEffect.COMMIT_FEED_FAIL:
      return commitFeedFail(state, action);
    case networkNonSideEffect.FETCH_FEED_FAIL:
      return fetchFeedFail(state, action);
    case networkNonSideEffect.STOP_RACE_START:
      return stopRaceStart(state);
    case networkNonSideEffect.STOP_RACE_SUCCESS:
      return stopRaceSuccess(state);
    case networkNonSideEffect.STOP_RACE_FAIL:
      return stopRaceFail(state, action);
    default:
      return state;
  }
};

export default reducer;
