import { TARGET_CAPACITY } from '../../constants';

export const getTargetOption = () => {
	const arr = [];
	for (let i = 0; i !== TARGET_CAPACITY; i += 1) {
		arr.push({
			isChecked : false,
		});
	}
	return arr;
};

export const covertToSelectOption = (list) => {
	const arr = [];
	list.map(ele => arr.push({ label: ele.get('name'), value: ele.get('id') }));
	return arr;
};

export const getSelectedTargetsToLocal = (
	selectedTargets,
	listTargetReducer,
) => {
	const arr = getTargetOption();
	if (selectedTargets !== null && selectedTargets.size !== 0) {
		selectedTargets.map((selectedTarget, index) => {
			const selectedTargetID = selectedTarget.get('target_id');
			const masterSelectedEl = listTargetReducer
				.get('list')
				.find(ele => ele.get('id') === selectedTargetID);
			arr[index] = {
				isChecked: true,
				selectedTarget: {
					label: masterSelectedEl.get('name'),
					value: selectedTarget.get('target_id'),
				},
				threshold: selectedTarget.get('threshold'),
			};
			return null;
		});
	}
	// console.table(arr);
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
