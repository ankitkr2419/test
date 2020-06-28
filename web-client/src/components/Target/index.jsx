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
		onNextClick,
		isLoginTypeAdmin,
		isLoginTypeOperator,
		isTargetListUpdated,
	} = props;

	const isTargetDisabled = (ele) => {
		if (ele.selectedTarget === undefined || ele.threshold === undefined) {
			return true;
		}
		return false;
	};

	const getTargetRows = useMemo(
		() => selectedTargetState.map((ele, index) => (
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
			selectedTargetState,
			isLoginTypeOperator,
			isLoginTypeAdmin,
		],
	);
	return (
		<>
			<div className="flex-70 scroll-y p-3">
				<TargetList className="list-target">
					<TargetListHeader>
						<Text className="mb-2 mr-2" />
						<Text className="flex-100 mb-2 px-4">Target</Text>
						<Text className="flex-40 mb-2 px-4">Threshold</Text>
					</TargetListHeader>
					{getTargetRows}
				</TargetList>
			</div>
			<div className="d-flex flex-10 align-items-end p-1">
				<Button
					color="primary"
					onClick={onSaveClick}
					className="mx-auto mb-3"
					disabled={isTargetListUpdated === false}
				>
          Save
				</Button>
			</div>
			{/* {isLoginTypeAdmin === true && (
				<div className="d-flex flex-20 align-items-end p-1">
					<Button
						color="primary"
						onClick={onNextClick}
						className="mx-auto mb-3"
						disabled={isTargetListUpdated}
					>
            View Stages
					</Button>
				</div>
			)} */}
			{isLoginTypeOperator === true && (
				<div className="d-flex flex-10 align-items-end p-1">
					<Button
						color="primary"
						onClick={onNextClick}
						className="mx-auto mb-3"
						disabled={isTargetListUpdated}
					>
            Next
					</Button>
				</div>
			)}
		</>
	);
};

TargetComponent.propTypes = {
	selectedTargetState: PropTypes.object.isRequired,
	onCheckedHandler: PropTypes.func.isRequired,
	onThresholdChange: PropTypes.func.isRequired,
	onSaveClick: PropTypes.func.isRequired,
	isLoginTypeAdmin: PropTypes.bool.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
	isTargetListUpdated: PropTypes.bool.isRequired,
	listTargetReducer: PropTypes.object,
	onTargetSelect: PropTypes.func,
	onNextClick: PropTypes.func,
};

export default TargetComponent;
