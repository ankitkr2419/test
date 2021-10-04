import React from "react";

import { Col } from "core-components";
import { ButtonIcon, Text } from "shared-components";

import { ProcessInnerBox } from "./Style";
import { useHistory } from "react-router";
import { useSelector } from "react-redux";
import { toast } from "react-toastify";

const Process = (props) => {
  const { iconName, processName, route } = props;
  const history = useHistory();

  const cartridge1DetailsReducer = useSelector(
    (state) => state.cartridge1DetailsReducer
  );
  const cartridge2DetailsReducer = useSelector(
    (state) => state.cartridge2DetailsReducer
  );

  const clickHandler = (route) => {
    // for piercing, if both cartridges are not defined then,
    // we will not allow redirect

    if (
      cartridge1DetailsReducer.cartridgeDetails === null &&
      cartridge2DetailsReducer.cartridgeDetails === null &&
      processName === "Piercing"
    ) {
      toast.error("Cartridges not configured for piercing process", {
        autoClose: false,
      });
      return;
    }

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
