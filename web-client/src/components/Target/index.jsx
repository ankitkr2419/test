import React from "react";
import { Button, CheckBox, Select } from "core-components";
import {
	TargetList,
	TargetListHeader,
	TargetListItem,
	Text
} from "shared-components";

export const TargetListContainer = (props) => {

  const Targets = [
		{
			selectionOption1: [
				{ value: "option 1", label: "Option 1" },
				{ value: "option 2", label: "Option 2" },
				{ value: "option 3", label: "Option 3" },
			],
			selectionOption2: [
				{ value: "0.1", label: "0.1" },
				{ value: "0.2", label: "0.2" },
				{ value: "0.3", label: "0.3" },
			],
		},
		{
			selectionOption1: [
				{ value: "option 1", label: "Option 1" },
				{ value: "option 2", label: "Option 2" },
				{ value: "option 3", label: "Option 3" },
			],
			selectionOption2: [
				{ value: "0.1", label: "0.1" },
				{ value: "0.2", label: "0.2" },
				{ value: "0.3", label: "0.3" },
			],
		},
		{
			selectionOption1: [
				{ value: "option 1", label: "Option 1" },
				{ value: "option 2", label: "Option 2" },
				{ value: "option 3", label: "Option 3" },
			],
			selectionOption2: [
				{ value: "0.1", label: "0.1" },
				{ value: "0.2", label: "0.2" },
				{ value: "0.3", label: "0.3" },
			],
		},
		{
			selectionOption1: [
				{ value: "option 1", label: "Option 1" },
				{ value: "option 2", label: "Option 2" },
				{ value: "option 3", label: "Option 3" },
			],
			selectionOption2: [
				{ value: "0.1", label: "0.1" },
				{ value: "0.2", label: "0.2" },
				{ value: "0.3", label: "0.3" },
			],
		},
		{
			selectionOption1: [
				{ value: "option 1", label: "Option 1" },
				{ value: "option 2", label: "Option 2" },
				{ value: "option 3", label: "Option 3" },
			],
			selectionOption2: [
				{ value: "0.1", label: "0.1" },
				{ value: "0.2", label: "0.2" },
				{ value: "0.3", label: "0.3" },
			],
		},
		{
			selectionOption1: [
				{ value: "option 1", label: "Option 1" },
				{ value: "option 2", label: "Option 2" },
				{ value: "option 3", label: "Option 3" },
			],
			selectionOption2: [
				{ value: "0.1", label: "0.1" },
				{ value: "0.2", label: "0.2" },
				{ value: "0.3", label: "0.3" },
			],
		},
	];

	return (
		<>
			<div className="flex-100 scroll-y p-4">
				<TargetList className="list-target">
					<TargetListHeader>
						<Text className="mb-2 mr-2" />
						<Text className="flex-100 mb-2 px-4">Target</Text>
						<Text className="flex-40 mb-2 px-4">Threshold</Text>
					</TargetListHeader>
					{Targets.map((target, i) => (
						<TargetListItem key={i}>
							<CheckBox className="mr-2" id={`target${i}`} />
							<Select
								className="flex-100 px-2"
								options={target.selectionOption1}
								placeholder=""
							/>
							<Select
								className="flex-40 pl-2"
								options={target.selectionOption2}
								placeholder=""
							/>
						</TargetListItem>
					))}
				</TargetList>
			</div>
			<div className="d-flex flex-30 align-items-end p-4">
				<Button color="primary" className="mx-auto mb-3" disabled>
					Save
				</Button>
			</div>
		</>
	);
};