import React, { useState } from "react";

import PropTypes from "prop-types";
import styled from "styled-components";
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
import { updateRecipe } from "action-creators/saveNewRecipeActionCreators";

const AddNewRecipesForm = styled.div`
    .recipe-name {
        width: 362px;
    }
    label {
        font-size: 0.813rem;
        line-height: 0.938rem;
    }
`;

const AddNewRecipesModal = (props) => {
    const {
        confirmationText,
        isOpen,
        toggleAddNewRecipesModal,
        deckName,
    } = props;

    const [recipeName, setRecipeName] = useState("");
    const [submitted, setSubmitted] = useState(false);

    const dispatch = useDispatch();

    const onChangeRecipeName = (e) => {
        setSubmitted(false);
        let val = e.target.value;
        setRecipeName(val);
    };

    const onCreateRecipeClicked = () => {
        if (recipeName) {
            setSubmitted(false);

            dispatch(
                updateRecipe({
                    name: recipeName,
                })
            );

            //go to labware page
            // history.push(ROUTES.labware);
        } else {
            setSubmitted(true);
            toast.error("Enter recipe name");
        }
    };

    return (
        <>
            {/* <Button color="primary" onClick={toggleAddNewRecipesModal}>
                Add New Recipes Modal
            </Button> */}
            <Modal
                isOpen={isOpen}
                toggle={toggleAddNewRecipesModal}
                centered
                size="sm"
            >
                <ModalBody className="p-0">
                    <div className="d-flex justify-content-center align-items-center modal-heading">
                        <Text className="mb-0 title font-weight-bold">
                            {deckName}
                        </Text>
                        <ButtonIcon
                            position="absolute"
                            placement="right"
                            top={0}
                            right={16}
                            size={36}
                            name="cross"
                            onClick={toggleAddNewRecipesModal}
                            className="border-0"
                        />
                    </div>
                    <div className="d-flex justify-content-center align-items-center flex-column h-100 pt-5 pb-4">
                        <Text
                            Tag="h5"
                            size={20}
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
                                        <FormError>
                                            Incorrect Recipe Name
                                        </FormError>
                                    ) : null}
                                </FormGroup>
                                <Center className="my-3">
                                    <Button
                                        color="primary"
                                        onClick={onCreateRecipeClicked}
                                    >
                                        Create Recipe
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

AddNewRecipesModal.propTypes = {
    confirmationText: PropTypes.string,
    isOpen: PropTypes.bool,
    // confirmationClickHandler: PropTypes.func,
};

AddNewRecipesModal.defaultProps = {
    confirmationText: "Name the Recipe",
    isOpen: false,
};

export default AddNewRecipesModal;
