import React, { useState } from "react";

import PropTypes from "prop-types";
import { Text, ProcessRemaining } from "shared-components";
import { Button } from "core-components";
import ActionButton from "./ActionButton";
import { DECKCARD_BTN, SELECT_PROCESS_PROPS } from "appConstants";
import { Progress } from "reactstrap";
import OperatorLoginModalContainer from "containers/OperatorLoginModalContainer";
import { useDispatch, useSelector } from "react-redux";
import { setActiveDeck } from "action-creators/loginActionCreators";
import { DeckCardBox, CardOverlay } from "./Styles";
import CommonTimerFields from "./CommonTimerFields";
import { toast } from "react-toastify";

const DeckCard = (props) => {
  const {
    deckName,
    recipeName,
    processNumber,
    processTotal,
    remainingTime,
    totalTime,
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
    processType,
  } = props;

  const [operatorLoginModalOpen, setOperatorLoginModalOpen] = useState(false);
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  let activeDeckName = activeDeckObj && activeDeckObj.name;
  const isDeckBlocked = activeDeckObj && activeDeckObj.isDeckBlocked;

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
          />
        );
      case DECKCARD_BTN.text.pause:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.pause}
            icon={DECKCARD_BTN.icon.pause}
            disabled={leftActionBtnDisabled}
          />
        );
      case DECKCARD_BTN.text.resume:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.resume}
            icon={DECKCARD_BTN.icon.resume}
            disabled={leftActionBtnDisabled}
          />
        );
      case DECKCARD_BTN.text.done:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.done}
            icon={DECKCARD_BTN.icon.done}
            disabled={leftActionBtnDisabled}
          />
        );
      case DECKCARD_BTN.text.startUv:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.startUv}
            icon={DECKCARD_BTN.icon.run}
            disabled={leftActionBtnDisabled}
          />
        );
      case DECKCARD_BTN.text.pauseUv:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.pauseUv}
            icon={DECKCARD_BTN.icon.pause}
            disabled={leftActionBtnDisabled}
          />
        );
      case DECKCARD_BTN.text.resumeUv:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.resumeUv}
            icon={DECKCARD_BTN.icon.resume}
            disabled={leftActionBtnDisabled}
          />
        );
      case DECKCARD_BTN.text.next:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.next}
            icon={DECKCARD_BTN.icon.next}
            disabled={leftActionBtnDisabled}
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
          />
        );
      case DECKCARD_BTN.text.abort:
        return (
          <ActionButton
            label={DECKCARD_BTN.text.abort}
            icon={DECKCARD_BTN.icon.abort}
            disabled={rightActionBtnDisabled}
          />
        );
      default:
        break;
    }
  };

  const setCurrentDeckActive = () => {
    /** one cannot switch between deck while adding/editing processes.*/
    if (deckName !== activeDeckName && isDeckBlocked) {
      toast.warning("Decks cannot be switched while adding/editing processes!");
      return;
    }
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
            ? { background: null, zIndex: isActiveDeck ? 4 : null }
            : { background: '#fff url("/images/deck-card-bg.svg") no-repeat' }
        }
      >
        {showProcess && (
          <ProcessRemaining
            processName={processName}
            processType={processType}
            processNumber={processNumber}
            processTotal={processTotal}
          />
        )}
        <div className="d-flex justify-content-between align-items-center">
          <div className="d-none1">
            {/* PROCESSES */}
            {showProcess && (
              <>
                <div
                  className="resume-button"
                  onClick={() => {
                    if (!leftActionBtnDisabled) {
                      handleLeftAction();
                    }
                  }}
                >
                  {getLeftActionBtn(leftActionBtn)}
                </div>
                <div
                  className="abort-button"
                  onClick={() => {
                    if (!rightActionBtnDisabled) {
                      handleRightAction();
                    }
                  }}
                >
                  {getRightActionBtn(rightActionBtn)}
                </div>

                <CommonTimerFields
                  recipeName={recipeName}
                  remainingTime={remainingTime}
                  totalTime={totalTime}
                />
              </>
            )}

            {/* {CLEAN-UP (UV)} */}
            {showCleanUp && (
              <>
                <div className="resume-button" onClick={handleLeftAction}>
                  {getLeftActionBtn(leftActionBtn)}
                </div>
                <div className="abort-button" onClick={handleRightAction}>
                  {getRightActionBtn(rightActionBtn)}
                </div>

                <CommonTimerFields
                  recipeName={"Clean Up"}
                  remainingTime={remainingTime}
                  totalTime={totalTime}
                />
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
          {loginBtn && (
            <OperatorLoginModalContainer
              operatorLoginModalOpen={operatorLoginModalOpen}
              toggleOperatorLoginModal={toggleOperatorLoginModal}
              deckName={deckName}
            />
          )}
        </div>

        {(showProcess || showCleanUp) && (
          <Progress
            value={progressPercentComplete}
            className={
              showProcess ? "custom-progress-bar" : "custom-progress-bar"
            }
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
  isDeckBlocked: false,
};

export default React.memo(DeckCard);
