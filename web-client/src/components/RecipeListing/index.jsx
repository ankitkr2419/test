import React, { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";

import { Card, CardBody, Button, Row, Col } from "core-components";
import { Icon, MlModal, VideoCard } from "shared-components";

import styled from "styled-components";
import AppFooter from "components/AppFooter";
import RecipeFlowModal from "components/modals/RecipeFlowModal";
// import ConfirmationModal from "components/modals/ConfirmationModal";
import TrayDiscardModal from "components/modals/TrayDiscardModal";
import RecipeCard from "components/RecipeListing/RecipeCard";
import {
  abortRecipeInitiated,
  pauseRecipeInitiated,
  resumeRecipeInitiated,
  runRecipeInitiated,
} from "action-creators/recipeActionCreators";
import { DECKCARD_BTN, MODAL_BTN, MODAL_MESSAGE } from "appConstants";

const TopContent = styled.div`
  margin-bottom: 2.25rem;
`;

const HeadingTitle = styled.label`
  font-size: 1.25rem;
  line-height: 1.438rem;
`;

const RecipeListingComponent = (props) => {
  const { allRecipeData } = props;

  const dispatch = useDispatch();

  const operatorLoginModalReducer = useSelector(
    (state) => state.operatorLoginModalReducer
  );
  const { deckName } = operatorLoginModalReducer.toJS();

  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  const {
    // runRecipeError,
    abortRecipeError,
    // pauseRecipeError,
    // resumeRecipeError,
    // recipeListingError,
    leftActionBtn,
    rightActionBtn,
    // isLoading,
  } = recipeActionReducer;

  const [confirmationModal, setConfirmationModal] = useState(false);
  const [recipeData, setRecipeData] = useState({});
  const [progressPercentComplete, setProgressPercentComplete] = useState(0);

  const [isOpen, setIsOpen] = useState(false);
  const toggle = (recipeId, recipeName, processCount) => {
    // const tempRecipeId = "bb7fcfa2-8337-4d79-829a-e9bd486add14";
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

  useEffect(() => {
    if ((abortRecipeError === false) && (confirmationModal === true)) {
      setConfirmationModal(false);
      setShowProcess(!showProcess);
    }
  }, [abortRecipeError, confirmationModal, showProcess]);

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

  const handleCancelAction = () => setShowProcess(!showProcess);
  const handleAbortAction = () => setConfirmationModal(true);

  const toggleConfirmModal = () => {
    const name = deckName === "Deck A" ? "A" : "B";
    dispatch(abortRecipeInitiated({ deckName: name }));
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

  console.log("DECKNAME: ", deckName);

  return (
    <div className="ml-content">
      <div className="landing-content px-2">
        <RecipeFlowModal
          isOpen={isOpen}
          toggle={toggle}
          toggleShowProcess={toggleShowProcess}
          recipeData={recipeData}
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
            <Button color="secondary" className="ml-auto">
              {" "}
              Clean Up
            </Button>
            <TrayDiscardModal />
          </div>
        </TopContent>

        {showProcess ? (
          <VideoCard />
        ) : (
          <Card>
            <CardBody className="p-5">
              <Row>
                {allRecipeData.length > 0 ? (
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
    </div>
  );
};

RecipeListingComponent.propTypes = {};

export default RecipeListingComponent;
