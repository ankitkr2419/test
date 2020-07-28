import React from 'react';
import { ButtonIcon, Icon } from 'shared-components';
import { Table, Button, CheckBox } from 'core-components';
import CounterPopover from './CounterPopover';

const CycleSteps = ({
	addCycleStep,
	editStep,
	cycleSteps,
	deleteStep,
	onStepRowClicked,
	selectedStepId,
}) => (
	<Table className='table-steps -cycle' size='sm' striped>
		<colgroup>
			<col width='132px' />
			<col width='102px' />
			<col width='160px' />
			<col width='128px' />
			<col width='125px' />
			<col width='133px' />
			<col />
		</colgroup>
		<thead>
			<tr>
				<th>Cycle Steps</th>
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
				<th className='th-counter'>
					Repeat counter <br />
					<CounterPopover />
				</th>
				<th>
					<Button
						color='primary'
						icon
						className='ml-auto'
						onClick={addCycleStep}
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
				<td>40</td>
				<td>30</td>
				<td>
					<CheckBox id='checkbox1' />
				</td>
				<td />
				<td className='td-actions'>
					<ButtonIcon size={16} name='pencil' />
					<ButtonIcon size={16} name='trash' />
				</td>
			</tr>
			<tr>
				<td>2</td>
				<td>2</td>
				<td>40</td>
				<td>30</td>
				<td>
					<CheckBox id='checkbox2' />
				</td>
				<td />
				<td className='td-actions'>
					<ButtonIcon size={16} name='pencil' />
					<ButtonIcon size={16} name='trash' />
				</td>
			</tr>
			<tr>
				<td>3</td>
				<td>2</td>
				<td>40</td>
				<td>30</td>
				<td>
					<CheckBox id='checkbox3' />
				</td>
				<td />
				<td className='td-actions'>
					<ButtonIcon size={16} name='pencil' />
					<ButtonIcon size={16} name='trash' />
				</td>
			</tr>
			<tr>
				<td>4</td>
				<td>2</td>
				<td>40</td>
				<td>30</td>
				<td>
					<CheckBox id='checkbox4' />
				</td>
				<td />
				<td className='td-actions'>
					<ButtonIcon size={16} name='pencil' />
					<ButtonIcon size={16} name='trash' />
				</td>
			</tr>
			<tr>
				<td>5</td>
				<td>2</td>
				<td>40</td>
				<td>30</td>
				<td>
					<CheckBox id='checkbox5' />
				</td>
				<td />
				<td className='td-actions'>
					<ButtonIcon size={16} name='pencil' />
					<ButtonIcon size={16} name='trash' />
				</td>
			</tr>
			<tr>
				<td>6</td>
				<td>2</td>
				<td>40</td>
				<td>30</td>
				<td>
					<CheckBox id='checkbox6' />
				</td>
				<td />
				<td className='td-actions'>
					<ButtonIcon size={16} name='pencil' />
					<ButtonIcon size={16} name='trash' />
				</td>
			</tr>
		</tbody>
	</Table>
);

CycleSteps.propTypes = {};

export default CycleSteps;
