import React from 'react';
import { ButtonIcon, Icon } from 'shared-components';
import { Table, Button } from 'core-components';

const HoldSteps = ({
	addHoldStep,
	editStep,
	holdSteps,
	deleteStep,
	onStepRowClicked,
	selectedStepId,
}) => (
	<div className='table-steps-wrapper -hold'>
		<Table className='table-steps' size='sm' striped>
			<colgroup>
				<col width='132px' />
				<col width='102px' />
				<col width='160px' />
				<col width='128px' />
				<col />
			</colgroup>
			<thead>
				<tr>
					<th>Hold Steps</th>
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
						<Button
							color='primary'
							icon
							className='ml-auto'
							onClick={addHoldStep}
						>
							<Icon size={40} name='plus-2' />
						</Button>
					</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>1</td>
					<td>2</td>
					<td>50</td>
					<td>30</td>
					<td className='td-actions'>
						<ButtonIcon size={16} name='pencil' />
						<ButtonIcon size={16} name='trash' />
					</td>
				</tr>
				<tr>
					<td>2</td>
					<td>2</td>
					<td>50</td>
					<td>30</td>
					<td className='td-actions'>
						<ButtonIcon size={16} name='pencil' />
						<ButtonIcon size={16} name='trash' />
					</td>
				</tr>
				<tr>
					<td>3</td>
					<td>2</td>
					<td>50</td>
					<td>30</td>
					<td className='td-actions'>
						<ButtonIcon size={16} name='pencil' />
						<ButtonIcon size={16} name='trash' />
					</td>
				</tr>
				<tr>
					<td>4</td>
					<td>2</td>
					<td>50</td>
					<td>30</td>
					<td className='td-actions'>
						<ButtonIcon size={16} name='pencil' />
						<ButtonIcon size={16} name='trash' />
					</td>
				</tr>
			</tbody>
		</Table>
	</div>
);

HoldSteps.propTypes = {};

export default HoldSteps;
