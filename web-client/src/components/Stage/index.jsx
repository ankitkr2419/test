import React, { useState } from 'react';
import { Button, Table } from 'core-components';
import {
	ButtonIcon,
	TableWrapper,
	TableWrapperFooter,
	Icon,
} from 'shared-components';
import AddStageModal from './AddStageModal';

const StageComponent = (props) => {
	const {
		templateID,
		stages,
		addStage,
		onStageRowClicked,
		selectedStageId,
		deleteStage,
		saveStage,
		goToStepWizard,
	} = props;
	// Local state to manage create stage modal
	const [isCreateStageModalVisible, setCreateStageModalVisibility] = useState(
		false,
	);
	// Local state to store stage name/count
	const [stageName, setStageName] = useState('');
	// Local state to store stage type
	const [stageType, setStageType] = useState('');
	// Local state to store stage repeat count
	const [stageRepeatCount, setStageRepeatCount] = useState('');
	// Local state to store selected stage for update
	const [updateStageId, setUpdateStageId] = useState(null);

	// helper method to toggle create template modal
	const toggleCreateStageModal = () => {
		setCreateStageModalVisibility(!isCreateStageModalVisible);
	};

	// Validate create stage form
	const validateStageForm = () => {
		if ((stageName !== '' && stageType !== '', stageRepeatCount !== '')) {
			return true;
		}
		return false;
	};

	const addClickHandler = () => {
		addStage({
			template_id: templateID,
			name: stageName,
			type: stageType.value,
			repeat_count: stageRepeatCount.value,
		});
		toggleCreateStageModal();
		// TODO show error notification
	};

	const saveClickHandler = (stageId) => {
		saveStage(stageId, {
			template_id: templateID,
			name: stageName,
			type: stageType.value,
			repeat_count: stageRepeatCount.value,
		});
		toggleCreateStageModal();
	};

	const editStage = (stage) => {
		const {
			id, name, type, repeat_count,
		} = stage.toJS();
		setUpdateStageId(id);
		setStageName(name);
		setStageType({
			label: type,
			value: type,
		});
		setStageRepeatCount({
			label: repeat_count,
			value: repeat_count,
		});
		toggleCreateStageModal();
	};

	const resetModalState = (stage) => {
		setStageName('');
		setStageType('');
		setStageRepeatCount('');
		setUpdateStageId(null);
	};

	return (
		<div className="d-flex flex-column flex-100">
			<TableWrapper>
				<Table striped>
					<colgroup>
						<col width="17%" />
						<col width="17%" />
						<col width="19%" />
						<col width="17%" />
						<col />
					</colgroup>
					<thead>
						<tr>
							<th>Stage</th>
							<th>Type</th>
							<th>Repeat Count</th>
							<th>Steps</th>
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
					<Button color="primary" icon onClick={toggleCreateStageModal}>
						<Icon size={40} name="plus-2" />
					</Button>
					{isCreateStageModalVisible && (
						<AddStageModal
							toggleCreateStageModal={toggleCreateStageModal}
							isCreateStageModalVisible={isCreateStageModalVisible}
							stageName={stageName}
							stageType={stageType}
							stageRepeatCount={stageRepeatCount}
							setStageName={setStageName}
							setStageType={setStageType}
							setStageRepeatCount={setStageRepeatCount}
							addClickHandler={addClickHandler}
							isFormValid={validateStageForm()}
							resetModalState={resetModalState}
							saveClickHandler={saveClickHandler}
							updateStageId={updateStageId}
						/>
					)}
					{/* <Button color="primary" className="ml-auto">
            Save
					</Button> */}
				</TableWrapperFooter>
			</TableWrapper>
		</div>
	);
};

StageComponent.propTypes = {};

export default StageComponent;
