import React from "react";
import { LineChart } from "core-components";
import { Text } from "shared-components";
import styled from "styled-components";
import PropTypes from "prop-types";
import {
  MIN_THRESHOLD,
  MAX_THRESHOLD
} from "components/Target/targetConstants";
import GraphFilters from "./GraphFilters";

const options = {
  legend: {
    display: false
  }
};

const WellGraph = ({
  data,
  experimentGraphTargetsList,
  onThresholdChangeHandler,
  toggleGraphFilterActive,
  isThresholdInvalid,
  setThresholdError,
  resetThresholdError
}) => (
  <div>
    <GraphCard>
      <LineChart data={data} options={options} />
    </GraphCard>
    <GraphFilters
      targets={experimentGraphTargetsList}
      onThresholdChangeHandler={onThresholdChangeHandler}
      toggleGraphFilterActive={toggleGraphFilterActive}
      setThresholdError={setThresholdError}
      resetThresholdError={resetThresholdError}
    />
    {isThresholdInvalid && (
      <Text Tag="p" size={14} className="text-danger px-2 mb-1">
        Threshold value should be between {MIN_THRESHOLD} - {MAX_THRESHOLD}
      </Text>
    )}
    <Text size={14} className="text-default text-center mb-0">
      Note: Click on the threshold number to change it.
    </Text>
  </div>
);

const GraphCard = styled.div`
  width: 830px;
  height: 272px;
  background: #ffffff 0% 0% no-repeat padding-box;
  border: 1px solid #707070;
  padding: 8px;
  margin: 0 0 32px 0;
`;

WellGraph.propTypes = {
  experimentGraphTargetsList: PropTypes.object.isRequired,
  onThresholdChangeHandler: PropTypes.func.isRequired,
  toggleGraphFilterActive: PropTypes.func.isRequired,
  data: PropTypes.object.isRequired
};

export default React.memo(WellGraph);
