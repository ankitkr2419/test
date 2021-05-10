import React from 'react';

import { 
  FormGroup, 
	Label, 
  FormError,
	Select
} from 'core-components';
import { CommmonCartridgeFields } from './CommmonCartridgeFields';

const CartridgeSelection = (props) => {
  const {
    handleOptionChange,
    options
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
