import React from 'react';
import PropTypes from "prop-types";
import styled from 'styled-components';
import SampleTarget from './SampleTarget';

const StyledSampleTargetList = styled.div`
  padding: 24px 0;
  margin: 0 0 8px;
`;

const SampleTargetList = ({className}) => {
  return (
    <StyledSampleTargetList className={className}>
      <SampleTarget title="Target 1" />
      <SampleTarget title="Target 1" />
      <SampleTarget title="Target 1" />
      <SampleTarget title="Target 1" />
      <SampleTarget title="Target 1" />
      <SampleTarget title="Target 1" />
    </StyledSampleTargetList>
  );
};

SampleTargetList.propTypes = {
	className: PropTypes.string,
};

SampleTargetList.defaultProps = {
	className: "",
};

export default SampleTargetList;