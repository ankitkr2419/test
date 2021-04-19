import React from "react";
import DeckCard from "shared-components/DeckCard";
import { DECKNAME } from "appConstants";

const AppFooter = (props) => {
  const {
    loginBtn,
    showProcess,
    showCleanUp,
    deckName,
    recipeName,
    processNumber,
    processTotal,
    handleLeftAction,
    handleRightAction,
    leftActionBtn,
    rightActionBtn,
    progressPercentComplete,
  } = props;

  return deckName === DECKNAME.DeckA ? (
    <div className="d-flex justify-content-center align-items-center">
      <DeckCard
        deckName={DECKNAME.DeckA}
        recipeName={recipeName}
        processNumber={processNumber}
        processTotal={processTotal}
        loginBtn={loginBtn}
        showProcess={showProcess}
        showCleanUp={showCleanUp}
        handleLeftAction={handleLeftAction}
        handleRightAction={handleRightAction}
        leftActionBtn={leftActionBtn}
        rightActionBtn={rightActionBtn}
        progressPercentComplete={progressPercentComplete}
      />
      <DeckCard deckName={DECKNAME.DeckB} loginBtn={true} />
    </div>
  ) : (
    <div className="d-flex justify-content-center align-items-center">
      <DeckCard deckName={DECKNAME.DeckA} loginBtn={true} />
      <DeckCard
        deckName={DECKNAME.DeckB}
        recipeName={recipeName}
        processNumber={processNumber}
        processTotal={processTotal}
        loginBtn={loginBtn}
        showProcess={showProcess}
        showCleanUp={showCleanUp}
        handleLeftAction={handleLeftAction}
        handleRightAction={handleRightAction}
        leftActionBtn={leftActionBtn}
        rightActionBtn={rightActionBtn}
        progressPercentComplete={progressPercentComplete}
      />
    </div>
  );
};

AppFooter.propTypes = {};

AppFooter.defaultProps = {
  loginBtn: false,
};

export default AppFooter;
