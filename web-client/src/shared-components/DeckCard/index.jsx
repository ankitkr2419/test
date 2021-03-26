import React, { useState } from "react";

import styled from "styled-components";
import PropTypes from "prop-types";
import { Text, Icon } from "shared-components";
import { Button } from "core-components";
import ActionButton from "./ActionButton";
import { DECKCARD_BTN } from "appConstants";
import { Progress } from "reactstrap";
import OperatorLoginModalContainer from "containers/OperatorLoginModalContainer";

const DeckCardBox = styled.div`
  width: 50%;
  // width: 32rem;
  height: 6.625rem;
  position: relative;
  box-shadow: 0px -3px 6px rgba(0, 0, 0, 0.16);
  &::before {
    content: "";
    position: absolute;
    background-image: linear-gradient(
      to right,
      #aedbd5,
      #a9dac5,
      #afd7b0,
      #bed29a,
      #d3ca87,
      #dcc278,
      #e7b96c,
      #f2ae64,
      #f2a453,
      #f29942,
      #f38d31,
      #f3811f
    );
    width: 100%;
    height: 2px;
    top: 0;
    left: 0;
    z-index: 1;
  }
  .deck-title {
    width: 2.563rem;
    height: 100%;
    font-size: 1.25rem;
    line-height: 1.688rem;
    font-weight: bold;
    color: #51575a;
    border: 1px solid transparent;
    box-shadow: 0 -3px 6px rgba(0, 0, 0, 0.16);
    > label {
      transform: rotate(-90deg);
      white-space: nowrap;
      margin-bottom: 0;
    }
    &.active {
      background-color: #b2dad1;
      border: 1px solid #ffffff;
    }
  }
  .deck-content {
    position: relative;
    background: #fff url("/images/deck-card-bg.svg") no-repeat;
    > button {
      min-width: 7.063rem;
      height: 2.5rem;
      line-height: 1.125rem;
    }
    .custom-progress-bar {
      border-radius: 7px;
      background-color: #b2dad131;
      border: 2px solid #b2dad131;
      .progress-bar {
        //background-color:#10907A;
        border-radius: 7px 0px 0px 7px;
        background-color: #72b5e6;
        animation: blink 1s linear infinite;
      }
    }
    // .uv-light-button{
    // 	position:absolute;
    // 	right:244px;
    // 	top:0;
    // }
    .resume-button {
      position: absolute;
      right: 123px;
      top: 0;
    }
    .abort-button {
      position: absolute;
      right: 21px;
      top: 0;
    }
    .hour-label {
      background-color: #f5e3d3;
      border-radius: 4px 0 0 4px;
      border-right: 2px solid #f38220;
      padding: 3px 4px;
      font-size: 0.875rem;
      line-height: 1rem;
    }
    .min-label {
      font-size: 0.875rem;
      line-height: 1rem;
    }
    .process-count-label {
      background-color: #f5e3d3;
      border-radius: 4px;
      padding: 3px 4px;
      font-size: 1.125rem;
      line-height: 1rem;
    }
    .process-total-count {
      font-size: 0.875rem;
      line-height: 1rem;
    }
    .process-remaining {
      font-size: 10px;
      line-height: 11px;
    }
    // add this class while login
    &.logged-in {
      background: #ffffff;
    }
  }
  @keyframes blink {
    0% {
      background-color: #9d9d9d;
    }
    50% {
      background-color: #72b5e6;
    }
    100% {
      background-color: #9d9d9d;
    }
  }
`;

