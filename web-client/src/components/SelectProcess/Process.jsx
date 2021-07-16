import React from "react";

import { Col } from "core-components";
import { ButtonIcon, Text } from "shared-components";

import { ProcessInnerBox } from "./Style";
import { useHistory } from "react-router";

const Process = (props) => {
  const { iconName, processName, route } = props;
  const history = useHistory();

  const clickHandler = (route) => {
    history.push(route);
  };
  return (
    <Col md={4}>
      <ProcessInnerBox>
        <div
          className="process-card bg-white d-flex align-items-center frame-icon"
          onClick={() => clickHandler(route)}
        >
          <ButtonIcon
            size={51}
            name={iconName}
            className="border-dark-gray text-primary"
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
