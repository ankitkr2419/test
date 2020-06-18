import React, { useReducer } from 'react';
import PropTypes from 'prop-types';
import { Table, Button } from 'core-components';
import {
	ButtonIcon,
	TableWrapper,
	TableWrapperBody,
	TableWrapperFooter,
	Icon,
} from 'shared-components';
import stepStateReducer, { stepStateInitialState } from 'components/Step/stepState';
import { convertStringToSeconds } from 'utils/helpers';
import AddStepModal from './AddStepModal';
import { stepStateActions } from './stepState';

const StepComponent = (props) => {
	const {
		stageId,
		steps, // list of steps
		addStep, // create api cal
		deleteStep,
		onStepRowClicked,
		selectedStepId,
		// saveStep, // update api cal
	} = props;

	// local state to save form data and modal state flag
	const [stepFormState, updateStepFormState] = useReducer(stepStateReducer, stepStateInitialState);

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

	// Validate create step form
	const validateStepForm = ({
		rampRate,
		targetTemperature,
		holdTime,
	}) => {
		if (rampRate !== '' && targetTemperature !== '' && holdTime !== '') {
			return true;
		}
		return false;
	};

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
			// hold_time: new Date(),
			hold_time: convertStringToSeconds(holdTime),
			data_capture: dataCapture,
		});
		toggleCreateStepModal();
	};

	const editStep = ({
		id,
		stage_id,
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
				holdTime: hold_time,
				dataCapture: data_capture,
			},
		});
		toggleCreateStepModal();
	};

	return (
		<div className="d-flex flex-column flex-100">
			<TableWrapper>
				<TableWrapperBody>
					<Table striped>
						<colgroup>
							<col width="16%" />
							<col width="12%" />
							<col />
							<col width="16%" />
							<col width="16%" />
							<col width="156px" />
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
								<th>
                  Data Capture <br />
                  (boolean flag)
								</th>
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
										<td>{step.get('hold_time')}</td>
										<td>{step.get('data_capture').toString()}</td>
										<td className="td-actions">
											<ButtonIcon
												onClick={() => {
													editStep(step.toJS());
												}}
												size={28}
												name="pencil"
											/>
											<ButtonIcon
												onClick={() => {
													deleteStep(stepId);
												}}
												size={28}
												name="trash"
											/>
										</td>
									</tr>
								);
							})}
						</tbody>
					</Table>
				</TableWrapperBody>
				<TableWrapperFooter>
					<Button color="primary" isIcon onClick={toggleCreateStepModal}>
						<Icon size={40} name="plus-2" />
					</Button>
					{isCreateStepModalVisible && (
						<AddStepModal
							isCreateStepModalVisible={isCreateStepModalVisible}
							toggleCreateStepModal={toggleCreateStepModal}
							stepFormState={stepFormStateJS}
							updateStepFormStateWrapper={updateStepFormStateWrapper}
							isFormValid={validateStepForm(stepFormStateJS)}
							addClickHandler={addClickHandler}
							resetFormValues={resetFormValues}
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
};

export default React.memo(StepComponent);
