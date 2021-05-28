import React, { useState } from "react";

import { Card, CardBody, Row, Col } from "core-components";
import SearchBox from "shared-components/SearchBox";
import PaginationBox from "shared-components/PaginationBox";
import RecipeCard from "components/RecipeListing/RecipeCard";
import MlModal from "shared-components/MlModal";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";

const RecipeListingCards = (props) => {
    const {
        isAdmin,
        deckName,
        searchRecipeText,
        onSearchRecipeTextChanged,
        fileteredRecipeData,
        handleCarousalModal,
        returnRecipeDetails,
        toggleRunRecipesModal,
        handlePublishModalClick,
        handleEditRecipe,
        handleDeleteRecipe,
    } = props;

    const [deleteRecipeId, setDeleteRecipeId] = useState(null);
    const [deleteModal, setDeleteModal] = useState(false);

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
            <CardBody className="p-5">
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
                    {fileteredRecipeData?.length ? (
                        fileteredRecipeData.map((recipe, index) => (
                            <Col md={6} key={index}>
                                <RecipeCard
                                    isAdmin={isAdmin}
                                    recipeId={recipe.id}
                                    recipeName={recipe.name}
                                    processCount={recipe.process_count}
                                    isPublished={recipe.is_published}
                                    handleCarousalModal={handleCarousalModal}
                                    returnRecipeDetails={returnRecipeDetails}
                                    toggleRunRecipesModal={
                                        toggleRunRecipesModal
                                    }
                                    handlePublishModalClick={(
                                        recipeId,
                                        isPublished
                                    ) =>
                                        handlePublishModalClick(
                                            recipeId,
                                            isPublished
                                        )
                                    }
                                    handleEditRecipe={() =>
                                        handleEditRecipe(recipe)
                                    }
                                    handleDeleteRecipe={() =>
                                        handleDeleteRecipeClick(recipe.id)
                                    }
                                />
                            </Col>
                        ))
                    ) : (
                        <h4>No recipes to show!</h4>
                    )}
                </Row>
            </CardBody>
        </Card>
    );
};

export default React.memo(RecipeListingCards);
