import React from 'react';
import PropTypes from 'prop-types';
import { ButtonIcon, Icon } from 'shared-components';
import { Table, Button, CheckBox } from 'core-components';
import CounterPopover from './CounterPopover';
import { CYCLE_STAGE } from './stepConstants';

const CycleSteps = ({
	addCycleStep,
	editStep,
	cycleSteps,
	deleteStep,
	onStepRowClicked,
	selectedStepId,
	cycleRepeatCount,
	repeatCounterState,
	updateRepeatCounterStateWrapper,
	saveRepeatCount,
}) => (
	<div className='table-steps-wrapper -cycle'>
		<Table className='table-steps' size='sm' striped>
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
					<th className='th-counter'>Repeat counter <br />
						{cycleRepeatCount !== 0 && (
							<CounterPopover
								cycleRepeatCount={cycleRepeatCount}
								repeatCounterState={repeatCounterState}
								updateRepeatCounterStateWrapper={updateRepeatCounterStateWrapper}
								saveRepeatCount={saveRepeatCount}
							/>
						)}
					</th>
					<th>
						<Button color='primary' icon className='ml-auto' onClick={addCycleStep}>
							<Icon size={40} name='plus-2' />
						</Button>
					</th>
				</tr>
			</thead>
			<tbody>
				{cycleSteps.map((step, index) => {
					const stepId = step.get('id');
					const classes = selectedStepId === stepId ? 'active' : '';
					return (
						<tr
							className={classes}
							key={stepId}
							onClick={() => {
								onStepRowClicked(stepId, CYCLE_STAGE);
							}}
						>
							<td>{index + 1}</td>
							<td>{step.get('ramp_rate')}</td>
							<td>{step.get('target_temp')}</td>
							<td>{(step.get('hold_time'))}</td>
							<td>
								<CheckBox
									id={`checkbox${index}`}
									checked={step.get('data_capture')}
									disabled={true}
								/>
							</td>
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

CycleSteps.propTypes = {
	addCycleStep: PropTypes.func.isRequired,
	editStep: PropTypes.func.isRequired,
	cycleSteps: PropTypes.object.isRequired,
	deleteStep: PropTypes.func.isRequired,
	onStepRowClicked: PropTypes.func.isRequired,
	selectedStepId: PropTypes.string,
	cycleRepeatCount: PropTypes.number.isRequired,
	repeatCounterState: PropTypes.object.isRequired,
	updateRepeatCounterStateWrapper: PropTypes.func.isRequired,
	saveRepeatCount: PropTypes.func.isRequired,
};

export default CycleSteps;
