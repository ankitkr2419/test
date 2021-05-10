import React, { useState } from "react";
// import styled from 'styled-components';
import { ButtonIcon, Center, ImageIcon } from "shared-components";
import { Button, Modal, ModalBody } from "core-components";
import Slide1 from "../../../assets/images/slide-1.jpg";
import Slide2 from "../../../assets/images/slide-2.jpg";
import Slide3 from "../../../assets/images/slide-3.jpg";
import Slider from "react-slick";

import { RecipeFlowSlider, NextButton } from './Style';

const RecipeFlowModal = (props) => {
  const [enableNext, setEnableNext] = useState(false);

  const { isOpen, toggle, toggleShowProcess, recipeData } = props;
  const { recipeId, recipeName, processCount } = recipeData;
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
        setEnableNext(true);
      } else {
        setEnableNext(false);
      }
    },
  };

  return (
    <>
      {/* <ButtonIcon
				size={34}
				name='external-link'
				className='mx-2'
				onClick={toggle}
			/> */}
      <Modal isOpen={isOpen} centered size="lg" className="recipe-flow-modal">
        <ModalBody className="py-5 px-0 recipe-flow-modal-body">
          <ButtonIcon
            position="absolute"
            placement="right"
            top={16}
            right={16}
            size={36}
            name="cross"
            onClick={() => {toggle(recipeId, recipeName, processCount); setEnableNext(false);}}
            className="ml-auto border-0"
          />
          <Center className="font-weight-bold mb-4">{recipeName}</Center>
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

          {(enableNext) && (
            <NextButton>
              <Button
                className="border-primary"
                color="outline-secondary"
                size="sm"
                onClick={() => {toggleShowProcess(); setEnableNext(false);}}
              >
                Next
              </Button>
            </NextButton>
          )}

        </ModalBody>
      </Modal>
    </>
  );
};

RecipeFlowModal.propTypes = {};

export default RecipeFlowModal;
