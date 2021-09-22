import React, { useState } from "react";

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
} from "core-components";
import { Center, Text } from "shared-components";

const DeleteCartridgeComponent = (props) => {
  const { handleDeleteBtn } = props;

  const [cartridgeId, setCartridgeId] = useState({
    value: "0",
    isInvalid: false,
  });

  const handleBlur = (value) => {
    if (value === "") {
      setCartridgeId({ ...cartridgeId, isInvalid: true });
    }
  };

  const isDeleteBtnDisabled = () => {
    if (
      parseInt(cartridgeId.value) < 0 ||
      cartridgeId.value === "" ||
      cartridgeId.value === "0" ||
      cartridgeId.value === null ||
      cartridgeId.isInvalid
    ) {
      return true;
    }
    return false;
  };

  return (
    <Card default className="my-3 w-100">
      <CardBody>
        <Form>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="id" className="font-weight-bold ml-3">
                  Cartridge ID
                </Label>
                <Input
                  className="ml-3"
                  type="number"
                  name="id"
                  id="id"
                  value={cartridgeId.value}
                  onChange={(event) =>
                    setCartridgeId({
                      ...cartridgeId,
                      value: event.target.value,
                    })
                  }
                  onBlur={(e) => handleBlur(e.target.value)}
                  onFocus={() =>
                    setCartridgeId({ ...cartridgeId, isInvalid: false })
                  }
                />
                {cartridgeId.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Invalid ID.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col>
              <Center className="text-center pt-4">
                <Button
                  className="w-auto"
                  onClick={() => handleDeleteBtn(parseInt(cartridgeId.value))}
                  disabled={isDeleteBtnDisabled()}
                  color="primary"
                >
                  Delete Cartridge
                </Button>
              </Center>
            </Col>
          </Row>
        </Form>
      </CardBody>
    </Card>
  );
};

export default React.memo(DeleteCartridgeComponent);
