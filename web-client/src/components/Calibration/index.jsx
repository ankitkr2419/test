import React, { useEffect, useState } from "react";
import {
  Button,
  Form,
  FormGroup,
  FormError,
  Input,
  Label,
  Card,
  CardBody,
  Row,
  Col,
} from "core-components";
import { Text } from "shared-components";

import {
  constants,
  isValidRoomTemp,
  isValidHomingTime,
  isValidNoOfHomingCycle,
  isValidCycleTime,
} from "./helper";
import { useCallback } from "react";

const CalibrationComponent = (props) => {
  let { configs, saveBtnClickHandler } = props;

  const [roomTemperature, setRoomTemperature] = useState(0);
  const [homingTime, setHomingTime] = useState(0);
  const [noOfHomingCycles, setNoOfHomingCycles] = useState(0);
  const [cycleTime, setCycleTime] = useState(0);

  //isInvalidData
  const [isInvalidRoomTemp, setIsInvalidRoomTemp] = useState(false);
  const [isInvalidHomingTime, setIsInvalidHomingTime] = useState(false);
  const [isInvalidNoOfHomingCycles, setIsInvalidNoOfHomingCycles] =
    useState(false);
  const [isInvalidCycleTime, setIsInvalidCycleTime] = useState(false);

  //store new data in local state
  useEffect(() => {
    if (configs?.room_temperature) {
      setRoomTemperature(configs.room_temperature);
    }
    if (configs?.homing_time) {
      setHomingTime(configs.homing_time);
    }
    if (configs?.no_of_homing_cycles) {
      setNoOfHomingCycles(configs.no_of_homing_cycles);
    }
    if (configs?.cycle_time) {
      setCycleTime(configs.cycle_time);
    }
  }, [configs]);

  //validations and api call
  const onSubmit = (e) => {
    e.preventDefault();
    if (
      roomTemperature !== null &&
      homingTime !== null &&
      noOfHomingCycles !== null &&
      cycleTime !== null
    ) {
      saveBtnClickHandler({
        roomTemperature,
        homingTime,
        noOfHomingCycles,
        cycleTime,
      });
    }
  };

  const blurHandlerRoomTemp = useCallback(
    (value) => {
      if (isValidRoomTemp(parseInt(value)) === false) {
        setIsInvalidRoomTemp(true);
      }
    },
    [setIsInvalidRoomTemp]
  );

  const blurHandlerHomingTime = useCallback(
    (value) => {
      if (isValidHomingTime(parseInt(value)) === false) {
        setIsInvalidHomingTime(true);
      }
    },
    [setIsInvalidHomingTime]
  );

  const blurHandlerNoOfHomingCycles = useCallback(
    (value) => {
      if (isValidNoOfHomingCycle(parseInt(value)) === false) {
        setIsInvalidNoOfHomingCycles(true);
      }
    },
    [setIsInvalidNoOfHomingCycles]
  );

  const blurHandlerCycleTime = useCallback(
    (value) => {
      if (isValidCycleTime(parseInt(value)) === false) {
        setIsInvalidCycleTime(true);
      }
    },
    [setIsInvalidCycleTime]
  );

  return (
    <div className="calibration-content px-5">
      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          <Form onSubmit={onSubmit}>
            <Row>
              <Col md={6}>
                <FormGroup>
                  <Label for="username">Room Temperature</Label>
                  <Input
                    type="number"
                    name="roomTemperature"
                    id="roomTemperature"
                    placeholder={`${constants.ROOM_TEMPERATURE.min} - ${constants.ROOM_TEMPERATURE.max}`}
                    value={roomTemperature}
                    onChange={(event) => {
                      setRoomTemperature(parseInt(event.target.value));
                    }}
                    onBlur={(event) =>
                      blurHandlerRoomTemp(parseInt(event.target.value))
                    }
                    onFocus={() => setIsInvalidRoomTemp(false)}
                  />
                  {(isInvalidRoomTemp || roomTemperature == null) && (
                    <div className="flex-70">
                      <Text Tag="p" size={14} className="text-danger">
                        {`It should be between ${constants.ROOM_TEMPERATURE.min} - ${constants.ROOM_TEMPERATURE.max}`}
                      </Text>
                    </div>
                  )}
                </FormGroup>
              </Col>

              <Col md={6}>
                <FormGroup>
                  <Label for="username">Homing Time</Label>
                  <Input
                    type="number"
                    name="homingTime"
                    id="homingTime"
                    placeholder={`${constants.HOMING_TIME.min} - ${constants.HOMING_TIME.max}`}
                    value={homingTime}
                    onChange={(event) => {
                      setHomingTime(parseInt(event.target.value));
                    }}
                    onBlur={(event) =>
                      blurHandlerHomingTime(parseInt(event.target.value))
                    }
                    onFocus={() => setIsInvalidHomingTime(false)}
                  />
                  {(isInvalidHomingTime || homingTime == null) && (
                    <div className="flex-70">
                      <Text Tag="p" size={14} className="text-danger">
                        {`It should be between ${constants.HOMING_TIME.min} - ${constants.HOMING_TIME.max}`}
                      </Text>
                    </div>
                  )}
                </FormGroup>
              </Col>

              <Col md={6}>
                <FormGroup>
                  <Label for="username">No. Of Homing Cycles</Label>
                  <Input
                    type="number"
                    name="noOfHomingCycles"
                    id="noOfHomingCycles"
                    placeholder={`${constants.NO_OF_HOMING_CYCLE.min} - ${constants.NO_OF_HOMING_CYCLE.max}`}
                    value={noOfHomingCycles}
                    onChange={(event) => {
                      setNoOfHomingCycles(parseInt(event.target.value));
                    }}
                    onBlur={(event) =>
                      blurHandlerNoOfHomingCycles(parseInt(event.target.value))
                    }
                    onFocus={() => setIsInvalidNoOfHomingCycles(false)}
                  />
                  {(isInvalidNoOfHomingCycles || noOfHomingCycles == null) && (
                    <div className="flex-70">
                      <Text Tag="p" size={14} className="text-danger">
                        {`It should be between ${constants.NO_OF_HOMING_CYCLE.min} - ${constants.NO_OF_HOMING_CYCLE.max}`}
                      </Text>
                    </div>
                  )}
                </FormGroup>
              </Col>
              
              <Col md={6}>
                <FormGroup>
                  <Label for="username">Cycle Time</Label>
                  <Input
                    type="number"
                    name="cycleTime"
                    id="cycleTime"
                    placeholder={`${constants.CYCLE_TIME.min} - ${constants.CYCLE_TIME.max}`}
                    value={cycleTime}
                    onChange={(event) => {
                      setCycleTime(parseInt(event.target.value));
                    }}
                    onBlur={(event) =>
                      blurHandlerCycleTime(parseInt(event.target.value))
                    }
                    onFocus={() => setIsInvalidCycleTime(false)}
                  />
                  {(isInvalidCycleTime || cycleTime == null) && (
                    <div className="flex-70">
                      <Text Tag="p" size={14} className="text-danger">
                        {`It should be between ${constants.CYCLE_TIME.min} - ${constants.CYCLE_TIME.max}`}
                      </Text>
                    </div>
                  )}
                </FormGroup>
              </Col>
            </Row>
            <div className="text-right pt-4 pb-1 mb-3">
              <Button
                color="primary"
                disabled={
                  roomTemperature == null ||
                  homingTime == null ||
                  noOfHomingCycles == null ||
                  cycleTime == null ||
                  isInvalidRoomTemp ||
                  isInvalidHomingTime ||
                  isInvalidNoOfHomingCycles ||
                  isInvalidCycleTime
                }
              >
                Save
              </Button>
            </div>
          </Form>
        </CardBody>
      </Card>
    </div>
  );
};

export default React.memo(CalibrationComponent);
