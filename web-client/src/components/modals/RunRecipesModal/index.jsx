import React from "react";

import {
    Button,
    Modal,
    ModalBody,
    Form,
    FormGroup,
    FormError,
} from "core-components";
import { Text, ButtonIcon, Center } from "shared-components";

import styled from "styled-components";
import Radio from "core-components/Radio";
import { RUN_RECIPE_TYPE } from "appConstants";

//For Run Recipes Type Form
const RunRecipeTypeForm = styled.div`
    .step-run-option,
    .continuous-run-option {
        border-radius: 1rem;
        width: 22.188rem;
        height: 3.75rem;
        padding: 0 5rem;
    }
    .step-run-option {
        border: 2px dashed #dbdbdb;
        margin-bottom: 30px;
    }
    .continuous-run-option {
        border: 2px solid #dbdbdb;
        margin-bottom: 30px;
    }
    .selected {
        border-color: #b2dad1;
        box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.06);
    }
    label {
        font-size: 0.813rem;
        line-height: 0.938rem;
    }
`;

const RunRecipesModal = (props) => {
    const {
        deckName,
        isOpen,
        toggleRunRecipesModal,
        runRecipeType,
        onChange,
        onConfirmed,
    } = props;

    return (
        <Modal
            isOpen={isOpen}
            toggle={toggleRunRecipesModal}
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
                        onClick={toggleRunRecipesModal}
                        className="border-0"
                    />
                </div>
                <div className="d-flex justify-content-center align-items-center flex-column h-100 mt-5 mb-3">
                    <Form>
                        <RunRecipeTypeForm>
                            <FormGroup className="step-run-option text-center d-flex align-items-center">
                                <Radio
                                    id="step-run"
                                    name="recipe-run-option"
                                    label="Step Run"
                                    checked={runRecipeType === RUN_RECIPE_TYPE.STEP_RUN}
                                    onChange={() => onChange(RUN_RECIPE_TYPE.STEP_RUN)}
                                />
                                <FormError>Incorrect Step Run</FormError>
                            </FormGroup>
                            <FormGroup className="continuous-run-option text-center d-flex align-items-center selected">
                                <Radio
                                    id="continuous-run"
                                    name="recipe-run-option"
                                    label="Continuous  Run"
                                    checked={runRecipeType=== RUN_RECIPE_TYPE.CONTINUOUS_RUN}
                                    onChange={() => onChange(RUN_RECIPE_TYPE.CONTINUOUS_RUN)}
                                />
                                <FormError>Incorrect Step Run</FormError>
                            </FormGroup>
                            <Center className="my-3">
                                <Button color="primary" onClick={onConfirmed}>Next</Button>
                            </Center>
                        </RunRecipeTypeForm>
                    </Form>
                </div>
            </ModalBody>
        </Modal>
    );
};

export default RunRecipesModal;
