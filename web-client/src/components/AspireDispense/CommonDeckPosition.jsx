import React from "react";

import { CreatableSelect, FormGroup, Label, Select } from "core-components";
import { ASPIRE_DISPENSE_DECK_POS_OPTNS } from "appConstants";
import { setFormikField } from "./helpers";

const CommonDeckPosition = (props) => {
  const { formik, isAspire, currentTab } = props;
  const formikValue = formik.values[isAspire ? "aspire" : "dispense"];
  const formikDeckPos = formikValue.deckPosition;
  const valueToSet = { value: formikDeckPos, label: formikDeckPos };
  const deckPositionValue = formikDeckPos === "" ? "" : valueToSet;

  const handleOnChange = (e) => {
    let value = "";
    if (e) {
      value = e.value;
    }
    setFormikField(formik, isAspire, currentTab, "deckPosition", value);
  };

  return (
    <FormGroup className="d-flex align-items-center mb-4">
      <Label for="deck-position" className="px-0 label-name">
        Deck position
      </Label>
      <div className="d-flex flex-column input-field">
        <CreatableSelect
          isClearable
          placeholder="Select Tip"
          className=""
          size="md"
          options={ASPIRE_DISPENSE_DECK_POS_OPTNS}
          value={deckPositionValue}
          onChange={(e) => handleOnChange(e)}
        />
      </div>
    </FormGroup>
  );
};

export default React.memo(CommonDeckPosition);
