import React from "react";

import {
  Button,
  Form,
  FormGroup,
  Row,
  Card,
  CardBody,
  Col,
  Input,
  Label,
  Select,
  CheckBox,
} from "core-components";
import { Center, Text } from "shared-components";

const DyeCalibration = (props) => {
  let { handleDyeCalibrationButton } = props;
  {/**TODO api integration: work in progress */}
  return (
    <Card default className="my-3">
      <CardBody>
        <Text
          Tag="h4"
          size={24}
          className="text-center text-gray text-bold mt-3 mb-4"
        >
          Dye Calibration
        </Text>

        <Form onSubmit={handleDyeCalibrationButton}>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="dye" className="font-weight-bold">
                  Dye
                </Label>
                <div>
                  <Select
                    placeholder="Select Type"
                    options={[{ label: "one", value: "one" }]}
                    value={{ label: "one", value: "one" }}
                    // onChange={(value) =>
                    //   handleOnChange("tipTubeType.value", value)
                    // }
                  />
                </div>
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="kitId" className="font-weight-bold">
                  ID
                </Label>
                <Input
                  type="number"
                  name="kitId"
                  id="kitId"
                  placeholder="Type Kit Id"
                  value={"1234"}
                  // onChange={(event) =>
                  //   handleOnChange(
                  //     "tipTubeId.value",
                  //     parseInt(event.target.value)
                  //   )
                  // }
                  // onBlur={(e) => handleBlurID(parseInt(e.target.value))}
                  // onFocus={() => handleOnChange("tipTubeId.isInvalid", false)}
                />
                {/* {tipTubeId.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      ID should be {MIN_TIPTUBE_ID} - {MAX_TIPTUBE_ID}
                    </Text>
                  </div>
                )} */}
              </FormGroup>
            </Col>
          </Row>
          <Center className="text-center pt-3">
            <Button /* disabled={}*/ color="primary">Start</Button>
          </Center>
        </Form>
      </CardBody>
    </Card>
  );
};

export default React.memo(DyeCalibration);
