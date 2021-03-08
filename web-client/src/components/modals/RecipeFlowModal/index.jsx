import React, { useState } from 'react';
// import styled from 'styled-components';
import { ButtonIcon} from 'shared-components';
import { Button, Modal, ModalBody} from 'core-components';
import {
  Carousel,
  CarouselItem,
  CarouselControl,
  CarouselIndicators,
  CarouselCaption,
  ModalHeader,
  ModalFooter
} from 'reactstrap';

const items = [
    {
      src: 'assets/images/slide-1.jpg',
      altText: 'Slide 1',
      caption: 'Slide 1'
    },
    {
      src: 'assets/images/slide-2.jpg',
      altText: 'Slide 2',
      caption: 'Slide 2'
    },
    {
      src: 'assets/images/slide-3.jpg',
      altText: 'Slide 3',
      caption: 'Slide 3'
    }
];



const RecipeFlowModal = (props) => {
 	const [activeIndex, setActiveIndex] = useState(0);
    const [animating, setAnimating] = useState(false);

    const next = () => {
        if (animating) return;
        const nextIndex = activeIndex === items.length - 1 ? 0 : activeIndex + 1;
        setActiveIndex(nextIndex);
    }

    const previous = () => {
        if (animating) return;
        const nextIndex = activeIndex === 0 ? items.length - 1 : activeIndex - 1;
        setActiveIndex(nextIndex);
    }

    const goToIndex = (newIndex) => {
        if (animating) return;
        setActiveIndex(newIndex);
    }

    const slides = items.map((item) => {
      return (
        <CarouselItem
          onExiting={() => setAnimating(true)}
          onExited={() => setAnimating(false)}
          key={item.src}
        >
          <img src={item.src} alt={item.altText} />
          <CarouselCaption captionText={item.caption} captionHeader={item.caption} />
        </CarouselItem>
      );
    });

    const [modal, setModal] = useState(false);
    const toggle = () => setModal(!modal);
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
			>
        <ModalHeader toggle={toggle}>Modal title</ModalHeader>
				<ModalBody>
          <Carousel
            activeIndex={activeIndex}
            next={next}
            previous={previous}
          >
            <CarouselIndicators items={items} activeIndex={activeIndex} onClickHandler={goToIndex} />
            {slides}
            <CarouselControl direction="prev" directionText="Previous" onClickHandler={previous} />
            <CarouselControl direction="next" directionText="Next" onClickHandler={next} />
          </Carousel>
				</ModalBody>

        <ModalFooter>
          <Button color="secondary" onClick={toggle}>Next</Button>
        </ModalFooter>
			</Modal>
		</>
	);
};

RecipeFlowModal.propTypes = {};

export default RecipeFlowModal;
