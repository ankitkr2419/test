import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import { Progress } from "reactstrap";
import { Icon, Text } from "shared-components";
import TemplatePopover from "components/Plate/Popover";
import { formatDate, formatTime } from "utils/helpers";
import { EXPERIMENT_STATUS } from "appConstants";

const StyledSubHeader = styled.div`
  background-color: #f2f2f2;
  height: 68px;
  padding: 8px 32px;
  color: #707070;

  h6 {
    font-size: 14px;
    line-height: 1.25;
  }
`;

const SubHeader = props => {
  const {
    data,
    experimentTemplate,
    experimentStatus,
    experimentDetails,
    temperatureData
  } = props;

  const { totalCycles, progressStatus, progress, remainingTime, totalTime } =
    data;

  let lidTemperature = 0;
  let temperature = 0;
  let cycleValue = 0;

  const TEMP_ARR_LEN = temperatureData.length;
  if (TEMP_ARR_LEN > 0) {
    // temperatureData is an array, and this array gets updated everytime new data is pushed.
    // we need the latest index of array
    const { lid_temp, temp, cycle } = temperatureData[TEMP_ARR_LEN - 1];

    lidTemperature = lid_temp;
    temperature = temp;
    cycleValue = cycle;
  }

  const { templateName } = experimentTemplate;
  const { start_time, end_time, well_count } = experimentDetails.toJS();

  let totalHours = 0;
  let totalMins = 0;
  let totalSecs = 0;

  if (totalTime) {
    totalHours = totalTime.hours;
    totalMins = totalTime.minutes;
    totalSecs = totalTime.seconds;
  }

  let remainHours = 0;
  let remainMins = 0;
  let remainSecs = 0;

  if (remainingTime) {
    remainHours = remainingTime.hours;
    remainMins = remainingTime.minutes;
    remainSecs = remainingTime.seconds;
  }

  const showProgressBar = () => {
    return experimentStatus === EXPERIMENT_STATUS.running;
    // return (
    //   experimentStatus === EXPERIMENT_STATUS.success ||
    //   experimentStatus === EXPERIMENT_STATUS.running ||
    //   experimentStatus === EXPERIMENT_STATUS.stopped
    // );
  };

  return (
    <StyledSubHeader className="plate-subheader d-flex flex-column justify-content-center">
      <div className="d-flex align-items-center">
        <Text Tag="h6" className="text-capitalize mb-0 mr-auto">
          {templateName}
        </Text>
        {TEMP_ARR_LEN > 0 && (
          <div className="d-flex align-items-center">
            <Text className="mb-0">
              Cycle - {cycleValue} / {totalCycles}{" "}
            </Text>
            <Text className="mb-0 mx-2">|</Text>
            <Text className="mb-0">
              Current temperature - {temperature.toFixed(2)} °C
            </Text>
            {/* <Text className="mb-0 mx-2">|</Text>
            <Text className="mb-0">Lid temperature - {lidTemperature}</Text> */}
          </div>
        )}
        {/* <TemplatePopover name={templateName} className="ml-auto" /> */}
      </div>
      <div className="d-flex align-items-center">
        {showProgressBar() &&
          (progressStatus === EXPERIMENT_STATUS.progressing ||
          progressStatus === EXPERIMENT_STATUS.progressComplete ? (
            <div className="progress-wrapper d-flex align-items-center">
              <div className="d-flex align-items-center flex-100 mr-3">
                <Progress
                  value={progress}
                  className="experiment-progress w-100"
                />
              </div>
              <div className="d-flex align-items-center">
                <Icon size={20} name="timer" className="text-primary" />
                <div className="time-wrapper d-flex align-items-center">
                  <Text>
                    {totalHours > 0 && `${totalHours} Hr `}
                    {`${totalMins} min ${totalSecs} sec`}
                  </Text>
                  <div className="separator"></div>
                  <Text>
                    {remainHours > 0 && `${remainHours} Hr `}
                    {`${remainMins} min ${remainSecs} sec`}
                  </Text>
                  <Text Tag="span">remaining</Text>
                </div>
              </div>
            </div>
          ) : (
            experimentStatus !== EXPERIMENT_STATUS.stopped && (
              <Text className="font-weight-bold mb-0">
                Homing is in Progress...
              </Text>
            )
          ))}
        {experimentStatus === EXPERIMENT_STATUS.success && (
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
