import React, { useEffect } from 'react';
import Plate from 'components/Plate';
import { useSelector, useDispatch } from 'react-redux';
import { getWells } from 'selectors/wellSelectors';
import { setWellSelected as setWellSelectedAction } from 'action-creators/wellActionCreators';
import { getExperimentTargets } from 'selectors/experimentTargetSelector';
import { fetchExperimentTargets } from 'action-creators/experimentTargetActionCreators';
import { getExperimentId } from 'selectors/experimentSelector';

const PlateContainer = () => {
	const dispatch = useDispatch();
	const experimentTargets = useSelector(getExperimentTargets);
	const wellListReducer = useSelector(getWells);
	const experimentId = useSelector(getExperimentId);

	// console.log('experimentId: ', experimentId);
	// console.log('wellListReducer: ', wellListReducer.toJS());
	const experimentTargetsList = experimentTargets.get('list');
	console.log('experimentTargetsList: ', experimentTargetsList.toJS());

	useEffect(() => {
		if (experimentId !== null) {
			dispatch(fetchExperimentTargets(experimentId));
		}
	}, [experimentId, dispatch]);

	const setWellSelected = (index, isWellSelected) => {
		dispatch(setWellSelectedAction(index, isWellSelected));
	};

	return (
		<Plate
			wells={wellListReducer.get('defaultList')}
			setWellSelected={setWellSelected}
			experimentTargetsList={experimentTargetsList}
		/>
	);
};

PlateContainer.propTypes = {};

export default PlateContainer;
