import React from "react";

import { Fade } from "reactstrap";
import { Text, Icon, ButtonIcon } from "shared-components";
import { RecipeCardStyle } from "./RecipeCardStyle";

const RecipeCard = (props) => {
  const {
    recipeId,
    recipeName,
    processCount,
    isAdmin,
    isPublished,
    handleCarousalModal,
    selectedRecipeData,
    returnRecipeDetails,
    toggleRunRecipesModal,
    handlePublishModalClick,
    handleEditRecipeNameModalClick,
    handleEditRecipe,
    handleDeleteRecipe,
    toggle,
  } = props;

  const handleClickOnCard = () => {
    if (isAdmin) {
      const param = toggle ? {} : { recipeId };
      returnRecipeDetails(param);
    } else {
      handleCarousalModal();
      returnRecipeDetails({ recipeId, recipeName, processCount });
    }
  };

  const handleRunRecipeByAdmin = (e) => {
    e.stopPropagation();
    toggleRunRecipesModal();
    returnRecipeDetails({ recipeId, recipeName, processCount, isAdmin });
  };

  return (
    <div onClick={handleClickOnCard}>
      <RecipeCardStyle
        className={toggle ? `${isAdmin ? "admin-" : ""}selected` : ""}
      >
        <div className="recipe-heading d-flex justify-content-between align-items-center">
          <div className="font-weight-bold">{recipeName}</div>
          {isAdmin && isPublished ? (
            <Text Tag="span">
              <Icon name="published" size={15} className="text-primary mr-3" />
            </Text>
          ) : null}
        </div>
        {isAdmin && toggle ? (
          <Fade in tag="h5" className="m-0">
            <div className="recipe-action d-flex justify-content-between align-items-center pt-2">
              <div className="d-flex justify-content-between align-items-center">
                <ButtonIcon
                  size={25}
                  name="play"
                  className="border-gray text-primary mr-2"
                  onClick={handleRunRecipeByAdmin}
                />
                <ButtonIcon
                  size={25}
                  name="edit-pencil"
                  className="border-gray text-primary mr-2"
                  onClick={handleEditRecipe}
                />
                <ButtonIcon
                  size={25}
                  name={isPublished ? "published" : `publish`}
                  className={`border-gray mr-2 ${
                    isPublished
                      ? "published-icon text-white bg-primary"
                      : "text-primary"
                  }`}
                  onClick={() => handlePublishModalClick(recipeId, isPublished)}
                />
                <ButtonIcon
                  size={25}
                  name="pencil"
                  className="border-gray text-primary mr-2"
                  onClick={() => handleEditRecipeNameModalClick(recipeId)}
                />
              </div>
              <ButtonIcon
                size={20}
                name="minus-1"
                className="border-gray text-primary"
                onClick={handleDeleteRecipe}
              />
            </div>
          </Fade>
        ) : (
          <>
            <Text Tag="span" className="recipe-name">
              Total Processes -
            </Text>
            <Text
              Tag="span"
              className="text-primary font-weight-bold recipe-value ml-2"
            >
              {processCount}
            </Text>
          </>
        )}
      </RecipeCardStyle>
    </div>
  );
};

export default React.memo(RecipeCard);
