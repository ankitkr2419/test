import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import StepComponent from 'components/Step';
import {
	addStep as addStepAction,
	deleteStep as deleteStepAction,
	updateStep as updateStepAction,
	fetchSteps,
	addStepReset,
	deleteStepReset,
	updateStepReset,
	fetchHoldSteps,
	fetchCycleSteps,
} from 'action-creators/stepActionCreators';
import { getHoldStepList, getCycleStepList } from 'selectors/stepSelector';

const StepContainer = (props) => {
	const { holdStageId, cycleStageId  } = props;
	const dispatch = useDispatch();
	// local state for storing step id for row selection
	const [selectedStepId, setSelectedStepId] = useState(null);
	// local state for storing stage id for creating steps in respective stages
	const [currentStageId, setCurrentStageId] = useState(null);

	// reading hold steps from redux
	const holdSteps = useSelector(getHoldStepList);
	// reading cycle steps from redux
	const cycleSteps = useSelector(getCycleStepList);

	// stageType = hold or cycle stage
	// const stageType = useSelector(state => getStageType(state, stageId));
	// isStepSaved = true means step created successfully
	const { isStepSaved, response } = useSelector(state => state.createStepReducer);
	// isStepDeleted = true means step deleted successfully
	const { isStepDeleted } = useSelector(state => state.deleteStepReducer);
	// isStepUpdated = true means step updated successfully
	const { isStepUpdated } = useSelector(state => state.updateStepReducer);

	useEffect(() => {
		// Once we create step will fetch updated step list
		if (isStepSaved === true) {
			// set the newly created step active
			setSelectedStepId(response.id);
			dispatch(addStepReset());
			// No need to fetch again as we have already added the created step to stepslist in reducer
		}
	}, [isStepSaved, dispatch, response]);

	useEffect(() => {
		// Once we delete hold step will fetch updated hold step list
		if (isStepDeleted === true) {
			dispatch(deleteStepReset());
			dispatch(fetchHoldSteps(holdStageId));
		}
	}, [isStepDeleted, holdStageId, dispatch]);

	useEffect(() => {
		// Once we delete cycle step will fetch updated cycle step list
		if (isStepDeleted === true) {
			dispatch(deleteStepReset());
			dispatch(fetchCycleSteps(cycleStageId));
		}
	}, [isStepDeleted, cycleStageId, dispatch]);

	useEffect(() => {
		// Once we update hold step will fetch updated hold step list
		if (isStepUpdated === true) {
			dispatch(updateStepReset());
			dispatch(fetchHoldSteps(holdStageId));
		}
	}, [isStepUpdated, holdStageId, dispatch]);

	useEffect(() => {
		// Once we update cycle step will fetch updated cycle step list
		if (isStepUpdated === true) {
			dispatch(updateStepReset());
			dispatch(fetchCycleSteps(cycleStageId));
		}
	}, [isStepUpdated, cycleStageId, dispatch]);

	useEffect(() => {
		// fetch updated hold and cycle step list from server
		dispatch(fetchHoldSteps(holdStageId));
		dispatch(fetchCycleSteps(cycleStageId));
	}, [holdStageId, cycleStageId, dispatch]);

	const addStep = (step) => {
		// creating step though api
		dispatch(addStepAction(step));
	};

	const deleteStep = (stepId) => {
		// deleting step though api
		dispatch(deleteStepAction(stepId));
	};

	const saveStep = (stepId, step) => {
		dispatch(updateStepAction(stepId, step));
	};

	// Here will update selected step id
	const onStepRowClicked = (stepId) => {
		// remove step id is if already selected
		if (stepId === selectedStepId) {
			setSelectedStepId(null);
		} else {
			setSelectedStepId(stepId);
		}
	};

	return (
		<StepComponent
			holdSteps={holdSteps.get('list')}
			cycleSteps={cycleSteps.get('list')}
			addStep={addStep}
			deleteStep={deleteStep}
			onStepRowClicked={onStepRowClicked}
			selectedStepId={selectedStepId}
			saveStep={saveStep}
			holdStageId={holdStageId}
			cycleStageId={cycleStageId}
			currentStageId={currentStageId}
			setCurrentStageId={setCurrentStageId}
		/>
	);
};

StepContainer.propTypes = {
	holdStageId: PropTypes.string.isRequired,
	cycleStageId: PropTypes.string.isRequired,
};

export default StepContainer;
