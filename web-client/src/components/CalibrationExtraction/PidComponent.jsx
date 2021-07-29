import React from "react";

import { Button } from "core-components";
import { Progress } from "reactstrap";
import { Text } from "shared-components";

const PidComponent = (props) => {
  const { progressData, handleBtnClick } = props;

  const { progressStatus, deckName, progress, remainingTime, totalTime } =
    progressData;

  return (
    <div className="d-flex">
      <Progress value={progress} />
      <Text>{21} hrs</Text>
      <Text> | </Text>
      <Text>{35} hrs remaining</Text>
      <Button className="primary mx-5" onClick={handleBtnClick}>
        PID
      </Button>
    </div>
  );
};

export default React.memo(PidComponent);
