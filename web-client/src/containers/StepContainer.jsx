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
} from 'action-creators/stepActionCreators';
import { getStepList } from 'selectors/stepSelector';

const StepContainer = (props) => {
	const { stageId, updateSelectedWizard } = props;
	const dispatch = useDispatch();
	// local state for storing step id for row selection
	const [selectedStepId, setSelectedStepId] = useState(null);

	// reading steps from redux
	const steps = useSelector(getStepList);

	// isStepSaved = true means step created successfully
	const { isStepSaved } = useSelector(state => state.createStepReducer);
	// isStepDeleted = true means step deleted successfully
	const { isStepDeleted } = useSelector(state => state.deleteStepReducer);
	// isStepUpdated = true means step updated successfully
	const { isStepUpdated } = useSelector(state => state.updateStepReducer);


	useEffect(() => {
		// Once we create step will fetch updated step list
		if (isStepSaved === true) {
			dispatch(addStepReset());
			dispatch(fetchSteps(stageId));
		}
	}, [isStepSaved, stageId, dispatch]);

	useEffect(() => {
		// Once we delete step will fetch updated step list
		if (isStepDeleted === true) {
			dispatch(deleteStepReset());
			dispatch(fetchSteps(stageId));
		}
	}, [isStepDeleted, stageId, dispatch]);

	useEffect(() => {
		// Once we update step will fetch updated step list
		if (isStepUpdated === true) {
			dispatch(updateStepReset());
			dispatch(fetchSteps(stageId));
		}
	}, [isStepUpdated, stageId, dispatch]);

	useEffect(() => {
		// fetch updated step list from server
		dispatch(fetchSteps(stageId));
	}, [stageId, dispatch]);

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

	const goToStageWizard = () => {
		updateSelectedWizard('stage');
	};

	return (
		<StepComponent
			stageId={stageId}
			steps={steps.get('list')}
			isStepsLoading={steps.get('isLoading')}
			addStep={addStep}
			deleteStep={deleteStep}
			onStepRowClicked={onStepRowClicked}
			selectedStepId={selectedStepId}
			saveStep={saveStep}
			goToStageWizard={goToStageWizard}
		/>
	);
};

StepContainer.propTypes = {
	stageId: PropTypes.string.isRequired,
};

export default StepContainer;
