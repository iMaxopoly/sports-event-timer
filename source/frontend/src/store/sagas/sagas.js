import { delay } from "redux-saga";
import { all, call, put, select } from "redux-saga/effects";
import axios from "axios";
import * as actions from "../actions";
import { getAthleteByIdentifier, getRaceStatus } from "../selectors/selector";
import { simulationType } from "../../simulationTypes";

// Helper error handler
const throwError = message => {
  throw message;
};

// The following generator helps with the athlete propagation in the race track
// Individual generators are spun using the 'all' redux-saga effects to concurrently
// increment the progress of athletes. When an athlete trips over a timepoint,
// their chip id, timepoint id and time taken to reach that timepoint, is sent
// to the server, whereby the server stores it in the database.
function* runningAthleteSaga(athleteIdentifier, timePoints, stopwatchStart) {
  while (true) {
    const raceStatus = yield select(getRaceStatus);
    if (!raceStatus) {
      break;
    }

    const runningAthlete = yield select(
      getAthleteByIdentifier,
      athleteIdentifier
    );
    if (runningAthlete === null) {
      return;
    }

    runningAthlete.location += Math.floor(Math.random() * 10) + 15;

    if (
      runningAthlete.location > timePoints[0].location &&
      runningAthlete.timeTakenToReachFinishCorridor === -1
    ) {
      // Athlete is in the finish corridor
      runningAthlete.timeTakenToReachFinishCorridor =
        new Date() - stopwatchStart;
      yield put(
        actions.commitFeed(
          runningAthlete.timeTakenToReachFinishCorridor,
          runningAthlete.identifier,
          timePoints[0].identifier,
          simulationType.CLIENT
        )
      );
      yield put(actions.updateRaceData([runningAthlete]));
    } else if (
      runningAthlete.location > timePoints[1].location &&
      runningAthlete.timeTakenToFinish === -1
    ) {
      // Athlete reached finish line
      runningAthlete.timeTakenToFinish = new Date() - stopwatchStart;
      yield put(
        actions.commitFeed(
          runningAthlete.timeTakenToFinish,
          runningAthlete.identifier,
          timePoints[1].identifier,
          simulationType.CLIENT
        )
      );
      yield put(actions.updateRaceData([runningAthlete]));
      break;
    } else {
      yield put(actions.updateRaceData([runningAthlete]));
      yield delay(Math.floor(Math.random() * 700) + 900);
    }
  }
}

// Starts the Race depending on whether it is client simulated or server simulated.
// In case of server simulated, the generator puts a new action generator to fetch live
// details of the progress of the race from the server that reiterates every second
// until the race is over or manually stopped by the user.
export function* startRaceSaga(action) {
  yield put(actions.startRaceStart());

  let url;
  switch (action.simulationType) {
    case simulationType.CLIENT:
      url = "client/race-simulation/start";
      break;
    case simulationType.SERVER:
      url = "server/race-simulation/start";
      break;
    default:
      break;
  }

  try {
    // Checking if simulation type requested is of type client or server.
    if (action.simulationType === simulationType.CLIENT) {
      const response = yield axios.get(url);

      const responseMessage = response.data.responseMessage;
      if (responseMessage !== "starting race") {
        throwError("Server responded with '" + responseMessage + "'");
      }

      const athletes = response.data.athletes;
      const timePoints = response.data.timePoints;

      let newAthletes = athletes.map(athlete => {
        const newAthlete = athlete;
        newAthlete.location = 0;
        return newAthlete;
      });

      yield put(actions.startRaceSuccess(newAthletes, timePoints));

      // Note the time
      const stopwatchStart = new Date();

      // Concurrently let each athlete run
      yield all(
        newAthletes.map(athlete =>
          call(
            runningAthleteSaga,
            athlete.identifier,
            timePoints,
            stopwatchStart
          )
        )
      );

      // Making sure not to send the stop request twice if the user has pushed
      // the stop button already.
      const raceStatus = yield select(getRaceStatus);
      if (!raceStatus) {
        return;
      }

      // Race finished normally without user interruption so we stop the race.
      yield put(actions.stopRace(simulationType.CLIENT));
    } else if (action.simulationType === simulationType.SERVER) {
      const response = yield axios.get(url);

      const responseMessage = response.data.responseMessage;
      if (responseMessage !== "starting race") {
        throwError("Server responded with '" + responseMessage + "'");
      }

      const athletes = response.data.athletes;

      // Instantiate each athlete received from the server with a location of 0
      // to be on the safe side. This may require optimization in the future.
      let newAthletes = athletes.map(athlete => {
        const newAthlete = athlete;
        newAthlete.location = 0;
        return newAthlete;
      });

      // Race is ready to begin, athletes get populated. We don't need timepoints for
      // server emulation.
      yield put(actions.startRaceSuccess(newAthletes, []));

      // We start the generator up to fetch live data from the server every second.
      yield put(actions.fetchFeed());
    }
  } catch (error) {
    console.log(error);
    if (error.response) {
      yield put(actions.startRaceFail(error.status));
    } else if (error.request) {
      yield put(
        actions.startRaceFail("we weren't able to connect to the backend")
      );
    } else if (error.message) {
      yield put(actions.startRaceFail(error.message));
    } else {
      yield put(actions.startRaceFail(error));
    }
  }
}

