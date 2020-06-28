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
import {
	fetchExperimentTargets,
	createExperimentTarget,
} from 'action-creators/experimentTargetActionCreators';
import { getExperimentTargets } from 'selectors/experimentTargetSelector';
import { getSelectedTargetExperiment } from 'components/Target/targetHelper';
import { Redirect } from 'react-router';

const TargetExperimentContainer = (props) => {
	// constants
	const {
		isLoginTypeAdmin,
		isLoginTypeOperator,
	} = props;
	const dispatch = useDispatch();
	// useSelector section
	// extracting experiment id
	const { id: experimentId } = useSelector(
		state => state.createExperimentReducer,
	);
	// list of experiment targets
	const listExperimentTargetsReducer = useSelector(getExperimentTargets);
	const experimentTargets = listExperimentTargetsReducer.get('list');

	// useReducer section
	// local state to manage selected target data
	const [selectedTargetState, updateTargetState] = useReducer(
		targetStateReducer,
		fromJS({ targetList: [], originalTargetList: [] }),
	);
	console.table(selectedTargetState.toJS().targetList);
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
		console.log('checkedTargets: ', checkedTargets);
		dispatch(createExperimentTarget(checkedTargets, experimentId));
	}, [selectedTargetState, experimentId, dispatch]);

	const onNextClick = () => {
		setRedirectToPlate(true);
	};

	if (isRedirectToPlate === true) {
		return <Redirect to="/plate" />;
	}

	return (
		<TargetComponent
			selectedTargetState={selectedTargetState.get('targetList')}
			onCheckedHandler={onTargetCheckedHandler}
			onThresholdChange={onThresholdChange}
			onSaveClick={onSaveClick}
			onNextClick={onNextClick}
			isLoginTypeAdmin={isLoginTypeAdmin}
			isLoginTypeOperator={isLoginTypeOperator}
			isTargetListUpdated={isTargetListUpdated(selectedTargetState)}
		/>
	);
};

TargetExperimentContainer.propTypes = {
	isLoginTypeAdmin: PropTypes.bool.isRequired,
	isLoginTypeOperator: PropTypes.bool.isRequired,
};

export default React.memo(TargetExperimentContainer);
