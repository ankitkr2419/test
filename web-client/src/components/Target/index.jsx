import React, { useMemo } from 'react';
import PropTypes from 'prop-types';
import {
	Button, CheckBox, Select, Input,
} from 'core-components';
import {
	TargetList,
	TargetListHeader,
	TargetListItem,
	Text,
} from 'shared-components';

import { covertToSelectOption } from 'utils/helpers';

const TargetComponent = (props) => {
	const {
		// master target list
		listTargetReducer,
		// local state target list used for manipulation
		selectedTargetState,
		onCheckedHandler,
		onTargetSelect,
		onThresholdChange,
		onSaveClick,
		getCheckedTargetCount,
		isLoginTypeAdmin,
		isLoginTypeOperator,
	} = props;

	// Extracting targetList from immutable target state
	const targetList = selectedTargetState.get('targetList');
	// selected target count will help us in validating save button
	const selectedTargetCount = getCheckedTargetCount();

	const isTargetDisabled = (ele) => {
		if (ele.selectedTarget === undefined || ele.threshold === undefined) {
			return true;
		}
		return false;
	};

	const getTargetRows = useMemo(
		() => targetList.map((ele, index) => (
			<TargetListItem key={index}>
				<CheckBox
					onChange={(event) => {
						onCheckedHandler(event, index);
					}}
					className="mr-2"
					id={`target${index}`}
					checked={ele.isChecked}
					disabled={isTargetDisabled(ele)}
				/>
				{isLoginTypeAdmin === true && (
					// if it's a admin he can select targets from master targets
					<Select
						className="flex-100 px-2"
						options={covertToSelectOption(listTargetReducer, 'name', 'id')}
						placeholder="Please select target."
						onChange={(selectedTarget) => {
							onTargetSelect(selectedTarget, index);
						}}
						value={ele.selectedTarget}
					/>
				)}
				{isLoginTypeOperator === true && (
					// for operator will show in disabled state
					<Input
						className="flex-100 mr-2"
						type="text"
						placeholder="Type here..."
						defaultValue={ele.selectedTarget && ele.selectedTarget.label}
						disabled
					/>
				)}
				<Input
					className="flex-40 pl-2"
					type="number"
					name={`threshold${index}`}
					index={`threshold${index}`}
					placeholder="Type here..."
					value={ele.threshold || ''}
					min="0"
					onChange={(event) => {
						onThresholdChange(event.target.value, index);
					}}
				/>
			</TargetListItem>
		)),
		[
			listTargetReducer,
			onCheckedHandler,
			onTargetSelect,
			onThresholdChange,
			targetList,
			isLoginTypeOperator,
			isLoginTypeAdmin,
		],
	);
	return (
		<>
			<div className="flex-100 scroll-y p-4">
				<TargetList className="list-target">
					<TargetListHeader>
						<Text className="mb-2 mr-2" />
						<Text className="flex-100 mb-2 px-4">Target</Text>
						<Text className="flex-40 mb-2 px-4">Threshold</Text>
					</TargetListHeader>
					{getTargetRows}
				</TargetList>
			</div>
			<div className="d-flex flex-30 align-items-end p-4">
				<Button
					color="primary"
					onClick={onSaveClick}
					className="mx-auto mb-3"
					// TODO Enabled it once operator new design is in place
					disabled={selectedTargetCount === 0 || isLoginTypeOperator === true}
				>
          Save
				</Button>
			</div>
		</>
	);
};

TargetComponent.propTypes = {
	listTargetReducer: PropTypes.object.isRequired,
	selectedTargetState: PropTypes.object.isRequired,
	onCheckedHandler: PropTypes.func.isRequired,
	onTargetSelect: PropTypes.func.isRequired,
	onThresholdChange: PropTypes.func.isRequired,
	onSaveClick: PropTypes.func.isRequired,
	getCheckedTargetCount: PropTypes.func.isRequired,
	isLoginTypeAdmin: PropTypes.bool.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
};

export default TargetComponent;
