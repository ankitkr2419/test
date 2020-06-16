import React, { useCallback, useReducer } from 'react';
// import PropTypes from "prop-types";
import { CardBody } from 'reactstrap';
import { Card } from 'core-components';
import Wizard from 'shared-components/Wizard';
import { TargetListContainer } from 'components/Target';
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
	const activeWidgetID = templateLayoutState.get('activeWidgetID');

	// Wizard click handler
	const updateSelectedWizard = useCallback((selectedWizard) => {
		// TODO add validation before updating wizard
		templateLayoutDispatch({
			type: templateLayoutActions.SET_ACTIVE_WIDGET,
			value: selectedWizard,
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
						<TemplateContainer updateSelectedWizard={updateSelectedWizard} />
					)}
					{activeWidgetID === 'target' && <TargetListContainer />}
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
