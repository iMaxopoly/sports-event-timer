import React from "react";
import ReactDOM from "react-dom";
import { applyMiddleware, createStore } from "redux";
import { Provider } from "react-redux";
import createSagaMiddleware from "redux-saga";
import axios from "axios";
import fontawesome from "@fortawesome/fontawesome";
import faSpinner from "@fortawesome/fontawesome-free-solid/faSpinner";
import faHourGlass from "@fortawesome/fontawesome-free-solid/faHourglass";
import faCheck from "@fortawesome/fontawesome-free-solid/faCheck";
import faFlag from "@fortawesome/fontawesome-free-solid/faFlagCheckered";
import faCaretDown from "@fortawesome/fontawesome-free-solid/faCaretDown";
import faMapMarker from "@fortawesome/fontawesome-free-solid/faMapMarkerAlt";

import reducer from "./store/reducers/reducer";
import watchNetwork from "./store/sagas";
import RaceEvent from "./containers/RaceEvent/RaceEvent";

import * as serviceWorker from "./serviceWorker";

import "bootstrap/scss/bootstrap.scss";
import "./index.scss";

// fontawesome setup
fontawesome.library.add(
  faSpinner,
  faHourGlass,
  faCheck,
  faFlag,
  faCaretDown,
  faMapMarker
);

// store and store middleware declaration
const sagaMiddleware = createSagaMiddleware();
const store = createStore(reducer, applyMiddleware(sagaMiddleware));

sagaMiddleware.run(watchNetwork);

// axios interceptors
axios.defaults.baseURL = "http://localhost:8082";
axios.defaults.headers.post["Content-Type"] =
  "application/x-www-form-urlencoded";

const main = (
  <Provider store={store}>
    <RaceEvent />
  </Provider>
);

ReactDOM.render(main, document.getElementById("root"));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
