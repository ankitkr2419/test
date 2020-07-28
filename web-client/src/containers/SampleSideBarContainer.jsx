import React, { useEffect, useReducer, useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import PropTypes from 'prop-types';
import SidebarSample from 'components/Plate/Sidebar/Sample/SidebarSample';
import { getSamples } from 'selectors/sampleSelectors';
import { getWellsSavedStatus } from 'selectors/wellSelectors';
import {
	fetchSamples as fetchSamplesAction,
	addSampleLocallyCreated,
} from 'action-creators/sampleActionCreators';
import createSampleStateReducer, {
	createSampleInitialState,
	createSampleActions,
	validate,
	getSampleRequestData,
} from 'components/Plate/Sidebar/Sample/createSampleState';
import { addWell, addWellReset } from 'action-creators/wellActionCreators';
import { taskOptions } from 'components/Plate/plateConstant';
import { getSampleTargetList, getInitSampleTargetList } from 'components/Plate/Sidebar/Sample/sampleHelper';

const SampleSideBarContainer = (props) => {
	// constant
	const {
		experimentTargetsList, positions, experimentId, updateWell,
	} = props;
	const dispatch = useDispatch();
	// useSelector
	const samplesListReducer = useSelector(getSamples);
	const {
		list: sampleList,
		isLoading: isSampleListLoading,
	} = samplesListReducer.toJS();
	const areWellsCreated = useSelector(getWellsSavedStatus);

	// useReducer
	const [sampleState, sampleStateDispatch] = useReducer(
		createSampleStateReducer,
		createSampleInitialState,
	);
	const isSampleStateValid = validate(sampleState);

	// helper function to update local state
	const updateCreateSampleWrapper = (key, value) => {
		sampleStateDispatch({
			type: createSampleActions.SET_VALUES,
			key,
			value,
		});
		// console log on sample sidebar opened or closed
		if (key === 'isSideBarOpen') {
			console.info(`Sample Drawer ${value ? 'Opened' : 'Closed'}`);
		}
	};

	// update targets to local state, so every time list will contain original target list
	const addTargetsToLocalState = useCallback(() => {
		if (experimentTargetsList !== null && experimentTargetsList.size !== 0) {
			updateCreateSampleWrapper('targets', getInitSampleTargetList(experimentTargetsList));
		}
	}, [experimentTargetsList]);

	// reset local state
	const resetLocalState = useCallback(() => {
		sampleStateDispatch({ type: createSampleActions.RESET_VALUES });
		// after local state reset update targets to local state, so each new well will contain
		// initially original target list
		addTargetsToLocalState();
	}, [addTargetsToLocalState]);

	useEffect(() => {
		// on page laod, Load target list to local
		addTargetsToLocalState();
	}, [addTargetsToLocalState]);

	useEffect(() => {
		// once well is created reset localState, addWellReducer and restore original target list
		if (areWellsCreated === true) {
			resetLocalState();
			dispatch(addWellReset());
			addTargetsToLocalState();
		}
	}, [areWellsCreated, addTargetsToLocalState, dispatch, resetLocalState]);

	useEffect(() => {
		// this effect will run when operator is trying to update well data
		if (updateWell !== null) {
			const {
				sample_name, sample_id, task, position, targets,
			} = updateWell;
			// set data to local state for update
			sampleStateDispatch({
				type: createSampleActions.UPDATE_STATE,
				value: {
					isEdit: true,
					position,
					isSideBarOpen: true,
					sample: {
						label: sample_name,
						value: sample_id,
					},
					targets: getSampleTargetList(targets, experimentTargetsList),
					task:{
						label: task,
						value: task,
					},
				},
			});
		}
	}, [updateWell, experimentTargetsList]);

	const fetchSamples = (text) => {
		dispatch(fetchSamplesAction(text));
	};

	const addNewLocalSample = (sample) => {
		dispatch(addSampleLocallyCreated(sample));
	};
	// helper function to select or unselect a target stored in targets list in local state
	const onTargetClickHandler = (index) => {
		sampleStateDispatch({
			type: createSampleActions.TOGGLE_TARGET,
			value: index,
		});
	};

	const addButtonClickHandler = () => {
		const requestObject = getSampleRequestData(sampleState, positions.toJS());
		dispatch(addWell(experimentId, requestObject));
	};

	return (
		<SidebarSample
			sampleState={sampleState}
			updateCreateSampleWrapper={updateCreateSampleWrapper}
			experimentTargetsList={experimentTargetsList}
			fetchSamples={fetchSamples}
			addNewLocalSample={addNewLocalSample}
			sampleOptions={sampleList}
			isSampleListLoading={isSampleListLoading}
			taskOptions={taskOptions}
			onTargetClickHandler={onTargetClickHandler}
			addButtonClickHandler={addButtonClickHandler}
			isSampleStateValid={isSampleStateValid}
			resetLocalState={resetLocalState}
			isDisabled={
				positions.size === 0 && sampleState.get('isSideBarOpen') === false
			}
		/>
	);
};

SampleSideBarContainer.propTypes = {
	experimentTargetsList: PropTypes.object.isRequired,
	positions: PropTypes.object.isRequired,
	experimentId: PropTypes.string.isRequired,
	// updated well will contain data of well which is to be updated
	updateWell: PropTypes.object,
};

export default SampleSideBarContainer;
