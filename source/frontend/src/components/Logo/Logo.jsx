import React from "react";
import logoImage from "../../assets/images/logo.png";

// Logo component, where the image is responsive using a Bootstrap CSS class.
const logo = () => (
  <div className="logo">
    <img src={logoImage} className="img-fluid" alt="Race Event" />
  </div>
);

export default logo;
