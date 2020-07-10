import React from 'react';
import PropTypes from 'prop-types';
import { Text } from 'shared-components';
import { FormGroup, Label, Input } from 'core-components';
import GraphLegend from './GraphLegend';

const GraphFilters = ({
	targets,
	onThresholdChangeHandler,
	toggleGraphFilterActive,
}) => (
	<div className='graph-filters d-flex mb-1'>
		<Text Tag='h4' size={19} className='flex-10 title mb-0 pr-3'>
			Filters
		</Text>
		<div className='d-flex flex-wrap flex-90'>
			{targets.map((ele, index) => (
				<FormGroup
					key={ele.get('target_id')}
					className={`d-flex flex-33 px-2 ${
						ele.get('isActive') ? 'active' : ''
					}`}
				>
					<Label
						className={`flex-60 mb-0 p-1 ${
							ele.get('isActive') ? 'active' : ''
						}`}
						onClick={() => {
							toggleGraphFilterActive(index, ele.get('isActive'));
						}}
					>
						<GraphLegend color={ele.get('lineColor')} />
						<Text
							Tag='span'
							className='flex-100 text-truncate font-weight-bold'
						>
							{ele.get('target_name')}
						</Text>
					</Label>
					<Input
						type='number'
						className={`p-1 ${ele.get('isActive') ? 'active' : ''}`}
						value={ele.get('threshold')}
						onChange={(event) => {
							onThresholdChangeHandler(event.target.value, index, ele.toJS());
						}}
					/>
					<Text Tag='span' size={12} className='floating-label'>
						Enter Threshold
					</Text>
				</FormGroup>
			))}
		</div>
	</div>
);

GraphFilters.propTypes = {
	className: PropTypes.string,
};

GraphFilters.defaultProps = {
	className: '',
};

export default GraphFilters;
