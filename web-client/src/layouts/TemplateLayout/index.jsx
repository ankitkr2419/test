import React, {
	useCallback, useReducer } from 'react';
import { CardBody, Card } from 'core-components';
import Wizard from 'shared-components/Wizard';
import TemplateContainer from 'containers/TemplateContainer';
import TargetContainer from 'containers/TargetContainer';
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
import { useSelector } from "react-redux";
import { ROUTES } from 'appConstants';
import { useHistory } from "react-router";

const TemplateLayout = (props) => {
	const history = useHistory();

	//get login reducer details
	const loginReducer = useSelector((state) => state.loginReducer);
	const loginReducerData = loginReducer.toJS();
	let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
	const { isAdmin, isLoggedIn } = activeDeckObj;

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
	// console.log('activeWidgetID: ',activeWidgetID);
	const templateID = templateLayoutState.get('templateID');
	// console.log('templateID: ', templateID)

	const wizardList = getWizardListByLoginType(
		templateLayoutState.get('wizardList'),
		isAdmin
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

	if(!isLoggedIn){
		// history.push(ROUTES.login);		
		history.push('splashscreen');
	}

	return (
		<div className='template-content'>
			<Wizard
				list={wizardList}
				onClickHandler={updateSelectedWizard}
				isAdmin={isAdmin}
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
							isLoginTypeOperator={!isAdmin}
							isLoginTypeAdmin={isAdmin}
							updateSelectedWizard={updateSelectedWizard}
							updateTemplateID={updateTemplateID}
							toggleTemplateModal={toggleTemplateModal}
							isCreateTemplateModalVisible={templateModalState.get('isCreateTemplateModalVisible')}
						/>
					)}
					{activeWidgetID === 'target' && (
						<TargetContainer
							isLoginTypeOperator={!isAdmin}
							isLoginTypeAdmin={isAdmin}
							updateSelectedWizard={updateSelectedWizard}
							templateID={templateID}
							setIsTemplateEdited={setIsTemplateEdited}
						/>
					)}
					{/**TODO remove comments when above code is working/tested */}
					{/*activeWidgetID === 'target-operator' && (
						<TargetExperimentContainer
							isLoginTypeOperator={!isAdmin}
							isLoginTypeAdmin={isAdmin}
							updateSelectedWizard={updateSelectedWizard}
							templateID={templateID}
						/>
					)*/}
					{/*activeWidgetID === 'step' && (
						<StepContainer />
					)*/}
				</CardBody>
			</Card>
		</div>
	);
};

export default React.memo(TemplateLayout);
