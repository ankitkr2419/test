import React, { useState } from "react";
import DeckCard from "shared-components/DeckCard";
import {
  DECKNAME,
  MODAL_BTN,
  MODAL_MESSAGE,
  DECKCARD_BTN,
} from "appConstants";
import { useSelector, useDispatch } from "react-redux";

import {
  abortRecipeInitiated,
  pauseRecipeInitiated,
  resumeRecipeInitiated,
  runRecipeInitiated,
  runRecipeReset,
  pauseRecipeReset,
  resumeRecipeReset,
  abortRecipeReset,
  updateRecipeReducerDataForDeck,
} from "action-creators/recipeActionCreators";
import {
  abortCleanUpActionInitiated,
  pauseCleanUpActionInitiated,
  resumeCleanUpActionInitiated,
  runCleanUpActionInitiated,
  runCleanUpActionReset,
} from "action-creators/cleanUpActionCreators";
import { MlModal } from "shared-components";
import { loginReset } from "action-creators/loginActionCreators";

const AppFooter = (props) => {
  const dispatch = useDispatch();

  const [
    isConfirmationModalVisibleDeckA,
    setIsConfirmationModalVisibleDeckA,
  ] = useState(false);
  const [
    isConfirmationModalVisibleDeckB,
    setIsConfirmationModalVisibleDeckB,
  ] = useState(false);

  const [abortConfirmationModalA, setAbortConfirmationModalA] = useState(false);
  const [abortConfirmationModalB, setAbortConfirmationModalB] = useState(false);

  //login reducer
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

  let loginDataOfA =
    loginReducerData &&
    loginReducerData.decks.find((deck) => deck.name === DECKNAME.DeckA);
  let isDeckALoggedIn = loginDataOfA.isLoggedIn;
  let isDeckAActive = loginDataOfA.isActive;

  let loginDataOfB =
    loginReducerData &&
    loginReducerData.decks.find((deck) => deck.name === DECKNAME.DeckB);
  let isDeckBLoggedIn = loginDataOfB.isLoggedIn;
  let isDeckBActive = loginDataOfB.isActive;

  //recipe reducer
  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  let recipeActionReducerForDeckA = recipeActionReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckA
  );
  let recipeActionReducerForDeckB = recipeActionReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckB
  );

  //cleanUp reducer
  const cleanUpReducer = useSelector((state) => state.cleanUpReducer);
  let cleanUpReducerForDeckA = cleanUpReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckA
  );
  let cleanUpReducerForDeckB = cleanUpReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckB
  );

  const [confirmAbortModal, setConfirmAbortModal] = useState(false);
  const [confirmDoneModal, setConfirmDoneModal] = useState(false);
  const [deckName, setDeckName] = useState("");

  const getLeftActionBtnHandler = (deckName) => {
    let recipeReducerData =
      deckName === DECKNAME.DeckA
        ? recipeActionReducerForDeckA
        : recipeActionReducerForDeckB;
    let showProcess = recipeReducerData.showProcess;

    let cleanUpReducerData =
      deckName === DECKNAME.DeckA
        ? cleanUpReducerForDeckA
        : cleanUpReducerForDeckB;
    let showCleanUp = cleanUpReducerData.showCleanUp;

    switch (
      showProcess
        ? recipeReducerData.leftActionBtn
        : cleanUpReducerData.leftActionBtn
    ) {
      case DECKCARD_BTN.text.run:
        return deckName === DECKNAME.DeckA
          ? handleRunActionDeckA
          : handleRunActionDeckB;
      case DECKCARD_BTN.text.pause:
        return deckName === DECKNAME.DeckA
          ? handlePauseActionDeckA
          : handlePauseActionDeckB;
      case DECKCARD_BTN.text.resume:
        return deckName === DECKNAME.DeckA
          ? handleResumeActionDeckA
          : handleResumeActionDeckB;

      case DECKCARD_BTN.text.startUv:
        return deckName === DECKNAME.DeckA
          ? handleRunActionDeckA
          : handleRunActionDeckB;
      case DECKCARD_BTN.text.pauseUv:
        return deckName === DECKNAME.DeckA
          ? handlePauseActionDeckA
          : handlePauseActionDeckB;
      case DECKCARD_BTN.text.resumeUv:
        return deckName === DECKNAME.DeckA
          ? handleResumeActionDeckA
          : handleResumeActionDeckB;
      case DECKCARD_BTN.text.done:
        return deckName === DECKNAME.DeckA
          ? handleDoneActionDeckA
          : handleDoneActionDeckB;

      default:
        break;
    }
  };

  const getRightActionBtnHandler = (deckName) => {
    let recipeReducerData =
      deckName === DECKNAME.DeckA
        ? recipeActionReducerForDeckA
        : recipeActionReducerForDeckB;
    let showProcess = recipeReducerData.showProcess;

    let cleanUpReducerData =
      deckName === DECKNAME.DeckA
        ? cleanUpReducerForDeckA
        : cleanUpReducerForDeckB;

    switch (
      showProcess
        ? recipeReducerData.rightActionBtn
        : cleanUpReducerData.rightActionBtn
    ) {
      case DECKCARD_BTN.text.abort:
        return deckName === DECKNAME.DeckA
          ? handleAbortActionDeckA
          : handleAbortActionDeckB;
      case DECKCARD_BTN.text.cancel:
        return deckName === DECKNAME.DeckA
          ? handleCancelActionDeckA
          : handleCancelActionDeckB;
      default:
        break;
    }
  };

  const handleRunActionDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      const { recipeId } = recipeReducerData.recipeData;
      dispatch(
        runRecipeInitiated({
          recipeId: recipeId,
          deckName: recipeReducerData.name, //deck A
        })
      );
    } else {
      dispatch(
        runCleanUpActionInitiated({
          time: `${cleanUpReducerForDeckA.hours}:${cleanUpReducerForDeckA.mins}:${cleanUpReducerForDeckA.secs}`,
          deckName: DECKNAME.DeckAShort,
        })
      );
    }
  };
  const handleRunActionDeckB = () => {
    let recipeReducerData = recipeActionReducerForDeckB;

    if (recipeReducerData.showProcess) {
      const { recipeId } = recipeReducerData.recipeData;
      dispatch(
        runRecipeInitiated({
          recipeId: recipeId,
          deckName: recipeReducerData.name, //deck B
        })
      );
    } else {
      dispatch(
        runCleanUpActionInitiated({
          time: `${cleanUpReducerForDeckB.hours}:${cleanUpReducerForDeckB.mins}:${cleanUpReducerForDeckB.secs}`,
          deckName: DECKNAME.DeckBShort,
        })
      );
    }
  };

  const handlePauseActionDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      dispatch(pauseRecipeInitiated({ deckName: recipeReducerData.name }));
    } else {
      dispatch(pauseCleanUpActionInitiated({ deckName: DECKNAME.DeckAShort }));
    }
  };
  const handlePauseActionDeckB = () => {
    let recipeReducerData = recipeActionReducerForDeckB;

    if (recipeReducerData.showProcess) {
      dispatch(pauseRecipeInitiated({ deckName: recipeReducerData.name }));
    } else {
      dispatch(pauseCleanUpActionInitiated({ deckName: DECKNAME.DeckBShort }));
    }
  };

  const handleResumeActionDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      dispatch(resumeRecipeInitiated({ deckName: recipeReducerData.name }));
    } else {
      dispatch(resumeCleanUpActionInitiated({ deckName: DECKNAME.DeckAShort }));
    }
  };

  const handleResumeActionDeckB = () => {
    let recipeReducerData = recipeActionReducerForDeckB;

    if (recipeReducerData.showProcess) {
      dispatch(resumeRecipeInitiated({ deckName: recipeReducerData.name }));
    } else {
      dispatch(resumeCleanUpActionInitiated({ deckName: DECKNAME.DeckBShort }));
    }
  };

  const handleDoneActionDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;
    if (recipeReducerData.showProcess) {
      toggleComfirmationModal(DECKNAME.DeckA);
    } else {
      setDeckName(DECKNAME.DeckA);
      setConfirmDoneModal(!confirmDoneModal);
    }
  };

  const handleDoneActionDeckB = () => {
    let recipeReducerData = recipeActionReducerForDeckB;
    if (recipeReducerData.showProcess) {
      toggleComfirmationModal(DECKNAME.DeckB);
    } else {
      setDeckName(DECKNAME.DeckB);
      setConfirmDoneModal(!confirmDoneModal);
    }
  };
  const handleDoneModalDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      dispatch(runRecipeReset(deckName));
    } else {
      dispatch(runCleanUpActionReset({ deckName: DECKNAME.DeckA }));
    }
    dispatch(loginReset(DECKNAME.DeckA));
    setConfirmDoneModal(!confirmDoneModal);
  };

  const handleDoneModalDeckB = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      dispatch(runRecipeReset(deckName));
    } else {
      dispatch(runCleanUpActionReset({ deckName: DECKNAME.DeckB }));
    }
    dispatch(loginReset(DECKNAME.DeckB));
    setConfirmDoneModal(!confirmDoneModal);
  };

  const handleCancelActionDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      dispatch(runRecipeReset(deckName));
    } else {
      dispatch(runCleanUpActionReset({ deckName: DECKNAME.DeckA }));
    }
  };
  const handleCancelActionDeckB = () => {
    let recipeReducerData = recipeActionReducerForDeckB;

    if (recipeReducerData.showProcess) {
      dispatch(runRecipeReset(deckName));
    } else {
      dispatch(runCleanUpActionReset({ deckName: DECKNAME.DeckB }));
    }
  };

  const handleAbortActionDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      toggleAbortConfirmationModal(DECKNAME.DeckA);
    } else {
      setConfirmAbortModal(!confirmAbortModal);
      setDeckName(DECKNAME.DeckA);
    }
  };

  const handleAbortActionDeckB = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
      toggleAbortConfirmationModal(DECKNAME.DeckB);
    } else {
      setConfirmAbortModal(!confirmAbortModal);
      setDeckName(DECKNAME.DeckB);
    }
  };

  const handleAbortModalDeckA = () => {
    let recipeReducerData = recipeActionReducerForDeckA;

    if (recipeReducerData.showProcess) {
    } else {
      setConfirmAbortModal(!confirmAbortModal);
      dispatch(abortCleanUpActionInitiated({ deckName: DECKNAME.DeckAShort }));
      dispatch(loginReset(DECKNAME.DeckA));
    }
  };

  const toggleComfirmationModal = (deckName) => {
    if (deckName === DECKNAME.DeckA) {
      setIsConfirmationModalVisibleDeckA(!isConfirmationModalVisibleDeckA);
    } else {
      setIsConfirmationModalVisibleDeckB(!isConfirmationModalVisibleDeckB);
    }
  };

  const toggleAbortConfirmationModal = (deckName) => {
    if (deckName === DECKNAME.DeckA) {
      setAbortConfirmationModalA(!abortConfirmationModalA);
    } else {
      setAbortConfirmationModalB(!abortConfirmationModalB);
    }
  };

  const onAbortConfirmed = (deckName) => {
    toggleAbortConfirmationModal(deckName);
    dispatch(abortRecipeInitiated({ deckName }));
  };

  const onNextClickedAfterDoneRecipe = (deckName) => {
    dispatch(loginReset(deckName));
  };

  const handleAbortModalDeckB = () => {
    setConfirmAbortModal(!confirmAbortModal);
    dispatch(abortCleanUpActionInitiated({ deckName: DECKNAME.DeckBShort }));
    dispatch(loginReset(DECKNAME.DeckB));
  };

  return (
    <div className="d-flex justify-content-center align-items-center">
      {confirmAbortModal && (
        <MlModal
          isOpen={confirmAbortModal}
          textHead={deckName}
          textBody={MODAL_MESSAGE.abortCleanupConfirmation}
          successBtn={MODAL_BTN.yes}
          failureBtn={MODAL_BTN.no}
          handleSuccessBtn={
            deckName === DECKNAME.DeckA
              ? handleAbortModalDeckA
              : handleAbortModalDeckB
          }
          handleCrossBtn={() => {
            setConfirmAbortModal(!confirmAbortModal);
          }}
        />
      )}

      <MlModal
        isOpen={isConfirmationModalVisibleDeckA}
        textHead={DECKNAME.DeckA}
        textBody={`Experiment was successful`}
        handleSuccessBtn={() => onNextClickedAfterDoneRecipe(DECKNAME.DeckA)}
        handleCrossBtn={() => toggleComfirmationModal(DECKNAME.DeckA)}
        successBtn={MODAL_BTN.next}
      />

      <MlModal
        isOpen={isConfirmationModalVisibleDeckB}
        textHead={DECKNAME.DeckB}
        textBody={`Experiment was successful`}
        handleSuccessBtn={() => onNextClickedAfterDoneRecipe(DECKNAME.DeckB)}
        handleCrossBtn={() => toggleComfirmationModal(DECKNAME.DeckB)}
        successBtn={MODAL_BTN.next}
      />

      <MlModal
        isOpen={abortConfirmationModalA}
        textHead={DECKNAME.DeckA}
        textBody={MODAL_MESSAGE.abortConfirmation}
        handleSuccessBtn={() => onAbortConfirmed(DECKNAME.DeckA)}
        handleCrossBtn={() => toggleAbortConfirmationModal(DECKNAME.DeckA)}
        successBtn={MODAL_BTN.yes}
        failureBtn={MODAL_BTN.no}
      />

      <MlModal
        isOpen={abortConfirmationModalB}
        textHead={DECKNAME.DeckB}
        textBody={MODAL_MESSAGE.abortConfirmation}
        handleSuccessBtn={() => onAbortConfirmed(DECKNAME.DeckB)}
        handleCrossBtn={() => toggleAbortConfirmationModal(DECKNAME.DeckB)}
        successBtn={MODAL_BTN.yes}
        failureBtn={MODAL_BTN.no}
      />

      {confirmDoneModal && (
        <MlModal
          isOpen={confirmDoneModal}
          textHead={deckName}
          textBody={MODAL_MESSAGE.uvSuccess}
          successBtn={MODAL_BTN.next}
          handleSuccessBtn={
            deckName === DECKNAME.DeckA
              ? handleDoneModalDeckA
              : handleDoneModalDeckB
          }
          handleCrossBtn={() => {
            setConfirmDoneModal(!confirmDoneModal);
          }}
        />
      )}

      {/**Deck A */}
      <DeckCard
        deckName={DECKNAME.DeckA}
        recipeName={
          recipeActionReducerForDeckA.recipeData &&
          recipeActionReducerForDeckA.recipeData.recipeName
            ? recipeActionReducerForDeckA.recipeData.recipeName
            : null
        }
        processNumber={
          recipeActionReducerForDeckA.runRecipeInProgress
            ? JSON.parse(recipeActionReducerForDeckA.runRecipeInProgress)
                .operation_details.current_step
            : 0
        }
        processTotal={
          recipeActionReducerForDeckA.recipeData &&
          recipeActionReducerForDeckA.recipeData.processCount
            ? recipeActionReducerForDeckA.recipeData.processCount
            : null
        }
        isActiveDeck={isDeckAActive}
        loginBtn={!isDeckALoggedIn}
        isAnotherDeckLoggedIn={isDeckBLoggedIn}
        leftActionBtn={
          isDeckALoggedIn
            ? recipeActionReducerForDeckA.showProcess
              ? recipeActionReducerForDeckA.leftActionBtn
              : cleanUpReducerForDeckA.leftActionBtn
            : ""
        }
        rightActionBtn={
          isDeckALoggedIn
            ? recipeActionReducerForDeckA.showProcess
              ? recipeActionReducerForDeckA.rightActionBtn
              : cleanUpReducerForDeckA.rightActionBtn
            : ""
        }
        showProcess={
          isDeckALoggedIn ? recipeActionReducerForDeckA.showProcess : false
        }
        hours={cleanUpReducerForDeckA.hours}
        mins={cleanUpReducerForDeckA.mins}
        secs={cleanUpReducerForDeckA.secs}
        progressPercentComplete={
          isDeckALoggedIn
            ? recipeActionReducerForDeckA.showProcess
              ? recipeActionReducerForDeckA.runRecipeInProgress &&
                JSON.parse(recipeActionReducerForDeckA.runRecipeInProgress)
                  .progress
              : cleanUpReducerForDeckA.cleanUpData &&
                JSON.parse(cleanUpReducerForDeckA.cleanUpData).progress
            : 0
        }
        showCleanUp={
          isDeckALoggedIn ? cleanUpReducerForDeckA.showCleanUp : false
        }
        handleLeftAction={getLeftActionBtnHandler(DECKNAME.DeckA)}
        handleRightAction={getRightActionBtnHandler(DECKNAME.DeckA)}
        leftActionBtnDisabled={cleanUpReducerForDeckA.leftActionBtnDisabled}
        rightActionBtnDisabled={cleanUpReducerForDeckA.rightActionBtnDisabled}
      />

      {/** Deck B */}
      <DeckCard
        deckName={DECKNAME.DeckB}
        recipeName={
          // recipeActionReducerForDeckB &&
          recipeActionReducerForDeckB.recipeData &&
          recipeActionReducerForDeckB.recipeData.recipeName
            ? recipeActionReducerForDeckB.recipeData.recipeName
            : null
        }
        processNumber={
          recipeActionReducerForDeckB.runRecipeInProgress
            ? JSON.parse(recipeActionReducerForDeckB.runRecipeInProgress)
                .operation_details.current_step
            : 0
        }
        processTotal={
          recipeActionReducerForDeckB.recipeData &&
          recipeActionReducerForDeckB.recipeData.processCount
            ? recipeActionReducerForDeckB.recipeData.processCount
            : null
        }
        isActiveDeck={isDeckBActive}
        loginBtn={!isDeckBLoggedIn}
        isAnotherDeckLoggedIn={isDeckALoggedIn}
        leftActionBtn={
          isDeckBLoggedIn
            ? recipeActionReducerForDeckB.showProcess
              ? recipeActionReducerForDeckB.leftActionBtn
              : cleanUpReducerForDeckB.leftActionBtn
            : ""
        }
        rightActionBtn={
          isDeckBLoggedIn
            ? recipeActionReducerForDeckB.showProcess
              ? recipeActionReducerForDeckB.rightActionBtn
              : cleanUpReducerForDeckB.rightActionBtn
            : ""
        }
        showProcess={
          isDeckBLoggedIn ? recipeActionReducerForDeckB.showProcess : false
        }
        hours={cleanUpReducerForDeckB.hours}
        mins={cleanUpReducerForDeckB.mins}
        secs={cleanUpReducerForDeckB.secs}
        progressPercentComplete={
          isDeckBLoggedIn
            ? recipeActionReducerForDeckB.showProcess
              ? recipeActionReducerForDeckB.runRecipeInProgress &&
                JSON.parse(recipeActionReducerForDeckB.runRecipeInProgress)
                  .progress
              : cleanUpReducerForDeckB.cleanUpData &&
                JSON.parse(cleanUpReducerForDeckB.cleanUpData).progress
            : 0
        }
        showCleanUp={cleanUpReducerForDeckB.showCleanUp}
        handleLeftAction={getLeftActionBtnHandler(DECKNAME.DeckB)}
        handleRightAction={getRightActionBtnHandler(DECKNAME.DeckB)}
        leftActionBtnDisabled={cleanUpReducerForDeckB.leftActionBtnDisabled}
        rightActionBtnDisabled={cleanUpReducerForDeckB.rightActionBtnDisabled}
      />
    </div>
  );
};

