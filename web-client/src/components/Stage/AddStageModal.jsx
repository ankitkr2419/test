import React, { useEffect } from 'react';
import {
	Button,
	Form,
	FormGroup,
	Row,
	Col,
	Input,
	Label,
	Modal,
	ModalBody,
	Select,
} from 'core-components';
import { ButtonGroup, ButtonIcon, Text } from 'shared-components';
import { stageTypeOptions, countTypeOptions } from './stageConstants';

const AddStageModal = (props) => {
	const {
		toggleCreateStageModal,
		isCreateStageModalVisible,
		stageName,
		stageType,
		stageRepeatCount,
		setStageName,
		setStageType,
		setStageRepeatCount,
		addClickHandler,
		isFormValid,
		resetModalState,
		saveClickHandler,
		updateStageId,
	} = props;

	const isUpdateForm = updateStageId !== null;

	useEffect(
		() => () => {
			// clear modal state on un-mount
			resetModalState();
		},
		[resetModalState],
	);

	return (
		<>
			<Modal
				isOpen={isCreateStageModalVisible}
				toggle={toggleCreateStageModal}
				centered
				size="lg"
			>
				<ModalBody>
					<Text
						tag="h4"
						className="modal-title text-center text-truncate font-weight-bold"
					>
            Add Stage
					</Text>
					<ButtonIcon
						position="absolute"
						placement="right"
						top={24}
						right={32}
						size={32}
						name="cross"
						onClick={toggleCreateStageModal}
					/>
					<Form>
						<Row form className="mb-5 pb-5">
							<Col sm={4}>
								<FormGroup>
									<Label for="stage" className="font-weight-bold">
                    Stage
									</Label>
									<Input
										type="text"
										name="stage"
										id="stage"
										placeholder="Type here"
										value={stageName}
										onChange={(event) => {
											setStageName(event.target.value);
										}}
									/>
								</FormGroup>
							</Col>
							<Col sm={4}>
								<FormGroup>
									<Label for="stageType" className="font-weight-bold">
                    Stage type
									</Label>
									<Select
										options={stageTypeOptions}
										onChange={(selectedStageType) => {
											setStageType(selectedStageType);
										}}
										value={stageType}
									/>
								</FormGroup>
							</Col>
							<Col sm={3}>
								<FormGroup>
									<Label for="repeatCount" className="font-weight-bold">
                    Repeat Count
									</Label>
									<Select
										options={countTypeOptions}
										onChange={(selectedRepeatCount) => {
											setStageRepeatCount(selectedRepeatCount);
										}}
										value={stageRepeatCount}
									/>
								</FormGroup>
							</Col>
						</Row>
						<ButtonGroup className="text-center p-0 m-0 pt-5">
							{isUpdateForm === false && (
								<Button
									color="primary"
									onClick={addClickHandler}
									disabled={isFormValid === false}
								>
                  Add
								</Button>
							)}
							{isUpdateForm === true && (
								<Button
									color="primary"
									onClick={() => saveClickHandler(updateStageId)}
									disabled={isFormValid === false}
								>
                  Save
								</Button>
							)}
						</ButtonGroup>
					</Form>
				</ModalBody>
			</Modal>
		</>
	);
};

AddStageModal.propTypes = {};

export default AddStageModal;
