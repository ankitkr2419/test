import React, { useEffect } from 'react';
import Plate from 'components/Plate';
import { useSelector, useDispatch } from 'react-redux';
import { getWells, getWellsPosition } from 'selectors/wellSelectors';
import {
	setSelectedWell as setSelectedWellAction,
	setMultiSelectedWell as setMultiSelectedWellAction,
	toggleMultiSelectOption as toggleMultiSelectOptionAction,
	resetSelectedWells as resetSelectedWellAction,
	fetchWells,
} from 'action-creators/wellActionCreators';
import { getExperimentTargets } from 'selectors/experimentTargetSelector';
import { fetchExperimentTargets } from 'action-creators/experimentTargetActionCreators';
import { getExperimentId, getExperimentTemplate } from 'selectors/experimentSelector';
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
	// running template details
	const experimentTemplate = useSelector(getExperimentTemplate);
	// selected wells positions i.e wells with isSelected/isMultiSelected flag
	const positions = getWellsPosition(wellListReducer);
	// activeWells means the well which are allowed to configure
	const activeWells = useSelector(getActiveLoadedWells);

	// set isPlateRoute true on mount and false on unmount
	useEffect(() => {
		// isPlateRoute use in appHeader to manage visibility of header buttons
		dispatch(setIsPlateRoute(true));
		return () => {
			dispatch(setIsPlateRoute(false));
		};
	}, [dispatch]);

	useEffect(() => {
		if (experimentId !== null) {
			// fetching configured wells data
			dispatch(fetchWells(experimentId));
			// fetching experiment targets to show while configuring sample and graph filter
			dispatch(fetchExperimentTargets(experimentId));
		}
		return () => {
			// isPlateRoute use in appHeader to manage visibility of header buttons
			dispatch(setIsPlateRoute(false));
		};
	}, [experimentId, dispatch]);

	const setSelectedWell = (index, isWellSelected) => {
		dispatch(setSelectedWellAction(index, isWellSelected));
	};

	const resetSelectedWells = () => {
		dispatch(resetSelectedWellAction());
	};

	const setMultiSelectedWell = (index, isWellSelected) => {
		dispatch(setMultiSelectedWellAction(index, isWellSelected));
	};

	const toggleMultiSelectOption = () => {
		// multi selection option for well selection to view it on graph
		dispatch(toggleMultiSelectOptionAction());
	};

	return (
		<Plate
			wells={wellListReducer.get('defaultList')}
			setSelectedWell={setSelectedWell}
			resetSelectedWells={resetSelectedWells}
			experimentTargetsList={experimentTargetsList}
			positions={positions}
			experimentId={experimentId}
			setMultiSelectedWell={setMultiSelectedWell}
			isMultiSelectionOptionOn={wellListReducer.get('isMultiSelectionOptionOn')}
			toggleMultiSelectOption={toggleMultiSelectOption}
			activeWells={activeWells}
			experimentTemplate={experimentTemplate}
		/>
	);
};

export default PlateContainer;
