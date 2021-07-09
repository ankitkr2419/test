import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import { Text } from "shared-components";
import TemplatePopover from "components/Plate/Popover";
import { formatDate, formatTime } from "utils/helpers";

const StyledSubHeader = styled.div`
  display: flex;
  align-items: center;
  height: 40px;
  padding: 8px 16px 8px 88px;
  color: #707070;

  h6 {
    font-size: 14px;
    line-height: 1.25;
  }
`;

const SubHeader = (props) => {
  const {
    experimentTemplate,
    isExperimentSucceeded,
    experimentDetails,
    experimentId //TODO remove if not needed anymore
  } = props;

  const { templateName } = experimentTemplate;
  const { start_time, end_time, well_count } = experimentDetails.toJS();

  return (
    <StyledSubHeader className="plate-subheader">
      <Text Tag="h6" className="mb-0 mx-5">
        {templateName}
      </Text>
      {isExperimentSucceeded === true && (
        <>
          <Text Tag="h6" className="mb-0 ml-5">
            {formatDate(start_time)}
          </Text>
          <Text Tag="h6" className="mb-0 ml-3">
            {`${formatTime(start_time)} to ${formatTime(end_time)}`}
          </Text>
          <Text Tag="h6" className="mb-0 ml-5">
            No. of wells - {well_count}
          </Text>
        </>
      )}
      <TemplatePopover className="ml-auto" />
    </StyledSubHeader>
  );
};

SubHeader.propTypes = {
  experimentTemplate: PropTypes.shape({
    templateId: PropTypes.string,
    templateName: PropTypes.string
  }).isRequired,
  isExperimentSucceeded: PropTypes.bool
};

export default SubHeader;
