export const getAthleteByIdentifier = (state, identifier) => {
  const result = state.athletes.filter(
    athlete => athlete.identifier === identifier
  );
  if (result && result.length === 1) {
    return result[0];
  }
  return null;
};

export const getRaceStatus = state => state.raceInProgress;
