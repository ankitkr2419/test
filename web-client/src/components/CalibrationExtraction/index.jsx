import React from "react";
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

const CalibrationExtractionComponent = (props) => {
  return (
    <div className="calibration-content px-5">
      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          <p>Extraction Flow: Calibration</p>
        </CardBody>
      </Card>
    </div>
  );
};

export default React.memo(CalibrationExtractionComponent);
