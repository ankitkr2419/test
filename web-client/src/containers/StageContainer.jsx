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

const StageContainer = (props) => {
	const { updateSelectedWizard, templateID, updateStageID } = props;
	const dispatch = useDispatch();
	const [selectedStageId, setSelectedStageId] = useState(null);

	// reading stages from redux
	const stages = useSelector(state => state.listStagesReducer);
	// reading stages from redux
	const { isStageSaved } = useSelector(state => state.createStageReducer);
	const { isStageDeleted } = useSelector(state => state.deleteStageReducer);
	const { isStageUpdated } = useSelector(state => state.updateStageReducer);

	useEffect(() => {
		if (isStageSaved === true) {
			dispatch(addStageReset());
			dispatch(fetchStages(templateID));
		}
	}, [isStageSaved, templateID, dispatch]);

	useEffect(() => {
		if (isStageDeleted === true) {
			dispatch(deleteStageReset());
			dispatch(fetchStages(templateID));
		}
	}, [isStageDeleted, templateID, dispatch]);

	useEffect(() => {
		if (isStageUpdated === true) {
			dispatch(updateStageReset());
			dispatch(fetchStages(templateID, dispatch));
		}
	}, [isStageUpdated, templateID, dispatch]);

	useEffect(() => {
		dispatch(fetchStages(templateID));
	}, [templateID, dispatch]);

	const addStage = (stage) => {
		// creating stage though api
		dispatch(addStageAction(stage));
	};

	const deleteStage = (stageId) => {
		console.log('in delete stageID: ', stageId);
		// deleting template though api
		dispatch(deleteStageAction(stageId));
	};

	const saveStage = (stageID, stage) => {
		console.log('saveStage: ', stageID, stage);
		dispatch(updateStageAction(stageID, { ...stage, template_id: templateID }));
	};

	const onStageRowClicked = (stageId) => {
		if (stageId === selectedStageId) {
			setSelectedStageId(null);
		} else {
			setSelectedStageId(stageId);
		}
	};

	const goToStepWizard = (stageId) => {
		updateStageID(stageId);
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
		/>
	);
};

StageContainer.propTypes = {
	updateSelectedWizard: PropTypes.func.isRequired,
	templateID: PropTypes.string.isRequired,
};

export default StageContainer;
