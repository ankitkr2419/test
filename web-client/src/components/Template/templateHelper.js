import { fromJS } from 'immutable';

export const getTemplateDetails = (response) => fromJS({
	id: response.id,
	name: response.name,
	description: response.description,
});
