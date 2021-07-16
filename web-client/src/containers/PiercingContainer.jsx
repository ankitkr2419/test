import React from "react";

import PiercingComponent from "components/Piercing";
import { useSelector } from "react-redux";

const PiercingContainer = () => {
  const editReducer = useSelector((state) => state.editProcessReducer);
  const editReducerData = editReducer.toJS();

  return <PiercingComponent editReducerData={editReducerData} />;
};

export default PiercingContainer;
