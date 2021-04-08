import React from "react";
import ReactDOM from "react-dom";

import { Provider } from "react-redux";
import { HashRouter as Router, Switch, Route } from "react-router-dom";

import store from "store";
import Routes from "Routes";

import * as serviceWorker from "./serviceWorker";

import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

ReactDOM.render(
  <Provider store={store}>
    <Router>
      <Switch>
        <Route path="/" component={Routes} />
        <ToastContainer position="top-right" autoClose={5000} />
      </Switch>
    </Router>
  </Provider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
