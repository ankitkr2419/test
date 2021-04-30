import React from "react";
import DeckCard from "shared-components/DeckCard";
import { DECKNAME } from "appConstants";
import { useSelector, useDispatch } from "react-redux";
import { DECKCARD_BTN } from "appConstants";
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

const AppFooter = (props) => {
    const dispatch = useDispatch();

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

    const getLeftActionBtnHandler = (deckName) => {
        let recipeReducerData =
            deckName === DECKNAME.DeckA
                ? recipeActionReducerForDeckA
                : recipeActionReducerForDeckB;
        let showProcess = recipeReducerData.showProcess;

        let cleanUpReducerData = { leftActionBtn: undefined }; //check //change (use cleanUpReducer data here)

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
            console.log("cleanUp in development");
            // dispatch(
            //     runCleanUpActionInitiated({
            //         time: `${recipeReducerData.hours}:${recipeReducerData.mins}:${recipeReducerData.secs}`,
            //         deckName: recipeReducerData.name,
            //     })
            // );
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
            console.log("cleanUp in development");
            // dispatch(
            //     runCleanUpActionInitiated({
            //         time: `${recipeReducerData.hours}:${recipeReducerData.mins}:${recipeReducerData.secs}`,
            //         deckName: recipeReducerData.name,
            //     })
            // );
        }
    };

    const handlePauseActionDeckA = () => {
        let recipeReducerData = recipeActionReducerForDeckA;

        if (recipeReducerData.showProcess) {
            dispatch(
                pauseRecipeInitiated({ deckName: recipeReducerData.name })
            );
        } else {
            //dispatch(pauseCleanUpActionInitiated({ deckName: recipeReducerData.name }));
        }
    };
    const handlePauseActionDeckB = () => {
        let recipeReducerData = recipeActionReducerForDeckB;

        if (recipeReducerData.showProcess) {
            dispatch(
                pauseRecipeInitiated({ deckName: recipeReducerData.name })
            );
        } else {
            //dispatch(pauseCleanUpActionInitiated({ deckName: recipeReducerData.name }));
        }
    };

    const handleResumeActionDeckA = () => {
        let recipeReducerData = recipeActionReducerForDeckA;

        if (recipeReducerData.showProcess) {
            dispatch(
                resumeRecipeInitiated({ deckName: recipeReducerData.name })
            );
        } else {
            //dispatch(resumeCleanUpActionInitiated({ deckName: recipeReducerData.name }));
        }
    };

    const handleResumeActionDeckB = () => {
        let recipeReducerData = recipeActionReducerForDeckB;

        if (recipeReducerData.showProcess) {
            dispatch(
                resumeRecipeInitiated({ deckName: recipeReducerData.name })
            );
        } else {
            //dispatch(resumeCleanUpActionInitiated({ deckName: recipeReducerData.name }));
        }
    };

    const handleDoneActionDeckA = () => {
        let recipeReducerData = recipeActionReducerForDeckA;
        dispatch(
            updateRecipeReducerDataForDeck(recipeReducerData.name, {
                showProcess: !recipeReducerData.showProcess,
            })
        );
    };
    const handleDoneActionDeckB = () => {
        let recipeReducerData = recipeActionReducerForDeckB;
        dispatch(
            updateRecipeReducerDataForDeck(recipeReducerData.name, {
                showProcess: !recipeReducerData.showProcess,
            })
        );
    };

    const getRightActionBtnHandler = (deckName) => {
        let recipeReducerData =
            deckName === DECKNAME.DeckA
                ? recipeActionReducerForDeckA
                : recipeActionReducerForDeckB;
        let showProcess = recipeReducerData.showProcess;

        let cleanUpReducerData = { leftActionBtn: undefined }; //check //change (use cleanUpReducer data here)

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
    const handleCancelActionDeckA = () => {
        let recipeReducerData = recipeActionReducerForDeckA;
        let deckName = recipeReducerData.name;
        dispatch(runRecipeReset(deckName));
        // setShowCleanUp(false); //check
    };
    const handleCancelActionDeckB = () => {
        let recipeReducerData = recipeActionReducerForDeckB;
        let deckName = recipeReducerData.name;
        dispatch(runRecipeReset(deckName));
        // setShowCleanUp(false);//check
    };

    const handleAbortActionDeckA = () => {
        //check
        // setConfirmationModal(true);
    };

    const handleAbortActionDeckB = () => {
        //check
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
                handleLeftAction={getLeftActionBtnHandler(DECKNAME.DeckA)}
                handleRightAction={getRightActionBtnHandler(DECKNAME.DeckA)}
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
