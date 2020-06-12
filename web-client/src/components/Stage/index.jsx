import React from 'react';
import Table from 'core-components/Table';
import { Button as IconButton } from 'reactstrap';
import Button from "core-components/Button";
import Icon from 'shared-components/Icon';
import {
	TableWrapper,
	TableWrapperFooter,
} from "shared-components/TableWrapper";
import ButtonIcon from "shared-components/ButtonIcon";

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
									<ButtonIcon>
										<Icon name="steps" />
									</ButtonIcon>
									<ButtonIcon>
										<Icon name="pencil" />
									</ButtonIcon>
									<ButtonIcon>
										<Icon name="trash" />
									</ButtonIcon>
								</td>
							</tr>
						))}
					</tbody>
				</Table>
				<TableWrapperFooter>
					<IconButton color="primary" className="btn-plus">
						<Icon name="plus-3" />
					</IconButton>
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