import React, { useEffect } from "react";

import LabWareComponent from "components/Labware";
import { useDispatch, useSelector } from "react-redux";
import {
  getCartridgeActionInitiated,
  getTipsActionInitiated,
  getTubesActionInitiated,
} from "action-creators/saveNewRecipeActionCreators";

const LabwareContainer = (props) => {
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const currentDeckName = activeDeckObj.name;
  const token = activeDeckObj.token;

  const params = { deckName: currentDeckName, token: token };
  useEffect(() => {
    dispatch(getCartridgeActionInitiated(params));
    dispatch(getTipsActionInitiated(params));
    dispatch(getTubesActionInitiated(params));
  }, [dispatch]);
  return <LabWareComponent />;
};

LabwareContainer.propTypes = {};

export default LabwareContainer;
