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
import { AddNewRecipesForm } from "./AddNewRecipesForm";

const EditRecipesNameModal = (props) => {
  const {
    confirmationText,
    isOpen,
    toggleEditRecipeNameModal,
    deckName,
    handleEditRecipeNameConfirmation,
  } = props;

  const [recipeName, setRecipeName] = useState("");
  const [submitted, setSubmitted] = useState(false);

  const onChangeRecipeName = (e) => {
    setSubmitted(false);
    let val = e.target.value;
    setRecipeName(val);
  };

  const onCreateRecipeClicked = () => {
    if (recipeName) {
      setSubmitted(false);

      handleEditRecipeNameConfirmation(recipeName);
      toggleEditRecipeNameModal();
    } else {
      setSubmitted(true);
      toast.error("Enter recipe name", { autoClose: false });
    }
  };

  return (
    <>
      <Modal
        isOpen={isOpen}
        toggle={toggleEditRecipeNameModal}
        centered
        size="sm"
      >
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
              onClick={toggleEditRecipeNameModal}
              className="border-0"
            />
          </div>
          <div className="d-flex justify-content-center align-items-center flex-column h-100 pt-5 pb-4">
            <Text
              Tag="h4"
              size={16}
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
                    Rename Recipe
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

EditRecipesNameModal.propTypes = {
  confirmationText: PropTypes.string,
  isOpen: PropTypes.bool,
  // confirmationClickHandler: PropTypes.func,
};

EditRecipesNameModal.defaultProps = {
  confirmationText: "Recipe Name",
  isOpen: false,
};

export default React.memo(EditRecipesNameModal);
