import React from 'react';
import { Table } from 'core-components';
import { ButtonIcon } from 'shared-components';
import './activity.scss';
import ActivityData from './ActivityData.json';

const headers = ActivityData.headers;
const experiments = ActivityData.experiments;

const ActivityComponent = (props) => {
	return (
		<div className='activity-content h-100 py-0'>
			<div className='table-responsive'>
				<Table striped className='table-log'>
					<colgroup>
						<col width='9%' />
						<col />
						<col width='12%' />
						<col width='10.5%' />
						<col width='9.5%' />
						<col width='8%' />
						<col width='8%' />
						<col width='12%' />
						<col width='15%' />
					</colgroup>
					<thead>
						<tr>
							{headers.map((header, i) => (
								<th key={i}>{header}</th>
							))}
							<th />
						</tr>
					</thead>
					<tbody>
						{experiments.map((experiment, i) => (
							<tr
								className={experiment.result === 'Inconclusive' ? 'active' : ''}
								key={i}
							>
								<td>{experiment.id}</td>
								<td>{experiment.template}</td>
								<td>{experiment.date}</td>
								<td>{experiment.start_time}</td>
								<td>{experiment.end_time}</td>
								<td>{experiment.no_of_wells}</td>
								<td>{experiment.repeat_cycles}</td>
								<td
									className={experiment.result === 'Error' ? 'text-danger' : ''}
								>
									{experiment.result}
								</td>
								<td className='td-actions'>
									<ButtonIcon size={28} name='expand' />
									<ButtonIcon size={28} name='trash' />
								</td>
							</tr>
						))}
					</tbody>
				</Table>
			</div>
		</div>
	);
};

ActivityComponent.propTypes = {};

export default ActivityComponent;
