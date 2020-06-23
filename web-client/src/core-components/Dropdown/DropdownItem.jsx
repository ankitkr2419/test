import React from 'react';
import styled from 'styled-components';
import { DropdownItem } from 'reactstrap';

const StyledDropdownItem = styled(DropdownItem)`
  font-size: 18px;
  line-height: 22px;
  color: #707070;
  
  :active {
    background-color: white;
    color: #707070;
    border: none;
  }
  :focus, :hover {
    border-radius: 0 0 24px 24px;
  }
`;

const CustomDropdownItem = props => <StyledDropdownItem {...props} />;

CustomDropdownItem.propTypes = {};

export default CustomDropdownItem;
