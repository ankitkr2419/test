import React from 'react';
import { Step, StepItem, StepLink } from "core-components/StepBar";
import Card from 'core-components/Card';


const TemplateContainer = props => {
  return (
		<>
			<Step>
				<StepItem>
					<StepLink to="/templates">Template</StepLink>
				</StepItem>
				<StepItem isDisable>
					<StepLink to="/login">View Target</StepLink>
				</StepItem>
			</Step>
			<Card></Card>
		</>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;