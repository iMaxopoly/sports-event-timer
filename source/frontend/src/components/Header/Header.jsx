import React from "react";
import Logo from "../Logo/Logo";

// The header component mainly holding the logo component within it,
// ideally created to hold nav components as well but the project didn't seem
// to require.
const header = () => (
  <div className="row">
    <div className="mx-auto text-center header">
      <Logo />
    </div>
  </div>
);

export default header;
