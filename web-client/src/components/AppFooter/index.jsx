import React from "react";
import DeckCard from "shared-components/DeckCard";
import { DECKNAME } from "appConstants";
import { useSelector } from "react-redux";
import { DECKCARD_BTN } from "appConstants";

const AppFooter = (props) => {
    //login reducer
    const loginReducer = useSelector((state) => state.loginReducer);
    const loginReducerData = loginReducer.toJS();
    let activeDeckObj =
        loginReducerData &&
        loginReducerData.decks.find((deck) => deck.isActive);
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
    const recipeActionReducer = useSelector(
        (state) => state.recipeActionReducer
    );
    let recipeActionReducerForDeckA = recipeActionReducer.decks.find(
        (deckObj) => deckObj.name === DECKNAME.DeckA
    );
    let recipeActionReducerForDeckB = recipeActionReducer.decks.find(
        (deckObj) => deckObj.name === DECKNAME.DeckB
    );

    const tempFunction = () => {//remove later//check
        alert("in dev");
    };

    const getLeftActionBtnHandler = (deckName) => {
        return tempFunction;
        // switch (
        //     showProcess
        //         ? recipeActionReducer.leftActionBtn
        //         : cleanUpReducer.leftActionBtn
        // ) {
        //     case DECKCARD_BTN.text.run:
        //         return handleRunAction;
        //     case DECKCARD_BTN.text.pause:
        //         return handlePauseAction;
        //     case DECKCARD_BTN.text.resume:
        //         return handleResumeAction;

        //     case DECKCARD_BTN.text.startUv:
        //         return handleRunAction;
        //     case DECKCARD_BTN.text.pauseUv:
        //         return handlePauseAction;
        //     case DECKCARD_BTN.text.resumeUv:
        //         return handleResumeAction;

        //     case DECKCARD_BTN.text.done:
        //         return handleDoneAction;

        //     default:
        //         break;
        // }
    };

    const handleRunAction = () => {
        // const name = deckName === "Deck A" ? "A" : "B";
        // if (showProcess) {
        //     const { recipeId } = recipeData;
        //     dispatch(
        //         runRecipeInitiated({
        //             recipeId: recipeId,
        //             deckName: name,
        //         })
        //     );
        // } else {
        //     dispatch(
        //         runCleanUpActionInitiated({
        //             time: `${hours}:${mins}:${secs}`,
        //             deckName: name,
        //         })
        //     );
        // }
    };

    const handlePauseAction = () => {
        // const name = deckName === "Deck A" ? "A" : "B";
        // if (showProcess) {
        //     dispatch(pauseRecipeInitiated({ deckName: name }));
        // } else {
        //     dispatch(pauseCleanUpActionInitiated({ deckName: name }));
        // }
    };

    const handleResumeAction = () => {
        // const name = deckName === "Deck A" ? "A" : "B";
        // if (showProcess) {
        //     dispatch(resumeRecipeInitiated({ deckName: name }));
        // } else {
        //     dispatch(resumeCleanUpActionInitiated({ deckName: name }));
        // }
    };

    const handleDoneAction = () => {
        // setShowProcess(!showProcess);
        // // setLeftActionBtn(DECKCARD_BTN.text.run);
        // //setRightActionBtn(DECKCARD_BTN.text.cancel);
    };

    const getRightActionBtnHandler = () => {
        return tempFunction;
        // switch (
        //     showProcess
        //         ? recipeActionReducer.rightActionBtn
        //         : cleanUpReducer.rightActionBtn
        // ) {
        //     case DECKCARD_BTN.text.abort:
        //         return handleAbortAction;
        //     case DECKCARD_BTN.text.cancel:
        //         return handleCancelAction;
        //     default:
        //         break;
        // }
    };

    const handleCancelAction = () => {
        // setShowProcess(false);
        // setShowCleanUp(false);
    };

    const handleAbortAction = () => {
        // setConfirmationModal(true);
    };

    return (
        <div className="d-flex justify-content-center align-items-center">
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
                        ? JSON.parse(
                              recipeActionReducerForDeckA.runRecipeInProgress
                          ).operation_details.current_step
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
                        ? recipeActionReducerForDeckA.leftActionBtn
                        : ""
                }
                rightActionBtn={
                    isDeckALoggedIn
                        ? recipeActionReducerForDeckA.rightActionBtn
                        : ""
                }
                showProcess={
                    isDeckALoggedIn
                        ? recipeActionReducerForDeckA.showProcess
                        : false
                }
                hours={recipeActionReducerForDeckA.hours}
                mins={recipeActionReducerForDeckA.mins}
                secs={recipeActionReducerForDeckA.secs}
                progressPercentComplete={
                    recipeActionReducerForDeckA.runRecipeInProgress
                        ? JSON.parse(
                              recipeActionReducerForDeckA.runRecipeInProgress
                          ).progress
                        : 0
                }
                showCleanUp={recipeActionReducerForDeckA.showCleanUp}
                handleLeftAction={getLeftActionBtnHandler(DECKNAME.DeckB)}
                handleRightAction={getRightActionBtnHandler(DECKNAME.DeckB)}
                leftActionBtnDisabled={
                    recipeActionReducerForDeckA.leftActionBtnDisabled
                }
                rightActionBtnDisabled={
                    recipeActionReducerForDeckA.rightActionBtnDisabled
                }
            />

            {/** Deck B */}
            <DeckCard
                deckName={DECKNAME.DeckB}
                recipeName={
                    recipeActionReducerForDeckB &&
                    recipeActionReducerForDeckB.recipeData &&
                    recipeActionReducerForDeckB.recipeData.recipeName
                        ? recipeActionReducerForDeckB.recipeData.recipeName
                        : null
                }
                processNumber={
                    recipeActionReducerForDeckB.runRecipeInProgress
                        ? JSON.parse(
                              recipeActionReducerForDeckB.runRecipeInProgress
                          ).operation_details.current_step
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
                    isDeckALoggedIn
                        ? recipeActionReducerForDeckB.leftActionBtn
                        : ""
                }
                rightActionBtn={
                    isDeckALoggedIn
                        ? recipeActionReducerForDeckB.rightActionBtn
                        : ""
                }
                showProcess={
                    isDeckALoggedIn
                        ? recipeActionReducerForDeckB.showProcess
                        : false
                }
                hours={recipeActionReducerForDeckB.hours}
                mins={recipeActionReducerForDeckB.mins}
                secs={recipeActionReducerForDeckB.secs}
                progressPercentComplete={
                    recipeActionReducerForDeckB.runRecipeInProgress
                        ? JSON.parse(
                              recipeActionReducerForDeckB.runRecipeInProgress
                          ).progress
                        : 0
                }
                showCleanUp={recipeActionReducerForDeckA.showCleanUp}
                handleLeftAction={getLeftActionBtnHandler(DECKNAME.DeckB)}
                handleRightAction={getRightActionBtnHandler(DECKNAME.DeckB)}
                leftActionBtnDisabled={
                    recipeActionReducerForDeckB.leftActionBtnDisabled
                }
                rightActionBtnDisabled={
                    recipeActionReducerForDeckB.rightActionBtnDisabled
                }
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
//         leftActionBtnDisabled={leftActionBtnDisabled}
//         rightActionBtnDisabled={rightActionBtnDisabled}
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
