import React from "react";
import { useHistory } from "react-router";

import { Button, Card, CardBody } from "core-components";
import { Text } from "shared-components";
import { ROOT_URL_PATH, ROUTES } from "appConstants";

const ShakerAndHeaterComponent = (props) => {
  const history = useHistory();

  return (
    <Card default className="my-3 w-100">
      <CardBody>
        <Text
          Tag="h4"
          size={24}
          className="text-center text-gray text-bold mt-3 mb-4"
        >
          {"Shaker & Heater"}
        </Text>

        {/* {Shaker Button} */}
        <Button
          className="mx-3"
          color={"primary"}
          onClick={() =>
            history.push(`${ROOT_URL_PATH}${ROUTES.calibration}/shaker`)
          }
        >
          Shaker
        </Button>

        {/* {Heater Button} */}
        <Button
          className="mx-3"
          color={"primary"}
          onClick={() =>
            history.push(`${ROOT_URL_PATH}${ROUTES.calibration}/heater`)
          }
        >
          Heater
        </Button>
      </CardBody>
    </Card>
  );
};

export default React.memo(ShakerAndHeaterComponent);
