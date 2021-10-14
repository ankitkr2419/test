import React from "react";

import styled from "styled-components";
import PropTypes from "prop-types";
import { Text, Icon } from "shared-components";
import { getIconName } from "shared-components/DeckCard/helpers";

const ProecssRemainingBox = styled.div`
  position: absolute;
  height: 38px;
  top: -40px;
  width: 100%;
  left: 0;
  background: #fff;
  border-radius: 12px 12px 0 0;
  box-shadow: -3px 3px 6px rgb(0 0 0 / 16%);
  .process-counter {
    font-size: 1.125rem;
    line-height: 1.313rem;
  }
  .total-process {
    font-size: 0.875rem;
    line-height: 1rem;
  }
  .process-remain-label {
    padding: 0 0.25rem !important;
  }
`;
const ProcessRemaining = (props) => {
  const { processName, processType, processNumber, processTotal } = props;

  const showProcessDetails = () => {
    if (
      (processNumber || processNumber === 0) &&
      (processTotal || processTotal === 0)
    ) {
      return true;
    }
    return false;
  };

  return (
    <ProecssRemainingBox>
      <Text Tag="label" className="d-flex align-items-center px-3 py-2 mb-0">
        <Icon
          name={getIconName(processType)}
          size={19}
          className="text-primary"
        />
        {showProcessDetails() && (
          <Text
            Tag="span"
            className="process-remain-label font-weight-bold ml-2"
          >
            <Text Tag="span" className="process-counter font-weight-bold">
              {" "}
              {processNumber}
            </Text>
            <Text Tag="span" className="total-process ml-1">
              /{processTotal}{" "}
            </Text>
          </Text>
        )}
        <Text Tag="span" className="min-label ml-2 font-weight-bold">
          {processName}
        </Text>
      </Text>
    </ProecssRemainingBox>
  );
};

ProcessRemaining.propTypes = {
  isUserLoggedIn: PropTypes.bool,
  processName: PropTypes.string,
};

ProcessRemaining.defaultProps = {
  isUserLoggedIn: false,
  processName: "Process Done",
};

export default ProcessRemaining;