export function* commitFeedSaga(action) {
  yield put(actions.commitFeedStart());

  let url = "client/race-simulation/push/current-standings";

  const postData = {
    requestCommand: "commit feed",
    athletes: [
      {
        timeElapsed: action.timeElapsed,
        identifier: action.athleteIdentifier,
        timePointIdentifier: action.timePointIdentifier
      }
    ]
  };

  try {
    const raceStatus = yield select(getRaceStatus);
    if (!raceStatus) {
      yield put(actions.commitFeedSuccess());
      return;
    }

    const response = yield axios.post(url, postData);

    const responseMessage = response.data.responseMessage;
    if (
      responseMessage !==
      "race is in process, data was committed, showing committed data"
    ) {
      throwError("Server responded with '" + responseMessage + "'");
    }
    yield put(actions.commitFeedSuccess());
  } catch (error) {
    console.log(error);
    if (error.response) {
      yield put(actions.commitFeedFail(error.status));
    } else if (error.request) {
      yield put(
        actions.commitFeedFail("we weren't able to connect to the backend")
      );
    } else if (error.message) {
      yield put(actions.commitFeedFail(error.message));
    } else {
      yield put(actions.commitFeedFail(error));
    }
  }
}

export function* fetchFeedSaga() {
  let url = "server/race-simulation/fetch/live-standings";

  try {
    while (true) {
      const raceStatus = yield select(getRaceStatus);
      if (!raceStatus) {
        return;
      }

      const response = yield axios.get(url);

      const responseMessage = response.data.responseMessage;
      const athletes = response.data.athletes;

      if (!responseMessage || !athletes) {
        console.log(response);
        throwError("Server encountered and error, please check console log.");
        break;
      }

      if (!athletes || athletes.length <= 0) {
        console.log(athletes);
        throwError("Invalid data acquired from the server");
        break;
      }

      const newAthletes = athletes.map(athlete => {
        const newAthlete = athlete;
        if (newAthlete.location) return newAthlete;
        newAthlete.location = 0;
        return newAthlete;
      });

      if (athletes.every(athlete => athlete.timeTakenToFinish !== -1)) {
        yield put(actions.setRaceData(newAthletes));
        break;
      }

      yield put(actions.setRaceData(newAthletes));

      if (responseMessage !== "race is in process, showing live data") {
        if (
          responseMessage ===
          "race is currently not in process, showing last race data"
        ) {
          break;
        }

        throwError("Server responded with '" + responseMessage + "'");
        break;
      }

      // Fetching race updates from the server every second
      yield delay(1000);
    }

    // End Race
    yield put(actions.stopRace(simulationType.SERVER));
  } catch (error) {
    console.log(error);
    if (error.response) {
      yield put(actions.fetchFeedFail(error.status));
    } else if (error.request) {
      yield put(
        actions.fetchFeedFail("we weren't able to connect to the backend")
      );
    } else if (error.message) {
      yield put(actions.fetchFeedFail(error.message));
    } else {
      yield put(actions.fetchFeedFail(error));
    }
  }
}

export function* stopRaceSaga(action) {
  yield put(actions.stopRaceStart());

  let url;
  switch (action.simulationType) {
    case simulationType.CLIENT:
      url = "/client/race-simulation/stop";
      break;
    case simulationType.SERVER:
      url = "/server/race-simulation/stop";
      break;
    default:
      break;
  }

  try {
    const response = yield axios.get(url);

    const responseMessage = response.data.responseMessage;
    if (
      responseMessage !== "stopping race" &&
      responseMessage !== "race is currently not in process"
    ) {
      throwError("Server responded with '" + responseMessage + "'");
    }

    yield put(actions.stopRaceSuccess());
  } catch (error) {
    console.log(error);
    if (error.response) {
      yield put(actions.stopRaceFail(error.status));
    } else if (error.request) {
      yield put(
        actions.stopRaceFail("we weren't able to connect to the backend")
      );
    } else if (error.message) {
      yield put(actions.stopRaceFail(error.message));
    } else {
      yield put(actions.stopRaceFail(error));
    }
  }
}
