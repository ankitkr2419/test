import wellGraphActions from 'actions/wellGraphActions';

export const wellGraphSucceeded = data => ({
	type: wellGraphActions.successAction,
	payload:{
		data,
	},
});
