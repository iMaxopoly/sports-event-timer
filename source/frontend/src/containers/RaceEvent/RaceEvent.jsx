import React, { Component } from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import * as actions from "../../store/actions";
import Layout from "../../components/Layout/Layout";
import RaceControl from "../../components/RaceControl/RaceControl";
import { simulationType } from "../../simulationTypes";

// RaceEvent
class RaceEvent extends Component {
  constructor(props) {
    super(props);
    this.startRace = this.startRace.bind(this);
    this.stopRace = this.stopRace.bind(this);
    this.changeSimulation = this.changeSimulation.bind(this);
  }

  startRace() {
    this.props.onStartRace(this.props.simulationType, this.props.timePoints);
  }

  stopRace() {
    this.props.onStopRace(this.props.simulationType);
  }

  changeSimulation(event) {
    if (event.target.value === simulationType.CLIENT.toString()) {
      this.props.onChangeSimulation(simulationType.CLIENT);
    } else {
      this.props.onChangeSimulation(simulationType.SERVER);
    }
  }

  render() {
    return (
      <Layout callback={this.changeSimulation} {...this.props}>
        {this.props.raceInProgress ? (
          <RaceControl
            callback={this.stopRace}
            text={this.props.loadingData ? "Loading..." : "Stop Race"}
            loadingData={this.props.loadingData}
            manualStop={this.props.manualStop}
            raceInProgress={this.props.raceInProgress}
            simulationType={this.props.simulationType}
          />
        ) : (
          <RaceControl
            callback={this.startRace}
            text={
              this.props.loadingData && !this.props.manualStop
                ? "Loading..."
                : "Start Race"
            }
            loadingData={this.props.loadingData}
            manualStop={this.props.manualStop}
          />
        )}
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  simulationType: state.simulationType,
  athletes: state.athletes,
  timePoints: state.timePoints,
  loadingData: state.loadingData,
  manualStop: state.manualStop,
  raceInProgress: state.raceInProgress,
  lastKnownError: state.lastKnownError
});

const mapDispatchToProps = dispatch => ({
  onStartRace: (simulationType, timePoints) =>
    dispatch(actions.startRace(simulationType, timePoints)),
  onStopRace: simulationType => dispatch(actions.stopRace(simulationType)),
  onChangeSimulation: simulationType =>
    dispatch(actions.changeSimulation(simulationType))
});

RaceEvent.propTypes = {
  loadingData: PropTypes.bool.isRequired,
  onChangeSimulation: PropTypes.any,
  onStartRace: PropTypes.any,
  onStopRace: PropTypes.any,
  raceInProgress: PropTypes.bool.isRequired,
  simulationType: PropTypes.symbol.isRequired,
  timePoints: PropTypes.array
};

export default connect(mapStateToProps, mapDispatchToProps)(RaceEvent);
