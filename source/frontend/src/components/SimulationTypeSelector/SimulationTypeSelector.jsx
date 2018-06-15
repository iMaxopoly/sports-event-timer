import React, { Fragment } from "react";
import PropTypes from "prop-types";
import { simulationType } from "../../simulationTypes";

// The component allows for easy switching between the server emulation demo and
// the client sided emulation demo. It remains dissabled while the race is on-going.
const SimulationTypeSelector = props => (
  <Fragment>
    <div className="text-center">
      <div className="form-check form-check-inline">
        <label className="form-check-label padding-radio">
          <input
            name="radio"
            className="form-check-input"
            type="radio"
            value={simulationType.CLIENT.toString()}
            onChange={props.changeSimulationTypeHandler}
            disabled={props.raceInProgress}
            checked={props.simulationType === simulationType.CLIENT}
          />{" "}
          Client Simulation
        </label>
        <label className="form-check-label padding-radio">
          <input
            name="radio"
            className="form-check-input"
            type="radio"
            value={simulationType.SERVER.toString()}
            onChange={props.changeSimulationTypeHandler}
            disabled={props.raceInProgress}
            checked={props.simulationType === simulationType.SERVER}
          />{" "}
          Server Simulation
        </label>
      </div>
    </div>
  </Fragment>
);

SimulationTypeSelector.propTypes = {
  changeSimulationTypeHandler: PropTypes.func.isRequired,
  raceInProgress: PropTypes.bool.isRequired,
  simulationType: PropTypes.symbol.isRequired
};

export default SimulationTypeSelector;
