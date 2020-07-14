import React, { useReducer, useEffect } from 'react';
import PropTypes from 'prop-types';
import { Table, Button } from 'core-components';
import {
	ButtonIcon,
	TableWrapper,
	TableWrapperBody,
	TableWrapperFooter,
	Icon,
} from 'shared-components';
import stepStateReducer, {
	stepStateInitialState,
} from 'components/Step/stepState';
import AddStepModal from './AddStepModal';
import { stepStateActions } from './stepState';
import { validateStepForm } from './stepHelper';

const StepComponent = (props) => {
	const {
		stageId,
		steps, // list of steps
		addStep, // create api cal
		deleteStep,
		onStepRowClicked,
		selectedStepId,
		saveStep, // update api call
		isStepsLoading,
		goToStageWizard,
		stageType,
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
			stage_id: stageId,
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
			stage_id: stageId,
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

	useEffect(() => {
		// make creat modal open if no data is available
		// isStagesLoading will tell us weather api calling is finish or not
		// stages.size = 0  will tell us there is no records present
		// isCreateStageModalVisible is check as we have to make modal visible only once
		if (
			isStepsLoading === false
			&& steps.size === 0
			&& isCreateStepModalVisible === false
		) {
			toggleCreateStepModal();
		}
		// isCreateStepModalVisible skipped in dependency because its causing issue with modal state
		// eslint-disable-next-line
	}, [isStepsLoading, steps]);

	return (
		<div className='d-flex flex-column flex-100'>
			<TableWrapper>
				<TableWrapperBody>
					<Table striped>
						<colgroup>
							<col width='16%' />
							<col width='12%' />
							<col />
							<col width='16%' />
							<col width='16%' />
							<col width='156px' />
						</colgroup>
						<thead>
							<tr>
								<th>
									Steps <br />
									(Count/Name)
								</th>
								<th>
									Ramp rate <br />
									(unit °C)
								</th>
								<th>
									Target Temperature <br />
									(unit °C)
								</th>
								<th>
									Hold Time <br />
									(unit seconds)
								</th>
								<th>Data Capture</th>
								<th />
							</tr>
						</thead>
						<tbody>
							{steps.map((step, index) => {
								const stepId = step.get('id');
								const classes = selectedStepId === stepId ? 'active' : '';
								return (
									<tr
										className={classes}
										key={stepId}
										onClick={() => {
											onStepRowClicked(stepId);
										}}
									>
										<td>{index + 1}</td>
										<td>{step.get('ramp_rate')}</td>
										<td>{step.get('target_temp')}</td>
										<td>{(step.get('hold_time'))}</td>
										{/* If the stage type is Hold show N/A for data capture property */}
										{stageType !== 'hold' ?
											<td>{step.get('data_capture') ? 'Yes' : 'No'}</td> : <td>N/A</td>}
										<td className='td-actions'>
											<ButtonIcon
												onClick={() => {
													editStep(step.toJS());
												}}
												size={28}
												name='pencil'
											/>
											<ButtonIcon
												onClick={() => {
													deleteStep(stepId);
												}}
												size={28}
												name='trash'
											/>
										</td>
									</tr>
								);
							})}
						</tbody>
					</Table>
				</TableWrapperBody>
				<TableWrapperFooter>
					<ButtonIcon
						name='angle-left'
						size={32}
						className='mr-auto border-0'
						onClick={goToStageWizard}
					/>
					<Button color='primary' icon onClick={toggleCreateStepModal}>
						<Icon size={40} name='plus-2' />
					</Button>
					{isCreateStepModalVisible && (
						<AddStepModal
							isCreateStepModalVisible={isCreateStepModalVisible}
							toggleCreateStepModal={toggleCreateStepModal}
							stepFormState={stepFormStateJS}
							updateStepFormStateWrapper={updateStepFormStateWrapper}
							isFormValid={validateStepForm(stepFormStateJS)}
							addClickHandler={addClickHandler}
							saveClickHandler={saveClickHandler}
							resetFormValues={resetFormValues}
							stageType={stageType}
						/>
					)}
				</TableWrapperFooter>
			</TableWrapper>
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
