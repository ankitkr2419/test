import React from "react";
import { LineChart } from "core-components";
import { Text } from "shared-components";
import styled from "styled-components";
import PropTypes from "prop-types";
import {
  MIN_THRESHOLD,
  MAX_THRESHOLD,
} from "components/Target/targetConstants";
import GraphFilters from "./GraphFilters";
import GraphRange from "./GraphRange";

const WellGraph = (props) => {
  const {
    data,
    headerData,
    experimentGraphTargetsList,
    onThresholdChangeHandler,
    toggleGraphFilterActive,
    isThresholdInvalid,
    setThresholdError,
    resetThresholdError,
    handleRangeChangeBtn,
    handleResetBtn,
    isInsidePreviewModal,
    isExpanded,
    options,
    isDataFromAPI,
  } = props;

  return (
    <div>
      <GraphCard>
        <LineChart data={data} options={options} redraw={isDataFromAPI} />
      </GraphCard>
      <GraphFilters
        targets={experimentGraphTargetsList}
        onThresholdChangeHandler={onThresholdChangeHandler}
        toggleGraphFilterActive={toggleGraphFilterActive}
        setThresholdError={setThresholdError}
        resetThresholdError={resetThresholdError}
      />

      {isInsidePreviewModal === false && (
        <GraphRange
          handleRangeChangeBtn={handleRangeChangeBtn}
          handleResetBtn={handleResetBtn}
          headerData={headerData}
          data={data}
          isExpanded={isExpanded}
        />
      )}

      {isThresholdInvalid && (
        <Text Tag="p" size={14} className="text-danger px-2 mb-1">
          Threshold value should be between {MIN_THRESHOLD} - {MAX_THRESHOLD}
        </Text>
      )}
    </div>
  );
};

const GraphCard = styled.div`
  width: 960px;
  height: 326px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 16px 0;
`;

WellGraph.propTypes = {
  experimentGraphTargetsList: PropTypes.object.isRequired,
  onThresholdChangeHandler: PropTypes.func.isRequired,
  toggleGraphFilterActive: PropTypes.func.isRequired,
  data: PropTypes.object.isRequired,
  isThresholdInvalid: PropTypes.bool,
};

export default React.memo(WellGraph);
