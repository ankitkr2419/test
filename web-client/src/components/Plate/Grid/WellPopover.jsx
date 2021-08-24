/* eslint-disable no-undef */
import React, { useState } from "react";
import PropTypes from "prop-types";
import { Button, Popover, PopoverHeader, PopoverBody } from "core-components";
import { Text, Center, ButtonIcon } from "shared-components";
import styled from "styled-components";
import { EXPERIMENT_STATUS } from "appConstants";

const StyledText = styled(Text)`
  color: ${(props) => {
    if (props.positive !== undefined) {
      return props.positive ? "#3FC13A" : "#F06666";
    }
    return "#707070";
  }} !important;
`;

const WellPopover = (props) => {
  const {
    status,
    text,
    index,
    sample,
    task,
    targets,
    onEditClickHandler,
    showGraphOfWell,
    isEditBtnDisabled,
    experimentStatus,
    ...rest
  } = props;

  const [showOnMap, setShowOnMap] = useState(true);

  const simulateOutSideClick = () => document.body.click();

  // on show graph button click handler
  const onShowClickHandler = () => {
    // close the popover
    simulateOutSideClick();
    showGraphOfWell(index, showOnMap);
    setShowOnMap(!showOnMap);
  };

  // on edit click handler
  const onEditClick = (event) => {
    // close the popover
    simulateOutSideClick();
    onEditClickHandler(event);
  };

  return (
    <Popover
      trigger="legacy"
      target={`PopoverWell${index}`}
      hideArrow
      placement="top-start"
      popperClassName="popover-well"
      status={status}
      {...rest}
    >
      <PopoverHeader status={status}>
        <Text Tag="span">{text}</Text>
        <ButtonIcon
          position="absolute"
          placement="right"
          top={0}
          right={0}
          size={32}
          name="cross"
          className="btn-close"
          onClick={simulateOutSideClick}
        />
      </PopoverHeader>
      <PopoverBody className="flex-100 scroll-y">
        <ul className="well-info flex-90 mx-auto mb-4 p-0">
          <li className="d-flex py-1">
            <Text className="flex-40 label mb-0">Sample</Text>
            <Text className="mb-0">{sample}</Text>
          </li>
          <li className="d-flex py-1">
            <Text className="flex-40 label mb-0">Target</Text>
            <div className="target-list">
              {targets === null && (
                <Text className={`mb-1 ${status}`}>---</Text>
              )}
              {targets !== null &&
                targets.map((ele) => (
                  <StyledText
                    key={ele.target_id}
                    className={"mb-1"}
                    positive={ele.ct === ""}
                  >
                    {ele.target_name || "target_name"}
                    {ele.ct === "" ? "" : `, CT ${ele.ct}`}
                  </StyledText>
                ))}
            </div>
          </li>
          <li className="d-flex py-1">
            <Text className="flex-40 label mb-0">Task</Text>
            <Text className="mb-0">{task}</Text>
          </li>
          {/* <li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Comment</Text>
						<Text className="mb-0">(No comments)</Text>
					</li> */}
        </ul>
        <Center>
          <Button
            className="mb-4"
            onClick={onShowClickHandler}
            disabled={experimentStatus === null}
          >
            {showOnMap ? "Show on Graph" : "Hide from Graph"}
          </Button>
          <Button
            onClick={onEditClick}
            disabled={
              experimentStatus === EXPERIMENT_STATUS.success ||
              experimentStatus === EXPERIMENT_STATUS.running ||
              experimentStatus === EXPERIMENT_STATUS.stopped
            }
          >
            Edit Info
          </Button>
        </Center>
      </PopoverBody>
    </Popover>
  );
};

WellPopover.propTypes = {
  status: PropTypes.string,
  text: PropTypes.string,
  index: PropTypes.number,
  sample: PropTypes.string,
  task: PropTypes.string,
  targets: PropTypes.array,
  onEditClickHandler: PropTypes.func,
  isDisabled: PropTypes.bool,
};

export default WellPopover;
