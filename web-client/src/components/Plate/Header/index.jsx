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
  } = props;

  const { templateName } = experimentTemplate;
  const { start_time, end_time, well_count } = experimentDetails.toJS();

  return (
    <StyledSubHeader className="plate-subheader d-flex flex-column">
      <div className="d-flex align-items-center mb-auto">
        <Text Tag="h6" className="text-capitalize mb-0 mr-auto">
          {templateName}
        </Text>
        <div className="d-flex align-items-center">
          <Text className="mb-0">Cycle Count - x</Text>
          <Text className="mb-0 mx-2">|</Text>
          <Text className="mb-0">Current temperature - x</Text>
          <Text className="mb-0 mx-2">|</Text>
          <Text className="mb-0">Lid temperature - x</Text>
        </div>
        {/* <TemplatePopover name={templateName} className="ml-auto" /> */}
      </div>
      <div className="d-flex align-items-center">
        <div className="progress-wrapper d-flex align-items-center">
          <div className="d-flex align-items-center flex-100 mr-3">
            <Progress value={50} className="experiment-progress w-100" />
          </div>
          <div className="d-flex align-items-center">
            <Icon size={20} name="timer" className="text-primary" />
            <div className="time-wrapper d-flex align-items-center">
              <Text>1 Hr</Text>
              <div className="separator"></div>
              <Text>8 min</Text>
              <Text Tag="span">remaining</Text>
            </div>
          </div>
        </div>
        {isExperimentSucceeded === true && (
          <div className="d-flex align-items-center ml-auto">
            <Text Tag="h6" className="mb-0 ml-5">
              {formatDate(start_time)}
            </Text>
            <Text Tag="h6" className="mb-0 ml-3">
              {`${formatTime(start_time)} to ${formatTime(end_time)}`}
            </Text>
            <Text Tag="h6" className="mb-0 ml-3">
              No. of wells - {well_count}
            </Text>
          </div>
        )}
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
