import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import styled from "styled-components";
import RecipeListingComponent from "components/RecipeListing";
import {
  duplicateRecipeInitiated,
  recipeListingInitiated,
  saveRecipeDataForDeck,
} from "action-creators/recipeActionCreators";
import { ROUTES, RUN_RECIPE_TYPE } from "appConstants";
import { Redirect } from "react-router-dom";
import { deckBlockReset } from "action-creators/loginActionCreators";
const RecipeListing = styled.div`
  .landing-content {
    padding: 1.25rem 4.5rem 0.875rem 4.5rem;
    &::after {
      height: 9.125rem;
    }
    .recipe-listing-cards {
      height: 30.75rem;
    }
  }
`;

const RecipeListingContainer = (props) => {
  const dispatch = useDispatch();

  const [
    isOperatorRunRecipeCarousalModalVisible,
    setOperatorRunRecipeCarousalModalVisible,
  ] = useState(false);
  const [selectedRecipeData, setSelectedRecipeData] = useState({});
  const [token, setToken] = useState();
  const [runRecipesmodal, setRunRecipesModal] = useState(false);
  const [runRecipeType, setRunRecipeType] = useState(
    RUN_RECIPE_TYPE.CONTINUOUS_RUN
  );

  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  const cleanUpReducer = useSelector((state) => state.cleanUpReducer);
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const [recipeFetched, setRecipeFetched] = useState(false);

  let deckName = activeDeckObj.name;
  let isAdmin = activeDeckObj.isAdmin;
  if (!token && activeDeckObj.token) setToken(activeDeckObj.token);

  const recipeReducerDataOfActiveDeck = recipeActionReducer.decks.find(
    (deck) => deck.name === deckName
  );
  const recipeData = recipeReducerDataOfActiveDeck.allRecipeData;
  const cleanUpReducerDataOfActiveDeck = cleanUpReducer.decks.find(
    (deck) => deck.name === deckName
  );

  const isProcessInProgress =
    recipeReducerDataOfActiveDeck.showProcess ||
    cleanUpReducerDataOfActiveDeck.showCleanUp;

  const handleCarousalModal = (
    prevState = isOperatorRunRecipeCarousalModalVisible
  ) => {
    //clear recipe data if run recipe closed
    if (prevState === true) {
      setSelectedRecipeData({});
    }
    setOperatorRunRecipeCarousalModalVisible(!prevState);
  };

  useEffect(() => {
    if (token && !recipeFetched) {
      dispatch(recipeListingInitiated(token, deckName));
      setRecipeFetched(true);
    }
  }, [token, recipeFetched, dispatch]);

  /** Blocked deck must be unblocked */
  useEffect(() => {
    if (activeDeckObj?.isDeckBlocked) {
      dispatch(deckBlockReset({ deckName: activeDeckObj.name }));
    }
  }, [activeDeckObj.isDeckBlocked]);

  /** Reset selectedRecipeData if deck is switched */
  useEffect(() => {
    if (
      selectedRecipeData?.deckName &&
      selectedRecipeData.deckName !== activeDeckObj.name
    ) {
      setSelectedRecipeData({});
    }
  });

  useEffect(() => {
    if (isProcessInProgress) {
      setSelectedRecipeData({});
    }
  }, [isProcessInProgress]);

  const returnRecipeDetails = (data) => {
    setSelectedRecipeData({ data, deckName });
  };

  const onConfirmedRecipeSelection = () => {
    let { data, deckName } = selectedRecipeData;
    if (!deckName) {
      console.error("deckName not found!");
      return;
    }
    dispatch(saveRecipeDataForDeck(data, deckName));
  };

  const onChangeRunRecipeType = (type) => {
    setRunRecipeType(type);
  };

  const onConfirmedRunRecipeByAdmin = () => {
    //save selected recipe data to reducer
    let { data, deckName } = selectedRecipeData;
    if (!deckName) {
      console.error("deckName not found!");
      return;
    }
    data = {
      ...data,
      runRecipeType,
    };
    dispatch(saveRecipeDataForDeck(data, deckName));

    //toggle modal
    toggleRunRecipesModal();
  };
  const toggleRunRecipesModal = () => setRunRecipesModal(!runRecipesmodal);

  const createDuplicateRecipe = (recipeId, recipeName) => {
    const token = activeDeckObj.token;
    dispatch(
      duplicateRecipeInitiated({ recipeId, token, deckName, recipeName })
    );
  };

  if (!activeDeckObj.isLoggedIn) return <Redirect to={`/${ROUTES.landing}`} />;

  return (
    <RecipeListing>
      {/* {(!isLoading) && <Loader/>} */}
      <RecipeListingComponent
        isProcessInProgress={isProcessInProgress}
        recipeData={recipeData}
        isOperatorRunRecipeCarousalModalVisible={
          isOperatorRunRecipeCarousalModalVisible
        }
        handleCarousalModal={handleCarousalModal}
        selectedRecipeData={selectedRecipeData}
        returnRecipeDetails={returnRecipeDetails}
        onConfirmedRecipeSelection={onConfirmedRecipeSelection}
        onConfirmedRunRecipeByAdmin={onConfirmedRunRecipeByAdmin}
        runRecipesmodal={runRecipesmodal}
        toggleRunRecipesModal={toggleRunRecipesModal}
        createDuplicateRecipe={createDuplicateRecipe}
        runRecipeType={runRecipeType}
        onChangeRunRecipeType={onChangeRunRecipeType}
        deckName={deckName}
        isAdmin={isAdmin}
      />
    </RecipeListing>
  );
};

export default RecipeListingContainer;
