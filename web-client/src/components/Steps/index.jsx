import React from "react";
import Table from "core-components/Table";
import Icon from "shared-components/Icon";
import {
	TableWrapper,
	TableWrapperBody,
	TableWrapperFooter,
} from "shared-components/TableWrapper";
import ButtonIcon from "shared-components/ButtonIcon";
import AddStepModal from "./AddStepModal";

const StepList = [
	{
		steps: "1",
		ramp_rate: "value",
		target_temp: "value",
		hold_time: "value",
		data_capture: "value",
	},
	{
		steps: "1",
		ramp_rate: "value",
		target_temp: "value",
		hold_time: "value",
		data_capture: "value",
	},
	{
		steps: "1",
		ramp_rate: "value",
		target_temp: "value",
		hold_time: "value",
		data_capture: "value",
	},
	{
		steps: "1",
		ramp_rate: "value",
		target_temp: "value",
		hold_time: "value",
		data_capture: "value",
	},
	{
		steps: "1",
		ramp_rate: "value",
		target_temp: "value",
		hold_time: "value",
		data_capture: "value",
	}
];

const Steps = (props) => {
	return (
		<div className="d-flex flex-column flex-100">
			<TableWrapper>
				<TableWrapperBody>
					<Table striped>
						<colgroup>
							<col width="16%" />
							<col width="12%" />
							<col />
							<col width="16%" />
							<col width="16%" />
							<col width="156px" />
						</colgroup>
						<thead>
							<tr>
								<th>
									Steps <br />
									(Count/Name)
								</th>
								<th>
									Ramp rate <br />
									(unit °C)
								</th>
								<th>
									Target Temperature <br />
									(unit °C)
								</th>
								<th>
									Hold Time <br />
									(unit seconds)
								</th>
								<th>
									Data Capture <br />
									(boolean flag)
								</th>
								<th />
							</tr>
						</thead>
						<tbody>
							{/* TODO: Add "active" class (ontouch) on <tr> tag to see action buttons */}
							{StepList.map((step, i) => (
								<tr key={i}>
									<td>{step.steps}</td>
									<td>{step.ramp_rate}</td>
									<td>{step.target_temp}</td>
									<td>{step.hold_time}</td>
									<td>{step.data_capture}</td>
									<td className="td-actions">
										<ButtonIcon>
											<Icon size={28} name="pencil" />
										</ButtonIcon>
										<ButtonIcon>
											<Icon size={28} name="trash" />
										</ButtonIcon>
									</td>
								</tr>
							))}
						</tbody>
					</Table>
				</TableWrapperBody>
				<TableWrapperFooter>
					<AddStepModal />
				</TableWrapperFooter>
			</TableWrapper>
		</div>
	);
};

Steps.propTypes = {};

export default Steps;
