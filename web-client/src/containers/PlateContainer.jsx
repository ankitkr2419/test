import React, { useEffect } from 'react';
import Plate from 'components/Plate';
import { useSelector, useDispatch } from 'react-redux';
import { getWells, getWellsPosition } from 'selectors/wellSelectors';
import {
	setSelectedWell as setSelectedWellAction,
	setMultiSelectedWell as setMultiSelectedWellAction,
	toggleMultiSelectOption as toggleMultiSelectOptionAction,
	fetchWells,
} from 'action-creators/wellActionCreators';
import { getExperimentTargets } from 'selectors/experimentTargetSelector';
import { fetchExperimentTargets } from 'action-creators/experimentTargetActionCreators';
import { getExperimentId } from 'selectors/experimentSelector';
import { setIsPlateRoute } from 'action-creators/loginActionCreators';

const PlateContainer = () => {
	const dispatch = useDispatch();
	const experimentTargets = useSelector(getExperimentTargets);
	const wellListReducer = useSelector(getWells);
	const experimentId = useSelector(getExperimentId);
	const positions = getWellsPosition(wellListReducer);
	const experimentTargetsList = experimentTargets.get('list');

	useEffect(() => {
		if (experimentId !== null) {
			dispatch(fetchWells(experimentId));
		}
		return () => {
			dispatch(setIsPlateRoute(false));
		};
	}, [experimentId, dispatch]);

	useEffect(() => {
		if (experimentId !== null) {
			dispatch(fetchExperimentTargets(experimentId));
		}
	}, [experimentId, dispatch]);

	const setSelectedWell = (index, isWellSelected) => {
		dispatch(setSelectedWellAction(index, isWellSelected));
	};

	const setMultiSelectedWell = (index, isWellSelected) => {
		dispatch(setMultiSelectedWellAction(index, isWellSelected));
	};

	const toggleMultiSelectOption = () => {
		dispatch(toggleMultiSelectOptionAction());
	};

	return (
		<Plate
			wells={wellListReducer.get('defaultList')}
			setSelectedWell={setSelectedWell}
			experimentTargetsList={experimentTargetsList}
			positions={positions}
			experimentId={experimentId}
			setMultiSelectedWell={setMultiSelectedWell}
			isMultiSelectionOptionOn={wellListReducer.get('isMultiSelectionOptionOn')}
			toggleMultiSelectOption={toggleMultiSelectOption}
		/>
	);
};

export default PlateContainer;
