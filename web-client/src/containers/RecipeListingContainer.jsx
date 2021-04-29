import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import styled from "styled-components";
import RecipeListingComponent from "components/RecipeListing";

import { recipeListingInitiated, saveRecipeDataForDeck } from "action-creators/recipeActionCreators";
import { ROUTES } from "appConstants";
import { Redirect } from "react-router-dom";
import { useHistory } from "react-router-dom";
import { toast } from "react-toastify";

// import { Loader } from 'shared-components'
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
  const history = useHistory();
  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  const [recipeFetched, setRecipeFetched] = useState(false)
  // const recipeData = [
  //   {
  //     id: "28101940-718b-4937-913d-39cb6b9057ba",
  //     name: "covid Extraction",
  //     description: "Covid Recipe",
  //     pos_1: 1,
  //     pos_2: 2,
  //     pos_3: 3,
  //     pos_4: 4,
  //     pos_5: 5,
  //     pos_cartridge_1: 1,
  //     pos_7: 6,
  //     pos_cartridge_2: 2,
  //     pos_9: 7,
  //     process_count: 33,
  //     created_at: "2021-04-09T19:00:55.233325Z",
  //     updated_at: "2021-04-09T19:00:55.233325Z",
  //     isPublished: true,
  //   },
  //   {
  //     id: "a1fbbacb-5078-4554-bf40-9cf07348e4fe",
  //     name: "covid PCR",
  //     description: "Covid Recipe",
  //     pos_1: 1,
  //     pos_2: 2,
  //     pos_3: 3,
  //     pos_4: 4,
  //     pos_5: 5,
  //     pos_cartridge_1: 1,
  //     pos_7: 6,
  //     pos_cartridge_2: 2,
  //     pos_9: 7,
  //     process_count: 8,
  //     created_at: "2021-04-09T19:01:44.405416Z",
  //     updated_at: "2021-04-09T19:01:44.405416Z",
  //     isPublished: false,
  //   },
  //   {
  //     id: "28101940-718b-4937-913d-39cb6b9057bc",
  //     name: "covid Extraction",
  //     description: "Covid Recipe",
  //     pos_1: 1,
  //     pos_2: 2,
  //     pos_3: 3,
  //     pos_4: 4,
  //     pos_5: 5,
  //     pos_cartridge_1: 1,
  //     pos_7: 6,
  //     pos_cartridge_2: 2,
  //     pos_9: 7,
  //     process_count: 33,
  //     created_at: "2021-04-09T19:00:55.233325Z",
  //     updated_at: "2021-04-09T19:00:55.233325Z",
  //     isPublished: false,
  //   },
  //   {
  //     id: "a1fbbacb-5078-4554-bf40-9cf07348e4ff",
  //     name: "covid PCR",
  //     description: "Covid Recipe",
  //     pos_1: 1,
  //     pos_2: 2,
  //     pos_3: 3,
  //     pos_4: 4,
  //     pos_5: 5,
  //     pos_cartridge_1: 1,
  //     pos_7: 6,
  //     pos_cartridge_2: 2,
  //     pos_9: 7,
  //     process_count: 8,
  //     created_at: "2021-04-09T19:01:44.405416Z",
  //     updated_at: "2021-04-09T19:01:44.405416Z",
  //     isPublished: true,
  //   },
  // ];

  const [isProcessInProgress, setIsProcessInProcess] = useState(false); //to be changed and data should come from Reducer
  const [
    isOperatorRunRecipeCarousalModalVisible,
    setOperatorRunRecipeCarousalModalVisible,
  ] = useState(false);

  const [selectedRecipeData, setSelectedRecipeData] = useState({}) 
  const [token, setToken] = useState()

  const handleCarousalModal = (
    prevState = isOperatorRunRecipeCarousalModalVisible
  ) => {
    setOperatorRunRecipeCarousalModalVisible(!prevState);
  };

  //useEffect(() => {
    // dispatch(recipeListingInitiated());
  //}, [dispatch]);
  useEffect(() => {
    if(token && !recipeFetched){
      dispatch(recipeListingInitiated(token, deckName))
      setRecipeFetched(true)
    }
  }, [token])

  const loginReducer = useSelector(
    (state) => state.loginReducer
  );

  const loginReducerData = loginReducer.toJS()
  let activeDeckObj = loginReducerData && loginReducerData.decks.find(deck => deck.isActive)
  if(!activeDeckObj.isLoggedIn){
    return <Redirect to={`/${ROUTES.landing}`} />;
  }
  let deckName  = activeDeckObj.name
  let isAdmin = activeDeckObj.isAdmin
  if(!token)
    setToken(activeDeckObj.token);

  const  recipeReducerDataOfActiveDeck = recipeActionReducer.decks.find(deck => deck.name === deckName);
  const recipeData = recipeReducerDataOfActiveDeck.allRecipeData;
  const isLoading = recipeReducerDataOfActiveDeck.isLoading;
  const returnRecipeDetails = (data) => {
    // let requiredData  =  {
    //   data,
    //   deckName //active deck
    // }
    // console.log("DATA returned--->", requiredData);
    // store data in reducer
    setSelectedRecipeData({data, deckName})
  };

  const onConfirmedRecipeSelection = () => {
    let { data, deckName } = selectedRecipeData;
    if(!deckName){
      console.error('deckName not found!')
      return;
    } 
    dispatch(saveRecipeDataForDeck(data, deckName))
    // history.push(ROUTES.landing);//go to landing page
    // return <Redirect to={`/${ROUTES.landing}`} />;
  }

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
        returnRecipeDetails={returnRecipeDetails}
        onConfirmedRecipeSelection={onConfirmedRecipeSelection}
        deckName={deckName}
        isAdmin={isAdmin}
      />
    </RecipeListing>
  );
};

RecipeListingContainer.propTypes = {};

export default RecipeListingContainer;
