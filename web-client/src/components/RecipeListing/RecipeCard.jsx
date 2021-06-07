import React, { useEffect, useState } from "react";

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
    handleEditRecipe,
    handleDeleteRecipe,
  } = props;

  const [toggle, setToggle] = useState(true);

  //maintain a var which tells when to show recipe as selected (orange background)
  const isSelected = (!isAdmin && selectedRecipeData?.data?.recipeId === recipeId) ? true : false;

  //hide menu if menu open and tab changed
  useEffect(() => {
    if(!toggle) {
      setToggle(true)
    }
  }, [isAdmin])

  const handleClickOnCard = () => {
    if (isAdmin) {
      setToggle(!toggle);
    } else {
      handleCarousalModal();
      returnRecipeDetails({ recipeId, recipeName, processCount });
    }
  };

  const handleRunRecipeByAdmin = () => {
    toggleRunRecipesModal();
    returnRecipeDetails({ recipeId, recipeName, processCount, isAdmin });
  };

  return (
    <div onClick={handleClickOnCard}>
      <RecipeCardStyle className={isSelected ? "selected": ""}>
        <div className="recipe-heading d-flex justify-content-between align-items-center">
          <div className="font-weight-bold">{recipeName}</div>
          {isAdmin && isPublished ? (
            <Text Tag="span">
              <Icon name="published" size={15} className="text-primary mr-3" />
            </Text>
          ) : null}
        </div>
        {toggle ? (
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
        ) : (
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
                  name={isPublished ? 'published' : `publish`}
                  className={`border-gray ${isPublished ? "published-icon text-white bg-primary": "text-primary"}`}
                  onClick={() => handlePublishModalClick(recipeId, isPublished)}
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
        )}
      </RecipeCardStyle>
    </div>
  );
};

export default React.memo(RecipeCard);
