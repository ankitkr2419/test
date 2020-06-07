import React from 'react';
import { Step, StepItem, StepLink } from "components/shared/StepBar/StepBar";
import MlCard from 'components/shared/Card/Card';


const TemplateContainer = props => {
  return (
		<>
			<Step>
				<StepItem>
					<StepLink to="/templates">Template</StepLink>
				</StepItem>
				<StepItem>
					<StepLink to="/login" className="is-disabled">View Target</StepLink>
				</StepItem>
			</Step>
      <MlCard></MlCard>
		</>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;