import React, { useCallback, useReducer } from 'react';
// import PropTypes from "prop-types";
import { CardBody } from 'reactstrap';
import { Card } from 'core-components';
import Wizard from 'shared-components/Wizard';
import TemplateContainer from 'containers/TemplateContainer';
import templateLayoutReducer, {
	templateInitialState,
	templateLayoutActions,
} from './templateState';

const TemplateLayout = (props) => {
	// Local state to manage selected wizard
	const [templateLayoutState, templateLayoutDispatch] = useReducer(
		templateLayoutReducer,
		templateInitialState,
	);

	// Here we have stored id for active widget
	const activeWidgetID = templateLayoutState
		.get('wizardList')
		.find(obj => obj.get('isDisable') === false).get('id');

	// Wizard click handler
	const onWizardIemClicked = useCallback((selectedWizard) => {
		// TODO add validation before updating wizard
		templateLayoutDispatch({
			type: templateLayoutActions.SET_ACTIVE_WIDGET,
			value: 'stage',
		});
	}, []);

	return (
		<div className="template-content">
			<Wizard
				list={templateLayoutState.get('wizardList')}
				onClickHandler={onWizardIemClicked}
			/>
			<Card>
				{/* TODO move CardBody to core-components */}
				<CardBody className="d-flex flex-unset overflow-hidden p-0">
					{activeWidgetID === 'template' && <TemplateContainer />}
					{/* <TargetListContainer /> */}
					{/* <Stage /> */}
					{/* <Steps /> */}
				</CardBody>
			</Card>
		</div>
	);
};

TemplateLayout.propTypes = {};

export default TemplateLayout;