const CardOverlay = styled.div`
  position: absolute;
  display: none;
  width: 50%;
  height: 6.625rem;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.28);
  z-index: 3;
  cursor: pointer;
`;
const DeckCard = (props) => {
  const {
    deckName,
    recipeName,
    processNumber,
    processTotal,
    loginBtn,
    showProcess,
    showCleanUp,
    handleRightAction,
    handleLeftAction,
    leftActionBtn,
    rightActionBtn,
    progressPercentComplete,
  } = props;

  const [operatorLoginModalOpen, setOperatorLoginModalOpen] = useState(false);
  const toggleOperatorLoginModal = () => {
    setOperatorLoginModalOpen(!operatorLoginModalOpen);
  };

  const getLeftActionBtn = (key) => {
    switch (key) {
      case DECKCARD_BTN.text.run:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.run}
            icon={DECKCARD_BTN.icon.run}
          />
        );
      case DECKCARD_BTN.text.pause:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.pause}
            icon={DECKCARD_BTN.icon.pause}
          />
        );
      case DECKCARD_BTN.text.resume:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.resume}
            icon={DECKCARD_BTN.icon.resume}
          />
        );
      case DECKCARD_BTN.text.done:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.done}
            icon={DECKCARD_BTN.icon.done}
          />
        );
      default:
        break;
    }
  };

  const getRightActionBtn = (key) => {
    switch (key) {
      case DECKCARD_BTN.text.cancel:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.cancel}
            icon={DECKCARD_BTN.icon.cancel}
          />
        );
      case DECKCARD_BTN.text.abort:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.abort}
            icon={DECKCARD_BTN.icon.abort}
          />
        );
      default:
        break;
    }
  };

  return (
    <DeckCardBox className="d-flex justify-content-start align-items-center">
      <CardOverlay />
      <div className="d-flex justify-content-center align-items-center deck-title">
        <Text Tag="label" size={20}>
          {deckName}
        </Text>
      </div>
      <div className="p-4 w-100 h-100 deck-content logged-in1">
        <div className="d-flex justify-content-between align-items-center">
          <div className="d-none1">
            {showProcess && (
              <>
                <div className="resume-button" onClick={handleLeftAction}>
                  {getLeftActionBtn(leftActionBtn)}
                </div>
                <div className="abort-button" onClick={handleRightAction}>
                  {getRightActionBtn(rightActionBtn)}
                </div>

                <div className="d-none1">
                  <Text
                    Tag="h5"
                    size={18}
                    className="mb-2 font-weight-bold recipe-name"
                  >
                    {recipeName}
                  </Text>

                  {/* <Text Tag="label" className="mb-1">Current Processes - (Process Name)</Text>
								<Text Tag="label" className="mb-1 d-flex align-items-center">
									<Icon name='timer' size={19} className="text-primary"/>
									<Text Tag="span" className="hour-label font-weight-bold ml-2"> 1 Hr </Text>
									<Text Tag="span" className="min-label ml-2 font-weight-bold">8 min</Text>
									<Text Tag="span" className="ml-1">remaining</Text>
								</Text> */}

                  <Text Tag="label" className="mb-1 d-flex align-items-center">
                    <Icon name="process" size={19} className="text-primary" />
                    <Text
                      Tag="span"
                      className="process-count-label font-weight-bold ml-2"
                    >
                      {" "}
                      {processNumber}
                      <Text
                        Tag="span"
                        className="process-total-count font-weight-bold"
                      >
                        /{processTotal}{" "}
                      </Text>{" "}
                    </Text>
                    <Text Tag="span" className="ml-1 process-remaining">
                      Processes remaining
                    </Text>
                  </Text>
                  <Progress
                    value={progressPercentComplete}
                    className="custom-progress-bar"
                  />
                </div>
              </>
            )}

            {showCleanUp && (
              <>
                <div className="resume-button">
                  <ActionButton />
                </div>
                <div className="abort-button">
                  <ActionButton />
                </div>

                <div className="d-none1">
                  <Text
                    Tag="h5"
                    size="18"
                    className="mb-2 font-weight-bold recipe-name"
                  >
                    Clean Up
                  </Text>
                  <Text Tag="label" className="mb-1 d-flex align-items-center">
                    <Icon name="timer" size={19} className="text-primary" />
                    <Text
                      Tag="span"
                      className="hour-label font-weight-bold ml-2"
                    >
                      {" "}
                      1 Hr{" "}
                    </Text>
                    <Text
                      Tag="span"
                      className="min-label ml-2 font-weight-bold"
                    >
                      8 min
                    </Text>
                    <Text Tag="span" className="ml-1">
                      remaining
                    </Text>
                  </Text>
                  <Progress value="2" className="custom-progress-bar" />
                </div>
              </>
            )}
          </div>

          {loginBtn && (
            <>
              <Button
                color="primary"
                className="ml-auto d-flex"
                size="sm"
                onClick={toggleOperatorLoginModal}
              >
                {" "}
                Login
              </Button>

              <OperatorLoginModalContainer
                operatorLoginModalOpen={operatorLoginModalOpen}
                toggleOperatorLoginModal={toggleOperatorLoginModal}
                deckName={deckName}
              />
            </>
          )}
        </div>
      </div>
    </DeckCardBox>
  );
};

DeckCard.propTypes = {
  isUserLoggedIn: PropTypes.bool,
  showProcess: PropTypes.bool,
  showCleanUp: PropTypes.bool,
  recipeName: PropTypes.string,
  processNumber: PropTypes.number,
  processTotal: PropTypes.number,
  progressPercentComplete: PropTypes.number,
};

DeckCard.defaultProps = {
  isUserLoggedIn: false,
  showProcess: false,
  showCleanUp: false,
  recipeName: "Recipe Name",
  processNumber: 0,
  processTotal: 10,
  progressPercentComplete: 0,
};

export default DeckCard;
