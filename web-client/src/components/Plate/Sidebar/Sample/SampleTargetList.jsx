import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import SampleTarget from './SampleTarget';

const StyledSampleTargetList = styled.div`
  padding: 24px 0;
  margin: 0 0 8px;
`;

const SampleTargetList = ({ list, onCrossClickHandler }) => (
	<StyledSampleTargetList>
		{list.map((ele, index) => (
			<SampleTarget
				key={ele.get('target_id')}
				onClickHandler={() => onCrossClickHandler(ele.get('target_id'))}
				label={ele.get('target_id')}
			/>
		))}
	</StyledSampleTargetList>
);

SampleTargetList.propTypes = {
	list: PropTypes.object.isRequired,
	onCrossClickHandler: PropTypes.func.isRequired,
};

export default SampleTargetList;
