import React, { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import { Redirect } from "react-router-dom";

import { Card, CardBody, Button, Row, Col } from "core-components";
import { Icon, MlModal, VideoCard } from "shared-components";

import styled from "styled-components";
import AppFooter from "components/AppFooter";
import RecipeFlowModal from "components/modals/RecipeFlowModal";
// import ConfirmationModal from "components/modals/ConfirmationModal";
import TrayDiscardModal from "components/modals/TrayDiscardModal";
import RecipeCard from "components/RecipeListing/RecipeCard";
import TimeModal from "components/modals/TimeModal";
import {
  abortRecipeInitiated,
  pauseRecipeInitiated,
  resumeRecipeInitiated,
  runRecipeInitiated,
  runRecipeReset,
  pauseRecipeReset,
  resumeRecipeReset,
  abortRecipeReset,
} from "action-creators/recipeActionCreators";
import { operatorLoginReset } from "action-creators/operatorLoginModalActionCreators";
import {
  discardTipAndHomingActionInitiated,
  discardTipAndHomingActionReset,
} from "action-creators/homingActionCreators";
import { DECKCARD_BTN, MODAL_BTN, MODAL_MESSAGE, ROUTES } from "appConstants";
import PaginationBox from "shared-components/PaginationBox";
import TipDiscardModal from "components/modals/TipDiscardModal";
import { cleanUpActionInitiated } from "action-creators/cleanUpActionCreators";

const RecipeListing = styled.div`
  .landing-content {
    padding: 1.25rem 4.5rem 0.875rem 4.5rem;
    &::after {
      height: 9.125rem;
    }
    .recipe-listing-cards {
      height: 30.75rem;
    }
  }
`;
const TopContent = styled.div`
  margin-bottom: 2.25rem;
  .icon-download-1 {
    font-size: 18px;
    color: #3c3c3c;
  }
  .btn-clean-up {
    width: 7.063rem;
  }
`;

const HeadingTitle = styled.label`
  font-size: 1.25rem;
  line-height: 1.438rem;
`;
const RecipeListingComponent = (props) => {
  const { allRecipeData } = props;
  const [timeModal, setTimeModal] = useState(false);
  const [trayDiscardModal, setTrayDiscardModal] = useState(false);
  const [hours, setHours] = useState(0);
  const [mins, setMins] = useState(0);
  const [secs, setSecs] = useState(0);

  const dispatch = useDispatch();

  const discardTipAndHomingReducer = useSelector(
    (state) => state.discardTipAndHomingReducer
  );
  const { error, serverErrors } = discardTipAndHomingReducer;
  const tipDiscardHomingError = error;
  const tipDiscardServerErrors = serverErrors;

  const operatorLoginModalReducer = useSelector(
    (state) => state.operatorLoginModalReducer
  );
  const { deckName } = operatorLoginModalReducer.toJS();

  const handleTimeModal = () => {
    setTimeModal(!timeModal);
  };

  const handleTrayDiscardModal = () => {
    setTrayDiscardModal(!trayDiscardModal);
  };

  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  const {
    runRecipeError,
    abortRecipeError,
    pauseRecipeError,
    resumeRecipeError,
    // recipeListingError,
    leftActionBtn,
    rightActionBtn,
    // isLoading,
  } = recipeActionReducer;

  const [redirect, setRedirect] = useState(false);
  const [confirmationModal, setConfirmationModal] = useState(false);
  const [tipDiscardModal, setTipDiscardModal] = useState(false);
  const [recipeData, setRecipeData] = useState({});
  const [progressPercentComplete, setProgressPercentComplete] = useState(0);

  const [isOpen, setIsOpen] = useState(false);
  const toggle = (recipeId, recipeName, processCount) => {
    const data = {
      recipeId: recipeId,
      recipeName: recipeName,
      processCount: processCount,
    };
    setRecipeData(data);
    setIsOpen(!isOpen);
  };

  const [showProcess, setShowProcess] = useState(false);
  const toggleShowProcess = () => {
    setShowProcess(!showProcess);
    setIsOpen(!isOpen);
  };

  //Do not change '===';
  useEffect(() => {
    if (tipDiscardHomingError === false) {
      dispatch(operatorLoginReset());
      setTipDiscardModal(false);
      setRedirect(true);
    } else if (tipDiscardHomingError === true) {
      //show toast error msg for tip discard and homing error
      console.log("Error in tip discard and homing: ", tipDiscardServerErrors);
    }

    if (abortRecipeError === false) {
      setConfirmationModal(false);
      setTipDiscardModal(true);
    }

    if (
      runRecipeError ||
      pauseRecipeError ||
      resumeRecipeError ||
      abortRecipeError
    ) {
      //show toast with error msg
    }

    dispatch(discardTipAndHomingActionReset());
    dispatch(abortRecipeReset());
    dispatch(runRecipeReset());
    dispatch(resumeRecipeReset());
    dispatch(pauseRecipeReset());
  }, [
    tipDiscardServerErrors,
    tipDiscardHomingError,
    runRecipeError,
    pauseRecipeError,
    resumeRecipeError,
    abortRecipeError,
    confirmationModal,
    dispatch,
  ]);

  const handleRunAction = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    const { recipeId } = recipeData;
    dispatch(runRecipeInitiated({ recipeId: recipeId, deckName: name }));
  };

  const handlePauseAction = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    dispatch(pauseRecipeInitiated({ deckName: name }));
  };

  const handleResumeAction = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    dispatch(resumeRecipeInitiated({ deckName: name }));
  };

  const handleDoneAction = () => {
    setShowProcess(!showProcess);
    // setLeftActionBtn(DECKCARD_BTN.text.run);
    // setRightActionBtn(DECKCARD_BTN.text.cancel);
    setProgressPercentComplete(0);
  };

  const handleChangeTime = (event) => {
    let name = event.target.name;
    if (name === "hours") {
      setHours(event.target.value);
    } else if (name === "minutes") {
      setMins(event.target.value);
    } else {
      setSecs(event.target.value);
    }
  };

  const submitTime = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    dispatch(
      cleanUpActionInitiated({
        time: `${hours}:${mins}:${secs}`,
        deckName: name,
      })
    );
  };

  const handleCancelAction = () => setShowProcess(!showProcess);
  const handleAbortAction = () => setConfirmationModal(true);

  const toggleConfirmModal = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    dispatch(abortRecipeInitiated({ deckName: name }));
  };

  const toggleTipDiscardModal = (discardTip) => {
    const name = deckName === "Deck A" ? "A" : "B";
    dispatch(
      discardTipAndHomingActionInitiated({
        deckName: name,
        discardTip: discardTip,
      })
    );
  };

  const getLeftActionBtnHandler = () => {
    switch (leftActionBtn) {
      case DECKCARD_BTN.text.run:
        return handleRunAction;
      case DECKCARD_BTN.text.pause:
        return handlePauseAction;
      case DECKCARD_BTN.text.resume:
        return handleResumeAction;
      case DECKCARD_BTN.text.done:
        return handleDoneAction;
      default:
        break;
    }
  };

  const getRightActionBtnHandler = () => {
    switch (rightActionBtn) {
      case DECKCARD_BTN.text.abort:
        return handleAbortAction;
      case DECKCARD_BTN.text.cancel:
        return handleCancelAction;
      default:
        break;
    }
  };

  if (redirect) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }
  return (
    <RecipeListing>
      <div className="landing-content px-2">
        <RecipeFlowModal
          isOpen={isOpen}
          toggle={toggle}
          toggleShowProcess={toggleShowProcess}
          recipeData={recipeData}
        />

        <TipDiscardModal
          isOpen={tipDiscardModal}
          handleSuccessBtn={toggleTipDiscardModal}
          deckName={deckName}
        />

        <MlModal
          isOpen={confirmationModal}
          textHead={deckName}
          textBody={MODAL_MESSAGE.abortConfirmation}
          handleSuccessBtn={toggleConfirmModal}
          handleCrossBtn={() => setConfirmationModal(!confirmationModal)}
          successBtn={MODAL_BTN.yes}
          failureBtn={MODAL_BTN.no}
        />
        {timeModal && (
          <TimeModal
            timeModal={timeModal}
            toggleTimeModal={handleTimeModal}
            hours={hours}
            mins={mins}
            secs={secs}
            handleChangeTime={handleChangeTime}
            submitTime={submitTime}
            deckName={deckName}
          />
        )}

        {trayDiscardModal && (
          <TrayDiscardModal
            trayDiscardModal={trayDiscardModal}
            toggleTrayDiscardModal={handleTrayDiscardModal}
            deckName={deckName}
          />
        )}

        <TopContent className="d-flex justify-content-between align-items-center mx-5">
          <div className="d-flex align-items-center">
            <Icon name="angle-left" size={32} className="text-white" />
            <HeadingTitle
              Tag="h5"
              className="text-white font-weight-bold ml-3 mb-0"
            >
              Select a Recipe for Deck B
            </HeadingTitle>
          </div>
          <div className="">
            <Icon name="download" size={19} className="text-white mr-3" />
            <Button
              color="secondary"
              className="ml-auto"
              onClick={handleTimeModal}
            >
              {" "}
              Clean Up
            </Button>

            <Button
              color="secondary"
              className="ml-2 border-primary btn-discard-tray"
              onClick={handleTrayDiscardModal}
            >
              Discard Tray
            </Button>
          </div>
        </TopContent>

        {showProcess ? (
          <VideoCard />
        ) : (
          <Card className="recipe-listing-cards">
            <CardBody className="p-5">
              <div className="d-flex justify-content-end">
                <PaginationBox />
              </div>
              <Row>
                {allRecipeData && allRecipeData.length > 0 ? (
                  allRecipeData.map((value, index) => (
                    <Col md={6} key={index}>
                      <RecipeCard
                        recipeId={value.id}
                        recipeName={value.name}
                        processCount={value.process_count}
                        toggle={toggle}
                      />
                    </Col>
                  ))
                ) : (
                  <h4>No recipes to show!</h4>
                )}
              </Row>
            </CardBody>
          </Card>
        )}
      </div>
      <AppFooter
        deckName={deckName}
        showProcess={showProcess}
        recipeName={recipeData.recipeName}
        processNumber={12}
        processTotal={recipeData.processCount}
        progressPercentComplete={progressPercentComplete}
        handleLeftAction={getLeftActionBtnHandler()}
        handleRightAction={getRightActionBtnHandler()}
        leftActionBtn={leftActionBtn}
        rightActionBtn={rightActionBtn}
      />
    </RecipeListing>
  );
};

RecipeListingComponent.propTypes = {};

export default RecipeListingComponent;
