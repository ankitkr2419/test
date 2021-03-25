import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import RecipeListingComponent from 'components/RecipeListing';
import { recipeListingInitiated } from 'action-creators/recipeActionCreators';
// import { Loader } from 'shared-components'

const RecipeListingContainer = (props) => {

	const dispatch = useDispatch();
	const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
	const { recipeData } = recipeActionReducer

	useEffect(() => dispatch(recipeListingInitiated()), [dispatch]);

	return  (
		<>
			{/* {(!isLoading) && <Loader/>} */}
			{(recipeData) && <RecipeListingComponent allRecipeData={recipeData}/>}
		</>
	);
};

RecipeListingContainer.propTypes = {};

export default RecipeListingContainer;
