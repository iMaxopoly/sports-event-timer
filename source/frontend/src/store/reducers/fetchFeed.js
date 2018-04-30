export const fetchFeedFail = (state, action) => ({
  ...state,
  loadingData: false,
  lastKnownError: action.lastKnownError
});
