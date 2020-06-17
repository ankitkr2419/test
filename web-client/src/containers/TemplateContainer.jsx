import React from "react";
import { Card, CardBody } from 'core-components';
import { Step, StepItem, StepLink } from "shared-components";
// import { TemplateListContainer } from "components/Template";
import { TargetListContainer } from "components/Target";
// import Stage from "components/Stage";
// import Steps from "components/Steps";

const StepList = [
	{
		name: "Template",
		disable: false,
	},
	{
		name: "View Target",
		disable: true,
	},
	{
		name: "Add Stage",
		disable: true,
	},
	{
		name: "Add Step",
		disable: true,
	},
];

const TemplateContainer = props => {
  return (
		<div className="template-content">
			<Step>
				{StepList.map((step, i) => (
					<StepItem key={i} isDisable={step.disable}>
						<StepLink>{step.name}</StepLink>
					</StepItem>
				))}
			</Step>
			<Card>
				<CardBody className="d-flex flex-unset overflow-hidden p-0">
					{/* <TemplateListContainer /> */}
					<TargetListContainer />
					{/* <Stage /> */}
					{/* <Steps /> */}
				</CardBody>
			</Card>
		</div>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;