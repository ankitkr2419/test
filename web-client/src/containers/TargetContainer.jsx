import React, { useEffect, useReducer, useCallback } from 'react';
import PropTypes from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import TargetComponent from 'components/Target';
import {
	fetchMasterTargets,
	fetchTargetsByTemplateID,
	saveTarget,
	resetSaveTarget,
} from 'action-creators/targetActionCreators';
import targetStateReducer, {
	targetInitialState,
	targetStateActions,
	isCheckable,
	getCheckedTargets,
} from 'components/Target/targetState';
import {
	getSelectedTargetsToLocal,
	isTargetAlreadySelected,
} from 'components/Target/targetHelper';
import { TARGET_CAPACITY } from '../constants';

const TargetContainer = (props) => {
	// constants
	const { updateSelectedWizard, templateID } = props;
	const dispatch = useDispatch();

	// useSelector section
	// listTargetReducer => master targets from server
	const listTargetReducer = useSelector(state => state.listTargetReducer);
	// listTargetReducer => selected targets from server
	const listTargetByTemplateIDReducer = useSelector(
		state => state.listTargetByTemplateIDReducer,
	);
	const selectedTargets = listTargetByTemplateIDReducer.get('selectedTargets');
	// isTargetSaved flag will get update when targets are saved successfully over server
	const { isTargetSaved } = useSelector(state => state.saveTargetReducer);

	// useReducer section
	// local state to manage selected target data
	const [selectedTargetState, updateTargetState] = useReducer(
		targetStateReducer,
		targetInitialState,
	);

	// useEffect section
	// below useEffect is use to navigate to next wizard when user will save targets
	useEffect(() => {
		if (isTargetSaved === true) {
			// isTargetSaved = true means targets saved successfully
			// reset save target reducer to avoid multiple re-renders
			dispatch(resetSaveTarget());
			// navigate to next wizard
			updateSelectedWizard('stage');
		}
	}, [dispatch, isTargetSaved, updateSelectedWizard]);

	useEffect(() => {
		// Update selected targets from server to local state
		// getSelectedTargetsToLocal will return merge list of selected targets
		updateTargetState({
			type: targetStateActions.UPDATE_LIST,
			value: getSelectedTargetsToLocal(selectedTargets, listTargetReducer),
		});
	}, [selectedTargets, listTargetReducer]);

	useEffect(() => {
		// get master targets from server
		dispatch(fetchMasterTargets());
	}, [dispatch]);

	useEffect(() => {
		// get selected targets from server
		dispatch(fetchTargetsByTemplateID(templateID));
	}, [dispatch, templateID]);

	// handler function section
	// checkbox handler for target list
	const onCheckedHandler = useCallback(
		(event, index) => {
			// isCheckable will check weather target and threshold values are present
			if (isCheckable(selectedTargetState, index)) {
				// save to local state
				updateTargetState({
					type: targetStateActions.SET_CHECKED_STATE,
					value: {
						checked: event.target.checked,
						index,
					},
				});
			}
		},
		[selectedTargetState],
	);

	const onTargetSelect = useCallback(
		(selectedTarget, index) => {
			if (!isTargetAlreadySelected(selectedTargetState, selectedTarget)) {
				// save to local state
				updateTargetState({
					type: targetStateActions.ADD_TARGET_ID,
					value: {
						targetId: selectedTarget,
						index,
					},
				});
			} else {
				// TODO set form error once design is ready
				console.error('already selected');
			}
		},
		[selectedTargetState],
	);

	const onThresholdSelect = useCallback((selectedThreshold, index) => {
		// save to local state
		updateTargetState({
			type: targetStateActions.ADD_THRESHOLD_VALUE,
			value: {
				threshold: selectedThreshold,
				index,
			},
		});
	}, []);

	// getCheckedTargetCount will return the count of selected checkbox
	const getCheckedTargetCount = useCallback(
		() => getCheckedTargets(selectedTargetState.get('targetList')).length,
		[selectedTargetState],
	);

	// onSaveClick Save data on server
	const onSaveClick = useCallback(() => {
		// get list of selected targets
		const checkedTargets = getCheckedTargets(
			selectedTargetState.get('targetList'),
			templateID,
		);
		// if Capacity exceeds for target selection will redirect to stage wizard
		if (
			(checkedTargets !== null
        && selectedTargets !== null
        && checkedTargets.length === TARGET_CAPACITY)
      || (selectedTargets !== null
        && checkedTargets.length === selectedTargets.size)
		) {
			updateSelectedWizard('stage');
		} else {
			// save call to server
			dispatch(saveTarget(templateID, checkedTargets));
		}
	}, [
		selectedTargetState,
		templateID,
		selectedTargets,
		updateSelectedWizard,
		dispatch,
	]);

	return (
		<TargetComponent
			listTargetReducer={listTargetReducer.get('list')}
			selectedTargetState={selectedTargetState}
			updateTargetState={updateTargetState}
			onCheckedHandler={onCheckedHandler}
			onTargetSelect={onTargetSelect}
			onThresholdSelect={onThresholdSelect}
			onSaveClick={onSaveClick}
			getCheckedTargetCount={getCheckedTargetCount}
		/>
	);
};

TargetContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	templateID: PropTypes.string.isRequired,
};

export default TargetContainer;
