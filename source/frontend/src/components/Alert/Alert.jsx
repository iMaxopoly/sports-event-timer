import React from "react";
import PropTypes from "prop-types";

const alert = props => (
  <div className="alert alert-light fade show text-center">
    <strong>{props.message}</strong>
  </div>
);

alert.propTypes = {
  message: PropTypes.string.isRequired
};

export default alert;
