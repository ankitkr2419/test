import React, { useCallback, useReducer } from 'react';
import { CardBody, Card } from 'core-components';
import Wizard from 'shared-components/Wizard';
import TemplateContainer from 'containers/TemplateContainer';
import TargetContainer from 'containers/TargetContainer';
import StageContainer from 'containers/StageContainer';
import StepContainer from 'containers/StepContainer';
import templateLayoutReducer, {
	templateInitialState,
	templateLayoutActions,
} from './templateState';

const TemplateLayout = () => {
	// Local state to manage selected wizard
	const [templateLayoutState, templateLayoutDispatch] = useReducer(
		templateLayoutReducer,
		templateInitialState,
	);
	// Here we have stored id for active widget
	const activeWidgetID = templateLayoutState.get('activeWidgetID');
	const templateID = templateLayoutState.get('templateID');
	const stageId = templateLayoutState.get('stageId');

	// Wizard click handler
	const updateSelectedWizard = useCallback((selectedWizard) => {
		// TODO add validation before updating wizard
		templateLayoutDispatch({
			type: templateLayoutActions.SET_ACTIVE_WIDGET,
			value: selectedWizard,
		});
	}, []);

	const updateTemplateID = useCallback((selectedTemplateID) => {
		// TODO add validation before updating wizard
		templateLayoutDispatch({
			type: templateLayoutActions.SET_TEMPLATE_ID,
			value: selectedTemplateID,
		});
	}, []);

	const updateStageID = useCallback((selectedStageId) => {
		// TODO add validation before updating wizard
		templateLayoutDispatch({
			type: templateLayoutActions.SET_STAGE_ID,
			value: selectedStageId,
		});
	}, []);

	return (
		<div className="template-content">
			<Wizard
				list={templateLayoutState.get('wizardList')}
				onClickHandler={updateSelectedWizard}
			/>
			<Card>
				{/* TODO move CardBody to core-components */}
				<CardBody className="d-flex flex-unset overflow-hidden p-0">
					{activeWidgetID === 'template' && (
						<TemplateContainer
							updateSelectedWizard={updateSelectedWizard}
							updateTemplateID={updateTemplateID}
						/>
					)}
					{activeWidgetID === 'target' && (
						<TargetContainer
							updateSelectedWizard={updateSelectedWizard}
							templateID={templateID}/>)
					}

					{activeWidgetID === 'stage' && (
						<StageContainer
							updateSelectedWizard={updateSelectedWizard}
							updateStageID={updateStageID}
							templateID={templateID}
						/>)
					}
					{activeWidgetID === 'step' && (
						<StepContainer
							updateSelectedWizard={updateSelectedWizard}
							stageId={stageId}
						/>)
					}
				</CardBody>
			</Card>
		</div>
	);
};

export default TemplateLayout;
