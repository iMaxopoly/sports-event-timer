export const stopRaceStart = state => ({
  ...state,
  loadingData: true,
  raceInProgress: true
});

export const stopRaceSuccess = state => ({
  ...state,
  loadingData: false,
  raceInProgress: false,
  lastKnownError: null
});

export const stopRaceFail = (state, action) => ({
  ...state,
  loadingData: true,
  lastKnownError: action.lastKnownError,
  manualStop: true
});
