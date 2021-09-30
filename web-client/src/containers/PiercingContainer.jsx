import React from "react";

import PiercingComponent from "components/Piercing";
import { useSelector } from "react-redux";

const PiercingContainer = () => {
  const editReducer = useSelector((state) => state.editProcessReducer);
  const editReducerData = editReducer.toJS();

  const cartridge1DetailsReducer = useSelector(
    (state) => state.cartridge1DetailsReducer
  );
  const cartridge2DetailsReducer = useSelector(
    (state) => state.cartridge2DetailsReducer
  );

  return (
    <PiercingComponent
      editReducerData={editReducerData}
      cartridge1Details={cartridge1DetailsReducer.cartridgeDetails}
      cartridge2Details={cartridge2DetailsReducer.cartridgeDetails}
    />
  );
};

export default PiercingContainer;
