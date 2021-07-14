import React, { useEffect, useState } from "react";
import {
  Button,
  Form,
  FormGroup,
  FormError,
  Input,
  Label,
  Row,
  Col,
} from "core-components";
const CalibrationComponent = (props) => {
  let { configs } = props;

  const [roomTemperature, setRoomTemperature] = useState(0);
  const [homingTime, setHomingTime] = useState(0);
  const [noOfHomingCycles, setNoOfHomingCycles] = useState(0);

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
  }, [configs]);

  const onSubmit = () => {
    //TODO
  };

  return (
    <div className="CalibrationPage">
      <Form onSubmit={onSubmit}>
        <Row>
          <Col>
            <FormGroup>
              <Label for="username">Room Temperature</Label>
              <Input
                type="number"
                name="roomTemperature"
                id="roomTemperature"
                placeholder="Type here"
                value={roomTemperature}
                onChange={(event) => {
                  setRoomTemperature(event.target.value);
                }}
              />
              <FormError>Incorrect Room Temperature</FormError>
            </FormGroup>
          </Col>

          <Col>
            <FormGroup>
              <Label for="username">Homing Time</Label>
              <Input
                type="number"
                name="homingTime"
                id="homingTime"
                placeholder="Type here"
                value={homingTime}
                onChange={(event) => {
                  setHomingTime(event.target.value);
                }}
              />
              <FormError>Incorrect Homing Time</FormError>
            </FormGroup>
          </Col>

          <Col>
            <FormGroup>
              <Label for="username">No. Of Homing Cycles</Label>
              <Input
                type="number"
                name="noOfHomingCycles"
                id="noOfHomingCycles"
                placeholder="Type here"
                value={noOfHomingCycles}
                onChange={(event) => {
                  setNoOfHomingCycles(event.target.value);
                }}
              />
              <FormError>Incorrect no Of Homing Cycles</FormError>
            </FormGroup>
          </Col>
        </Row>
        <div className="text-right pt-4 pb-1 mb-3">
          <Button color="primary">Save</Button>
        </div>
      </Form>
    </div>
  );
};

export default React.memo(CalibrationComponent);
