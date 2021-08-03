import React from 'react';
import PropTypes from 'prop-types';
import { ButtonIcon, Icon } from 'shared-components';
import { Table, Button } from 'core-components';
import { HOLD_STAGE } from './stepConstants';

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
						<Button color='primary' icon className='ml-auto' onClick={addHoldStep}>
							<Icon size={40} name='plus-2' />
						</Button>
					</th>
				</tr>
			</thead>
			<tbody>
				{holdSteps.map((step, index) => {
					const stepId = step.get('step_id');
					const classes = selectedStepId === stepId ? 'active' : '';
					return (
						<tr
							className={classes}
							key={stepId}
							onClick={() => {
								onStepRowClicked(stepId, HOLD_STAGE);
							}}
						>
							<td>{index + 1}</td>
							<td>{step.get('ramp_rate')}</td>
							<td>{step.get('target_temp')}</td>
							<td>{(step.get('hold_time'))}</td>
							<td className='td-actions'>
								<ButtonIcon
									size={16}
									name='pencil'
									onClick={() => {
										editStep(step.toJS());
									}}
								/>
								<ButtonIcon
									size={16}
									name='trash'
									onClick={() => {
										deleteStep(stepId);
									}}
								/>
							</td>
						</tr>
					);
				})}
			</tbody>
		</Table>
	</div>
);

HoldSteps.propTypes = {
	addHoldStep: PropTypes.func.isRequired,
	editStep: PropTypes.func.isRequired,
	holdSteps: PropTypes.object.isRequired,
	deleteStep: PropTypes.func.isRequired,
	onStepRowClicked: PropTypes.func.isRequired,
	selectedStepId: PropTypes.string,
};

export default HoldSteps;
