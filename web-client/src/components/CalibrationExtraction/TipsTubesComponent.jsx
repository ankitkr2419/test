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
import {
  MIN_TIPTUBE_ID,
  MAX_TIPTUBE_ID,
  MIN_TIPTUBE_VOLUME,
  MAX_TIPTUBE_VOLUME,
  MIN_TIPTUBE_HEIGHT,
  MAX_TIPTUBE_HEIGHT,
  MIN_TIPTUBE_TTBASE,
  MAX_TIPTUBE_TTBASE,
} from "appConstants";
import { Center, Text } from "shared-components";
import {
  formikToArray,
  isTipsTubesButtonDisabled,
  tipTubeTypeOptions,
} from "./helpers";
import { CheckBoxGroup } from "components/AppHeader/CheckBoxGroup";

const TipsTubesComponent = (props) => {
  const { handleTipesTubesButton, formik } = props;
  const {
    tipTubeId,
    tipTubeName,
    tipTubeType,
    allowedPositions,
    volume,
    height,
    ttBase,
  } = formik.values;

  const handleBlurID = (value) => {
    if (!value || value < MIN_TIPTUBE_ID || value > MAX_TIPTUBE_ID) {
      formik.setFieldValue("tipTubeId.isInvalid", true);
    }
  };

  const handleBlurName = (value) => {
    if (!value) {
      formik.setFieldValue("tipTubeName.isInvalid", true);
    }
  };

  const handleBlurVolume = (value) => {
    if (!value || value < MIN_TIPTUBE_VOLUME || value > MAX_TIPTUBE_VOLUME) {
      formik.setFieldValue("volume.isInvalid", true);
    }
  };

  const handleBlurHeight = (value) => {
    if (!value || value < MIN_TIPTUBE_HEIGHT || value > MAX_TIPTUBE_HEIGHT) {
      formik.setFieldValue("height.isInvalid", true);
    }
  };

  const handleBlurTTBase = (value) => {
    if (!value || value < MIN_TIPTUBE_TTBASE || value > MAX_TIPTUBE_TTBASE) {
      formik.setFieldValue("ttBase.isInvalid", true);
    }
  };

  const onAllowedPositionChanged = (key, value) => {
    //clear old validation
    handleOnChange("allowedPositions.isInvalid", false);

    //from available positions, change selected position and save to formik
    let newPositions = {
      ...allowedPositions.value,
      [key]: !value, //toggle value
    };
    handleOnChange("allowedPositions.value", newPositions);

    //validations of new values
    let arrayOfAllowedPositions = formikToArray({ value: newPositions });
    if (arrayOfAllowedPositions.length === 0) {
      handleOnChange("allowedPositions.isInvalid", true);
    }
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  return (
    <Card default className="my-3">
      <CardBody>
        <Form onSubmit={handleTipesTubesButton}>
          <Row form>
            <Col sm={4}>
              <FormGroup>
                <Label for="tipTubeId" className="font-weight-bold">
                  ID
                </Label>
                <Input
                  type="number"
                  name="tipTubeId"
                  id="tipTubeId"
                  placeholder={`${MIN_TIPTUBE_ID} - ${MAX_TIPTUBE_ID}`}
                  value={tipTubeId.value || ""}
                  max={MAX_TIPTUBE_ID}
                  min={MIN_TIPTUBE_ID}
                  onChange={(event) =>
                    handleOnChange(
                      "tipTubeId.value",
                      parseInt(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurID(parseInt(e.target.value))}
                  onFocus={() => handleOnChange("tipTubeId.isInvalid", false)}
                />
                {tipTubeId.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      ID should be {MIN_TIPTUBE_ID} - {MAX_TIPTUBE_ID}
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="tipTubeName" className="font-weight-bold">
                  Name
                </Label>
                <Input
                  name="tipTubeName"
                  id="tipTubeName"
                  placeholder={"Type here"}
                  value={tipTubeName.value || ""}
                  onChange={(event) =>
                    handleOnChange("tipTubeName.value", event.target.value)
                  }
                  onBlur={(e) => handleBlurName(e.target.value)}
                  onFocus={() => handleOnChange("tipTubeName.isInvalid", false)}
                />
                {tipTubeName.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Enter valid name
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="tipTubeType" className="font-weight-bold">
                  Type
                </Label>
                <div>
                  <Select
                    placeholder="Select Type"
                    options={tipTubeTypeOptions}
                    value={tipTubeType.value}
                    onChange={(value) =>
                      handleOnChange("tipTubeType.value", value)
                    }
                  />
                </div>
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="volume" className="font-weight-bold">
                  Volume
                </Label>
                <Input
                  type="number"
                  name="volume"
                  id="volume"
                  step="0.1"
                  placeholder={`${MIN_TIPTUBE_VOLUME} - ${MAX_TIPTUBE_VOLUME}`}
                  value={volume.value || ""}
                  max={MAX_TIPTUBE_VOLUME}
                  min={MIN_TIPTUBE_VOLUME}
                  onChange={(event) =>
                    handleOnChange(
                      "volume.value",
                      parseFloat(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurVolume(parseFloat(e.target.value))}
                  onFocus={() => handleOnChange("volume.isInvalid", false)}
                />
                {volume.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Volume should be {MIN_TIPTUBE_VOLUME} -{" "}
                      {MAX_TIPTUBE_VOLUME}
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="height" className="font-weight-bold">
                  Height
                </Label>
                <Input
                  type="number"
                  name="height"
                  id="height"
                  placeholder={`${MIN_TIPTUBE_HEIGHT} - ${MAX_TIPTUBE_HEIGHT}`}
                  value={height.value || ""}
                  max={MAX_TIPTUBE_HEIGHT}
                  min={MIN_TIPTUBE_HEIGHT}
                  onChange={(event) =>
                    handleOnChange("height.value", parseInt(event.target.value))
                  }
                  onBlur={(e) => handleBlurHeight(parseInt(e.target.value))}
                  onFocus={() => handleOnChange("height.isInvalid", false)}
                />
                {height.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Height should be {MIN_TIPTUBE_HEIGHT} to{" "}
                      {MAX_TIPTUBE_HEIGHT}.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={4}>
              <FormGroup>
                <Label for="ttBase" className="font-weight-bold">
                  TTBase
                </Label>
                <Input
                  type="number"
                  name="ttBase"
                  id="ttBase"
                  step="0.1"
                  placeholder={`${MIN_TIPTUBE_TTBASE} - ${MAX_TIPTUBE_TTBASE}`}
                  value={ttBase.value || ""}
                  max={MAX_TIPTUBE_TTBASE}
                  min={MIN_TIPTUBE_TTBASE}
                  onChange={(event) =>
                    handleOnChange(
                      "ttBase.value",
                      parseFloat(event.target.value)
                    )
                  }
                  onBlur={(e) => handleBlurTTBase(parseFloat(e.target.value))}
                  onFocus={() => handleOnChange("ttBase.isInvalid", false)}
                />
                {ttBase.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      TTBase should be {MIN_TIPTUBE_TTBASE} to{" "}
                      {MAX_TIPTUBE_TTBASE}.
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>

            <Col sm={12}>
              <FormGroup>
                <Label className="font-weight-bold">Allowed Positions</Label>
                <CheckBoxGroup className="d-flex" style={{ height: "auto" }}>
                  {Object.entries(allowedPositions?.value).map(
                    ([key, value]) => {
                      return (
                        <CheckBox
                          key={key}
                          id={key}
                          name={key}
                          label={key}
                          className="mr-4"
                          checked={value}
                          onChange={() => onAllowedPositionChanged(key, value)}
                        />
                      );
                    }
                  )}
                </CheckBoxGroup>
                {allowedPositions.isInvalid && (
                  <div className="flex-70">
                    <Text Tag="p" size={14} className="text-danger">
                      Atleast one position required
                    </Text>
                  </div>
                )}
              </FormGroup>
            </Col>
          </Row>
          <Center className="text-center pt-3">
            <Button
              disabled={isTipsTubesButtonDisabled(formik.values)}
              color="primary"
            >
              Create {tipTubeType.value.label}
            </Button>
          </Center>
        </Form>
      </CardBody>
    </Card>
  );
};

export default React.memo(TipsTubesComponent);
