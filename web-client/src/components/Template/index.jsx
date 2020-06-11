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
			<div className="d-flex">
				<Button color="primary" className="mx-auto">
					Next
				</Button>
			</div>
		</div>
	);
};
