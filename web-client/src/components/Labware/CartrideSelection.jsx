import React from 'react';

import { 
  FormGroup, 
	Label, 
  FormError,
	Select
} from 'core-components';
import {

} from 'shared-components';

import styled from 'styled-components';

const CommmonCartridgeFields = styled.div`
  .label-name{
    width:9.125rem;
  }
  .input-field{
    width:14.125rem;
    height:2.25rem;
    .height-icon-box{
      position:absolute;
      top:3px;
      right:0.75rem;
    }
  }
`;

const CartridgeSelection = (props) => {
  const {
    handleOptionChange,
    options,
    value
  } = props;
	return (
		<>
      <CommmonCartridgeFields>
        <FormGroup>
          <Label className="mb-3 font-weight-bold px-0">
            Select Cartridge
          </Label>
        </FormGroup>
        <FormGroup className="d-flex align-items-center mb-2">
          <Label className="px-0 label-name">
          Cartridge Type
          </Label>
          <div className="d-flex flex-column input-field position-relative">
            <Select
              placeholder="Select Option"
              className=""
              size="sm"
              options={options}
              value={value}
              onChange={handleOptionChange}
            />
            <FormError>Incorrect Cartridge Type</FormError>
          </div>
        </FormGroup>
    </CommmonCartridgeFields>
		</>
	);
};

CartridgeSelection.propTypes = {};

export default CartridgeSelection;
