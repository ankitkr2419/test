import React, { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { VideoCard } from "shared-components";
import MlModal from "shared-components/MlModal";
import TimeModal from "components/modals/TimeModal";
import OperatorRunRecipeCarousalModal from "components/modals/OperatorRunRecipeCarousalModal";
import AppFooter from "components/AppFooter";
import { useHistory } from "react-router-dom";
import { DECKNAME, MODAL_BTN, ROUTES, MODAL_MESSAGE } from "appConstants";
import { loginReset } from "action-creators/loginActionCreators";
import {
  setCleanUpHours,
  setCleanUpMins,
  setCleanUpSecs,
  setShowCleanUp,
} from "action-creators/cleanUpActionCreators";
import TrayDiscardModal from "components/modals/TrayDiscardModal";
import { discardDeckInitiated } from "action-creators/discardDeckActionCreators";
import { restoreDeckInitiated } from "action-creators/restoreDeckActionCreators";
import AddNewRecipesModal from "components/modals/AddNewRecipesModal";
import RunRecipesModal from "components/modals/RunRecipesModal";
import { publishRecipeInitiated } from "action-creators/recipeActionCreators";
import TopContentComponent from "./TopContentComponent";
import RecipeListingCards from "./RecipeListingCards";

const RecipeListingComponent = (props) => {
  const {
    isProcessInProgress,
    isAdmin,
    deckName,
    recipeData,
    isOperatorRunRecipeCarousalModalVisible,
    handleCarousalModal,
    returnRecipeDetails,
    onConfirmedRecipeSelection,
    onConfirmedRunRecipeByAdmin,
    runRecipesmodal,
    toggleRunRecipesModal,
    onChangeRunRecipeType,
    runRecipeType,
  } = props;

  const [isLogoutModalVisible, setLogoutModalVisibility] = useState(false);
  const dispatch = useDispatch();
  const history = useHistory();

  const [timeModal, setTimeModal] = useState(false);
  const [trayDiscardModal, setTrayDiscardModal] = useState(false);
  const [nextModal, setNextModal] = useState(true);
  const [addNewRecipesModal, setAddNewRecipesModal] = useState(false);
  const [searchRecipeText, setSearchRecipeText] = useState("");
  const [recipeIdToPublish, setRecipeIdToPublish] = useState("");
  const [publishModal, setPublishModal] = useState(false);

  useEffect(() => {
    setSearchRecipeText("");
  }, [deckName]);

  const onSearchRecipeTextChanged = (e) => {
    const value = e.target.value;
    setSearchRecipeText(value);
  };

  const toggleAddNewRecipesModal = () => {
    setAddNewRecipesModal(!addNewRecipesModal);
  };

  const onLogoutClicked = () => {
    toggleLogoutModalVisibility();
    dispatch(loginReset(deckName));
    history.push(ROUTES.landing);
  };

  const toggleLogoutModalVisibility = () => {
    setLogoutModalVisibility(!isLogoutModalVisible);
  };

  const toggleTimeModal = () => {
    setTimeModal(!timeModal);
  };

  const toggleTrayDiscardModal = () => {
    setTrayDiscardModal(!trayDiscardModal);
    setNextModal(true);
  };

  const togglePublishModal = () => {
    setPublishModal(!publishModal);
  };

  const handlePublishModalClick = (recipeId) => {
    setRecipeIdToPublish(recipeId);
    if (recipeId) togglePublishModal();
  };

  const handlePublishConfirmation = () => {
    togglePublishModal();
    if (recipeIdToPublish)
      dispatch(
        publishRecipeInitiated({ recipeId: recipeIdToPublish, deckName })
      );
    else console.error("recipeId not found!");
  };

  const handleSuccessBtn = () => {
    if (nextModal) {
      dispatch(
        discardDeckInitiated({
          deckName:
            deckName === DECKNAME.DeckA
              ? DECKNAME.DeckAShort
              : DECKNAME.DeckBShort,
        })
      );
      setNextModal(!nextModal);
    } else {
      dispatch(
        restoreDeckInitiated({
          deckName:
            deckName === DECKNAME.DeckA
              ? DECKNAME.DeckAShort
              : DECKNAME.DeckBShort,
        })
      );
      setTrayDiscardModal(!trayDiscardModal);
      setNextModal(true);
    }
  };

  const submitTime = () => {
    setTimeModal(!timeModal);
    dispatch(setShowCleanUp({ deckName: deckName }));
  };

  const handleChangeTime = (event) => {
    let name = event.target.name;

    if (name === "hours") {
      dispatch(
        setCleanUpHours({
          deckName: deckName,
          hours: event.target.value,
        })
      );
    } else if (name === "minutes") {
      dispatch(
        setCleanUpMins({ deckName: deckName, mins: event.target.value })
      );
    } else if (name === "seconds") {
      dispatch(
        setCleanUpSecs({ deckName: deckName, secs: event.target.value })
      );
    }
  };

  const fileteredRecipeData = recipeData.filter((recipeObj) =>
    recipeObj.name.toLowerCase().includes(searchRecipeText.toLowerCase())
  );

  const handleEditRecipe = (recipe) => {
    let recipeId = recipe?.id;
    if (!recipeId) {
      console.error("recipeId not found");
      return;
    }

    //TODO: save recipe in reducer to edit

    //go to processList page of recipe
    history.push(ROUTES.processListing);
  };

  const getLogoutTextBody = () => {
    return `Are you sure you want to sign out of ${
      isAdmin ? "Admin" : "Operator"
    } role?`;
  };

  return (
    <>
      <div className="landing-content px-2">
        {/* The following modal is displayed when an operator begins to run a recipe */}
        {isOperatorRunRecipeCarousalModalVisible && (
          <OperatorRunRecipeCarousalModal
            isOpen={isOperatorRunRecipeCarousalModalVisible}
            handleCarousalModal={handleCarousalModal}
            onConfirmedRecipeSelection={onConfirmedRecipeSelection}
          />
        )}

        {runRecipesmodal && isAdmin && (
          <RunRecipesModal
            isOpen={runRecipesmodal}
            deckName={deckName}
            toggleRunRecipesModal={toggleRunRecipesModal}
            runRecipeType={runRecipeType}
            onChange={(type) => onChangeRunRecipeType(type)}
            onConfirmed={onConfirmedRunRecipeByAdmin}
          />
        )}

        {timeModal && (
          <TimeModal
            timeModal={timeModal}
            toggleTimeModal={toggleTimeModal}
            handleChangeTime={handleChangeTime}
            submitTime={submitTime}
            deckName={deckName}
          />
        )}

        {trayDiscardModal && (
          <TrayDiscardModal
            trayDiscardModal={trayDiscardModal}
            toggleTrayDiscardModal={toggleTrayDiscardModal}
            handleSuccessBtn={handleSuccessBtn}
            nextModal={nextModal}
            deckName={deckName}
          />
        )}

        <MlModal
          isOpen={isLogoutModalVisible}
          textHead={deckName}
          textBody={getLogoutTextBody()}
          handleSuccessBtn={onLogoutClicked}
          handleCrossBtn={toggleLogoutModalVisibility}
          successBtn={MODAL_BTN.yes}
          failureBtn={MODAL_BTN.no}
        />

        {addNewRecipesModal && (
          <AddNewRecipesModal
            isOpen={addNewRecipesModal}
            toggleAddNewRecipesModal={toggleAddNewRecipesModal}
            deckName={deckName}
            confirmationText={"Name the Recipe"}
          />
        )}

        {/** publish confirmation modal */}
        {publishModal && (
          <MlModal
            isOpen={publishModal}
            textHead={deckName}
            textBody={MODAL_MESSAGE.publishConfirmation}
            handleSuccessBtn={handlePublishConfirmation}
            handleCrossBtn={togglePublishModal}
            successBtn={MODAL_BTN.yes}
            failureBtn={MODAL_BTN.no}
          />
        )}

        {/** Sub - Menu above recipe listings (like addNewRecipe/ cleanUp/ etc) */}
        <TopContentComponent
          isProcessInProgress={isProcessInProgress}
          onLogoutClicked={onLogoutClicked}
          deckName={deckName}
          isAdmin={isAdmin}
          toggleAddNewRecipesModal={toggleAddNewRecipesModal}
          toggleTimeModal={toggleTimeModal}
          toggleTrayDiscardModal={toggleTrayDiscardModal}
          toggleLogoutModalVisibility={toggleLogoutModalVisibility}
        />

        {/**
         * Show Video if some process is going on, like runRecipe
         * else show Recipe list
         *
         * RecipeListingCards: pagination, searchRecipe, recipeList, etc
         */}
        <>
          {isProcessInProgress ? (
            <VideoCard />
          ) : (
            <RecipeListingCards
              isAdmin={isAdmin}
              searchRecipeText={searchRecipeText}
              onSearchRecipeTextChanged={onSearchRecipeTextChanged}
              fileteredRecipeData={fileteredRecipeData}
              handleCarousalModal={handleCarousalModal}
              returnRecipeDetails={returnRecipeDetails}
              toggleRunRecipesModal={toggleRunRecipesModal}
              handlePublishModalClick={(recipeId) =>
                handlePublishModalClick(recipeId)
              }
              handleEditRecipe={(recipe) => handleEditRecipe(recipe)}
            />
          )}
        </>
      </div>
      <AppFooter />
    </>
  );
};
export default React.memo(RecipeListingComponent);
