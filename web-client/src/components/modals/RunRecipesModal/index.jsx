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

import Radio from "core-components/Radio";
import { RUN_RECIPE_TYPE } from "appConstants";
import { RunRecipeTypeForm } from './RunRecipeTypeForm';

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

export default React.memo(RunRecipesModal);
