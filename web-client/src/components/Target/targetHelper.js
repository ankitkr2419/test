import { TARGET_CAPACITY } from 'appConstants';
import { createSelector } from 'reselect';

export const getTargetOption = () => {
	const arr = [];
	for (let i = 0; i !== TARGET_CAPACITY; i += 1) {
		arr.push({});
	}
	return arr;
};

export const getSelectedTargetsToLocal = (
	selectedTargets, // Selected targets from server
	listTargetReducer, // list targets from server
	isLoginTypeOperator,
) => {
	// getTargetOption will return array[size=TARGET_CAPACITY] of object with initial isChecked=false
	let arr = getTargetOption(); // Default all option for admin
	// if it's Operator will show all available option to edit target
	if (isLoginTypeOperator === true) {
		arr = [];
	}
	// check before merging if selectedTargets are present
	if (selectedTargets !== null && selectedTargets.size !== 0) {
		selectedTargets.map((selectedTarget, index) => {
			const selectedTargetID = selectedTarget.get('target_id');
			// check selected target is present in master list
			// It will present for sure
			const masterSelectedEl = listTargetReducer
				.get('list')
				.find(ele => ele.get('id') === selectedTargetID);
			// extract name from masterSelectedEl add add to our arr
			arr[index] = {
				// isChecked: true,
				selectedTarget: {
					label: masterSelectedEl.get('name'),
					value: selectedTarget.get('target_id'),
				},
				threshold: selectedTarget.get('threshold'),
			};
			return null;
		});
	}
	return arr;
};

export const isTargetAlreadySelected = (selectedTargetState, selectedTarget) => {
	const list = selectedTargetState.get('targetList');
	if (list === null) {
		return false;
	}
	const foundElement = list.find(
		ele => (ele.selectedTarget && ele.selectedTarget.value === selectedTarget.value),
	);
	return foundElement !== undefined;
};

export const isTargetsModified = (checkedTargets, selectedTargets, isLoginTypeAdmin) => {
	// if selected targets is equal to TARGET_CAPACITY
	if (checkedTargets !== null
		&& selectedTargets !== null
		&& checkedTargets.length === TARGET_CAPACITY) {
		return false;
	}

	// this condition is to verify that nothing is changed
	if (selectedTargets !== null
			&& checkedTargets.length === selectedTargets.size) {
		return false;
	}

	return true;
};

export const getSelectedTargetExperiment = (
	selectedTargets, // Selected targets from server
) => {
	if (selectedTargets === null || selectedTargets.size === 0) {
		return null;
	}
	const arr = [];
	// check before merging if selectedTargets are present
	if (selectedTargets !== null && selectedTargets.size !== 0) {
		selectedTargets.map((selectedTarget, index) => {
			arr[index] = {
				isChecked: true,
				selectedTarget: {
					label: selectedTarget.get('target_name'),
					value: selectedTarget.get('target_id'),
				},
				threshold: selectedTarget.get('threshold'),
				experiment_id: selectedTarget.get('experiment_id'),
				template_id: selectedTarget.get('template_id'),
			};
			return null;
		});
	}
	return arr;
};

export const checkIfIdPresentInList = (id, selectedTargetState) => {
	const list = selectedTargetState.filter(ele => ele.selectedTarget
		&& ele.selectedTarget.value === id);
	return list.size !== 0;
};

// return true if no targets are selected
export const isNoTargetSelected = createSelector(
	targetList => targetList.toJS(),
	(targetList) => {
		const selectedTargetList = targetList.filter(ele => ele.isChecked);
		return selectedTargetList.length === 0;
	},
);
// validate the threshold value
export const validateThreshold = (threshold) => {
	if (threshold <= 10 && threshold >= 0.5) {
		return true;
	}
	return false;
};

// If any threshold value in selected target list is invalid return true
export const isAnyThresholdInvalid = (selectedTargets) => {
	const list = selectedTargets.toJS().filter(
		ele => ele.thresholdError && ele.thresholdError === true,
	);
	return list.length !== 0;
};
