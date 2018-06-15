import React from "react";
import PropTypes from "prop-types";
import FontAwesomeIcon from "@fortawesome/react-fontawesome";

// This component is actually a wrapper for the button in the project to start and stop the race.
const RaceControl = props => (
  <button
    className="btn btn-primary btn-lg btn-block"
    onClick={props.callback}
    disabled={props.loadingData && !props.manualStop}
  >
    {props.loadingData && !props.manualStop ? (
      <FontAwesomeIcon icon="spinner" spin className="spin-margin-right" />
    ) : null}
    {props.text}
  </button>
);

RaceControl.propTypes = {
  callback: PropTypes.func.isRequired,
  loadingData: PropTypes.bool.isRequired,
  manualStop: PropTypes.bool.isRequired,
  text: PropTypes.string.isRequired
};

export default RaceControl;
