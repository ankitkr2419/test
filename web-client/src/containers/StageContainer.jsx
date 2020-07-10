import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import StageComponent from 'components/Stage';
import { useDispatch, useSelector } from 'react-redux';
import {
	fetchStages,
	addStage as addStageAction,
	addStageReset,
	deleteStage as deleteStageAction,
	deleteStageReset,
	updateStage as updateStageAction,
	updateStageReset,
} from 'action-creators/stageActionCreators';
import { getStageList } from 'selectors/stageSelector';

const StageContainer = (props) => {
	const { updateSelectedWizard, templateID, updateStageID } = props;
	const dispatch = useDispatch();
	// local state for storing stage id for row selection
	const [selectedStageId, setSelectedStageId] = useState(null);

	// reading stages from redux
	const stages = useSelector(getStageList);

	const isStagesLoading = stages.get('isLoading');

	// isStageSaved = true means stage created successfully
	const { isStageSaved, response } = useSelector(state => state.createStageReducer);
	// isStageDeleted = true means stage deleted successfully
	const { isStageDeleted } = useSelector(state => state.deleteStageReducer);
	// isStageSaved = true means stage updated successfully
	const { isStageUpdated } = useSelector(state => state.updateStageReducer);

	useEffect(() => {
		// Once we create stage will fetch updated stage list
		if (isStageSaved === true) {
			// set the newly created stage active
			setSelectedStageId(response.id);
			dispatch(addStageReset());
			dispatch(fetchStages(templateID));
		}
	}, [isStageSaved, templateID, dispatch, response]);

	useEffect(() => {
		// Once we delete stage will fetch updated stage list
		if (isStageDeleted === true) {
			dispatch(deleteStageReset());
			dispatch(fetchStages(templateID));
		}
	}, [isStageDeleted, templateID, dispatch]);

	useEffect(() => {
		// Once we update stage will fetch updated stage list
		if (isStageUpdated === true) {
			dispatch(updateStageReset());
			dispatch(fetchStages(templateID, dispatch));
		}
	}, [isStageUpdated, templateID, dispatch]);

	useEffect(() => {
		// fetch updated stage list from server
		dispatch(fetchStages(templateID));
	}, [templateID, dispatch]);

	const addStage = (stage) => {
		// creating stage through api
		dispatch(addStageAction(stage));
	};

	const deleteStage = (stageId) => {
		// deleting stage through api
		dispatch(deleteStageAction(stageId));
	};

	const saveStage = (stageID, stage) => {
		// updating stage through api
		dispatch(updateStageAction(stageID, { ...stage, template_id: templateID }));
	};

	// Here will update selected stage id
	const onStageRowClicked = (stageId) => {
		// remove stage id is if already selected
		if (stageId === selectedStageId) {
			setSelectedStageId(null);
		} else {
			setSelectedStageId(stageId);
		}
	};

	const goToStepWizard = (stageId) => {
		// stageId will be required for step wizard, So before navigating set it.
		updateStageID(stageId);
		// navigation to step wizard
		updateSelectedWizard('step');
	};

	return (
		<StageComponent
			templateID={templateID}
			stages={stages.get('list')}
			addStage={addStage}
			deleteStage={deleteStage}
			onStageRowClicked={onStageRowClicked}
			selectedStageId={selectedStageId}
			saveStage={saveStage}
			goToStepWizard={goToStepWizard}
			isStagesLoading={isStagesLoading}
		/>
	);
};

StageContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	templateID: PropTypes.string.isRequired,
};

export default  React.memo(StageContainer);
