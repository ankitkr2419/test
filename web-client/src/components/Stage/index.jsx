import React from "react";
import { Button, Table } from "core-components";
import {
	ButtonIcon,
	TableWrapper,
	TableWrapperFooter,
} from "shared-components";
import AddStageModal from "./AddStageModal";

const StageList = [
	{
		stage: "1",
		type: "value",
		count: "value",
		steps: "value",
	},
	{
		stage: "1",
		type: "value",
		count: "value",
		steps: "value",
	},
	{
		stage: "1",
		type: "value",
		count: "value",
		steps: "value",
	},
	{
		stage: "1",
		type: "value",
		count: "value",
		steps: "value",
	},
	{
		stage: "1",
		type: "value",
		count: "value",
		steps: "value",
	},
];

const Stage = props => {
  return (
		<div className="d-flex flex-column flex-100">
			<TableWrapper>
				<Table striped>
					<colgroup>
						<col width="17%" />
						<col width="17%" />
						<col width="19%" />
						<col width="17%" />
						<col />
					</colgroup>
					<thead>
						<tr>
							<th>Stage</th>
							<th>Type</th>
							<th>Repeat Count</th>
							<th>Steps</th>
							<th />
						</tr>
					</thead>
					<tbody>
						{StageList.map((stage, i) => (
							<tr key={i}>
								<td>{stage.stage}</td>
								<td>{stage.type}</td>
								<td>{stage.count}</td>
								<td>{stage.steps}</td>
								<td className="td-actions">
									<ButtonIcon size={28} name="steps" />
									<ButtonIcon size={28} name="pencil" />
									<ButtonIcon size={28} name="trash" />
								</td>
							</tr>
						))}
					</tbody>
				</Table>
				<TableWrapperFooter>
					<AddStageModal />
					<Button color="primary" className="ml-auto">
						Save
					</Button>
				</TableWrapperFooter>
			</TableWrapper>
		</div>
	);
};

Stage.propTypes = {};

export default Stage;