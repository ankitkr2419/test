import React from "react";
import { useSelector } from "react-redux";

import AspireDispenseComponent from "components/AspireDispense";

const AspireDispenseContainer = () => {
  const cartridge1DetailsReducer = useSelector(
    (state) => state.cartridge1DetailsReducer
  );
  const cartridge2DetailsReducer = useSelector(
    (state) => state.cartridge2DetailsReducer
  );

  return (
    <AspireDispenseComponent
      cartridge1Details={cartridge1DetailsReducer.cartridgeDetails}
      cartridge2Details={cartridge2DetailsReducer.cartridgeDetails}
    />
  );
};

export default AspireDispenseContainer;
