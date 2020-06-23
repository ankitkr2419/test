import React from 'react';
import styled from 'styled-components';
import { DropdownItem } from 'reactstrap';

const StyledDropdownItem = styled(DropdownItem)`
  font-size: 18px;
  line-height: 22px;
	color: #707070;
`;

const CustomDropdownItem = (props) => {
  return <StyledDropdownItem {...props} />;
};

CustomDropdownItem.propTypes = {};

export default CustomDropdownItem;