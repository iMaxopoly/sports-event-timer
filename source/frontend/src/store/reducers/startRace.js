export const startRaceStart = state => ({
  ...state,
  loadingData: true,
  manualStop: false,
  lastKnownError: null
});

export const startRaceSuccess = (state, action) => ({
  ...state,
  loadingData: false,
  lastKnownError: null,
  raceInProgress: true,
  athletes: action.athletes,
  timePoints: action.timePoints
});

export const startRaceFail = (state, action) => ({
  ...state,
  loadingData: false,
  lastKnownError: action.lastKnownError
});
