import React from "react";
import Button from "core-components/Button";
import {
	TemplateList,
	TemplateListItem,
	Template,
} from "shared-components/Template";

const Templates = [
	{
		name: "Template Name",
		active: false,
		editable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
	},
	{
		name: "Template Name",
		active: true,
		editable: true,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
	},
	{
		name: "Template Name",
		active: false,
		editable: false,
	},
];

export const TemplateListContainer = (props) => {
	return (
		<div className="d-flex flex-100 flex-column pt-4">
			<TemplateList>
        {Templates.map((template, i) => 
            <TemplateListItem key={i}>
              <Template title={template.name} isActive={template.active} isEditable={template.editable} />
            </TemplateListItem>
        )}
			</TemplateList>
			<div className="d-flex">
				<Button color="primary" className="mx-auto">
					Next
				</Button>
			</div>
		</div>
	);
};