/**old code for reference */
// const AppFooter = (props) => {
//   const {
//     loginBtn,
//     showProcess,
//     showCleanUp,
//     deckName,
//     recipeName,
//     processNumber,
//     processTotal,
//     hours,
//     mins,
//     secs,
//     handleLeftAction,
//     handleRightAction,
//     leftActionBtn,
//     rightActionBtn,
//     progressPercentComplete,
//     leftActionBtnDisabled,
//     rightActionBtnDisabled,
//   } = props;

//   return deckName === DECKNAME.DeckA ? (
//     <div className="d-flex justify-content-center align-items-center">
//       <DeckCard
//         deckName={DECKNAME.DeckA}//
//         recipeName={recipeName}//
//         processNumber={processNumber}//
//         processTotal={processTotal}//
//         hours={hours}//
//         mins={mins}//
//         secs={secs}//
//         loginBtn={loginBtn}//
//         showProcess={showProcess}//
//         showCleanUp={showCleanUp}//
//         handleLeftAction={handleLeftAction}//
//         handleRightAction={handleRightAction}//
//         leftActionBtn={leftActionBtn}//
//         rightActionBtn={rightActionBtn}//
//         progressPercentComplete={progressPercentComplete}//
//         leftActionBtnDisabled={leftActionBtnDisabled}//
//         rightActionBtnDisabled={rightActionBtnDisabled}//
//       />
//       <DeckCard deckName={DECKNAME.DeckB} loginBtn={true} />
//     </div>
//   ) : (
//     <div className="d-flex justify-content-center align-items-center">
//       <DeckCard deckName={DECKNAME.DeckA} loginBtn={true} />
//       <DeckCard
//         deckName={DECKNAME.DeckB}
//         recipeName={recipeName}
//         processNumber={processNumber}
//         processTotal={processTotal}
//         hours={hours}
//         mins={mins}
//         secs={secs}
//         loginBtn={loginBtn}
//         showProcess={showProcess}
//         showCleanUp={showCleanUp}
//         handleLeftAction={handleLeftAction}
//         handleRightAction={handleRightAction}
//         leftActionBtn={leftActionBtn}
//         rightActionBtn={rightActionBtn}
//         progressPercentComplete={progressPercentComplete}
//         leftActionBtnDisabled={leftActionBtnDisabled}
//         rightActionBtnDisabled={rightActionBtnDisabled}
//       />
//     </div>
//   );
// };

AppFooter.propTypes = {};

AppFooter.defaultProps = {
  loginBtn: false,
};

export default AppFooter;
