import React from "react";
import PropTypes from "prop-types";

const Alert = props => (
  <div className="alert alert-light fade show text-center">
    <strong>{props.message}</strong>
  </div>
);

Alert.propTypes = {
  message: PropTypes.string.isRequired
};

export default Alert;
