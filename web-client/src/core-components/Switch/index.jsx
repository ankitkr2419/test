import React from 'react';
import styled from 'styled-components';
import { CustomInput } from 'reactstrap';

const StyledSwitch = styled(CustomInput)``;

const Switch = (props) => {
	return <StyledSwitch type='switch' {...props} />;
};

Switch.propTypes = {};

export default Switch;
