import React from "react";

import { Col, FormGroup, Label, Select } from "core-components";

import { ExtractionBox } from "./Style";
import { WellComponent } from "./WellComponent";
import { DiscardRadioComponent } from "./DiscardRadioComponent";

const Extraction = (props) => {
  const extractionLabelArray = ["1", "2", "3", "4", "5", "6", "7", "8"];
  const pcrLabelArray = ["1", "2", "3", "4", "5"];

  return (
    <>
      <ExtractionBox>
        <div className="process-box mx-auto">
          <div className="mb-3 border-bottom-line">
            <FormGroup row className="align-items-center">
              <Label for="tip-selection" md={2}>
                Tip Selection
              </Label>
              <Col md={4}>
                <Select
                  placeholder="Select Tip"
                  className=""
                  size="sm"
                  //options={taskOptions}
                  //value={task}
                  //onChange={handleTaskChange}
                />
              </Col>
            </FormGroup>
          </div>
          {/* <WellComponent labelArray={extractionLabelArray}/> */}
          <DiscardRadioComponent />
        </div>
      </ExtractionBox>
    </>
  );
};

Extraction.propTypes = {};

export default Extraction;
