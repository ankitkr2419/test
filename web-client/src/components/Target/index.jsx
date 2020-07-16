import React, { useMemo } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
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
import TargetHeader from './TargetHeader';
import { checkIfIdPresentInList } from './targetHelper';

const TargetActions = styled.div`
  justify-content: space-between;
  padding: 39px 48px 37px 16px;
`;

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
		selectedTemplateDetails,
		isViewStagesEnabled,
		navigateToStageWizard,
		editTemplate,
		isNoTargetSelected,
	} = props;

	const isTargetDisabled = (ele) => {
		if (ele.selectedTarget === undefined || ele.threshold === undefined) {
			return true;
		}
		return false;
	};

	const getFilteredOptionsList = useMemo(
		() => {
			if (isLoginTypeAdmin === true) {
				return listTargetReducer
					.filter(ele => !checkIfIdPresentInList(ele.get('id'), selectedTargetState));
			}
		},
		[listTargetReducer, selectedTargetState, isLoginTypeAdmin],
	);

	const getTargetRows = useMemo(
		() => selectedTargetState.map((ele, index) => (
			<TargetListItem key={index}>
				{isLoginTypeOperator === true && (
					<CheckBox
						onChange={(event) => {
							onCheckedHandler(event, index);
						}}
						className="mr-2"
						id={`target${index}`}
						checked={ele.isChecked}
						disabled={isTargetDisabled(ele)}
					/>
				)}
				{isLoginTypeAdmin === true && (
					// if it's a admin he can select targets from master targets
					<Select
						className="flex-100 px-2"
						options={covertToSelectOption(getFilteredOptionsList, 'name', 'id')}
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
					value={ele.threshold === undefined ? '' : ele.threshold}
					min="0"
					onChange={(event) => {
						onThresholdChange(event.target.value, index);
					}}
				/>
			</TargetListItem>
		)),
		[
			onCheckedHandler,
			onTargetSelect,
			onThresholdChange,
			selectedTargetState,
			isLoginTypeOperator,
			isLoginTypeAdmin,
			getFilteredOptionsList,
		],
	);
	return (
		<div className="d-flex flex-column overflow-hidden flex-100 py-5">
			<TargetHeader
				isLoginTypeAdmin={isLoginTypeAdmin}
				isLoginTypeOperator={isLoginTypeOperator}
				selectedTemplateDetails={selectedTemplateDetails}
				editTemplate={editTemplate}
			/>
			<div className="d-flex overflow-hidden">
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
				<TargetActions className="d-flex flex-column flex-30">
					{isLoginTypeAdmin === true && (
						<Button
							color="primary"
							onClick={navigateToStageWizard}
							className="mx-auto mb-3"
							disabled={isTargetListUpdated === true || isViewStagesEnabled === false}
						>
              View Stages
						</Button>
					)}
					{isLoginTypeOperator === true && (
						<Button
							color="primary"
							onClick={onNextClick}
							className="mx-auto mb-3"
							disabled={isTargetListUpdated}
						>
              Next
						</Button>
					)}
					<Button
						color="primary"
						onClick={onSaveClick}
						className="mx-auto"
						disabled={isTargetListUpdated === false || isNoTargetSelected}
					>
            Save
					</Button>
				</TargetActions>
			</div>
		</div>
	);
};

TargetComponent.propTypes = {
	selectedTargetState: PropTypes.object.isRequired,
	selectedTemplateDetails: PropTypes.object.isRequired,
	onCheckedHandler: PropTypes.func.isRequired,
	onThresholdChange: PropTypes.func.isRequired,
	onSaveClick: PropTypes.func.isRequired,
	isLoginTypeAdmin: PropTypes.bool.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
	isTargetListUpdated: PropTypes.bool.isRequired,
	listTargetReducer: PropTypes.object,
	onTargetSelect: PropTypes.func,
	onNextClick: PropTypes.func,
	isViewStagesEnabled: PropTypes.bool,
	navigateToStageWizard: PropTypes.func,
	editTemplate: PropTypes.func,
};

export default TargetComponent;
