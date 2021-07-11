import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import { Progress } from "reactstrap";
import { Icon, Text } from "shared-components";
import TemplatePopover from "components/Plate/Popover";
import { formatDate, formatTime } from "utils/helpers";

const StyledSubHeader = styled.div`
  background-color: #f2f2f2;
  height: 68px;
  padding: 10px 16px 10px 76px;
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
    <StyledSubHeader className="plate-subheader d-flex flex-column">
      <div className="progress-wrapper d-flex align-items-center mb-auto">
        <div className="d-flex align-items-center">
          <Icon size={20} name="timer" className="text-primary" />
          <div className="time-wrapper d-flex align-items-center">
            <Text>1 Hr</Text>
            <div className="separator"></div>
            <Text>8 min</Text>
            <Text Tag="span">remaining</Text>
          </div>
        </div>
        <div className="d-flex align-items-center flex-100 ml-3">
          <Progress value={50} className="experiment-progress w-100" />
        </div>
      </div>
      <div className="d-flex align-items-center">
        <Text Tag="h6" className="text-capitalize mb-0">
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
        <TemplatePopover name={templateName} className="ml-auto" />
      </div>
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
