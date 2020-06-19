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
		onThresholdSelect,
		onSaveClick,
		getCheckedTargetCount,
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

	const getTargetRows = useMemo(() => targetList.map((ele, index) => (
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
			<Select
				className="flex-100 px-2"
				options={covertToSelectOption(listTargetReducer, 'name', 'id')}
				placeholder="Please select target."
				onChange={(selectedTarget) => {
					onTargetSelect(selectedTarget, index);
				}}
				value={ele.selectedTarget}
			/>
			<Input
				className="flex-40 pl-2"
				type="number"
				name="threshold"
				id="threshold"
				placeholder="Type here..."
				value={ele.threshold || ''}
				onChange={(event) => {
					onThresholdSelect(event.target.value, index);
				}}
			/>
		</TargetListItem>
	)), [listTargetReducer, onCheckedHandler, onTargetSelect, onThresholdSelect, targetList]);

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
					disabled={selectedTargetCount === 0}
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
	onThresholdSelect: PropTypes.func.isRequired,
	onSaveClick: PropTypes.func.isRequired,
	getCheckedTargetCount: PropTypes.func.isRequired,
};

export default TargetComponent;
