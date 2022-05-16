import React, { useState } from "react";

import PropTypes from "prop-types";
import {
  Modal,
  ModalBody,
  Button,
  Form,
  FormGroup,
  FormError,
  Input,
} from "core-components";
import { Center, Text, ButtonIcon } from "shared-components";
import { toast } from "react-toastify";
import { useDispatch } from "react-redux";
import { saveNewRecipe } from "action-creators/saveNewRecipeActionCreators";
import { ROUTES } from "appConstants";
import { useHistory } from "react-router";
import { AddNewRecipesForm } from "./DuplicateRecipesForm";

const DuplicateRecipeModal = (props) => {
  const {
    createDuplicateRecipe,
    confirmationCopyText,
    confirmationText,
    isOpen,
    toggleCopyRecipeModel,
    deckName,
    recipeId,
    fileteredRecipeData,
  } = props;
  const [recipeName, setRecipeName] = useState("");
  const [submitted, setSubmitted] = useState(false);

  const dispatch = useDispatch();
  const history = useHistory();

  const onChangeRecipeName = (e) => {
    setSubmitted(false);
    let val = e.target.value;
    setRecipeName(val);
  };

  const onCreateRecipeClicked = () => {
    if (fileteredRecipeData.some((recipe) => recipe.name === recipeName)) {
      toast.error("Recipe with same name already exists");
      return;
    }
    if (recipeName) {
      setSubmitted(false);
      createDuplicateRecipe(recipeId, recipeName);
      toggleCopyRecipeModel();
      // history.push(ROUTES.labware);
    } else {
      setSubmitted(true);
      toast.error("Enter recipe name", { autoClose: false });
    }
  };

  return (
    <>
      <Modal isOpen={isOpen} toggle={toggleCopyRecipeModel} centered size="sm">
        <ModalBody className="p-0">
          <div className="d-flex justify-content-center align-items-center modal-heading">
            <Text className="mb-0 title font-weight-bold">{deckName}</Text>
            <ButtonIcon
              position="absolute"
              placement="right"
              top={0}
              right={16}
              size={36}
              name="cross"
              onClick={toggleCopyRecipeModel}
              className="border-0"
            />
          </div>
          <div className="d-flex justify-content-center align-items-center flex-column h-100 pt-5 pb-4">
            <Text
              Tag="h4"
              size={15}
              className="text-center font-weight-bold mb-2"
            >
              {confirmationCopyText}
            </Text>
            <Text
              Tag="h4"
              size={15}
              className="text-center font-weight-bold mb-4"
            >
              {confirmationText}
            </Text>
            <Form>
              <AddNewRecipesForm>
                <FormGroup
                  row
                  className="d-flex align-items-center justify-content-center mb-5"
                >
                  <Input
                    type="text"
                    name="Type here"
                    id="Type here"
                    placeholder="Type here"
                    value={recipeName}
                    onChange={onChangeRecipeName}
                    className="recipe-name"
                  />
                  {submitted && !recipeName ? (
                    <FormError>Incorrect Recipe Name</FormError>
                  ) : null}
                </FormGroup>
                <Center className="my-3">
                  <Button color="primary" onClick={onCreateRecipeClicked}>
                    Copy Recipe
                  </Button>
                </Center>
              </AddNewRecipesForm>
            </Form>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
};

DuplicateRecipeModal.propTypes = {
  confirmationText: PropTypes.string,
  isOpen: PropTypes.bool,
  // confirmationClickHandler: PropTypes.func,
};

DuplicateRecipeModal.defaultProps = {
  confirmationText: "Recipe Name",
  isOpen: false,
};

export default React.memo(DuplicateRecipeModal);
