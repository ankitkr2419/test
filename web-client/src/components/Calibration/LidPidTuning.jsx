import React from "react";
import { Spinner } from "reactstrap";

import { Button, Card, CardBody } from "core-components";
import { PID_STATUS } from "appConstants";
import { Text } from "shared-components";

const LidPidTuning = (props) => {
  let { lidPidStatus, handleButtonClick } = props;

  let lidPidRunning =
    lidPidStatus === PID_STATUS.running ||
    lidPidStatus === PID_STATUS.progressing;
  let lidPidCompleted = lidPidStatus === PID_STATUS.progressComplete;

  return (
    <Card default className="my-3">
      <CardBody>
        <div className="d-flex align-items-center mr-3">
          <Text
            Tag="h4"
            size={24}
            className="text-left text-gray text-bold mt-3 mb-4"
          >
            Lid PID Tuning
          </Text>
          <Button
            color={lidPidRunning ? "secondary" : "primary"}
            onClick={handleButtonClick}
            className="ml-3"
          >
            {lidPidRunning === true && lidPidCompleted === false ? (
              <div className="d-flex">
                <Spinner size="sm" />
                <Text className="ml-5">Abort</Text>
              </div>
            ) : (
              "Start"
            )}
          </Button>
        </div>
      </CardBody>
    </Card>
  );
};

export default React.memo(LidPidTuning);
