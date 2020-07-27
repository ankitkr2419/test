import React, { useReducer, useEffect } from 'react';
import PropTypes from 'prop-types';
import stepStateReducer, {
	stepStateInitialState,
} from 'components/Step/stepState';
import AddStepModal from './AddStepModal';
import { stepStateActions } from './stepState';
import { validateStepForm } from './stepHelper';
import HoldSteps from './HoldSteps';
import CycleSteps from './CycleSteps';

const StepComponent = (props) => {
	const {
		currentStageId,
		setCurrentStageId,
		holdStageId,
		cycleStageId,
		holdSteps,	// list of hold steps
		cycleSteps, // list of cycle steps
		addStep, // create api cal
		deleteStep,
		onStepRowClicked,
		selectedStepId,
		saveStep, // update api call
	} = props;

	// local state to save form data and modal state flag
	const [stepFormState, updateStepFormState] = useReducer(
		stepStateReducer,
		stepStateInitialState,
	);

	// immutable => js
	const stepFormStateJS = stepFormState.toJS();
	const { isCreateStepModalVisible } = stepFormStateJS;

	// helper function to update local state
	const updateStepFormStateWrapper = (key, value) => {
		updateStepFormState({
			type: stepStateActions.SET_VALUES,
			key,
			value,
		});
	};

	// resetFormValues will clear out form values
	const resetFormValues = () => {
		updateStepFormState({
			type: stepStateActions.RESET_VALUES,
		});
	};

	// helper method to toggle create template modal
	const toggleCreateStepModal = () => {
		updateStepFormStateWrapper(
			'isCreateStepModalVisible',
			!isCreateStepModalVisible,
		);
	};

	// create step handler
	const addClickHandler = () => {
		const {
			rampRate,
			targetTemperature,
			holdTime,
			dataCapture,
		} = stepFormStateJS;
		addStep({
			stage_id: currentStageId,
			ramp_rate: parseFloat(rampRate),
			target_temp: parseFloat(targetTemperature),
			hold_time: parseInt(holdTime, 10),
			data_capture: dataCapture,
		});
		toggleCreateStepModal();
	};

	// update step handler
	const saveClickHandler = () => {
		const {
			stepId,
			rampRate,
			targetTemperature,
			holdTime,
			dataCapture,
		} = stepFormStateJS;
		saveStep(stepId, {
			stage_id: currentStageId,
			ramp_rate: parseFloat(rampRate),
			target_temp: parseFloat(targetTemperature),
			hold_time: parseInt(holdTime, 10),
			data_capture: dataCapture,
		});
		toggleCreateStepModal();
	};

	const editStep = ({
		id,
		ramp_rate,
		target_temp,
		hold_time,
		data_capture,
	}) => {
		// updating local state with stage details
		// For edit modal view
		updateStepFormState({
			type: stepStateActions.UPDATE_STATE,
			value: {
				stepId: id,
				rampRate: ramp_rate,
				targetTemperature: target_temp,
				holdTime: hold_time.toString(),
				dataCapture: data_capture,
			},
		});
		toggleCreateStepModal();
	};

	const addHoldStep = () => {
		setCurrentStageId(holdStageId);
		toggleCreateStepModal();
	};

	const addCycleStep = () => {
		setCurrentStageId(cycleStageId);
		toggleCreateStepModal();
	};

	// useEffect(() => {
	// 	// make creat modal open if no data is available
	// 	// isStagesLoading will tell us weather api calling is finish or not
	// 	// stages.size = 0  will tell us there is no records present
	// 	// isCreateStageModalVisible is check as we have to make modal visible only once
	// 	if (
	// 		isStepsLoading === false
	// 		&& steps.size === 0
	// 		&& isCreateStepModalVisible === false
	// 	) {
	// 		toggleCreateStepModal();
	// 	}
	// 	// isCreateStepModalVisible skipped in dependency because its causing issue with modal state
	// 	// eslint-disable-next-line
	// }, [isStepsLoading, steps]);

	return (
		<div className='d-flex flex-column flex-100'>
			<HoldSteps
				editStep={editStep}
				holdSteps={holdSteps}
				deleteStep={deleteStep}
				onStepRowClicked={onStepRowClicked}
				selectedStepId={selectedStepId}
				addHoldStep={addHoldStep}
			/>
			<CycleSteps
				editStep={editStep}
				cycleSteps={cycleSteps}
				deleteStep={deleteStep}
				onStepRowClicked={onStepRowClicked}
				selectedStepId={selectedStepId}
				addCycleStep={addCycleStep}
			/>
			{isCreateStepModalVisible && (
				<AddStepModal
					isCreateStepModalVisible={isCreateStepModalVisible}
					toggleCreateStepModal={toggleCreateStepModal}
					updateStepFormStateWrapper={updateStepFormStateWrapper}
					isFormValid={validateStepForm(stepFormStateJS)}
					stepFormState={stepFormStateJS}
					addClickHandler={addClickHandler}
					saveClickHandler={saveClickHandler}
					resetFormValues={resetFormValues}
					stageType={'cycle'}
				/>
			)}
		</div>
	);
};

StepComponent.propTypes = {
	stageId: PropTypes.string.isRequired,
	steps: PropTypes.object.isRequired,
	addStep: PropTypes.func.isRequired,
	deleteStep: PropTypes.func.isRequired,
	onStepRowClicked: PropTypes.func.isRequired,
	selectedStepId: PropTypes.string,
	saveStep: PropTypes.func.isRequired,
	isStepsLoading: PropTypes.bool.isRequired,
	stageType: PropTypes.string.isRequired,
};

export default React.memo(StepComponent);
