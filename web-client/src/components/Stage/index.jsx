import React, { useReducer, useEffect } from 'react';
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
		isStagesLoading,
		goToTargetWizard,
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
	const validateStageForm = ({ stageType, stageRepeatCount }) => {
		if (stageType.value === 'cycle' && stageRepeatCount !== '' && parseInt(stageRepeatCount, 10) > 0) {
			return true;
		}
		if (stageType !== '' && stageType.value === 'hold') {
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
			repeat_count: parseInt(stageRepeatCount, 10),
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
			repeat_count: parseInt(stageRepeatCount, 10),
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
				stageRepeatCount: repeat_count,
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

	useEffect(() => {
		// make creat modal open if no data is available
		// isStagesLoading will tell us weather api calling is finish or not
		// stages.size = 0  will tell us there is no records present
		// isCreateStageModalVisible is check as we have to make modal visible only once
		if (
			isStagesLoading === false
			&& stages.size === 0
			&& isCreateStageModalVisible === false
		) {
			updateStageFormStateWrapper('isCreateStageModalVisible', true);
		}
		// isCreateStageModalVisible skipped in dependency because its causing issue with modal state
		// eslint-disable-next-line
	}, [isStagesLoading, stages]);

	return (
		<div className='d-flex flex-column flex-100'>
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
									<td>{stage.get('type')}</td>
									<td>{stage.get('repeat_count')}</td>
									<td>{stage.get('step_count') || '-'}</td>
									<td className='td-actions'>
										<ButtonIcon
											onClick={() => {
												goToStepWizard(stageId);
											}}
											size={28}
											name='steps'
										/>
										<ButtonIcon
											onClick={() => {
												editStage(stage);
											}}
											size={28}
											name='pencil'
										/>
										<ButtonIcon
											onClick={() => {
												deleteStage(stageId);
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
				<TableWrapperFooter>
					<ButtonIcon
						name='angle-left'
						size={32}
						className='mr-auto border-0'
						onClick={goToTargetWizard}
					/>
					<Button color='primary' icon onClick={toggleCreateStageModal}>
						<Icon size={40} name='plus-2' />
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
	selectedStageId: PropTypes.string,
	deleteStage: PropTypes.func.isRequired,
	saveStage: PropTypes.func.isRequired,
	goToStepWizard: PropTypes.func.isRequired,
	isStagesLoading: PropTypes.bool.isRequired,
};

export default StageComponent;
