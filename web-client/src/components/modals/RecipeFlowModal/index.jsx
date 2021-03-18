import React, { useState } from 'react';
// import styled from 'styled-components';
import { ButtonIcon, Center, ImageIcon} from 'shared-components';
import { Button, Modal, ModalBody} from 'core-components';
import Slide1 from '../../../assets/images/slide-1.jpg';
import Slide2 from '../../../assets/images/slide-2.jpg';
import Slide3 from '../../../assets/images/slide-3.jpg';
import styled from 'styled-components';
import Slider from "react-slick";

// const items = [
//     {
//       src: Slide1,
//       altText: '',
//       caption: ''
//     },
//     {
//       src: Slide2,
//       altText: '',
//       caption: ''
//     },
//     {
//       src: Slide3,
//       altText: '',
//       caption: ''
//     }
// ];
const RecipeFlowSlider = styled.div`
  .slides{
    .slides-inner-box{
      width:43.5rem;
      height:100%;
      // height:25rem;
      margin:0 auto;
      overflow: hidden !important;
      img{
      border-radius:1.5rem !important;
      box-shadow:0px 3px 6px rgba(0,0,0,0.16) !important;
      }
    }
  }
  .slick-dots{
    bottom:-2.5rem !important;
    li button:before{
      font-size:0.75rem;
    }
    li.slick-active button:before{
      transform:scale(1.5);
      color:#F38220;
    }
  }
  .center {
    .slick-list{
      padding-top: 30px !important;
      padding-bottom: 30px !important;
    }
    .slick-center .slides-inner-box {
      transform: scale(1.12);
      overflow: hidden;
      border-radius: 1.5rem;
    }
    .slides{
      -webkit-transition: all 0.3s ease-out;
      transition: all 0.3s ease-out;
    }
  }
`;
const NextButton = styled.div`
  position: absolute;
  bottom: 1.5rem;
  right: 6rem;
`;


const RecipeFlowModal = (props) => {
    const [modal, setModal] = useState(false);
    const toggle = () => setModal(!modal);
    const recipeFlowsettings = {
      className: "center",
      centerMode: true,
      centerPadding: "65px",
      dots: true,
      infinite: true,
      speed: 500,
      slidesToShow: 1,
      slidesToScroll: 1,
      arrows: false,
    };
	return (
		<>
			<ButtonIcon
				size={34}
				name='external-link'
				className='mx-2'
				onClick={toggle}
			/>
      <Modal
				isOpen={true}
				toggle={toggle}
				centered
				size='lg'
        className="recipe-flow-modal"
			>
				<ModalBody className="py-5 px-0 recipe-flow-modal-body">
            <ButtonIcon
            position="absolute"
            placement="right"
            top={16}
            right={16}
            size={36}
            name="cross"
            onClick={toggle}
            className="ml-auto border-0"
            />
            <Center className="font-weight-bold mb-4">Name Name Name Name Name Name Name</Center>
            <RecipeFlowSlider className="mb-4">
              <Slider {...recipeFlowsettings}>
                  <div className="slides">
                    <div className="slides-inner-box">
                      <ImageIcon src={Slide1} alt=""/>
                    </div>
                  </div>
                  <div className="slides">
                    <div className="slides-inner-box">
                      <ImageIcon src={Slide2} alt=""/>
                    </div>
                  </div>
                  <div className="slides position-relative">
                    <div className="slides-inner-box">
                      <ImageIcon src={Slide3} alt=""/>
                    </div>
                  </div>
                </Slider>  
            </RecipeFlowSlider>
            <NextButton>
              <Button
              color="outline-secondary"
              size="sm"
              >
							Next
						</Button>
					</NextButton>          
				</ModalBody>
			</Modal>
		</>
	);
};

RecipeFlowModal.propTypes = {};

export default RecipeFlowModal;