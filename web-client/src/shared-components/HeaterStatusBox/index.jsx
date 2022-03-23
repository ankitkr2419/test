import React from "react";

import styled from "styled-components";
import PropTypes from "prop-types";
import { Text, Icon } from "shared-components";
import { getIconName } from "shared-components/DeckCard/helpers";
import { useSelector } from "react-redux";

const ProecssRemainingBox = styled.div`
  position: absolute;
  height: 40px;
  top: -32px;
  width: 100%;
  left: 0px;
  background: rgb(255, 255, 255);
  border-radius: 12px 12px 0px 0px;
  box-shadow: rgb(0 0 0 / 16%) -3px 3px 6px;
  z-index: -1;
  display: flex;
  justify-content: space-evenly;
  align-content: center;
  flex-wrap: wrap;
`;
const HeaterStatusBox = () => {
  const heaterReducer = useSelector((state) => state.heaterProgressReducer);
  const heaterProgressReducerData = heaterReducer.toJS();
  const { heater_on, shaker_1_temp, shaker_2_temp } =
    heaterProgressReducerData.data;

  return (
    <ProecssRemainingBox>
      <span
        style={{ fontSize: "14px", fontWeight: "bold" }}
      >{`Heater Temperature 1 : ${shaker_1_temp}`}</span>{" "}
      <span
        style={{ fontSize: "14px", fontWeight: "bold" }}
      >{`Heater Temperature 2 : ${shaker_2_temp}`}</span>{" "}
    </ProecssRemainingBox>
  );
};

export default HeaterStatusBox;
