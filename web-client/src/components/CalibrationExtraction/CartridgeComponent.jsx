import React, { useState } from "react";
import { useFormik } from "formik";

import {
  Button,
  Row,
  Card,
  CardBody,
  Col,
  Input,
  Label,
  Select,
} from "core-components";
import { Center, MlModal, Text } from "shared-components";
import {
  cartridgeFormikInitialState,
  checkIsCartridgeFieldInvalid,
  isAddWellsBtnDisabled,
  isCreateCartridgesBtnDisabled,
  getRequestBody,
} from "./helpers";
import {
  CARTRIDGE_TYPE_OPTIONS,
  CARTRIDGE_WELLS,
  MODAL_BTN,
  MODAL_MESSAGE,
} from "appConstants";
const {
  MAX_DISTANCE,
  MIN_DISTANCE,
  MAX_HEIGHT,
  MIN_HEIGHT,
  MAX_VOLUME,
  MIN_VOLUME,
} = CARTRIDGE_WELLS;

const CartridgeComponent = (props) => {
  const { handleCreateCartridgeBtn } = props;

  const [noOfWellToShow, setNoOfWells] = useState(null);
  const [showWarningModal, setWarningModal] = useState(false);

  const formik = useFormik({
    initialValues: cartridgeFormikInitialState,
    enableReinitialize: true,
  });

  const { id, type, wellsCount, description, distance, height, volume } =
    formik.values;

  const handleBlur = (key, value) => {
    const isInvalid = checkIsCartridgeFieldInvalid(key, value);
    formik.setFieldValue(`${key}.isInvalid`, isInvalid);
  };

  const handleFocus = (name) => {
    formik.setFieldValue(`${name}.isInvalid`, false);
  };

  const handleOnChange = (key, value) => {
    // if user changes wells count afterwards
    if (key === "wellsCount.value" && noOfWellToShow !== null) {
      setWarningModal(true);
      return;
    }

    let valueToSet = value;
    if (key === "type.value" && !value) {
      valueToSet = "";
    }
    formik.setFieldValue(key, valueToSet);
  };

  const handleAddWellsBtn = () => {
    const wellsCountInt = parseInt(wellsCount.value);

    // if the user clicks on button after wells are configured
    if (noOfWellToShow !== null) {
      setWarningModal(true);
      return;
    }

    // set formik values
    const arrInit = [...Array(wellsCountInt)].map(() => ({
      value: null,
      isInvalid: false,
    }));

    formik.setFieldValue("distance", arrInit);
    formik.setFieldValue("height", arrInit);
    formik.setFieldValue("volume", arrInit);

    setNoOfWells(wellsCountInt);
  };

  const handleSaveBtn = (state) => {
    const requestBody = getRequestBody(state);
    handleCreateCartridgeBtn(requestBody);
  };

  const handleModalSuccessBtn = () => {
    // set isInvalid to false
    if (wellsCount.isInvalid === true) {
      formik.setFieldValue(`wellsCount.isInvalid`, false);
    }

    // reset wells config
    formik.setFieldValue("distance", []);
    formik.setFieldValue("height", []);
    formik.setFieldValue("volume", []);

    setNoOfWells(null);
    setWarningModal(false);
  };

  const handleModalCrossBtn = () => {
    setWarningModal(false);
  };

  return (
    <>
      {showWarningModal && (
        <MlModal
          isOpen={showWarningModal}
          textBody={"Are you sure you want to reset wells configurations?"}
          successBtn={MODAL_BTN.yes}
          failureBtn={MODAL_BTN.no}
          handleSuccessBtn={handleModalSuccessBtn}
          handleCrossBtn={handleModalCrossBtn}
        />
      )}

      <Card default className="my-3">
        <CardBody>
          <Text
            Tag="h4"
            size={24}
            className="text-center text-gray text-bold mt-3 mb-4"
          >
            {"Create Cartridge"}
          </Text>
          <Row>
            <Col className="mb-4" md={3}>
              <Label for="id" className="font-weight-bold">
                ID
              </Label>
              <Input
                name="id"
                id="id"
                placeholder={"Type here"}
                value={id.value}
                onChange={(event) =>
                  handleOnChange("id.value", event.target.value)
                }
                onBlur={(e) => handleBlur(e.target.name, e.target.value)}
                onFocus={(e) => handleFocus(e.target.name)}
              />
              {id.isInvalid && (
                <div className="flex-auto">
                  <Text Tag="p" size={14} className="text-danger">
                    {"Enter valid id"}
                  </Text>
                </div>
              )}
            </Col>

            <Col className="mb-4" md={3}>
              <Label for="type" className="font-weight-bold">
                Type
              </Label>
              <Select
                name="type"
                placeholder="Select Cartridge"
                className=""
                size="md"
                options={CARTRIDGE_TYPE_OPTIONS}
                value={{ value: type.value, label: type.value }}
                onChange={(e) => handleOnChange("type.value", e?.value)}
                // onBlur={(e) => handleBlur("type", e.target.value)}
                // onFocus={(e) => handleFocus("type")}
                isSearchable={false}
              />
              {type.isInvalid && (
                <div className="flex-auto">
                  <Text Tag="p" size={14} className="text-danger">
                    {"Invalid"}
                  </Text>
                </div>
              )}
            </Col>

            <Col className="mb-4" md={3}>
              <Label for="id" className="font-weight-bold">
                Description
              </Label>
              <Input
                name="desc"
                id="desc"
                placeholder={"Type here"}
                value={description.value}
                onChange={(event) =>
                  handleOnChange("description.value", event.target.value)
                }
              />
            </Col>

            <Col className="mb-4" md={3}>
              <Label for="wellsCount" className="font-weight-bold">
                Wells Count
              </Label>
              <Input
                name="wellsCount"
                id="wellsCount"
                placeholder={"Type here"}
                value={wellsCount.value}
                onChange={(event) =>
                  handleOnChange("wellsCount.value", event.target.value)
                }
                onBlur={(e) => handleBlur(e.target.name, e.target.value)}
                onFocus={(e) => handleFocus(e.target.name)}
              />
              {wellsCount.isInvalid && (
                <div className="flex-auto">
                  <Text Tag="p" size={14} className="text-danger">
                    {"Enter valid input"}
                  </Text>
                </div>
              )}
            </Col>
          </Row>

          <Row>
            <Col>
              <Center className="text-center pt-3">
                <Button
                  className="w-auto"
                  onClick={() => handleAddWellsBtn()}
                  disabled={isAddWellsBtnDisabled(formik.values)}
                  color="primary"
                >
                  Next Add Wells Details
                </Button>
              </Center>
            </Col>
          </Row>

          {noOfWellToShow &&
            parseInt(noOfWellToShow) > 0 &&
            [...Array(parseInt(noOfWellToShow))].map((ele, index) => (
              <Row className="my-3">
                <Col className="mb-4 md-1">
                  <Label for="wellNumber" className="font-weight-bold">
                    Well Number
                  </Label>
                  <Input disabled value={index + 1} />
                </Col>

                <Col className="mb-4 md-2">
                  <Label for="id" className="font-weight-bold">
                    Cartridge ID
                  </Label>
                  <Input disabled value={id.value} />
                </Col>

                <Col className="mb-4 md-3">
                  <Label for="distance" className="font-weight-bold">
                    Distance
                  </Label>
                  <Input
                    name="distance"
                    id="distance"
                    step="0.1"
                    placeholder={"Type here"}
                    value={distance[index].value}
                    onChange={(event) =>
                      handleOnChange(
                        `distance.${index}.value`,
                        event.target.value
                      )
                    }
                    onBlur={(e) =>
                      handleBlur(`distance.${index}`, e.target.value)
                    }
                    onFocus={() => handleFocus(`distance.${index}`)}
                  />
                  {distance[index].isInvalid && (
                    <div className="flex-auto">
                      <Text Tag="p" size={14} className="text-danger">
                        {`Distance must be between ${MIN_DISTANCE} and ${MAX_DISTANCE}`}
                      </Text>
                    </div>
                  )}
                </Col>

                <Col className="mb-4 md-3">
                  <Label for="height" className="font-weight-bold">
                    Height
                  </Label>
                  <Input
                    name="height"
                    id="height"
                    step="0.1"
                    placeholder={"Type here"}
                    value={height[index].value}
                    onChange={(event) =>
                      handleOnChange(
                        `height.${index}.value`,
                        event.target.value
                      )
                    }
                    onBlur={(e) =>
                      handleBlur(`height.${index}`, e.target.value)
                    }
                    onFocus={() => handleFocus(`height.${index}`)}
                  />
                  {height[index].isInvalid && (
                    <div className="flex-auto">
                      <Text Tag="p" size={14} className="text-danger">
                        {`Height must be between ${MIN_HEIGHT} and ${MAX_HEIGHT}`}
                      </Text>
                    </div>
                  )}
                </Col>

                <Col className="mb-4 md-3">
                  <Label for="volume" className="font-weight-bold">
                    Volume
                  </Label>
                  <Input
                    name="volume"
                    id="volume"
                    step="1"
                    placeholder={"Type here"}
                    value={volume[index].value}
                    onChange={(event) =>
                      handleOnChange(
                        `volume.${index}.value`,
                        event.target.value
                      )
                    }
                    onBlur={(e) =>
                      handleBlur(`volume.${index}`, e.target.value)
                    }
                    onFocus={() => handleFocus(`volume.${index}`)}
                  />
                  {volume[index].isInvalid && (
                    <div className="flex-auto">
                      <Text Tag="p" size={14} className="text-danger">
                        {`Volume must be between ${MIN_VOLUME} and ${MAX_VOLUME}`}
                      </Text>
                    </div>
                  )}
                </Col>
              </Row>
            ))}

          {noOfWellToShow && parseInt(noOfWellToShow) > 0 && (
            <Row>
              <Col>
                <Center className="text-center pt-3">
                  <Button
                    className="w-auto"
                    onClick={() => handleSaveBtn(formik.values)}
                    disabled={isCreateCartridgesBtnDisabled(formik.values)}
                    color="primary"
                  >
                    Create Cartridges
                  </Button>
                </Center>
              </Col>
            </Row>
          )}
        </CardBody>
      </Card>
    </>
  );
};

export default React.memo(CartridgeComponent);
