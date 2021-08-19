import React from "react";
import PropTypes from "prop-types";
import { Select } from "core-components";
import { Text } from "shared-components";

const Filters = (props) => {
  let { targetOptions, selectedTarget, onTargetChanged } = props;
  return (
    <>
      {/** Target selector */}
      <div className="graph-filters d-flex">
        <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
          Target
        </Text>
        <div style={{ width: "350px" }}>
          <Select
            placeholder="Select Target"
            className="mb-4"
            options={targetOptions}
            value={selectedTarget}
            onChange={onTargetChanged}
            // isDisabled={isDisabled}
          />
        </div>
      </div>
      <div className="range-filters d-flex">
        <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3"></Text>
      </div>
    </>
  );
};

Filters.propTypes = {
  targetOptions: PropTypes.array.isRequired,
  selectedTarget: PropTypes.object.isRequired,
  onTargetChanged: PropTypes.func.isRequired,
};

export default React.memo(Filters);
