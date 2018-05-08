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

    // This is a copy of the athlete from the store considering it
    // is returned by the filter function. Mutations on this array should not
    // affect the store.
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
      yield put.resolve(
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
      yield put.resolve(
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

  try {
    let url, response, responseMessage, athletes, newAthletes;

    // Checking if simulation type requested is of type client or server.
    switch (action.simulationType) {
      case simulationType.CLIENT:
        url = "client/race-simulation/start";
        response = yield axios.get(url);

        responseMessage = response.data.responseMessage;
        if (responseMessage !== "starting race") {
          throwError("Server responded with '" + responseMessage + "'");
        }

        athletes = response.data.athletes;
        const timePoints = response.data.timePoints;

        newAthletes = athletes.map(athlete => {
          const newAthlete = {...athlete};
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
        break;
      case simulationType.SERVER:
        url = "server/race-simulation/start";
        response = yield axios.get(url);

        responseMessage = response.data.responseMessage;
        if (responseMessage !== "starting race") {
          throwError("Server responded with '" + responseMessage + "'");
        }

        athletes = response.data.athletes;

        // Instantiate each athlete received from the server with a location of 0
        // to be on the safe side. This may require optimization in the future.
        newAthletes = athletes.map(athlete => {
          const newAthlete = athlete;
          newAthlete.location = 0;
          return newAthlete;
        });

        // Race is ready to begin, athletes get populated. We don't need timepoints for
        // server emulation.
        yield put(actions.startRaceSuccess(newAthletes, []));

        // We start the generator up to fetch live data from the server every second.
        yield put(actions.fetchFeed());
        break;
      default:
        break;
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

// Pushes current athlete standings to the server which stores the data into the database.
// The saga is strictly for client-based simulation and pushes data onto the server whenever
// an athlete trips a timepoint.
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
      if (responseMessage === "race is currently not in process") {
        yield put(actions.commitFeedSuccess());
        return;
      }
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

// fetchFeedSaga fetches live race data from the server periodically.
// This is essentially only for server simulated race.
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

      if (athletes.length <= 0) {
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

// stopRaceSaga stops the race and flags all underlying concurrent activites to stop.
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
