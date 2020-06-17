import React from "react";
import { Button } from "core-components";
import { ButtonGroup, TemplateList, TemplateListItem, Template } from "shared-components";
import CreateTemplateModal from "./CreateTemplateModal";

const Templates = [
	{
		name: "Template Name",
		active: false,
		editable: false,
		deletable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
		deletable: false,
	},
	{
		name: "Template Name",
		active: true,
		editable: true,
		deletable: true,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
		deletable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
		deletable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
		deletable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
		deletable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
		deletable: false,
	},
];

export const TemplateListContainer = (props) => {
	return (
		<div className="d-flex flex-100 flex-column p-4 mt-3">
			<TemplateList>
				{Templates.map((template, i) => 
					<TemplateListItem key={i}>
						<Template title={template.name} isActive={template.active} isEditable={template.editable} isDeletable={template.deletable} />
					</TemplateListItem>
				)}
			</TemplateList>
			<ButtonGroup className="text-center">
				<Button color="primary">
					Next
				</Button>
				<CreateTemplateModal />
			</ButtonGroup>
		</div>
	);
};
