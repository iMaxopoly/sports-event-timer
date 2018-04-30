const changeSimulation = (state, action) => ({
  ...state,
  simulationType: action.simulationType
});

export default changeSimulation;
