import React from "react";

import TipPositionComponent from "components/TipPosition";
import { useSelector } from "react-redux";

const TipPositionContainer = (props) => {
  const editReducer = useSelector((state) => state.editProcessReducer);
  const editReducerData = editReducer.toJS();

  return <TipPositionComponent editReducerData={editReducerData} />;
};

TipPositionContainer.propTypes = {};

export default TipPositionContainer;
