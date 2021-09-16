import React from "react";

import { Button, Card, CardBody } from "core-components";
import { Text } from "shared-components";

const TECOperations = (props) => {
  let { handleResetTEC, handleAutoTuneTEC } = props;

  return (
    <Card default className="my-3">
      <CardBody>
        <div className="d-flex align-items-center mr-3">
          <Text
            Tag="h4"
            size={24}
            className="text-left text-gray text-bold mt-3 mb-4"
          >
            TEC Operations
          </Text>
          <Button color={"secondary"} className="ml-3" onClick={handleResetTEC}>
            Reset TEC
          </Button>
          <Button
            color={"secondary"}
            className="ml-3"
            onClick={handleAutoTuneTEC}
          >
            Auto Tune TEC
          </Button>
        </div>
      </CardBody>
    </Card>
  );
};

export default React.memo(TECOperations);
