import React, { useEffect, useState, useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import StepComponent from 'components/Step';
import {
	addStep as addStepAction,
	deleteStep as deleteStepAction,
	updateStep as updateStepAction,
	addStepReset,
	deleteStepReset,
	updateStepReset,
	fetchHoldSteps,
	fetchCycleSteps,
} from 'action-creators/stepActionCreators';
import { getHoldStepList, getCycleStepList } from 'selectors/stepSelector';
import {
	getHoldStageId,
	getCycleStageId,
	getCycleRepeatCount,
	getCycleStage,
} from 'selectors/stageSelector';
import {
	updateStage as updateStageAction,
	fetchStages,
	updateStageReset,
} from 'action-creators/stageActionCreators';
import { HOLD_STAGE, CYCLE_STAGE } from 'components/Step/stepConstants';


const StepContainer = (props) => {
	const dispatch = useDispatch();
	// local state for storing step id for row selection
	const [selectedStepId, setSelectedStepId] = useState(null);

	// local state for storing stage type ie. hold or cycle stage while creating steps
	const [stageType, setStageType] = useState('');
	// reading holdStageId from redux
	const holdStageId = useSelector(getHoldStageId);
	// reading cycleStageId from redux
	const cycleStageId = useSelector(getCycleStageId);
	// reading hold steps from redux
	const holdSteps = useSelector(getHoldStepList);
	// reading cycle steps from redux
	const cycleSteps = useSelector(getCycleStepList);
	// reading cycle stage from redux
	const cycleStage = useSelector(getCycleStage);
	// reading cycle stage repeat count from redux
	const cycleRepeatCount = useSelector(getCycleRepeatCount);

	// isStepSaved = true means step created successfully
	const { isStepSaved, response } = useSelector(state => state.createStepReducer);
	// isStepDeleted = true means step deleted successfully
	const { isStepDeleted } = useSelector(state => state.deleteStepReducer);
	// isStepUpdated = true means step updated successfully
	const { isStepUpdated } = useSelector(state => state.updateStepReducer);
	// isStepUpdated = true means stage updated successfully
	const { isStageUpdated } = useSelector(state => state.updateStageReducer);

	// fetch the steps for current stage type stored in local state
	const fetchUpdatedSteps = useCallback(() => {
		// if the stage type is hold fetch hold steps
		if (stageType === HOLD_STAGE) {
			dispatch(fetchHoldSteps(holdStageId));
		}
		// if the stage type is cycle fetch cycle steps
		if (stageType === CYCLE_STAGE) {
			dispatch(fetchCycleSteps(cycleStageId));
		}
	}, [holdStageId, cycleStageId, stageType, dispatch]);

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
	const onStepRowClicked = (stepId, stage) => {
		// set stage type of clicked step row
		setStageType(stage);
		// remove step id is if already selected
		if (stepId === selectedStepId) {
			setSelectedStepId(null);
		} else {
			setSelectedStepId(stepId);
		}
	};

	// helper function to save repeat count through api
	const saveRepeatCount = (repeatCount) => {
		// updating stage through api
		dispatch(updateStageAction(cycleStageId, {
			template_id: cycleStage.get('template_id'),
			name: 'Cycle',
			type: cycleStage.get('type'),
			repeat_count: parseInt(repeatCount, 10),
		}));
	};

	// useEffect section
	// useEffect for fetching updated stage list from server after updating repeat count
	useEffect(() => {
		if (isStageUpdated === true) {
			dispatch(updateStageReset());
			dispatch(fetchStages(cycleStage.get('template_id')));
		}
	}, [isStageUpdated, cycleStage, dispatch]);

	useEffect(() => {
		// fetch hold and cycle steps list from server on mount
		dispatch(fetchHoldSteps(holdStageId));
		dispatch(fetchCycleSteps(cycleStageId));
	}, [holdStageId, cycleStageId, dispatch]);

	useEffect(() => {
		// Once we create step will fetch updated step list
		if (isStepSaved === true) {
			// set the newly created step active
			setSelectedStepId(response.id);
			dispatch(addStepReset());
			fetchUpdatedSteps();
		}
	}, [isStepSaved, response, fetchUpdatedSteps, dispatch]);

	useEffect(() => {
		// Once we delete step will fetch updated steps list
		if (isStepDeleted === true) {
			dispatch(deleteStepReset());
			fetchUpdatedSteps();
		}
	}, [isStepDeleted, dispatch, fetchUpdatedSteps]);

	useEffect(() => {
		// Once we update hold step will fetch updated hold step list
		if (isStepUpdated === true) {
			dispatch(updateStepReset());
			fetchUpdatedSteps();
		}
	}, [isStepUpdated, dispatch, fetchUpdatedSteps]);

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
			stageType={stageType}
			setStageType={setStageType}
			saveRepeatCount={saveRepeatCount}
			cycleRepeatCount={cycleRepeatCount}
		/>
	);
};

export default StepContainer;
