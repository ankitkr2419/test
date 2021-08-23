import React from "react";
import PropTypes from "prop-types";

import { LineChart } from "core-components";
import { options } from "./GraphOptions";
import Filters from "./Filters";

import { GraphCard } from "./GraphCard";

const AnalyseDataGraphComponent = (props) => {
  let { data, targetOptions, selectedTarget, onTargetChanged } = props;
  return (
    <div>
      <GraphCard>
        <LineChart data={data} options={options} />
      </GraphCard>
      <Filters
        targetOptions={targetOptions}
        selectedTarget={selectedTarget}
        onTargetChanged={onTargetChanged}
      />
    </div>
  );
};

AnalyseDataGraphComponent.propTypes = {
  data: PropTypes.object.isRequired,
  targetOptions: PropTypes.array.isRequired,
  selectedTarget: PropTypes.object.isRequired,
  onTargetChanged: PropTypes.func.isRequired,
};

export default React.memo(AnalyseDataGraphComponent);
