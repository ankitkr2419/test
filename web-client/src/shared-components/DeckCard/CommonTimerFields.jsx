import React from "react";

import { Text, Icon } from "shared-components";

const CommonTimerFields = (props) => {
  const { recipeName, remainingTime, totalTime } = props;

  return (
    <div className="d-none1">
      <Text Tag="h5" size={18} className="mb-2 font-weight-bold recipe-name">
        {recipeName}
      </Text>
      <Text Tag="label" className="mb-1 d-flex align-items-center">
        <Icon name="timer" size={19} className="text-primary" />
        <Text Tag="span" className="hour-label font-weight-bold ml-2">
          {" "}
          {/* {hours} Hr{" "} {mins} min {secs} sec */}
          {totalTime}
        </Text>
        <Text Tag="span" className="min-label ml-2 font-weight-bold">
          {/* {mins} min {secs} sec */}
          {remainingTime}
        </Text>
        <Text Tag="span" className="ml-1">
          remaining
        </Text>
      </Text>
    </div>
  );
};

export default React.memo(CommonTimerFields);
