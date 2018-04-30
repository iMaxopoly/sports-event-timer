import React, { Component, Fragment } from "react";
import PropTypes from "prop-types";
import FontAwesomeIcon from "@fortawesome/react-fontawesome";

// Athlete is a stateful component which can arguably be a stateless component.
// It seems the 3rd party animation dependency requires Athlete to be stateful for
// it's proper functioning as it uses refs.
// Athlete holds the table rows for individual athletes in the demo, indicating
// their attributes during the race.
class Athlete extends Component {
  render() {
    return (
      <tr className={this.props.style}>
        <th scope="row">{this.props.startNumber}</th>
        <td>
          <strong>{this.props.fullName}</strong>
        </td>
        <td>{this.props.location} m</td>
        <td>
          {this.props.timeTakenToReachFinishCorridor === -1 ? (
            <FontAwesomeIcon icon="hourglass" spin />
          ) : (
            <FontAwesomeIcon icon="check" />
          )}
        </td>
        <td>
          {this.props.timeTakenToFinish === -1 ? (
            <FontAwesomeIcon icon="hourglass" spin />
          ) : (
            <FontAwesomeIcon icon="check" />
          )}
        </td>
        <td>
          {this.props.timeTakenToFinish === -1 ? (
            <FontAwesomeIcon icon="hourglass" spin />
          ) : (
            <Fragment>
              <span>{(this.props.timeTakenToFinish / 1000).toFixed(2)} s </span>
              <FontAwesomeIcon icon="flag-checkered" />
            </Fragment>
          )}
        </td>
      </tr>
    );
  }
}

Athlete.propTypes = {
  fullName: PropTypes.string.isRequired,
  location: PropTypes.number.isRequired,
  startNumber: PropTypes.number.isRequired,
  style: PropTypes.string.isRequired,
  timeTakenToFinish: PropTypes.number.isRequired,
  timeTakenToReachFinishCorridor: PropTypes.number.isRequired
};

export default Athlete;
