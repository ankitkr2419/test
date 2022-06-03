import React, { useState } from "react";

import { ButtonIcon, Center } from "shared-components";
import { Button, Modal, ModalBody } from "core-components";
import Slide1 from "../../../assets/images/slide-1.jpg";
import Slide2 from "../../../assets/images/slide-2.jpg";
import Slide3 from "../../../assets/images/slide-3.jpg";

import { NextButton, RecipeFlowSlider } from "./Style";
import OperatorCarousal from "./OperatorCarousal";

const OperatorRunRecipeCarousalModal = (props) => {
  const { isOpen, handleCarousalModal, onConfirmedRecipeSelection } = props;

  const [isNextButtonVisible, setNextButtonVisbile] = useState(false);

  const toggleModalView = () => handleCarousalModal(isOpen);
  const images = [Slide1, Slide2, Slide3];

  return (
    <Modal isOpen={isOpen} centered size="lg" className="recipe-flow-modal">
      <ModalBody className="py-5 px-0 recipe-flow-modal-body">
        <ButtonIcon
          position="absolute"
          placement="right"
          top={16}
          right={16}
          size={36}
          name="cross"
          onClick={toggleModalView}
          className="ml-auto border-0"
        />
        {/* <Center className="font-weight-bold mb-4">Recipe Name</Center> */}
        <RecipeFlowSlider className="mb-4">
          <OperatorCarousal
            images={images}
            setNextButtonVisbile={setNextButtonVisbile}
          />
        </RecipeFlowSlider>

        {/** next button will always visible but label will be 'skip' instead 'next' for previous slides */}
        <NextButton>
          <Button
            className="border-primary"
            color="outline-secondary"
            size="sm"
            onClick={() => {
              handleCarousalModal(); //hide/show
              onConfirmedRecipeSelection(); //save data in reducer
            }}
          >
            {isNextButtonVisible ? "Next" : "Skip"}
          </Button>
        </NextButton>
      </ModalBody>
    </Modal>
  );
};

export default React.memo(OperatorRunRecipeCarousalModal);
