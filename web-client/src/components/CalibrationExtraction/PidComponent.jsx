import React from "react";

import { Button } from "core-components";
import { Progress } from "reactstrap";
import { Icon, Text } from "shared-components";
import { PID_STATUS } from "appConstants";

const PidComponent = (props) => {
  const { progressData, handleBtnClick } = props;

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
      {true && (
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
      )}
      <Button
        color={progressIsRunning ? "secondary" : "primary"}
        disabled={progressIsRunning}
        // outline={!(progressIsRunning || progressStatus === null)}
        className="mx-5"
        onClick={handleBtnClick}
      >
        {progressIsRunning ? "PID Progressing..." : "Start PID"}
      </Button>
    </div>
  );
};

export default React.memo(PidComponent);
