import React, {
	useCallback, useReducer } from 'react';
import { CardBody, Card } from 'core-components';
import Wizard from 'shared-components/Wizard';
import TemplateContainer from 'containers/TemplateContainer';
import TargetContainer from 'containers/TargetContainer';
import StageContainer from 'containers/StageContainer';
import StepContainer from 'containers/StepContainer';
import TargetExperimentContainer from 'containers/TargetExperimentContainer';
import TemplateModalContainer from 'containers/TemplateModalContainer';
import templateModalReducer, {
	templateModalInitialState,
	templateModalActions,
} from 'components/TemplateModal/templateModalState';
import templateLayoutReducer, {
	templateInitialState,
	templateLayoutActions,
	getWizardListByLoginType,
} from './templateState';


const TemplateLayout = (props) => {
	const { isLoginTypeOperator, isLoginTypeAdmin } = props;
	// Local state to manage selected wizard
	const [templateLayoutState, templateLayoutDispatch] = useReducer(
		templateLayoutReducer,
		templateInitialState,
	);

	// Local state to manage template Modal
	const [templateModalState, templateModalDispatch] = useReducer(
		templateModalReducer,
		templateModalInitialState,
	);

	// Here we have stored id for active widget
	const activeWidgetID = templateLayoutState.get('activeWidgetID');
	const templateID = templateLayoutState.get('templateID');
	const stageId = templateLayoutState.get('stageId');

	const wizardList = getWizardListByLoginType(
		templateLayoutState.get('wizardList'),
		isLoginTypeAdmin,
		isLoginTypeOperator,
	);

	// helper method to toggle template modal
	const toggleTemplateModal = useCallback(() => {
		templateModalDispatch({
			type: templateModalActions.TOGGLE_TEMPLATE_MODAL_VISIBLE,
		});
	}, []);

	// helper method to set isTemplateEdited flag true
	const setIsTemplateEdited = useCallback(() => {
		templateModalDispatch({
			type:templateModalActions.SET_IS_TEMPLATE_EDITED,
		});
	}, []);

	// Wizard click handler
	const updateSelectedWizard = useCallback((selectedWizard) => {
		templateLayoutDispatch({
			type: templateLayoutActions.SET_ACTIVE_WIDGET,
			value: selectedWizard,
		});
	}, []);

	const updateTemplateID = useCallback((selectedTemplateID) => {
		templateLayoutDispatch({
			type: templateLayoutActions.SET_TEMPLATE_ID,
			value: selectedTemplateID,
		});
		// reset wizard list to disable already enabled wizard stages
		templateLayoutDispatch({
			type: templateLayoutActions.RESET_WIZARD_LIST,
		});
	}, []);

	const updateStageID = useCallback((selectedStageId) => {
		templateLayoutDispatch({
			type: templateLayoutActions.SET_STAGE_ID,
			value: selectedStageId,
		});
	}, []);

	return (
		<div className='template-content'>
			<Wizard
				list={wizardList}
				onClickHandler={updateSelectedWizard}
				isLoginTypeAdmin={isLoginTypeAdmin}
			/>
			<Card>
				<CardBody className='d-flex flex-unset overflow-hidden p-0'>
					{/* TemplateModal container that provides template modal to create
							and edit the template from template and target wizards */}
					<TemplateModalContainer
						templateModalState={templateModalState}
						templateModalDispatch={templateModalDispatch}
						templateID={templateID}
						toggleTemplateModal={toggleTemplateModal}
					/>
					{activeWidgetID === 'template' && (
						<TemplateContainer
							isLoginTypeOperator={isLoginTypeOperator}
							isLoginTypeAdmin={isLoginTypeAdmin}
							updateSelectedWizard={updateSelectedWizard}
							updateTemplateID={updateTemplateID}
							toggleTemplateModal={toggleTemplateModal}
						/>
					)}
					{activeWidgetID === 'target' && (
						<TargetContainer
							isLoginTypeOperator={isLoginTypeOperator}
							isLoginTypeAdmin={isLoginTypeAdmin}
							updateSelectedWizard={updateSelectedWizard}
							templateID={templateID}
							setIsTemplateEdited={setIsTemplateEdited}
						/>
					)}
					{activeWidgetID === 'target-operator' && (
						<TargetExperimentContainer
							isLoginTypeOperator={isLoginTypeOperator}
							isLoginTypeAdmin={isLoginTypeAdmin}
							updateSelectedWizard={updateSelectedWizard}
							templateID={templateID}
						/>
					)}

					{activeWidgetID === 'stage' && (
						<StageContainer
							updateSelectedWizard={updateSelectedWizard}
							updateStageID={updateStageID}
							templateID={templateID}
						/>
					)}
					{activeWidgetID === 'step' && (
						<StepContainer
							updateSelectedWizard={updateSelectedWizard}
							stageId={stageId}
						/>
					)}
				</CardBody>
			</Card>
		</div>
	);
};

export default React.memo(TemplateLayout);
