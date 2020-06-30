import React from 'react';
import { Table } from 'core-components';
import './activity.scss';

const ActivityComponent = (props) => {
	return (
		<div className='activity-content h-100 py-0'>
			<div className='table-responsive'>
				<Table striped bordered className='table-log'>
					<thead>
						<tr>
							<th>Exp ID</th>
							<th>Template</th>
							<th>Date</th>
							<th>Start time</th>
							<th>End Time</th>
							<th>No. Of Wells</th>
							<th>Repeat Cycles</th>
							<th>Result</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>ID001</td>
							<td>Template Name</td>
							<td>23/05/2020</td>
							<td>23:50 PM</td>
							<td>01:21 AM</td>
							<td>60</td>
							<td>5</td>
							<td>Successful</td>
						</tr>
						<tr>
							<td>ID001</td>
							<td>Template Name</td>
							<td>23/05/2020</td>
							<td>23:50 PM</td>
							<td>01:21 AM</td>
							<td>60</td>
							<td>5</td>
							<td>Successful</td>
						</tr>
						<tr>
							<td>ID001</td>
							<td>Template Name</td>
							<td>23/05/2020</td>
							<td>23:50 PM</td>
							<td>01:21 AM</td>
							<td>60</td>
							<td>5</td>
							<td>Successful</td>
						</tr>
						<tr>
							<td>ID001</td>
							<td>Template Name</td>
							<td>23/05/2020</td>
							<td>23:50 PM</td>
							<td>01:21 AM</td>
							<td>60</td>
							<td>5</td>
							<td>Successful</td>
						</tr>
					</tbody>
				</Table>
			</div>
		</div>
	);
};

ActivityComponent.propTypes = {};

export default ActivityComponent;
