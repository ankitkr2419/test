import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import SampleTarget from './SampleTarget';

const StyledSampleTargetList = styled.div`
	padding: 24px 0;
	margin: 0 0 8px;
`;

const SampleTargetList = ({ list, onTargetClickHandler, isDisabled }) => (
	<StyledSampleTargetList>
		{list.map((ele, index) => (
			<SampleTarget
				key={ele.get('target_id')}
				onClickHandler={() => onTargetClickHandler(index)}
				label={ele.get('target_name')}
				isSelected={ele.get('isSelected')}
				isDisabled={isDisabled}
			/>
		))}
	</StyledSampleTargetList>
);

SampleTargetList.propTypes = {
	list: PropTypes.object.isRequired,
	onTargetClickHandler: PropTypes.func.isRequired,
	isDisabled: PropTypes.bool,
};

export default SampleTargetList;
