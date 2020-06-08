import React from 'react';
import { Step, StepItem, StepLink } from "core-components/StepBar";
import Card from 'core-components/Card';
import { TemplateList, TemplateListItem, Template } from "components/Template";
import Button from 'core-components/Button';


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
				<TemplateList>
					<TemplateListItem>
						<Template>Template Name</Template>
					</TemplateListItem>
					<TemplateListItem>
						<Template>Template Name</Template>
					</TemplateListItem>
					<TemplateListItem>
						<Template isActive>
							<Button icon className="btn-edit">
								E
							</Button>
							Template 3
							<Button icon className="btn-delete">
								D
							</Button>
						</Template>
					</TemplateListItem>
					<TemplateListItem>
						<Template>Template Name</Template>
					</TemplateListItem>
					<TemplateListItem>
						<Template>Template Name</Template>
					</TemplateListItem>
					<TemplateListItem>
						<Template>Template Name</Template>
					</TemplateListItem>
					<TemplateListItem>
						<Template>Template Name</Template>
					</TemplateListItem>
					<TemplateListItem>
						<Template>Template Name</Template>
					</TemplateListItem>
				</TemplateList>
				<Button color="primary" className="mx-auto">
					Next
				</Button>
			</Card>
		</>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;