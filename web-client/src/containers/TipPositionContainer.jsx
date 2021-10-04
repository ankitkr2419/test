import React from "react";

import TipPositionComponent from "components/TipPosition";
import { useSelector } from "react-redux";

const TipPositionContainer = (props) => {
  const editReducer = useSelector((state) => state.editProcessReducer);
  const editReducerData = editReducer.toJS();

  const cartridge1DetailsReducer = useSelector(
    (state) => state.cartridge1DetailsReducer
  );
  const cartridge2DetailsReducer = useSelector(
    (state) => state.cartridge2DetailsReducer
  );

  return (
    <TipPositionComponent
      editReducerData={editReducerData}
      cartridge1Details={cartridge1DetailsReducer.cartridgeDetails}
      cartridge2Details={cartridge2DetailsReducer.cartridgeDetails}
    />
  );
};

TipPositionContainer.propTypes = {};

export default TipPositionContainer;
