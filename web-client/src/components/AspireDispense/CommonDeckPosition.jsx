import React from "react";

import { FormGroup, Label, Select } from "core-components";
import { ASPIRE_DISPENSE_DECK_POS_OPTNS } from "appConstants";

const CommonDeckPosition = (props) => {
  const { formik, type } = props;
  return (
    <FormGroup className="d-flex align-items-center mb-4">
      <Label for="deck-position" className="px-0 label-name">
        Deck position
      </Label>
      <div className="d-flex flex-column input-field">
        <Select
          placeholder="Select Tip"
          className=""
          size="md"
          options={ASPIRE_DISPENSE_DECK_POS_OPTNS}
          onChange={(e) => {
            formik.setFieldValue(`${type}.deckPosition`, e.value);
          }}
        />
      </div>
    </FormGroup>
  );
};

export default React.memo(CommonDeckPosition);
