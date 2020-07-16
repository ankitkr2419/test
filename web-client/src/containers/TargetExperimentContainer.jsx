import React, {
	useReducer, useEffect, useCallback, useState,
} from 'react';
import { fromJS } from 'immutable';
import PropTypes from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import TargetComponent from 'components/Target';
import targetStateReducer, {
	getCheckedExperimentTargets,
	targetStateActions,
	isTargetListUpdated,
} from 'components/Target/targetState';
import { getTemplateById } from 'reducers/templateReducer';
import {
	fetchExperimentTargets,
	createExperimentTarget,
	createExperimentTargetReset,
} from 'action-creators/experimentTargetActionCreators';
import { getExperimentTargets } from 'selectors/experimentTargetSelector';
import { getSelectedTargetExperiment, isNoTargetSelected } from 'components/Target/targetHelper';
import { Redirect } from 'react-router';
import { getExperimentId } from 'selectors/experimentSelector';
import { setIsPlateRoute } from 'action-creators/loginActionCreators';

const TargetExperimentContainer = (props) => {
	// constants
	const { isLoginTypeAdmin, isLoginTypeOperator, templateID } = props;
	const dispatch = useDispatch();
	// useSelector section
	// extracting experiment id
	const experimentId = useSelector(getExperimentId);
	// list of experiment targets
	const listExperimentTargetsReducer = useSelector(getExperimentTargets);
	const experimentTargets = listExperimentTargetsReducer.get('list');

	// extracting selected template
	const templates = useSelector(state => state.listTemplatesReducer);
	const selectedTemplateDetails = getTemplateById(templates, templateID);

	// extracting selected targets
	const { isExperimentTargetSaved } = useSelector(state => state.createExperimentTargetReducer);

	// useReducer section
	// local state to manage selected target data
	const [selectedTargetState, updateTargetState] = useReducer(
		targetStateReducer,
		fromJS({ targetList: [], originalTargetList: [] }),
	);
	const [isRedirectToPlate, setRedirectToPlate] = useState(false);

	useEffect(() => {
		// fetching list of experiment targets
		dispatch(fetchExperimentTargets(experimentId));
	}, [dispatch, experimentId]);

	useEffect(() => {
		// converting list of experiment targets to local state
		const value = getSelectedTargetExperiment(experimentTargets);
		if (value !== null) {
			// update local state list
			updateTargetState({
				type: targetStateActions.UPDATE_LIST,
				value,
			});
		}
	}, [experimentTargets]);

	useEffect(() => {
		if (isExperimentTargetSaved === true) {
			// fetching list of experiment targets
			// dispatch(fetchExperimentTargets(experimentId));
			dispatch(createExperimentTargetReset());
			updateTargetState({
				type: targetStateActions.UPDATE_LIST,
				value: selectedTargetState.get('targetList'),
			});
		}
	}, [isExperimentTargetSaved, dispatch, selectedTargetState]);

	// handler function section
	// checkbox handler for target list
	const onTargetCheckedHandler = useCallback((event, index) => {
		// save to local state
		updateTargetState({
			type: targetStateActions.SET_CHECKED_STATE,
			value: {
				checked: event.target.checked,
				index,
			},
		});
	}, []);

	// threshold change handler for target list
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

	// onSaveClick Save data on server
	const onSaveClick = useCallback(() => {
		// get list of selected targets
		const checkedTargets = getCheckedExperimentTargets(
			selectedTargetState.get('targetList'),
		);
		dispatch(createExperimentTarget(checkedTargets, experimentId));
	}, [selectedTargetState, experimentId, dispatch]);

	const onNextClick = () => {
		dispatch(setIsPlateRoute(true));
		setRedirectToPlate(true);
	};

	if (isRedirectToPlate === true) {
		return <Redirect to='/plate' />;
	}

	return (
		<TargetComponent
			selectedTemplateDetails={selectedTemplateDetails}
			selectedTargetState={selectedTargetState.get('targetList')}
			onCheckedHandler={onTargetCheckedHandler}
			onThresholdChange={onThresholdChange}
			onSaveClick={onSaveClick}
			onNextClick={onNextClick}
			isLoginTypeAdmin={isLoginTypeAdmin}
			isLoginTypeOperator={isLoginTypeOperator}
			isTargetListUpdated={isTargetListUpdated(selectedTargetState)}
			isNoTargetSelected={isNoTargetSelected(selectedTargetState.get('targetList'))}
		/>
	);
};

TargetExperimentContainer.propTypes = {
	isLoginTypeAdmin: PropTypes.bool.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
};

export default React.memo(TargetExperimentContainer);
