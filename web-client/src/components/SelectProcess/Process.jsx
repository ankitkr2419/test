import React from "react";

import { Col } from "core-components";
import { ButtonIcon, Text } from "shared-components";

import { ProcessInnerBox } from "./Style";

const Process = ({ iconName, processName }) => {
  return (
    <Col md={4}>
      <ProcessInnerBox>
        <div className="process-card bg-white d-flex align-items-center frame-icon">
          <ButtonIcon
            size={51}
            name={iconName}
            className="border-dark-gray text-primary"
            //onClick={toggleExportDataModal}
          />
          <Text Tag="span" className="ml-2 process-name">
            {processName}
          </Text>
        </div>
      </ProcessInnerBox>
    </Col>
  );
};

Process.propTypes = {};

export default Process;
