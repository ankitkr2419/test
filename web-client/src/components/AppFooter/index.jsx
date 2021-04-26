import React from "react";
import DeckCard from "shared-components/DeckCard";
import { DECKNAME } from "appConstants";
import { useSelector } from "react-redux";

const AppFooter = (props) => {
    const loginReducer = useSelector((state) => state.loginReducer);
    const loginReducerData = loginReducer.toJS();
    let activeDeckObj =
        loginReducerData &&
        loginReducerData.decks.find((deck) => deck.isActive);
    let loginDataOfA =
        loginReducerData &&
        loginReducerData.decks.find((deck) => deck.name === DECKNAME.DeckA);
    let isDeckALoggedIn = loginDataOfA.isLoggedIn;
    let isDeckAActive = loginDataOfA.isActive
    let loginDataOfB =
        loginReducerData &&
        loginReducerData.decks.find((deck) => deck.name === DECKNAME.DeckB);
    let isDeckBLoggedIn = loginDataOfB.isLoggedIn;
    let isDeckBActive = loginDataOfB.isActive

    const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
    let recipeActionReducerForDeckA = recipeActionReducer.decks.find(deckObj => deckObj.name === DECKNAME.DeckA)
    let recipeActionReducerForDeckB = recipeActionReducer.decks.find(deckObj => deckObj.name === DECKNAME.DeckB)
    
    return (
        <div className="d-flex justify-content-center align-items-center">
            <DeckCard
                deckName={DECKNAME.DeckA}
                isActiveDeck={isDeckAActive}
                loginBtn={!isDeckALoggedIn}
                isAnotherDeckLoggedIn={isDeckBLoggedIn}
                leftActionBtn={isDeckALoggedIn ? recipeActionReducerForDeckA.leftActionBtn : ''}
                rightActionBtn={isDeckALoggedIn ? recipeActionReducerForDeckA.rightActionBtn: ''}
            />
            <DeckCard
                deckName={DECKNAME.DeckB}
                isActiveDeck={isDeckBActive}
                loginBtn={!isDeckBLoggedIn}
                isAnotherDeckLoggedIn={isDeckALoggedIn}
                leftActionBtn={recipeActionReducerForDeckB.leftActionBtn}
                rightActionBtn={recipeActionReducerForDeckB.rightActionBtn}
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
//         deckName={DECKNAME.DeckA}
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
