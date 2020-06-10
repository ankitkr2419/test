import React from "react";
import { CardBody } from "reactstrap";
import { Step, StepItem, StepLink } from "core-components/StepBar";
import Card from 'core-components/Card';
import Button from 'core-components/Button';
import Select from 'core-components/Select';
import CheckBox from "core-components/Checkbox";
import {
	TemplateList,
	TemplateListItem,
	Template,
} from "shared-components/Template";
import { TargetList, TargetListHeader, TargetListItem } from "shared-components/Target";
import Text from "shared-components/Text";

const TemplateListContainer = props => {
	return (
		<div className="d-flex flex-column pt-4">
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
			<div className="d-flex">
				<Button color="primary" className="mx-auto">
					Next
				</Button>
			</div>
		</div>
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
		<>
			<div className="flex-100 scroll-y p-1">
				<TargetList className="list-target">
					<TargetListHeader>
						<Text tag="p" className="mb-2 mr-2" />
						<Text tag="p" className="flex-100 mb-2 px-4">
							Target
						</Text>
						<Text tag="p" className="flex-40 mb-2 px-4">
							Threshold
						</Text>
					</TargetListHeader>
					<TargetListItem>
						<CheckBox className="mr-2" id="target1" />
						<Select
							className="flex-100 px-2"
							options={TargetOptions}
							placeholder=""
						/>
						<Select
							className="flex-40 pl-2"
							options={ThresholdOptions}
							placeholder=""
						/>
					</TargetListItem>
					<TargetListItem>
						<CheckBox className="mr-2" id="target2" />
						<Select
							className="flex-100 px-2"
							options={TargetOptions}
							placeholder=""
						/>
						<Select
							className="flex-40 pl-2"
							options={ThresholdOptions}
							placeholder=""
						/>
					</TargetListItem>
					<TargetListItem>
						<CheckBox className="mr-2" id="target3" />
						<Select
							className="flex-100 px-2"
							options={TargetOptions}
							placeholder=""
						/>
						<Select
							className="flex-40 pl-2"
							options={ThresholdOptions}
							placeholder=""
						/>
					</TargetListItem>
					<TargetListItem>
						<CheckBox className="mr-2" id="target4" />
						<Select
							className="flex-100 px-2"
							options={TargetOptions}
							placeholder=""
						/>
						<Select
							className="flex-40 pl-2"
							options={ThresholdOptions}
							placeholder=""
						/>
					</TargetListItem>
					<TargetListItem>
						<CheckBox className="mr-2" id="target5" />
						<Select
							className="flex-100 px-2"
							options={TargetOptions}
							placeholder=""
						/>
						<Select
							className="flex-40 pl-2"
							options={ThresholdOptions}
							placeholder=""
						/>
					</TargetListItem>
					<TargetListItem>
						<CheckBox className="mr-2" id="target6" />
						<Select
							className="flex-100 px-2"
							options={TargetOptions}
							placeholder=""
						/>
						<Select
							className="flex-40 pl-2"
							options={ThresholdOptions}
							placeholder=""
						/>
					</TargetListItem>
				</TargetList>
			</div>
			<div className="d-flex flex-30 align-items-end p-1">
				<Button color="primary" className="mx-auto mb-3" disabled>
					Save
				</Button>
			</div>
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
			<Card>
				<CardBody className="d-flex flex-unset overflow-hidden">
					{/* <TemplateListContainer /> */}
					<TargetListContainer />
				</CardBody>
			</Card>
		</>
	);
};

TemplateContainer.propTypes = {};

export default TemplateContainer;