import React, { useEffect } from "react";

import { Button } from "core-components";
import { Progress, Spinner } from "reactstrap";
import { HomingModal, Icon, Text } from "shared-components";
import { showHomingModal as showHomingModalAction } from "action-creators/homingActionCreators";
import { PID_STATUS } from "appConstants";
import { useDispatch } from "react-redux";

const PidProgressComponent = (props) => {
  const { pidStatus, handleBtnClick } = props;

  const dispatch = useDispatch();

  // if progress is aborted then open homing modal
  useEffect(() => {
    if (pidStatus === PID_STATUS.aborted) {
      dispatch(showHomingModalAction());
    }
  }, [pidStatus]);

  return (
    <>
      <HomingModal />
      <div className="d-flex align-items-center mr-3">
        <Button
          color={pidStatus === PID_STATUS.running ? "secondary" : "primary"}
          onClick={handleBtnClick}
        >
          {pidStatus === PID_STATUS.running ? (
            <div className="d-flex">
              <Spinner size="sm" />
              <Text className="m-auto">Abort</Text>
            </div>
          ) : (
            "Start PID"
          )}
        </Button>
      </div>
    </>
  );
};

export default React.memo(PidProgressComponent);
