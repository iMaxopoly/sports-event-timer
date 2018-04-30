import React from "react";
import { storiesOf } from "@storybook/react";
import Alert from "../components/Alert/Alert";

storiesOf("Alert", module).add("with message", () => <Alert message="hello" />);
