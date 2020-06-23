import React from 'react';
import styled from 'styled-components';
import { Dropdown } from 'reactstrap';

const StyledDropdown = styled(Dropdown)``;

const CustomDropdown = (props) => {
  return <StyledDropdown {...props} />;
};

CustomDropdown.propTypes = {};

export default CustomDropdown;