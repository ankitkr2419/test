import React, { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import { Redirect } from "react-router-dom";
import { toast } from "react-toastify";

import { Card, CardBody, Button, Row, Col } from "core-components";
import { Icon, MlModal, VideoCard, ButtonIcon } from "shared-components";

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
import {
  DECKCARD_BTN,
  MODAL_BTN,
  MODAL_MESSAGE,
  ROUTES,
  TOAST_MESSAGE,
} from "appConstants";
import PaginationBox from "shared-components/PaginationBox";
import TipDiscardModal from "components/modals/TipDiscardModal";
import {
  abortCleanUpActionReset,
  pauseCleanUpActionInitiated,
  resumeCleanUpActionInitiated,
  runCleanUpActionInitiated,
} from "action-creators/cleanUpActionCreators";
import {
  discardDeckInitiated,
  discardDeckReset,
} from "action-creators/discardDeckActionCreators";
import {
  restoreDeckInitiated,
  restoreDeckReset,
} from "action-creators/restoreDeckActionCreators";

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
    font-size: 1.125rem;
    color: #3c3c3c;
  }
  .btn-clean-up {
    width: 7.063rem;
  }
  .btn-discard-tray {
    width: 10rem;
  }
  .icon-logout {
    font-size: 1rem;
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
  const [confirmationModal, setConfirmationModal] = useState(false);
  const [tipDiscardModal, setTipDiscardModal] = useState(false);
  const [recipeFlowModal, setRecipeFlowModal] = useState(false);

  const [hours, setHours] = useState(0);
  const [mins, setMins] = useState(0);
  const [secs, setSecs] = useState(0);

  const [redirect, setRedirect] = useState(false);
  const [recipeData, setRecipeData] = useState({});

  const [showProcess, setShowProcess] = useState(false);
  const [showCleanUp, setShowCleanUp] = useState(false);

  const dispatch = useDispatch();

  const discardTipAndHomingReducer = useSelector(
    (state) => state.discardTipAndHomingReducer
  );
  const { discardTipAndHomingError } = discardTipAndHomingReducer;
  const discardTipServerErrors = discardTipAndHomingReducer.serverErrors;

  // const operatorLoginModalReducer = useSelector(
  //   (state) => state.operatorLoginModalReducer
  // );
  // const { deckName } = operatorLoginModalReducer.toJS();
  const loginReducer = useSelector((state) => state.loginReducer);

  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  let deckName = activeDeckObj ? activeDeckObj.name : "";

  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  const {
    runRecipeError,
    abortRecipeError,
    pauseRecipeError,
    resumeRecipeError,
    runRecipeInProgress,
  } = recipeActionReducer;

  const cleanUpReducer = useSelector((state) => state.cleanUpReducer);
  const { cleanUpApiError } = cleanUpReducer;
  const discardDeckReducer = useSelector((state) => state.discardDeckReducer);
  const { discardDeckError } = discardDeckReducer;
  const discardDeckServerErrors = discardDeckReducer.serverErrors;

  const restoreDeckReducer = useSelector((state) => state.restoreDeckReducer);
  const { restoreDeckError } = restoreDeckReducer;
  const restoreDeckServerErrors = restoreDeckReducer.serverErrors;

  const [
    switchTrayDiscardModalContents,
    setSwitchTrayDiscardModalContents,
  ] = useState(true);

  const toggle = (recipeId, recipeName, processCount) => {
    const data = {
      recipeId: recipeId,
      recipeName: recipeName,
      processCount: processCount,
    };
    setRecipeData(data);
    setRecipeFlowModal(!recipeFlowModal);
  };

  const toggleShowProcess = () => {
    setShowProcess(true);
    setShowCleanUp(false);
    setRecipeFlowModal(false);
  };

  //Do not change '===';
  useEffect(() => {
    if (cleanUpApiError === false) {
      setTimeModal(false);
      setShowCleanUp(true);
      setShowProcess(false);
    }

    if (discardDeckError === false) {
      setSwitchTrayDiscardModalContents(false);
    } else if (discardDeckError === true) {
      toast.error(`${discardDeckServerErrors}`);
    }

    if (restoreDeckError === false) {
      setTrayDiscardModal(false);
      setRedirect(true);
    } else if (restoreDeckError === true) {
      toast.error(`${restoreDeckServerErrors}`);
    }

    if (discardTipAndHomingError === false) {
      dispatch(operatorLoginReset());
      setTipDiscardModal(false);
      setRedirect(true);
    } else if (discardTipAndHomingError === true) {
      toast.error(`${discardTipServerErrors}`);
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
      toast.error(TOAST_MESSAGE.error);
    }

    dispatch(restoreDeckReset());
    dispatch(discardDeckReset());
    dispatch(discardTipAndHomingActionReset());
    dispatch(abortRecipeReset());
    dispatch(runRecipeReset());
    dispatch(resumeRecipeReset());
    dispatch(pauseRecipeReset());
  }, [
    restoreDeckServerErrors,
    discardDeckServerErrors,
    restoreDeckError,
    discardDeckError,
    discardTipServerErrors,
    discardTipAndHomingError,
    runRecipeError,
    pauseRecipeError,
    resumeRecipeError,
    abortRecipeError,
    confirmationModal,
    runRecipeInProgress,
    cleanUpApiError,
    dispatch,
  ]);

  const handleRunAction = () => {
    const name = deckName === "Deck A" ? "A" : "B";

    if (showProcess) {
      const { recipeId } = recipeData;
      dispatch(
        runRecipeInitiated({
          recipeId: recipeId,
          deckName: name,
        })
      );
    } else {
      dispatch(
        runCleanUpActionInitiated({
          time: `${hours}:${mins}:${secs}`,
          deckName: name,
        })
      );
    }
  };

  const handlePauseAction = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    if (showProcess) {
      dispatch(pauseRecipeInitiated({ deckName: name }));
    } else {
      dispatch(pauseCleanUpActionInitiated({ deckName: name }));
    }
  };

  const handleResumeAction = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    if (showProcess) {
      dispatch(resumeRecipeInitiated({ deckName: name }));
    } else {
      dispatch(resumeCleanUpActionInitiated({ deckName: name }));
    }
  };

  const handleDoneAction = () => {
    setShowProcess(!showProcess);
    // setLeftActionBtn(DECKCARD_BTN.text.run);
    // setRightActionBtn(DECKCARD_BTN.text.cancel);
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
    setTimeModal(false);
    setShowCleanUp(true);
  };

  const handleSuccessBtn = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    if (switchTrayDiscardModalContents) {
      dispatch(discardDeckInitiated({ deckName: name }));
    } else {
      dispatch(restoreDeckInitiated({ deckName: name }));
    }
  };

  const handleCancelAction = () => {
    setShowProcess(false);
    setShowCleanUp(false);
  };
  const handleAbortAction = () => setConfirmationModal(true);

  const toggleConfirmModal = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    if (showProcess) {
      dispatch(abortRecipeInitiated({ deckName: name }));
    } else {
      dispatch(operatorLoginReset());
      dispatch(abortCleanUpActionReset());
      setRedirect(true);
    }
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

  const handleTimeModal = () => {
    setTimeModal(!timeModal);
  };

  const handleTrayDiscardModal = () => {
    setTrayDiscardModal(!trayDiscardModal);
    setSwitchTrayDiscardModalContents(true);
  };

  const getLeftActionBtnHandler = () => {
    switch (
      showProcess
        ? recipeActionReducer.leftActionBtn
        : cleanUpReducer.leftActionBtn
    ) {
      case DECKCARD_BTN.text.run:
        return handleRunAction;
      case DECKCARD_BTN.text.pause:
        return handlePauseAction;
      case DECKCARD_BTN.text.resume:
        return handleResumeAction;

      case DECKCARD_BTN.text.startUv:
        return handleRunAction;
      case DECKCARD_BTN.text.pauseUv:
        return handlePauseAction;
      case DECKCARD_BTN.text.resumeUv:
        return handleResumeAction;

      case DECKCARD_BTN.text.done:
        return handleDoneAction;

      default:
        break;
    }
  };

  const getRightActionBtnHandler = () => {
    switch (
      showProcess
        ? recipeActionReducer.rightActionBtn
        : cleanUpReducer.rightActionBtn
    ) {
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
          isOpen={recipeFlowModal}
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
            handleSuccessBtn={handleSuccessBtn}
            switchModalContent={switchTrayDiscardModalContents}
            deckName={deckName}
          />
        )}

        <TopContent className="d-flex justify-content-between align-items-center mx-5">
          {showProcess || showCleanUp ? null : (
            <div className="d-flex align-items-center">
              <Icon name="angle-left" size={32} className="text-white" />
              <HeadingTitle
                Tag="h5"
                className="text-white font-weight-bold ml-3 mb-0"
              >
                Select a Recipe for Deck B
              </HeadingTitle>
            </div>
          )}
          {showProcess || showCleanUp ? null : (
            <div className="d-flex align-items-center ml-auto">
              <ButtonIcon
                name="download-1"
                size={28}
                className="bg-white border-primary"
              />
              <Button
                color="secondary"
                className="ml-2 border-primary btn-clean-up bg-white"
                onClick={handleTimeModal}
              >
                {" "}
                Clean Up
              </Button>

              <Button
                color="secondary"
                className="ml-2 border-primary btn-discard-tray bg-white"
                onClick={handleTrayDiscardModal}
              >
                Discard Tray
              </Button>

              <ButtonIcon
                name="logout"
                size={28}
                className="ml-2 bg-white border-primary"
              />
            </div>
          )}
        </TopContent>

        {showProcess || showCleanUp ? (
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
        recipeName={showProcess ? recipeData.recipeName : "Clean Up"}
        showProcess={showProcess}
        processNumber={
          runRecipeInProgress
            ? JSON.parse(runRecipeInProgress).operation_details.current_step
            : 0
        }
        processTotal={recipeData.processCount}
        showCleanUp={showCleanUp}
        hours={hours}
        mins={mins}
        secs={secs}
        progressPercentComplete={
          runRecipeInProgress ? JSON.parse(runRecipeInProgress).progress : 0
        }
        handleLeftAction={getLeftActionBtnHandler()}
        handleRightAction={getRightActionBtnHandler()}
        leftActionBtn={
          showProcess
            ? recipeActionReducer.leftActionBtn
            : cleanUpReducer.leftActionBtn
        }
        rightActionBtn={
          showProcess
            ? recipeActionReducer.rightActionBtn
            : cleanUpReducer.rightActionBtn
        }
        rightActionBtnDisabled={
          cleanUpReducer.leftActionBtn === DECKCARD_BTN.text.done ||
          recipeActionReducer.leftActionBtn === DECKCARD_BTN.text.done
        }
      />
    </RecipeListing>
  );
};

RecipeListingComponent.propTypes = {};

export default RecipeListingComponent;
