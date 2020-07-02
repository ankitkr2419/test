import React from 'react';
import PropTypes from 'prop-types';
import { Text } from 'shared-components';
import { FormGroup, Label, Input } from 'core-components';

const GraphFilters = ({ targets, onThresholdChangeHandler }) => (
	<div className="graph-filters d-flex mb-1">
		<Text Tag="h4" size={19} className="flex-15 title mb-0 pr-3">
      Filters
		</Text>
		<div className="d-flex flex-wrap flex-85">
			{targets.map((ele, index) => (
				<FormGroup key={ele.get('target_id')} className="d-flex flex-33 px-3">
					<Label className="flex-75 text-truncate mb-0 p-1">
						{ele.get('target_name')}
					</Label>
					<Input
						type="number"
						className="p-1"
						value={ele.get('threshold')}
						onChange={(event) => {
							onThresholdChangeHandler(event.target.value, index, ele.toJS());
						}}
					/>
					<Text Tag="span" size={12} className="floating-label">
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
