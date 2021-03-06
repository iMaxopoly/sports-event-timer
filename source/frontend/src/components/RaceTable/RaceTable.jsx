import React, { Component } from "react";
import PropTypes from "prop-types";
import FontAwesomeIcon from "@fortawesome/react-fontawesome";
import FlipMove from "react-flip-move";
import Alert from "../Alert/Alert";
import Athlete from "../Athlete/Athlete";

// RaceTable wraps the Athlete container in a table showing the entirety of the
// race in an animated manner.
class RaceTable extends Component {
  constructor(props) {
    super(props);

    this.tableHeaderStyle = "table-active";
    this.athleteRowStyles = [
      "table-primary",
      "table-secondary",
      "table-info",
      "table-dark",
      "table-light"
    ];
  }

  // componentDidUpdate(prevProps, prevState) {
  //   console.log('test', prevProps, prevState)
  // }

  lastErrorRenderer = error => <Alert message={error} />;

  raceNotInProgressRenderer = () => (
    <Alert message="Race is currently not underway." />
  );

  athleteSortComparison = (a, b) => (b < a ? 1 : a < b ? -1 : 0);

  // Sorts the athletes based on time taken to finish the race.
  athletesSorter = athletes => {
    if (athletes.every(athlete => athlete.timeTakenToFinish !== -1)) {
      return athletes
        .map(athlete => ({ ...athlete }))
        .sort((a, b) =>
          this.athleteSortComparison(a.timeTakenToFinish, b.timeTakenToFinish)
        );
    }

    // Sorts the athletes based on time taken to reach the corridor.
    if (
      athletes.every(athlete => athlete.timeTakenToReachFinishCorridor !== -1)
    ) {
      return athletes
        .map(athlete => ({ ...athlete }))
        .sort((a, b) =>
          this.athleteSortComparison(
            a.timeTakenToReachFinishCorridor,
            b.timeTakenToReachFinishCorridor
          )
        );
    }

    // Sorts the athletes based on their location during the race.
    return athletes
      .map(athlete => ({ ...athlete }))
      .sort((a, b) => this.athleteSortComparison(a.location, b.location))
      .reverse();
  };

  // Populates the table with athlete attributes in an on-going race
  tablePopulator = athletes =>
    athletes.map(athlete => (
      <Athlete
        key={athlete.startNumber}
        startNumber={athlete.startNumber}
        style={this.athleteRowStyles[athlete.startNumber - 1]}
        fullName={athlete.fullName}
        location={athlete.location}
        timeTakenToReachFinishCorridor={athlete.timeTakenToReachFinishCorridor}
        timeTakenToFinish={athlete.timeTakenToFinish}
      />
    ));

  // Renders the table structure
  tableRenderer = tableData => (
    <div className="table-responsive">
      <table className="table">
        <FlipMove typeName="tbody" enterAnimation="fade" leaveAnimation="fade">
          <tr className={this.tableHeaderStyle}>
            <th>
              <FontAwesomeIcon icon="caret-down" /> Start Number
            </th>
            <th>
              <FontAwesomeIcon icon="caret-down" /> Full Name
            </th>
            <th>
              <FontAwesomeIcon icon="caret-down" /> Distance Raced
            </th>
            <th>
              <FontAwesomeIcon icon="caret-down" /> In Finish Corridor?
            </th>
            <th>
              <FontAwesomeIcon icon="caret-down" /> Has Finished?
            </th>
            <th>
              <FontAwesomeIcon icon="caret-down" /> Finish Time
            </th>
          </tr>
          {tableData}
        </FlipMove>
      </table>
    </div>
  );

  isRaceInProgress = () =>
    this.props.raceInProgress ||
    this.props.athletes.length !== 0 ||
    this.props.lastKnownError;

  isTableReadyForDisplay = () =>
    (this.props.athletes.length !== 0 && !this.props.lastKnownError) ||
    (this.props.athletes.length !== 0 && this.props.manualStop);

  isErrorPresent = () => this.props.lastKnownError && !this.props.manualStop;

  render() {
    return (
      <div className="card">
        <div className="card-body">
          <div className="card-text">
            {!this.isRaceInProgress()
              ? this.raceNotInProgressRenderer()
              : this.isTableReadyForDisplay()
                ? this.tableRenderer(
                    this.tablePopulator(
                      this.athletesSorter(this.props.athletes)
                    )
                  )
                : this.isErrorPresent()
                  ? this.lastErrorRenderer(this.props.lastKnownError)
                  : null}
          </div>
        </div>
      </div>
    );
  }
}

RaceTable.propTypes = {
  athletes: PropTypes.array,
  lastKnownError: PropTypes.string,
  manualStop: PropTypes.bool,
  raceInProgress: PropTypes.bool.isRequired
};

export default RaceTable;
