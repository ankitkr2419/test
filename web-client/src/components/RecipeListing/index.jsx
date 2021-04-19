import React from "react";
// import { useSelector, useDispatch } from "react-redux";

import styled from "styled-components";
import { Card, CardBody, Button, Row, Col } from "core-components";
import { Icon, VideoCard, ButtonIcon } from "shared-components";

import SearchBox from "shared-components/SearchBox";
import PaginationBox from "shared-components/PaginationBox";

import RecipeCard from "components/RecipeListing/RecipeCard";
import OperatorRunRecipeCarousalModal from "components/modals/OperatorRunRecipeCarousalModal";

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
    recipeData,
    isOperatorRunRecipeCarousalModalVisible,
    handleCarousalModal,
    returnRecipeDetails,
  } = props;

  return (
    <>
      <div className="landing-content px-2">
        {/* The following modal is displayed when an operator begins to run a recipe */}
        {isOperatorRunRecipeCarousalModalVisible && (
          <OperatorRunRecipeCarousalModal
            isOpen={isOperatorRunRecipeCarousalModalVisible}
            handleCarousalModal={handleCarousalModal}
          />
        )}

        <TopContent className="d-flex justify-content-between align-items-center mx-5">
          {isProcessInProgress ? null : (
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

          {isProcessInProgress ? null : (
            <div className="d-flex align-items-center ml-auto">
              {isAdmin ? (
                <Button
                  color="secondary"
                  className="ml-2 border-primary btn-discard-tray bg-white"
                  //   onClick={handleTrayDiscardModal}
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
                    // onClick={handleTimeModal}
                  >
                    {" "}
                    Clean Up
                  </Button>
                  <Button
                    color="secondary"
                    className="ml-2 border-primary btn-discard-tray bg-white"
                    // onClick={handleTrayDiscardModal}
                  >
                    Discard Tray
                  </Button>
                </>
              )}
              <ButtonIcon
                name="logout"
                size={28}
                className="ml-2 bg-white border-primary"
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
                  {isAdmin ? <SearchBox /> : null}
                  <div className="d-flex justify-content-end">
                    <PaginationBox />
                  </div>
                </div>

                <Row>
                  {recipeData?.length ? (
                    recipeData.map((recipe, index) => (
                      <Col md={6} key={index}>
                        <RecipeCard
                          isAdmin={isAdmin}
                          recipeId={recipe.id}
                          recipeName={recipe.name}
                          processCount={recipe.process_count}
                          isPublished={recipe.isPublished}
                          handleCarousalModal={handleCarousalModal}
                          returnRecipeDetails={returnRecipeDetails}
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
    </>
  );
};
export default React.memo(RecipeListingComponent);
