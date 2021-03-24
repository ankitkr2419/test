import React, { useState } from "react";
import { useSelector, useDispatch } from "react-redux";

import { Card, CardBody, Button, Row, Col } from "core-components";
import { Icon, VideoCard } from "shared-components";

import styled from "styled-components";
import AppFooter from "components/AppFooter";
import RecipeFlowModal from "components/modals/RecipeFlowModal";
// import ConfirmationModal from "components/modals/ConfirmationModal";
import TrayDiscardModal from "components/modals/TrayDiscardModal";
import RecipeCard from "components/RecipeListing/RecipeCard";
import { runRecipeInitiated } from "action-creators/recipeActionCreators";

const TopContent = styled.div`
  margin-bottom: 2.25rem;
`;

const HeadingTitle = styled.label`
  font-size: 1.25rem;
  line-height: 1.438rem;
`;

const RecipeListingComponent = (props) => {
  const { allRecipeData } = props;

  const dispatch = useDispatch()

  const operatorLoginModalReducer = useSelector(
    (state) => state.operatorLoginModalReducer
  );
  const { deckName } = operatorLoginModalReducer.toJS();

  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  console.log(recipeActionReducer);

  const [recipeData, setRecipeData] = useState({});

  const [isOpen, setIsOpen] = useState(false);
  const toggle = (recipeId, recipeName, processCount) => {
    //will be deleted
    const tempRecipeId = "bb7fcfa2-8337-4d79-829a-e9bd486add14";
    const data = {
      recipeId: tempRecipeId,
      recipeName: recipeName,
      processCount: processCount
    }
    setRecipeData(data);
    setIsOpen(!isOpen);
  };

  const [showProcess, setShowProcess] = useState(false);
  const toggleShowProcess = () => {
    setShowProcess(!showProcess);
    setIsOpen(!isOpen);
  };

  const handleCancelAction = () => setShowProcess(!showProcess);

  const handleRunAction = () => {
    const name = (deckName === "Deck A") ? "A" : "B";
    const { recipeId } = recipeData;
    dispatch(runRecipeInitiated({recipeId:recipeId, deckName:name}));
  };

  return (
    <div className="ml-content">
      <div className="landing-content px-2">
        {/* <ConfirmationModal isOpen={true} /> */}

        <RecipeFlowModal
          isOpen={isOpen}
          toggle={toggle}
          toggleShowProcess={toggleShowProcess}
          recipeData={recipeData}
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
              {allRecipeData.map((value, index) => (
                <Col md={6} key={index}>
                  <RecipeCard
                    recipeId={value.id}
                    recipeName={value.name}
                    processCount={value.process_count}
                    toggle={toggle}
                  />
                </Col>
              ))}
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
        handleCancelAction={handleCancelAction}
        handleRunAction={handleRunAction}
      />
    </div>
  );
};

RecipeListingComponent.propTypes = {};

export default RecipeListingComponent;
