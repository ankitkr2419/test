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
import { getActiveLoadedWells } from 'selectors/activeWellSelector';

const PlateContainer = () => {
	const dispatch = useDispatch();
	// experiment targets
	const experimentTargets = useSelector(getExperimentTargets);
	const experimentTargetsList = experimentTargets.get('list');
	// get wells data from server
	const wellListReducer = useSelector(getWells);
	// running experiment id
	const experimentId = useSelector(getExperimentId);
	// selected wells positions
	const positions = getWellsPosition(wellListReducer);
	// activeWells means the well which are allowed to configure
	const activeWells = useSelector(getActiveLoadedWells);

	useEffect(() => {
		if (experimentId !== null) {
			dispatch(fetchWells(experimentId));
			dispatch(fetchExperimentTargets(experimentId));
		}
		return () => {
			dispatch(setIsPlateRoute(false));
		};
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
			activeWells={activeWells}
		/>
	);
};

export default PlateContainer;
