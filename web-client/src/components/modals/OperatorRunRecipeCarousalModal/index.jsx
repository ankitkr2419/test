import React, { useState } from "react";
import Slider from "react-slick";

import { ButtonIcon, Center, ImageIcon } from "shared-components";
import { Button, Modal, ModalBody } from "core-components";

import Slide1 from "../../../assets/images/slide-1.jpg";
import Slide2 from "../../../assets/images/slide-2.jpg";
import Slide3 from "../../../assets/images/slide-3.jpg";

import { RecipeFlowSlider, NextButton } from './Style';

const OperatorRunRecipeCarousalModal = (props) => {
  const { isOpen, handleCarousalModal, onConfirmedRecipeSelection } = props;

  const [isNextButtonVisible, setNextButtonVisbile] = useState(false);

  const LAST_SLIDE = 2;

  const recipeFlowsettings = {
    className: "center",
    centerMode: true,
    centerPadding: "65px",
    dots: true,
    infinite: false,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    arrows: true,
    afterChange: (currentSlide) => {
      if (currentSlide === LAST_SLIDE) {
        setNextButtonVisbile(true);
      } else {
        setNextButtonVisbile(false);
      }
    },
  };

  const toggleModalView = () => {
    handleCarousalModal(isOpen);
  };

  // const handleStartRunRecipe = () => {
  //   toggleModalView();
  // };

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
        <Center className="font-weight-bold mb-4">Recipe Name</Center>
        <RecipeFlowSlider className="mb-4">
          <Slider {...recipeFlowsettings}>
            <div className="slides">
              <div className="slides-inner-box">
                <ImageIcon src={Slide1} alt="" />
              </div>
            </div>
            <div className="slides">
              <div className="slides-inner-box">
                <ImageIcon src={Slide2} alt="" />
              </div>
            </div>
            <div className="slides position-relative">
              <div className="slides-inner-box">
                <ImageIcon src={Slide3} alt="" />
              </div>
            </div>
          </Slider>
        </RecipeFlowSlider>

        {isNextButtonVisible && (
          <NextButton>
            <Button
              className="border-primary"
              color="outline-secondary"
              size="sm"
              onClick={() => {
                handleCarousalModal();//hide/show
                onConfirmedRecipeSelection();//save data in reducer
              }}
            >
              Next
            </Button>
          </NextButton>
        )}
      </ModalBody>
    </Modal>
  );
};

export default OperatorRunRecipeCarousalModal;
