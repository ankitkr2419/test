import React, { useEffect } from "react";

import LabWareComponent from "components/Labware";
import { useDispatch, useSelector } from "react-redux";
import {
  getCartridgeActionInitiated,
  getTipsActionInitiated,
  getTipsAndTubesActionInitiated,
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

  useEffect(() => {
    dispatch(
      getTipsAndTubesActionInitiated({
        deckName: currentDeckName,
        token: token,
      })
    );
    dispatch(
      getCartridgeActionInitiated({ deckName: currentDeckName, token: token })
    );
    dispatch(
      getTipsActionInitiated({ deckName: currentDeckName, token: token })
    );
    dispatch(
      getTubesActionInitiated({ deckName: currentDeckName, token: token })
    );
  }, [dispatch]);
  return <LabWareComponent />;
};

LabwareContainer.propTypes = {};

export default LabwareContainer;
