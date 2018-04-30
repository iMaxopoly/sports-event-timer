export const updateRaceData = (state, action) => ({
  ...state,
  athletes: state.athletes
    .slice()
    .map(
      a => (a.identifier === action.athletes.identifier ? action.athletes : a)
    )
});

export default updateRaceData;
