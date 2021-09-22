import React from "react";

import { Button } from "core-components";
import { Progress, Spinner } from "reactstrap";
import { Icon, Text } from "shared-components";
import { PID_STATUS } from "appConstants";

const PidProgressComponent = (props) => {
  const { pidStatus, abortStatus, handleBtnClick } = props;
  let abortBtnLabel = "Abort";
  if (abortStatus === "aborting" || abortStatus === "aborted") {
    abortBtnLabel = "Aborting...";
  }

  return (
    <div className="d-flex align-items-center mr-3">
      <Button
        color={pidStatus === PID_STATUS.running ? "secondary" : "primary"}
        onClick={handleBtnClick}
      >
        {pidStatus === PID_STATUS.running ? (
          <div className="d-flex">
            <Spinner size="sm" />
            <Text className="m-auto">{abortBtnLabel}</Text>
          </div>
        ) : (
          "Start PID"
        )}
      </Button>
    </div>
  );
};

export default React.memo(PidProgressComponent);
