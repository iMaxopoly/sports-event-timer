import * as actionTypes from "./actionTypes";

export const fetchFeed = () => ({
  type: actionTypes.networkSideEffect.FETCH_FEED
});

export const fetchFeedFail = lastKnownError => ({
  type: actionTypes.networkNonSideEffect.FETCH_FEED_FAIL,
  lastKnownError
});
