import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import RecipeListingComponent from 'components/RecipeListing';
import { recipeListingInitiated } from 'action-creators/recipeActionCreators';
import { Loader } from 'shared-components'

const RecipeListingContainer = (props) => {

	const dispatch = useDispatch();
	const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
	const { isLoading, response } = recipeActionReducer

	useEffect(() => dispatch(recipeListingInitiated()), []);

	console.log("RECIPE ACTION REDUCER: ", recipeActionReducer.response);
	console.log("isLoading: ", isLoading);

	return  (
		<>
			{/* {(!isLoading) && <Loader/>} */}
			{<RecipeListingComponent recipeData={response}/>}
		</>
	);
};

RecipeListingContainer.propTypes = {};

export default RecipeListingContainer;
