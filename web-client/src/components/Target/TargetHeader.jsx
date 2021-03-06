import React from "react";
import PropTypes from "prop-types";
import { Text, ButtonIcon } from "shared-components";
import { StyledTargetHeader } from "./Style";
import { Input } from "core-components";

const TargetHeader = (props) => {
  const {
    isLoginTypeAdmin,
    isLoginTypeOperator,
    selectedTemplateDetails,
    editTemplate,
    onVolumeChange,
  } = props;
  return (
    <StyledTargetHeader className="target-header">
      {/* {isLoginTypeOperator === true && (
        <Text
          Tag="h6"
          size={18}
          className="text-default font-weight-light mb-0"
        >
          {selectedTemplateDetails.get("name")}
        </Text>
      )} */}
      {/* {isLoginTypeAdmin === true && ( */}
      {
        <div className="d-flex">
          {isLoginTypeAdmin === true && (
            <ButtonIcon
              name="pencil"
              size={28}
              className="px-0 border-0"
              onClick={editTemplate}
            />
          )}
          <Text
            size={14}
            className="flex-15 text-default text-truncate-multi-line font-weight-light mb-0 pl-3 pr-2 py-1"
          >
            {selectedTemplateDetails.get("name")}
          </Text>
          <Text
            size={14}
            className="flex-100 text-default text-truncate-multi-line font-weight-light mb-0 px-2 py-1"
          >
            {selectedTemplateDetails.get("description")}
          </Text>

          <Text
            size={14}
            className="flex-25 text-default text-truncate-multi-line font-weight-light mb-0 px-2 py-1"
          >
            {`Volume: ${selectedTemplateDetails.get("volume")} μL.`}
          </Text>

          <Text
            size={14}
            className="flex-25 text-default text-truncate-multi-line font-weight-light mb-0 px-2 py-1"
          >
            {`Lid Temperature: ${selectedTemplateDetails.get("lid_temp")} °C.`}
          </Text>

          {/* <div className="d-flex"> */}
          {/* <Input
              className="flex-100"
              type="number"
              name={`threshold`}
              placeholder={`µ units`}
              onChange={(event) => onVolumeChange(event.target.value)}
            /> */}
          {/* </div> */}
        </div>
      }
    </StyledTargetHeader>
  );
};

TargetHeader.propTypes = {
  isLoginTypeAdmin: PropTypes.bool.isRequired,
  isLoginTypeOperator: PropTypes.bool.isRequired,
  editTemplate: PropTypes.func,
};

export default TargetHeader;
