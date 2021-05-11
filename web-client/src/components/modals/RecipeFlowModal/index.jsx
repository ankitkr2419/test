import React, { useState } from "react";
// import styled from 'styled-components';
import { ButtonIcon, Center, ImageIcon } from "shared-components";
import { Button, Modal, ModalBody } from "core-components";
import Slide1 from "../../../assets/images/slide-1.jpg";
import Slide2 from "../../../assets/images/slide-2.jpg";
import Slide3 from "../../../assets/images/slide-3.jpg";
import styled from "styled-components";
import Slider from "react-slick";

const RecipeFlowSlider = styled.div`
  .slides {
    .slides-inner-box {
      width: 43.5rem;
      height: 100%;
      // height:25rem;
      margin: 0 auto;
      overflow: hidden !important;
      img {
        border-radius: 1.5rem !important;
        box-shadow: 0px 3px 6px rgba(0, 0, 0, 0.16) !important;
      }
    }
  }
  .slick-dots {
    bottom: -2.5rem !important;
    li button:before {
      font-size: 0.75rem;
    }
    li.slick-active button:before {
      transform: scale(1.5);
      color: #9AD0C8;
    }
  }
  .center {
    .slick-list {
      padding-top: 1.875rem !important;
      padding-bottom: 1.875rem !important;
    }
    .slick-center .slides-inner-box {
      transform: scale(1.12);
      overflow: hidden;
      border-radius: 1.5rem;
    }
    .slides {
      -webkit-transition: all 0.3s ease-out;
      transition: all 0.3s ease-out;
    }
    .slick-next, .slick-prev{
      background-color:#9AD0C8;
      z-index:1;
      width:3rem;
      height:6.063rem;
      box-shadow:0px 3px 6px rgba(0,0,0,0.16);
    }
    .slick-next{
      right:-1px;
      border-radius:3.125rem 0 0 3.125rem;
      &::before{
        background: url("/images/next-arrow.svg") no-repeat;
        background-position:top center;
        position: relative;
        left: 5px;
        background-size:contain;
        color:transparent;
      }
    }
    .slick-prev{
      left:-1px;
      border-radius:0 3.125rem 3.125rem 0;
      &::before{
        background: url("/images/prev-arrow.svg") no-repeat;
        background-position:top center;
        position: relative;
        right: 5px;
        background-size:contain;
        color:transparent;
      
      }
    }
  }
`;
const NextButton = styled.div`
  position: absolute;
  bottom: 1.5rem;
  right: 6rem;
`;

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

export default React.memo(RecipeFlowModal);
