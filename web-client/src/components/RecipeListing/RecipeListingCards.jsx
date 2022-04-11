import React, { useState, useEffect, useCallback } from "react";

import { Card, CardBody, Row, Col } from "core-components";
import SearchBox from "shared-components/SearchBox";
import PaginationBox from "shared-components/PaginationBox";
import RecipeCard from "components/RecipeListing/RecipeCard";
import MlModal from "shared-components/MlModal";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import {
  paginator,
  initialPaginationStateRecipeList,
} from "utils/paginationHelper";
import { CardOverlay } from "./RecipeCardStyle";
import { useSelector, useDispatch } from "react-redux";
import {
  nextRecipePageBtn,
  prevRecipePageBtn,
} from "action-creators/PageActionCreators";

const RecipeListingCards = (props) => {
  const {
    isAdmin,
    deckName,
    searchRecipeText,
    onSearchRecipeTextChanged,
    fileteredRecipeData,
    handleCarousalModal,
    selectedRecipeData,
    returnRecipeDetails,
    toggleRunRecipesModal,
    handlePublishModalClick,
    handleEditRecipe,
    handleDeleteRecipe,
    handleEditRecipeNameModalClick,
  } = props;

  const dispatch = useDispatch();
  const [deleteRecipeId, setDeleteRecipeId] = useState(null);
  const [deleteModal, setDeleteModal] = useState(false);
  const [paginatedData, setPaginatedData] = useState(
    initialPaginationStateRecipeList
  );
  const pageReducer = useSelector((state) => state.pageReducer);
  const page =
    deckName == "Deck A"
      ? pageReducer.recipePageDeckA
      : pageReducer.recipePageDeckB;
  /**reset pagination when recipeList changed */
  useEffect(() => {
    findAndSetPagination();
  }, [fileteredRecipeData]);

  /**reset pagination when page changed */
  useEffect(() => {
    findAndSetPagination();
  }, [page]);

  useEffect(() => {
    //if we dont have data on this page but having on previous page then go to previous page
    if (paginatedData?.list?.length === 0 && paginatedData?.total !== 0) {
      handlePrev();
    }
  }, [paginatedData]);

  const findAndSetPagination = () => {
    const data = paginator(
      fileteredRecipeData,
      page,
      paginatedData.perPageItems
    );
    const newData = {
      ...paginatedData,
      page: data.page,
      prevPage: data.prePage,
      nextPage: data.nextPage,
      total: data.total,
      list: data.list,
      from: data.from,
      to: data.to,
    };
    setPaginatedData(newData);
  };

  const handleNext = useCallback(() => {
    if (paginatedData.nextPage) {
      dispatch(nextRecipePageBtn(deckName));
    }
  });

  const handlePrev = useCallback(() => {
    if (paginatedData.prevPage) {
      dispatch(prevRecipePageBtn(deckName));
    }
  });

  const handleDeleteRecipeClick = (id) => {
    setDeleteRecipeId(id);
    toggleDeleteModal();
  };

  const toggleDeleteModal = () => {
    setDeleteModal(!deleteModal);
  };

  const onDeleteRecipeConfirmed = () => {
    toggleDeleteModal();
    handleDeleteRecipe(deleteRecipeId);
  };

  return (
    <Card className="recipe-listing-cards">
      {isAdmin && selectedRecipeData?.data?.recipeId && <CardOverlay />}
      <CardBody className="p-5">
        <div className="d-flex justify-content-between align-items-center">
          {isAdmin ? (
            <SearchBox
              value={searchRecipeText}
              onChange={onSearchRecipeTextChanged}
            />
          ) : null}
          <div className="d-flex justify-content-end ml-auto">
            <PaginationBox
              firstIndexOnPage={paginatedData.from}
              lastIndexOnPage={paginatedData.to}
              totalPages={paginatedData?.total || 0}
              handlePrev={handlePrev}
              handleNext={handleNext}
            />
          </div>
        </div>

        {deleteModal && (
          <MlModal
            isOpen={deleteModal}
            textHead={deckName}
            textBody={MODAL_MESSAGE.deleteRecipeConfirmation}
            handleSuccessBtn={onDeleteRecipeConfirmed}
            handleCrossBtn={toggleDeleteModal}
            successBtn={MODAL_BTN.yes}
            failureBtn={MODAL_BTN.no}
          />
        )}

        <Row>
          {paginatedData?.list?.length > 0 ? (
            paginatedData?.list?.map((recipe, index) => (
              <Col md={6} key={index}>
                <RecipeCard
                  isAdmin={isAdmin}
                  recipeId={recipe.id}
                  recipeName={recipe.name}
                  processCount={recipe.process_count}
                  isPublished={recipe.is_published}
                  handleCarousalModal={handleCarousalModal}
                  selectedRecipeData={selectedRecipeData}
                  toggle={
                    selectedRecipeData?.data?.recipeId &&
                    selectedRecipeData.data.recipeId === recipe.id
                  }
                  returnRecipeDetails={returnRecipeDetails}
                  toggleRunRecipesModal={toggleRunRecipesModal}
                  handlePublishModalClick={(recipeId, isPublished) =>
                    handlePublishModalClick(recipeId, isPublished)
                  }
                  handleEditRecipeNameModalClick={(recipeId) =>
                    handleEditRecipeNameModalClick(recipeId)
                  }
                  handleEditRecipe={() => handleEditRecipe(recipe)}
                  handleDeleteRecipeClick={(recipeId) => {
                    handleDeleteRecipeClick(recipeId);
                  }}
                />
              </Col>
            ))
          ) : (
            <Col md={12}>
              <div className="text-center">
                <h4>No recipes found</h4>
              </div>
            </Col>
          )}
        </Row>
      </CardBody>
    </Card>
  );
};

export default React.memo(RecipeListingCards);
