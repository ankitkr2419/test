import React from 'react';
import { ButtonIcon, Icon } from 'shared-components';
import { Table, Button } from 'core-components';

const HoldSteps = (props) => {
	return (
		<Table className='table-steps -hold' size='sm' striped>
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
						<Button color='primary' icon className='ml-auto'>
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
	);
};

HoldSteps.propTypes = {};

export default HoldSteps;
