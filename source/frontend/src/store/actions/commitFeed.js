import * as actionTypes from "./actionTypes";

export const commitFeed = (
  timeElapsed,
  athleteIdentifier,
  timePointIdentifier,
  simulationType
) => ({
  type: actionTypes.networkSideEffect.COMMIT_FEED,
  timeElapsed,
  athleteIdentifier,
  timePointIdentifier,
  simulationType
});

export const commitFeedStart = () => ({
  type: actionTypes.networkNonSideEffect.COMMIT_FEED_START
});

export const commitFeedSuccess = () => ({
  type: actionTypes.networkNonSideEffect.COMMIT_FEED_SUCCESS
});

export const commitFeedFail = lastKnownError => ({
  type: actionTypes.networkNonSideEffect.COMMIT_FEED_FAIL,
  lastKnownError
});
