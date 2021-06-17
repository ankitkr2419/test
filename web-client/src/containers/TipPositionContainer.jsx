import React from "react";

import TipPositionComponent from "components/TipPosition";
import { useSelector } from "react-redux";

const TipPositionContainer = (props) => {
  // const editReducer = useSelector((state) => state.editProcessReducer);
  // const editReducerData = editReducer.toJS();

  const editReducerData = {
    id: "08b17de4-e041-40cd-9fc0-21e396c61932",
    process_id: "55b1b4cd-533d-41f8-8ded-ddc9fba33cb2",
    created_at: "2021-06-04T20:15:43.860643Z",
    updated_at: "2021-06-04T21:29:25.333753Z",
    type: "cartridge_2",
    deck_position: "sample_tube",
    position: 2,
    height: 5.5,
  };

  return <TipPositionComponent editReducerData={null} />;
};

TipPositionContainer.propTypes = {};

export default TipPositionContainer;
