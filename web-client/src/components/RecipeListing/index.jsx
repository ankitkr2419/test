import React, { useEffect, useState } from "react";
import { useDispatch } from "react-redux";

import styled from "styled-components";
import { Card, CardBody, Button, Row, Col } from "core-components";
import { Icon, VideoCard, ButtonIcon } from "shared-components";

import SearchBox from "shared-components/SearchBox";
import PaginationBox from "shared-components/PaginationBox";
import MlModal from "shared-components/MlModal";
import TimeModal from "components/modals/TimeModal";
import RecipeCard from "components/RecipeListing/RecipeCard";
import OperatorRunRecipeCarousalModal from "components/modals/OperatorRunRecipeCarousalModal";
import AppFooter from "components/AppFooter";
import { useHistory } from "react-router-dom";
import { DECKNAME, MODAL_BTN, ROUTES } from "appConstants";
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
      dispatch(loginReset(deckName));
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
          textBody={`Are you sure you want to sign out of ${
            isAdmin ? "Admin" : "Operator"
          } role?`}
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

        <TopContent className="d-flex justify-content-between align-items-center mx-5">
          {isProcessInProgress ? null : (
            <div className="d-flex align-items-center">
              <div style={{ cursor: "pointer" }} onClick={onLogoutClicked}>
                <Icon name="angle-left" size={32} className="text-white" />
              </div>
              <HeadingTitle
                Tag="h5"
                className="text-white font-weight-bold ml-3 mb-0"
              >
                {`Select a Recipe for ${deckName}`}
              </HeadingTitle>
            </div>
          )}

          {isProcessInProgress ? null : (
            <div className="d-flex align-items-center ml-auto">
              {isAdmin ? (
                <Button
                  color="secondary"
                  className="ml-2 border-primary btn-discard-tray bg-white"
                  onClick={toggleAddNewRecipesModal}
                >
                  Add Recipe
                </Button>
              ) : (
                <>
                  <ButtonIcon
                    name="download-1"
                    size={28}
                    className="bg-white border-primary"
                  />
                  <Button
                    color="secondary"
                    className="ml-2 border-primary btn-clean-up bg-white"
                    onClick={toggleTimeModal}
                  >
                    {" "}
                    Clean Up
                  </Button>
                  <Button
                    color="secondary"
                    className="ml-2 border-primary btn-discard-tray bg-white"
                    onClick={toggleTrayDiscardModal}
                  >
                    Discard Tray
                  </Button>
                </>
              )}
              <ButtonIcon
                name="logout"
                size={28}
                className="ml-2 bg-white border-primary"
                onClick={toggleLogoutModalVisibility}
              />
            </div>
          )}
        </TopContent>
        <>
          {isProcessInProgress ? (
            <VideoCard />
          ) : (
            <Card className="recipe-listing-cards">
              <CardBody className="p-5">
                {/* Search Functionality Input not working */}
                <div className="d-flex justify-content-between align-items-center">
                  {isAdmin ? (
                    <SearchBox
                      value={searchRecipeText}
                      onChange={onSearchRecipeTextChanged}
                    />
                  ) : null}
                  <div className="d-flex justify-content-end">
                    <PaginationBox />
                  </div>
                </div>

                <Row>
                  {fileteredRecipeData?.length ? (
                    fileteredRecipeData.map((recipe, index) => (
                      <Col md={6} key={index}>
                        <RecipeCard
                          isAdmin={isAdmin}
                          recipeId={recipe.id}
                          recipeName={recipe.name}
                          processCount={recipe.process_count}
                          isPublished={recipe.isPublished}
                          handleCarousalModal={handleCarousalModal}
                          returnRecipeDetails={returnRecipeDetails}
                          toggleRunRecipesModal={toggleRunRecipesModal}
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
        </>
      </div>
      <AppFooter />
    </>
  );
};
export default React.memo(RecipeListingComponent);
