import React from "react";
import Slider from "react-slick";

import { Center, ImageIcon, Text } from "shared-components";

const OperatorCarousal = (props) => {
  const { images, setNextButtonVisbile } = props;

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

  const msg = [
    "Deck Consumable Details",
    "Extraction Cartridge",
    "Pcr Cartridge",
  ];

  const slides = images.map((image, index) => {
    return (
      <div key={index} className="slides">
        <Center className="mb-5">
          <Text Tag="h3" size={16} className="font-weight-bold">
            {msg[index]}
          </Text>
        </Center>
        <div className="slides-inner-box">
          <ImageIcon src={image} alt="" />
        </div>
      </div>
    );
  });

  return <Slider {...recipeFlowsettings}>{slides}</Slider>;
};

export default React.memo(OperatorCarousal);
