import React from "react";
import { Row, Col } from "reactstrap";
import { Step, StepItem, StepLink } from "core-components/StepBar";
import Card from 'core-components/Card';
import { TemplateList, TemplateListItem, Template } from "components/Template";
import Button from 'core-components/Button';
import Select from 'core-components/Select';
import { TargetList, TargetListItem } from "components/Target";
import CheckBox from "core-components/Checkbox";

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

const TargetListContainer = (props) => {
	const TargetOptions = [
		{ value: "option 1", label: "Option 1" },
		{ value: "option 2", label: "Option 2" },
		{ value: "option 3", label: "Option 3" },
	];

	const ThresholdOptions = [
		{ value: "0.1", label: "0.1" },
		{ value: "0.2", label: "0.2" },
		{ value: "0.3", label: "0.3" },
	];

	return (
		<Row className="mh-100">
			<Col className="mh-100">
				<TargetList className="mh-100">
					<TargetListItem>
						<p />
						<p className="-target">Target</p>
						<p className="-threshold">Threshold</p>
					</TargetListItem>
					<TargetListItem>
						<CheckBox id="target1" />
						<Select wrapperClassName="-target" options={TargetOptions} />
						<Select wrapperClassName="-threshold" options={ThresholdOptions} />
					</TargetListItem>
					<TargetListItem>
						<CheckBox id="target2" />
						<Select wrapperClassName="-target" options={TargetOptions} />
						<Select wrapperClassName="-threshold" options={ThresholdOptions} />
					</TargetListItem>
					<TargetListItem>
						<CheckBox id="target3" />
						<Select wrapperClassName="-target" options={TargetOptions} />
						<Select wrapperClassName="-threshold" options={ThresholdOptions} />
					</TargetListItem>
					<TargetListItem>
						<CheckBox id="target4" />
						<Select wrapperClassName="-target" options={TargetOptions} />
						<Select wrapperClassName="-threshold" options={ThresholdOptions} />
					</TargetListItem>
					<TargetListItem>
						<CheckBox id="target5" />
						<Select wrapperClassName="-target" options={TargetOptions} />
						<Select wrapperClassName="-threshold" options={ThresholdOptions} />
					</TargetListItem>
				</TargetList>
			</Col>
			<Col sm={4}>
				<div className="d-flex align-items-end h-100">
					<Button color="primary" className="mx-auto" disabled>
						Save
					</Button>
				</div>
			</Col>
		</Row>
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
				{/* <TemplateListContainer /> */}
				<TargetListContainer />
			</Card>
		</>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;