import React from 'react';
import Plate from 'components/Plate';
import { useSelector, useDispatch } from 'react-redux';
import { getWells } from 'selectors/wellSelectors';
import {
	setWellSelected as  setWellSelectedAction,
} from 'action-creators/wellActionCreators';

const PlateContainer = () => {
	const wellListReducer = useSelector(getWells);
	// console.log('wellListReducer: ', wellListReducer.toJS());
	const dispatch = useDispatch();

	const setWellSelected = (index, isWellSelected) => {
		dispatch(setWellSelectedAction(index, isWellSelected));
	};

	return <Plate wells={wellListReducer.get('defaultList')} setWellSelected={setWellSelected}/>;
};

PlateContainer.propTypes = {};

export default PlateContainer;
