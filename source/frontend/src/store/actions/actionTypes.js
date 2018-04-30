const internal = "internal";
const network = "network";

export const internalNonSideEffect = {
  UPDATE_RACE_DATA: `${internal}/UPDATE_RACE_DATA`,
  SET_RACE_DATA: `${internal}/SET_RACE_DATA`,
  CHANGE_SIMULATION_TYPE: `${internal}/CHANGE_SIMULATION_TYPE`
};

export const networkNonSideEffect = {
  COMMIT_FEED_START: `${network}/COMMIT_FEED_START`,
  COMMIT_FEED_SUCCESS: `${network}/COMMIT_FEED_SUCCESS`,
  COMMIT_FEED_FAIL: `${network}/COMMIT_FEED_FAIL`,

  FETCH_FEED_FAIL: `${network}/FETCH_FEED_FAIL`,

  START_RACE_START: `${network}/START_RACE_START`,
  START_RACE_SUCCESS: `${network}/START_RACE_SUCCESS`,
  START_RACE_FAIL: `${network}/START_RACE_FAIL`,

  STOP_RACE_START: `${network}/STOP_RACE_START`,
  STOP_RACE_SUCCESS: `${network}/STOP_RACE_SUCCESS`,
  STOP_RACE_FAIL: `${network}/STOP_RACE_FAIL`
};

export const networkSideEffect = {
  COMMIT_FEED: `${network}/COMMIT_FEED`,
  FETCH_FEED: `${network}/FETCH_FEED`,
  START_RACE: `${network}/START_RACE`,
  STOP_RACE: `${network}/STOP_RACE`
};
