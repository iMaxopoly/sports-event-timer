import React from "react";
import ReactDOM from "react-dom";
import { createStore } from "redux";
import { Provider } from "react-redux";
import reducer from "./store/reducers/reducer";
import App from "./App";

it("renders without crashing", () => {
  const store = createStore(reducer);
  const main = (
    <Provider store={store}>
      <App />
    </Provider>
  );
  const div = document.createElement("div");
  ReactDOM.render(main, div);
  ReactDOM.unmountComponentAtNode(div);
});
