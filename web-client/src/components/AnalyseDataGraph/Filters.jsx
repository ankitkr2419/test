import React from "react";
import PropTypes from "prop-types";
import { Select } from "core-components";
import { Text } from "shared-components";

const Filters = (props) => {
  let { dyeOptions, selectedDye, onDyeChanged } = props;
  return (
    <>
      {/** dye selector */}
      <div className="graph-filters d-flex">
        <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
          Dye
        </Text>
        <div style={{ width: "350px" }}>
          <Select
            placeholder="Select Dye"
            className="mb-4"
            options={dyeOptions}
            value={selectedDye}
            onChange={onDyeChanged}
            // isDisabled={isDisabled}
          />
        </div>
      </div>
      <div className="range-filters d-flex">
        <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
        </Text>
      </div>
    </>
  );
};

Filters.propTypes = {
  dyeOptions: PropTypes.array.isRequired,
  selectedDye: PropTypes.string.isRequired,
  onDyeChanged: PropTypes.func.isRequired,
};

export default React.memo(Filters);
