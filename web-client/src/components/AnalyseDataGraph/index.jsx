import React from "react";
import PropTypes from "prop-types";

import { LineChart } from "core-components";
import { options } from "./GraphOptions";
import Filters from "./Filters";

import { GraphCard } from "./GraphCard";

const AnalyseDataGraphComponent = (props) => {
  let {
    data,
    targetOptions,
    selectedTarget,
    onTargetChanged,
    analyseDataGraphFilters,
    isInsidePreviewModal,
    onFiltersChanged,
    onResetThresholdFilter,
    onResetBaselineFilter,
  } = props;

  return (
    <div>
      <GraphCard>
        <LineChart data={data} options={options} />
      </GraphCard>
      {isInsidePreviewModal === false && (
        <Filters
          targetOptions={targetOptions}
          selectedTarget={selectedTarget}
          onTargetChanged={onTargetChanged}
          analyseDataGraphFilters={analyseDataGraphFilters}
          onFiltersChanged={onFiltersChanged}
          onResetThresholdFilter={onResetThresholdFilter}
          onResetBaselineFilter={onResetBaselineFilter}
        />
      )}
    </div>
  );
};

AnalyseDataGraphComponent.propTypes = {
  data: PropTypes.object.isRequired,
  targetOptions: PropTypes.array.isRequired,
  selectedTarget: PropTypes.object.isRequired,
  onTargetChanged: PropTypes.func.isRequired,
  analyseDataGraphFilters: PropTypes.object.isRequired,
  onFiltersChanged: PropTypes.func.isRequired,
  onResetThresholdFilter: PropTypes.func.isRequired,
  onResetBaselineFilter: PropTypes.func.isRequired,
};

export default React.memo(AnalyseDataGraphComponent);
