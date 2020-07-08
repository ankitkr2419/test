
import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { UncontrolledPopover } from 'reactstrap';

const StyledPopover = styled(UncontrolledPopover)``;

const CustomPopover = (props) => {
	const { children, ...rest } = props;

	return (
		<StyledPopover  {...rest}>
			{children}
		</StyledPopover>
	);
};

CustomPopover.propTypes = {
	children: PropTypes.any,
};

CustomPopover.defaultProps = {};

export default CustomPopover;
