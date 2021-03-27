import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import RecipeListingComponent from "components/RecipeListing";
import { recipeListingInitiated } from "action-creators/recipeActionCreators";
// import { Loader } from 'shared-components'

const RecipeListingContainer = (props) => {
  const dispatch = useDispatch();
  // const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  // const { recipeData } = recipeActionReducer
  const recipeData = [
    {
      id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de4",
      name: "covid",
      description: "Recipe for covid extraction",
      pos_1: 1,
      pos_2: 2,
      pos_3: 3,
      pos_4: 4,
      pos_5: 5,
      pos_cartridge_1: 1,
      pos_7: 6,
      pos_cartridge_2: 2,
      pos_9: 7,
      process_count: 19,
      created_at: "2021-03-11T17:34:16.101723Z",
      updated_at: "2021-03-11T17:34:16.101723Z",
    },
    {
      id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de5",
      name: "tip_docking",
      description: "Recipe for tip_docking position on deck",
      pos_1: 1,
      pos_2: 2,
      pos_3: 3,
      pos_4: 4,
      pos_5: 5,
      pos_cartridge_1: 1,
      pos_7: 6,
      pos_cartridge_2: 2,
      pos_9: 7,
      process_count: 1,
      created_at: "2021-03-11T17:34:16.21166Z",
      updated_at: "2021-03-11T17:34:16.21166Z",
    },
  ];

  useEffect(() => dispatch(recipeListingInitiated()), [dispatch]);

  return (
    <>
      {/* {(!isLoading) && <Loader/>} */}
      {recipeData && <RecipeListingComponent allRecipeData={recipeData} />}
    </>
  );
};

RecipeListingContainer.propTypes = {};

export default RecipeListingContainer;
