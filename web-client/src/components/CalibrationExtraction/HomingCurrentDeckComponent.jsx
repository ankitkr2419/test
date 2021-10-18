import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import { Button, Form, Card, CardBody } from "core-components";
import { Center, Text } from "shared-components";
import {
  hideHomingModal,
  showHomingModal,
} from "action-creators/homingActionCreators";

const HomingCurrentDeckComponent = (props) => {
  const dispatch = useDispatch();
  const { heading, handleHomingDeck } = props;
  const [isHomingDisabled, setHomingDisability] = useState(false);

  //enable button when homing completed
  const homingReducer = useSelector((state) => state.homingReducer);
  const { isHomingActionCompleted } = homingReducer;

  // enable btn if homing is completed
  useEffect(() => {
    if (isHomingActionCompleted === true && isHomingDisabled === true) {
      setHomingDisability(false);
    }
  }, [isHomingActionCompleted, isHomingDisabled]);

  //start homing and disable button
  const handleStartHoming = (e) => {
    e.preventDefault();
    // handleHomingDeck();
    setHomingDisability(true);
    dispatch(showHomingModal());
  };

  return (
    <Card default className="my-3">
      <CardBody>
        <Text
          Tag="h4"
          size={24}
          className="text-center text-gray text-bold mt-3 mb-4"
        >
          {heading}
        </Text>
        <Form>
          <Center className="text-center pt-3">
            <Button
              color="primary"
              disabled={isHomingDisabled}
              onClick={handleStartHoming}
            >
              Start
            </Button>
          </Center>
        </Form>
      </CardBody>
    </Card>
  );
};

export default HomingCurrentDeckComponent;
