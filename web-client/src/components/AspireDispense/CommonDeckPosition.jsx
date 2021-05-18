import React from "react";

import { FormGroup, Label, Select } from "core-components";

const CommonDeckPosition = (props) => {
  return (
    <FormGroup className="d-flex align-items-center mb-4">
      <Label for="deck-position" className="px-0 label-name">
        Deck position
      </Label>
      <div className="d-flex flex-column input-field">
        <Select
          placeholder="Select Tip"
          className=""
          size="sm"
          //options={taskOptions}
          //value={task}
          //onChange={handleTaskChange}
        />
      </div>
    </FormGroup>
  );
};

export default React.memo(CommonDeckPosition);
