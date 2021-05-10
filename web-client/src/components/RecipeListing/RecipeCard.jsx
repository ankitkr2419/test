import React, { useState } from "react";

import { Fade } from "reactstrap";
import { Text, Icon, ButtonIcon } from "shared-components";
import styled from "styled-components";

const RecipeCardStyle = styled.div`
  padding: 0.8rem 0.5rem;
  border: 1px solid #e3e3e3;
  border-radius: 0.5rem;
  margin-bottom: 0.688rem;
  box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.04);
  // width:27.5rem;
  // height: 5.563rem;
  .recipe-heading {
    padding-bottom: 0.5rem;
  }
  .recipe-card-body {
    padding-top: 0.25rem;
    border-top: 1px solid #d9d9d9;

    .recipe-name {
      font-size: 0.875rem;
      line-height: 1rem;
    }
    .recipe-value {
      font-size: 1.125rem;
      line-height: 1.313rem;
    }
    .recipe-action {
      button {
        width: 33px !important;
        height: 33px !important;
        border: 1px solid #696969 !important;
        &:not(:first-child) {
          margin-left: 12px;
        }
      }
    }
  }
  &:focus,
  &:hover {
    background-color: rgba(243, 130, 32, 0.3);
  }
`;

const RecipeCard = (props) => {
  const {
    recipeId,
    recipeName,
    processCount,
    isAdmin,
    isPublished,
    handleCarousalModal,
    returnRecipeDetails,
    toggleRunRecipesModal,
    handlePublishModalClick
  } = props;

  const [toggle, setToggle] = useState(true);

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
  }

  return (
    <div onClick={handleClickOnCard}>
      <RecipeCardStyle>
        <>
          <div className="recipe-heading d-flex justify-content-between align-items-center">
            <div className="font-weight-bold">{recipeName}</div>{" "}
            {isAdmin && isPublished ? (
              <Text Tag="span">
                <Icon
                  name="published"
                  size={15}
                  className="text-primary mr-3"
                />{" "}
              </Text>
            ) : null}
          </div>
        </>
        <>
          {toggle ? (
            <>
              <Text Tag="span" className="recipe-name">
                Total Processes -
              </Text>
              <Text
                Tag="span"
                className="text-primary font-weight-bold recipe-value ml-2"
              >
                {processCount}{" "}
              </Text>
            </>
          ) : (
            <Fade in={true} tag="h5" className="m-0">
              <div className="recipe-action d-flex justify-content-between align-items-center">
                <div className="d-flex justify-content-between align-items-center">
                  <ButtonIcon
                    size={14}
                    name="play"
                    className="border-gray text-primary"
                    onClick={handleRunRecipeByAdmin}
                  />
                  <ButtonIcon
                    size={14}
                    name="edit-pencil"
                    className="border-gray text-primary"
                    // onClick={toggleExportDataModal}
                  />
                  <ButtonIcon
                    size={14}
                    name="publish"
                    className="border-gray text-primary"
                    onClick={() => handlePublishModalClick(recipeId)}
                  />
                </div>
                <ButtonIcon
                  size={20}
                  name="minus-1"
                  className="border-gray text-primary"
                  // onClick={toggleExportDataModal}
                />
              </div>
            </Fade>
          )}
        </>
      </RecipeCardStyle>
    </div>
  );
};

export default RecipeCard;
