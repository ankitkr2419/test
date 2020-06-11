import React from "react";
import { CardBody } from "reactstrap";
import Card from 'core-components/Card';
import { Step, StepItem, StepLink } from "shared-components/StepBar";
// import { TemplateListContainer } from "components/Template";
import { TargetListContainer } from "components/Target";

const Steps = [
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
				{Steps.map((step, i) => (
					<StepItem key={i} isDisable={step.disable}>
						<StepLink href="/">{step.name}</StepLink>
					</StepItem>
				))}
			</Step>
			<Card>
				<CardBody className="d-flex flex-unset overflow-hidden">
					{/* <TemplateListContainer /> */}
					<TargetListContainer />
				</CardBody>
			</Card>
		</div>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;