import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { PopoverBody } from 'reactstrap';

const StyledPopoverBody = styled(PopoverBody)``;

const CustomPopoverBody = (props) => {

const { children, ...rest } = props;

return (
  <StyledPopoverBody {...rest}>
    {children}
    </StyledPopoverBody>
  );
};

CustomPopoverBody.propTypes = {
	children: PropTypes.any,
};

CustomPopoverBody.defaultProps = {};

export default CustomPopoverBody;