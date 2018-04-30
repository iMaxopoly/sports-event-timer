import React, { Component } from "react";
import RaceEvent from "./containers/RaceEvent/RaceEvent";

import "bootstrap/scss/bootstrap.scss";
import "./index.scss";

// App bootstraps our RaceEvent container
class App extends Component {
  render() {
    return <RaceEvent />;
  }
}

export default App;
