import React from "react";

import { Card, CardBody, Row, Col } from "core-components";
import SearchBox from "shared-components/SearchBox";
import PaginationBox from "shared-components/PaginationBox";
import RecipeCard from "components/RecipeListing/RecipeCard";

const RecipeListingCards = (props) => {
    const {
        isAdmin,
        searchRecipeText,
        onSearchRecipeTextChanged,
        fileteredRecipeData,
        handleCarousalModal,
        returnRecipeDetails,
        toggleRunRecipesModal,
        handlePublishModalClick,
        handleEditRecipe,
    } = props;

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
                                    handlePublishModalClick={(recipeId, isPublished) =>
                                        handlePublishModalClick(recipeId, isPublished)
                                    }
                                    handleEditRecipe={() =>
                                        handleEditRecipe(recipe)
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
