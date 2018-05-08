import React from "react";
import Logo from "../Logo/Logo";

// The Header component mainly holding the Logo component within it,
// ideally created to hold nav components as well but the project didn't seem
// to require.
const Header = () => (
  <div className="row">
    <div className="mx-auto text-center header">
      <Logo />
    </div>
  </div>
);

export default Header;
