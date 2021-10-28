import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import { CARTRIDGE_1, CARTRIDGE_2 } from "appConstants";
import SelectProcessPageComponent from "components/SelectProcess";
import {
  getCartridge1ActionInitiated,
  getCartridge2ActionInitiated,
} from "action-creators/saveNewRecipeActionCreators";

const SelectProcessContainer = (props) => {
  // const { cartridgeApiCall = false } = props;
  const dispatch = useDispatch();

  /** get Active deck details */
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);

  /* get Recipe details */
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const { recipeDetails } = recipeDetailsReducer;

  /** get isApiCalled field */
  const cartridge1DetailsReducer = useSelector(
    (state) => state.cartridge1DetailsReducer
  );
  const cartridge2DetailsReducer = useSelector(
    (state) => state.cartridge2DetailsReducer
  );

  /** API calls for cartridge details */
  useEffect(() => {
    const { token } = activeDeckObj;
    const { name, pos_cartridge_1, pos_cartridge_2 } = recipeDetails;

    /** Checks if API is already called */
    const { isApiCalled: cartridge1ApiCalled } = cartridge1DetailsReducer;
    const { isApiCalled: cartridge2ApiCalled } = cartridge2DetailsReducer;

    let params = { deckName: name, token: token };
    if (pos_cartridge_1 && cartridge1ApiCalled === false) {
      params = { ...params, id: pos_cartridge_1, type: CARTRIDGE_1 };
      dispatch(getCartridge1ActionInitiated(params));
    }
    if (pos_cartridge_2 && cartridge2ApiCalled === false) {
      params = { ...params, id: pos_cartridge_2, type: CARTRIDGE_2 };
      dispatch(getCartridge2ActionInitiated(params));
    }
  }, [recipeDetails]);

  return <SelectProcessPageComponent />;
};

export default SelectProcessContainer;
