import React from "react";

import { Button } from "core-components";
import { Progress, Spinner } from "reactstrap";
import { Icon, Text } from "shared-components";
import { PID_STATUS } from "appConstants";

const PidComponent = (props) => {
  const { pidStatus, progressData, handleBtnClick } = props;

  const { progressStatus, progress, remainingTime, totalTime } = progressData;

  const totalHours = totalTime?.hours || 0;
  const totalMinutes = totalTime?.minutes || 0;
  const totalSeconds = totalTime?.seconds || 0;

  const remainingHours = remainingTime?.hours || 0;
  const remainingMinutes = remainingTime?.minutes || 0;
  const remainingSeconds = remainingTime?.seconds || 0;

  let progressIsRunning =
    progressStatus === PID_STATUS.progressing ||
    progressStatus === PID_STATUS.progressComplete;

  return (
    <div className="d-flex align-items-center">
      {/* {progressStatus === PID_STATUS.progressing && (
        <Text>PID Progressing...</Text>
      )}
      {progressStatus === PID_STATUS.progressComplete && (
        <Text>PID Progress Complete</Text>
      )} */}
      {/* {true && (
        <div className="progress-wrapper d-flex align-items-center">
          <div className="d-flex align-items-center flex-100 mr-3">
            <Progress value={progress} className="experiment-progress w-100" />
          </div>
          <div className="d-flex align-items-center">
            <Icon size={20} name="timer" className="text-primary" />
            <div className="time-wrapper d-flex align-items-center">
              <Text>
                {totalHours > 0 && `${totalHours} Hr `}
                {`${totalMinutes} min ${totalSeconds} sec`}
              </Text>
              <div className="separator"></div>
              <Text>
                {remainingHours > 0 && `${remainingHours} Hr `}
                {`${remainingMinutes} min ${remainingSeconds} sec`}
              </Text>
              <Text Tag="span">remaining</Text>
            </div>
          </div>
        </div>
      )} */}
      <Button
        color={pidStatus === PID_STATUS.running ? "secondary" : "primary"}
        // disabled={progressStatus === PID_STATUS.progressComplete}
        onClick={handleBtnClick}
      >
        {pidStatus === PID_STATUS.running ? (
          <div className="d-flex">
            <Spinner size="sm" />
            <Text className="ml-5">Abort</Text>
          </div>
        ) : (
          "Start PID"
        )}
      </Button>
    </div>
  );
};

export default React.memo(PidComponent);
