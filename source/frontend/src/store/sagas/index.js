import { all, takeEvery, takeLatest } from "redux-saga/effects";
import * as actionTypes from "../actions/actionTypes";
import {
  commitFeedSaga,
  fetchFeedSaga,
  startRaceSaga,
  stopRaceSaga
} from "./sagas";

function* watchNetwork() {
  yield all([
    takeLatest(actionTypes.networkSideEffect.START_RACE, startRaceSaga),
    takeEvery(actionTypes.networkSideEffect.COMMIT_FEED, commitFeedSaga),
    takeLatest(actionTypes.networkSideEffect.FETCH_FEED, fetchFeedSaga),
    takeLatest(actionTypes.networkSideEffect.STOP_RACE, stopRaceSaga)
  ]);
}

export default watchNetwork;
