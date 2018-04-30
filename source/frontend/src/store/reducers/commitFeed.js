export const commitFeedStart = state => ({
  ...state
});

export const commitFeedSuccess = state => ({
  ...state,
  lastKnownError: null
});

export const commitFeedFail = (state, action) => ({
  ...state,
  lastKnownError: action.lastKnownError
});
