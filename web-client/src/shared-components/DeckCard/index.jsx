import React, { useState } from "react";

import styled from "styled-components";
import PropTypes from "prop-types";
import { Text, Icon } from "shared-components";
import { Button } from "core-components";
import ActionButton from "./ActionButton";
import { DECKCARD_BTN, PROCESS_ICON_CONSTANTS } from "appConstants";
import { Progress } from "reactstrap";
import OperatorLoginModalContainer from "containers/OperatorLoginModalContainer";
import { useDispatch, useSelector } from "react-redux";
import { setActiveDeck } from "action-creators/loginActionCreators";

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
    // background-color: #b2dad1;
    // border: 1px solid #ffffff;
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
    // background: #fff url("/images/deck-card-bg.svg") no-repeat;
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
      .icon-pause {
        font-size: 0.938rem;
      }
      .icon-resume {
        font-size: 1.25rem;
      }
    }
    .abort-button {
      position: absolute;
      right: 21px;
      top: 0;
      .semi-circular-button {
        border: 1px solid transparent;
        background-color: #ffffff;
        color: #3c3c3c;
      }
      .icon-cancel {
        font-size: 0.875rem;
      }
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
  .marquee {
    width: 80%;
    white-space: nowrap;
    overflow: hidden;
    box-sizing: border-box;
    .recipe-name {
      display: inline-block;
      padding-left: 100%;
      animation: marquee 10s linear infinite;
    }
    @keyframes marquee {
      0% {
        transform: translate(0, 0);
      }
      100% {
        transform: translate(-100%, 0);
      }
    }
  }
`;

const CardOverlay = styled.div`
  position: absolute;
  // display: none;
  width: 100%;
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
    hours,
    mins,
    secs,
    loginBtn,
    showProcess,
    showCleanUp,
    handleRightAction,
    handleLeftAction,
    leftActionBtn,
    rightActionBtn,
    leftActionBtnDisabled,
    rightActionBtnDisabled,
    progressPercentComplete,
    isActiveDeck,
    isAnotherDeckLoggedIn,
    processName,
    processType
  } = props;

  const [operatorLoginModalOpen, setOperatorLoginModalOpen] = useState(false);
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  let activeDeckName = activeDeckObj && activeDeckObj.name;

  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  let recipeActionReducerForDeck = recipeActionReducer.decks.find(
    (deckObj) => deckObj.name === deckName
  );

  const cleanUpReducer = useSelector((state) => state.cleanUpReducer);
  let cleanUpReducerForDeck = cleanUpReducer.decks.find(
    (deckObj) => deckObj.name === deckName
  );

  const toggleOperatorLoginModal = () => {
    setCurrentDeckActive();
    setOperatorLoginModalOpen(!operatorLoginModalOpen);
  };

  const getLeftActionBtn = (key) => {
    switch (key) {
      case DECKCARD_BTN.text.run:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.run}
            icon={DECKCARD_BTN.icon.run}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.pause:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.pause}
            icon={DECKCARD_BTN.icon.pause}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.resume:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.resume}
            icon={DECKCARD_BTN.icon.resume}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.done:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.done}
            icon={DECKCARD_BTN.icon.done}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.startUv:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.startUv}
            icon={DECKCARD_BTN.icon.run}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.pauseUv:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.pauseUv}
            icon={DECKCARD_BTN.icon.pause}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.resumeUv:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.resumeUv}
            icon={DECKCARD_BTN.icon.resume}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.next:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.next}
            icon={DECKCARD_BTN.icon.next}
            disabled={leftActionBtnDisabled}
            showCardOverLay={showCardOverLay}
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
            disabled={rightActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      case DECKCARD_BTN.text.abort:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.abort}
            icon={DECKCARD_BTN.icon.abort}
            disabled={rightActionBtnDisabled}
            showCardOverLay={showCardOverLay}
          />
        );
      default:
        break;
    }
  };

  const setCurrentDeckActive = () => {
    /**
     *  set active deck to current deck if:
     *  activeDeckName not found or not equal to current deck
     */
    if (!activeDeckName || deckName !== activeDeckName)
      dispatch(setActiveDeck(deckName));
  };

  const showCardOverLay = () => {
    return (
      (isAnotherDeckLoggedIn && loginBtn && !isActiveDeck) ||
      (isAnotherDeckLoggedIn && !loginBtn && !isActiveDeck) ||
      (!isAnotherDeckLoggedIn && !loginBtn && !isActiveDeck)
    );
  };

  const isProcessRunning = () => {
    return (
      recipeActionReducerForDeck.showProcess ||
      cleanUpReducerForDeck.showCleanUp
    );
  };


  /* get icon by process type
    * if process type not found use 'default'
    * if process type found but icon not found, use 'default'
  */
  const getIconName = () => {
    let processTypeText = processType ? processType : 'default'
    
    let obj = PROCESS_ICON_CONSTANTS.find(obj => obj.processType === processTypeText)
    
    let iconName = (obj?.iconName)
      ? obj.iconName 
      : PROCESS_ICON_CONSTANTS.find(obj => obj.processType === 'default').iconName
    return iconName;
  }

  return (
    <DeckCardBox
      className="d-flex justify-content-start align-items-center"
      onClick={setCurrentDeckActive}
    >
      <CardOverlay className={showCardOverLay() ? "" : "d-none"} />
      <div
        className="d-flex justify-content-center align-items-center deck-title"
        style={
          isProcessRunning()
            ? { backgroundColor: "#B2DAD1", border: "1px solid #ffffff" }
            : null
        }
      >
        <Text Tag="label" size={20}>
          {deckName}
        </Text>
      </div>
      <div
        className="p-4 w-100 h-100 deck-content logged-in1"
        style={
          isProcessRunning()
            ? { background: null }
            : { background: '#fff url("/images/deck-card-bg.svg") no-repeat' }
        }
      >
        <div className="d-flex justify-content-between align-items-center">
          <div className="d-none1">
            {showProcess && (
              <>
                <div className="resume-button" onClick={() => {
                  if(!leftActionBtnDisabled){
                    handleLeftAction()
                  }
                }}>
                  {getLeftActionBtn(leftActionBtn)}
                </div>
                <div className="abort-button" onClick={() => {
                  if(!rightActionBtnDisabled){
                    handleRightAction()
                  }
                }}>
                  {getRightActionBtn(rightActionBtn)}
                </div>

                <div className="marquee">
                  <Text
                    Tag="h5"
                    size={18}
                    className="mb-2 font-weight-bold recipe-name"
                  >
                    {recipeName}
                  </Text>
                  {/* TODO : Remove this commented code after clean up process developement done */}
                  {/* <Text Tag="label" className="mb-1">Current Processes - (Process Name)</Text>
								<Text Tag="label" className="mb-1 d-flex align-items-center">
									<Icon name='timer' size={19} className="text-primary"/>
									<Text Tag="span" className="hour-label font-weight-bold ml-2"> 1 Hr </Text>
									<Text Tag="span" className="min-label ml-2 font-weight-bold">8 min</Text>
									<Text Tag="span" className="ml-1">remaining</Text>
								</Text> */}

                  <Text Tag="label" className="mb-1 d-flex align-items-center">
                    <Icon name={getIconName()} size={19} className="text-primary" />
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
                      {processName ? processName : 'Processes remaining'}
                    </Text>
                  </Text>
                </div>
              </>
            )}

            {showCleanUp && (
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
                  <Text Tag="label" className="mb-1 d-flex align-items-center">
                    <Icon name="timer" size={19} className="text-primary" />
                    <Text
                      Tag="span"
                      className="hour-label font-weight-bold ml-2"
                    >
                      {" "}
                      {hours} Hr{" "}
                    </Text>
                    <Text
                      Tag="span"
                      className="min-label ml-2 font-weight-bold"
                    >
                      {mins} min {secs} sec
                    </Text>
                    <Text Tag="span" className="ml-1 mt-1 process-remaining">
                      remaining
                    </Text>
                  </Text>
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
            </>
          )}
          <OperatorLoginModalContainer
            operatorLoginModalOpen={operatorLoginModalOpen}
            toggleOperatorLoginModal={toggleOperatorLoginModal}
            deckName={deckName}
          />
        </div>

        {(showProcess || showCleanUp) && (
          <Progress
            value={progressPercentComplete}
            className={showProcess ? "custom-progress-bar" : "mt-3 custom-progress-bar"}
          />
        )}
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
  hours: PropTypes.number,
  mins: PropTypes.number,
  secs: PropTypes.number,
  processTotal: PropTypes.number,
  progressPercentComplete: PropTypes.number,
  leftActionBtnDisabled: PropTypes.bool,
  rightActionBtnDisabled: PropTypes.bool,
};

DeckCard.defaultProps = {
  isUserLoggedIn: false,
  showProcess: false,
  showCleanUp: false,
  recipeName: "Recipe Name",
  processNumber: 0,
  processTotal: 10,
  hours: 0,
  mins: 0,
  secs: 0,
  progressPercentComplete: 0,
  leftActionBtnDisabled: false,
  rightActionBtnDisabled: false,
};

export default React.memo(DeckCard);