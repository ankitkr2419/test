import React, { useReducer, useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import {
	stepStateInitialState,
	repeatCounterInitialState,
	repeatCounterStateReducer,
	stepStateReducer,
	repeatCounterStateActions,
} from 'components/Step/stepState';
import AddStepModal from './AddStepModal';
import { stepStateActions } from './stepState';
import { validateStepForm } from './stepHelper';
import HoldSteps from './HoldSteps';
import CycleSteps from './CycleSteps';
import { HOLD_STAGE, CYCLE_STAGE } from './stepConstants';

const StepComponent = (props) => {
	const {
		holdSteps,	// list of hold steps
		cycleSteps, // list of cycle steps
		addStep, // create api cal
		deleteStep,
		onStepRowClicked,
		selectedStepId,
		saveStep, // update api call
		holdStageId,
		cycleStageId,
		stageType,
		setStageType,
		saveRepeatCount, // update repeat count api call
		cycleRepeatCount, // cycle repeat count stored over server
	} = props;

	// local state to save form data and modal state flag
	const [stepFormState, updateStepFormState] = useReducer(
		stepStateReducer,
		stepStateInitialState,
	);

	// local state to save repeat count and repeatCountError flag
	const [repeatCounterState, updateRepeatCounterState] = useReducer(
		repeatCounterStateReducer,
		repeatCounterInitialState,
	);

	// local state to show cycle step form conditionally
	const [showCycleStepForm, setShowCycleStepForm] = useState(false);

	// immutable => js
	const stepFormStateJS = stepFormState.toJS();
	const { isCreateStepModalVisible } = stepFormStateJS;

	// get stage Id for currently opened create step modal
	const getStageId = () => (stageType === HOLD_STAGE ? holdStageId : cycleStageId);

	// helper function to update step form local state
	const updateStepFormStateWrapper = (key, value) => {
		updateStepFormState({
			type: stepStateActions.SET_VALUES,
			key,
			value,
		});
	};

	// helper function to update repeat counter local state
	const updateRepeatCounterStateWrapper = (key, value) => {
		updateRepeatCounterState({
			type: repeatCounterStateActions.SET_VALUES,
			key,
			value,
		});
	};

	// resetFormValues will clear out form values
	const resetFormValues = () => {
		updateStepFormState({
			type: stepStateActions.RESET_VALUES,
		});
		// reset the repeat count stored in local state
		updateRepeatCounterStateWrapper('repeatCount', '');
		// if user closes the modal without saving the repeat count in first use case scenario
		// reset the showCycleStepForm to false
		if (cycleRepeatCount === 0 && showCycleStepForm === true) {
			setShowCycleStepForm(false);
		}
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
			stage_id: getStageId(),
			ramp_rate: parseFloat(rampRate),
			target_temp: parseFloat(targetTemperature),
			hold_time: parseInt(holdTime, 10),
			data_capture: dataCapture,
		});
		// if its cycle stage and repeat count is initial zero then
		// its the first use scenario and we need to save repeat count
		if (stageType === CYCLE_STAGE && cycleRepeatCount === 0) {
			saveRepeatCount(repeatCounterState.get('repeatCount'));
		}
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
			stage_id: getStageId(),
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
		// set stage type as hold
		setStageType(HOLD_STAGE);
		toggleCreateStepModal();
	};

	const addCycleStep = () => {
		// set stage type as cycle
		setStageType(CYCLE_STAGE);
		toggleCreateStepModal();
	};

	useEffect(() => {
		// set the showCycleStepForm true if repeat count is not the initial zero
		if (cycleRepeatCount !== 0) {
			// show cycle step form
			setShowCycleStepForm(true);
			// store the cycle repeat count from server in local state
			updateRepeatCounterStateWrapper('repeatCount', cycleRepeatCount);
		}
	}, [cycleRepeatCount, setShowCycleStepForm]);

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
				cycleRepeatCount={cycleRepeatCount}
				repeatCounterState={repeatCounterState.toJS()}
				updateRepeatCounterStateWrapper={updateRepeatCounterStateWrapper}
				saveRepeatCount={saveRepeatCount}
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
					stageType={stageType}
					saveRepeatCount={saveRepeatCount}
					repeatCounterState={repeatCounterState.toJS()}
					updateRepeatCounterStateWrapper={updateRepeatCounterStateWrapper}
					showCycleStepForm={showCycleStepForm}
					setShowCycleStepForm={setShowCycleStepForm}
					cycleRepeatCount={cycleRepeatCount}
				/>
			)}
		</div>
	);
};

StepComponent.propTypes = {
	holdStageId: PropTypes.string.isRequired,
	cycleStageId: PropTypes.string.isRequired,
	stageType: PropTypes.string.isRequired,
	holdSteps: PropTypes.object.isRequired,
	cycleSteps: PropTypes.object.isRequired,
	addStep: PropTypes.func.isRequired,
	deleteStep: PropTypes.func.isRequired,
	onStepRowClicked: PropTypes.func.isRequired,
	setStageType: PropTypes.func.isRequired,
	selectedStepId: PropTypes.string,
	saveStep: PropTypes.func.isRequired,
	saveRepeatCount: PropTypes.func.isRequired,
	cycleRepeatCount: PropTypes.number.isRequired,
};

export default React.memo(StepComponent);
