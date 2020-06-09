import React from 'react';
import { Step, StepItem, StepLink } from "core-components/StepBar";
import Card from 'core-components/Card';
import { TemplateList, TemplateListItem, Template } from "components/Template";
import Button from 'core-components/Button';

const TemplateListContainer = props => {
	return (
		<>
			<TemplateList>
				<TemplateListItem>
					<Template title="Template Name" />
				</TemplateListItem>
				<TemplateListItem>
					<Template title="Template Name" />
				</TemplateListItem>
				<TemplateListItem>
					<Template title="Template 3" isActive isEditable />
				</TemplateListItem>
				<TemplateListItem>
					<Template title="Template Name" />
				</TemplateListItem>
				<TemplateListItem>
					<Template title="Template Name" />
				</TemplateListItem>
				<TemplateListItem>
					<Template title="Template Name" />
				</TemplateListItem>
				<TemplateListItem>
					<Template title="Template Name" />
				</TemplateListItem>
				<TemplateListItem>
					<Template title="Template Name" />
				</TemplateListItem>
			</TemplateList>
			<Button color="primary" className="mx-auto">
				Next
			</Button>
		</>
	);
};

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
			<Card className="card-template">
				<TemplateListContainer />
			</Card>
		</>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;