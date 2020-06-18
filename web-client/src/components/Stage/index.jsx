import React, { useReducer } from 'react';
import PropTypes from 'prop-types';
import { Button, Table } from 'core-components';
import {
	ButtonIcon,
	TableWrapper,
	TableWrapperFooter,
	Icon,
} from 'shared-components';
import AddStageModal from './AddStageModal';
import stageStateReducer, {
	stageStateInitialState,
	stageStateActions,
} from './stageState';
import { stageTableHeader } from './stageConstants';

const StageComponent = (props) => {
	const {
		templateID,
		stages, // list of stages
		addStage,
		onStageRowClicked,
		selectedStageId,
		deleteStage,
		saveStage,
		goToStepWizard,
	} = props;

	// local state to save form data and modal state flag
	const [stageFormState, updateStageFormState] = useReducer(
		stageStateReducer,
		stageStateInitialState,
	);
	// immutable => js
	const stageFormStateJS = stageFormState.toJS();
	const { isCreateStageModalVisible } = stageFormStateJS;

	// helper function to update local state
	const updateStageFormStateWrapper = (key, value) => {
		updateStageFormState({
			type: stageStateActions.SET_STAGE_VALUES,
			key,
			value,
		});
	};

	// helper method to toggle create template modal
	const toggleCreateStageModal = () => {
		updateStageFormStateWrapper(
			'isCreateStageModalVisible',
			!isCreateStageModalVisible,
		);
	};

	// Validate create stage form
	const validateStageForm = ({ stageName, stageType, stageRepeatCount }) => {
		if ((stageName !== '' && stageType !== '', stageRepeatCount !== '')) {
			return true;
		}
		return false;
	};

	// create stage handler
	const addClickHandler = () => {
		const { stageName, stageType, stageRepeatCount } = stageFormStateJS;
		addStage({
			template_id: templateID,
			name: stageName,
			type: stageType.value,
			repeat_count: stageRepeatCount.value,
		});
		toggleCreateStageModal();
		// TODO show error notification
	};

	// update stage handler
	const saveClickHandler = () => {
		const {
			stageId,
			stageName,
			stageType,
			stageRepeatCount,
		} = stageFormStateJS;
		saveStage(stageId, {
			template_id: templateID,
			name: stageName,
			type: stageType.value,
			repeat_count: stageRepeatCount.value,
		});
		toggleCreateStageModal();
	};

	// edit stage handler
	const editStage = (stage) => {
		const {
			id, name, type, repeat_count,
		} = stage.toJS();
		updateStageFormState({
			type: stageStateActions.UPDATE_STAGE_STATE,
			value: {
				stageId: id,
				stageName: name,
				stageType: { label: type, value: type },
				stageRepeatCount: { label: repeat_count, value: repeat_count },
			},
		});
		toggleCreateStageModal();
	};

	// resetModalState will clear out form values
	const resetModalState = () => {
		updateStageFormState({
			type: stageStateActions.RESET_STAGE_VALUES,
		});
	};

	return (
		<div className="d-flex flex-column flex-100">
			<TableWrapper>
				<Table striped>
					<colgroup>
						{stageTableHeader.map(ele => (
							<col key={ele.name} width={ele.width} />
						))}
						<col />
					</colgroup>

					<thead>
						<tr>
							{stageTableHeader.map(ele => (
								<th key={ele.name}>{ele.name}</th>
							))}
							<th />
						</tr>
					</thead>
					<tbody>
						{stages.map((stage, i) => {
							const stageId = stage.get('id');
							const classes = selectedStageId === stageId ? 'active' : '';
							return (
								<tr
									className={classes}
									onClick={() => {
										onStageRowClicked(stageId);
									}}
									key={stageId}
								>
									<td>{i + 1}</td>
									<td>{stage.get('name')}</td>
									<td>{stage.get('type')}</td>
									<td>{stage.get('repeat_count')}</td>
									<td className="td-actions">
										<ButtonIcon
											onClick={() => {
												goToStepWizard(stageId);
											}}
											size={28}
											name="steps"
										/>
										<ButtonIcon
											onClick={() => {
												editStage(stage);
											}}
											size={28}
											name="pencil"
										/>
										<ButtonIcon
											onClick={() => {
												deleteStage(stageId);
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
				<TableWrapperFooter>
					<Button color="primary" isIcon onClick={toggleCreateStageModal}>
						<Icon size={40} name="plus-2" />
					</Button>
					{isCreateStageModalVisible && (
						<AddStageModal
							toggleCreateStageModal={toggleCreateStageModal}
							isCreateStageModalVisible={isCreateStageModalVisible}
							stageFormStateJS={stageFormStateJS}
							updateStageFormStateWrapper={updateStageFormStateWrapper}
							addClickHandler={addClickHandler}
							isFormValid={validateStageForm(stageFormStateJS)}
							resetModalState={resetModalState}
							saveClickHandler={saveClickHandler}
						/>
					)}
				</TableWrapperFooter>
			</TableWrapper>
		</div>
	);
};

StageComponent.propTypes = {
	templateID: PropTypes.string.isRequired,
	stages: PropTypes.object.isRequired,
	addStage: PropTypes.func.isRequired,
	onStageRowClicked: PropTypes.func.isRequired,
	selectedStageId: PropTypes.string.isRequired,
	deleteStage: PropTypes.func.isRequired,
	saveStage: PropTypes.func.isRequired,
	goToStepWizard: PropTypes.func.isRequired,
};

export default StageComponent;
