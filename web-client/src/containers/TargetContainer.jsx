/* eslint-disable arrow-body-style */
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
	isTargetListUpdatedAdmin,
} from 'components/Target/targetState';
import { getTemplateById } from 'reducers/templateReducer';
import {
	getSelectedTargetsToLocal,
	isTargetAlreadySelected,
	isAnyThresholdInvalid,
} from 'components/Target/targetHelper';

const TargetContainer = (props) => {
	// constants
	const {
		updateSelectedWizard, templateID, isLoginTypeAdmin, isLoginTypeOperator,
		setIsTemplateEdited,
	} = props;
	const dispatch = useDispatch();

	// useSelector section
	// listTargetReducer => master targets from server
	const listTargetReducer = useSelector(state => state.listTargetReducer);
	// listTargetReducer => selected targets from server
	const listTargetByTemplateIDReducer = useSelector(
		state => state.listTargetByTemplateIDReducer,
	);
	const selectedTargets = listTargetByTemplateIDReducer.get('selectedTargets');

	// extracting selected template
	const templates = useSelector(state => state.listTemplatesReducer);
	const selectedTemplateDetails = getTemplateById(templates, templateID);

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
			dispatch(fetchTargetsByTemplateID(templateID));
		}
	}, [dispatch, isTargetSaved, updateSelectedWizard, templateID]);

	useEffect(() => {
		// Update selected targets from server to local state
		// getSelectedTargetsToLocal will return merge list of selected targets
		const value = getSelectedTargetsToLocal(
			selectedTargets,
			listTargetReducer,
			isLoginTypeOperator,
		);
		updateTargetState({
			type: targetStateActions.UPDATE_LIST,
			value,
		});
	}, [selectedTargets, listTargetReducer, isLoginTypeOperator]);

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

	const onThresholdChange = useCallback((selectedThreshold, index) => {
		// save to local state
		updateTargetState({
			type: targetStateActions.ADD_THRESHOLD_VALUE,
			value: {
				threshold: selectedThreshold,
				index,
			},
		});
	}, []);

	// set thresholdError value maintained in local state
	const setThresholdError = useCallback((thresholdError, index) => {
		updateTargetState({
			type: targetStateActions.SET_THRESHOLD_ERROR,
			value: {
				thresholdError,
				index,
			},
		});
	}, []);

	// onSaveClick Save data on server
	const onSaveClick = useCallback(() => {
		// get list of selected targets
		const checkedTargets = getCheckedTargets(
			selectedTargetState.get('targetList'),
			templateID,
		);
		dispatch(saveTarget(templateID, checkedTargets));
	}, [
		selectedTargetState,
		templateID,
		dispatch,
	]);

	const navigateToStepWizard = useCallback(() => {
		return updateSelectedWizard('step');
	}, [updateSelectedWizard]);

	// onEditing a template
	const editTemplate = useCallback(() => {
		// Set the template editing flag on to open the template modal
		setIsTemplateEdited(true);
	}, [setIsTemplateEdited]);

	// this function will check weather our local state is changed or not
	const getIsTargetListUpdatedAdmin = useCallback(() => {
		return isTargetListUpdatedAdmin(selectedTargetState);
	}, [selectedTargetState]);

	// getIsViewStagesEnabled check if we have at least one selected target
	const getIsViewStagesEnabled = useCallback(() => {
		const selectedTargetList = listTargetByTemplateIDReducer.get('selectedTargets');
		if (selectedTargetList.size === 0) {
			return false;
		}
		return true;
	}, [listTargetByTemplateIDReducer]);

	return (
		<TargetComponent
			selectedTemplateDetails={selectedTemplateDetails}
			listTargetReducer={listTargetReducer.get('list')}
			selectedTargetState={selectedTargetState.get('targetList')}
			updateTargetState={updateTargetState}
			onCheckedHandler={onCheckedHandler}
			onTargetSelect={onTargetSelect}
			onThresholdChange={onThresholdChange}
			onSaveClick={onSaveClick}
			isLoginTypeAdmin={isLoginTypeAdmin}
			isLoginTypeOperator={isLoginTypeOperator}
			isTargetListUpdated={getIsTargetListUpdatedAdmin()}
			isViewStagesEnabled={getIsViewStagesEnabled()}
			navigateToStepWizard={navigateToStepWizard}
			editTemplate={editTemplate}
			setThresholdError={setThresholdError}
			isThresholdInvalid={isAnyThresholdInvalid(selectedTargetState.get('targetList'))}
		/>
	);
};

TargetContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	templateID: PropTypes.string.isRequired,
	setIsTemplateEdited: PropTypes.func.isRequired,
};

export default TargetContainer;
