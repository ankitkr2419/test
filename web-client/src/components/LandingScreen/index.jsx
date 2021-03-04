import React from 'react';

import { Card, CardBody } from 'core-components';
import {
	ImageIcon,
	Center,
	Icon
} from 'shared-components';

// import SearchBox from 'shared-components/SearchBox';
// import ButtonBar from 'shared-components/ButtonBar';
import DeckCard from 'shared-components/DeckCard';
import imgNoTemplate from 'assets/images/video-thumbnail.png';
import styled from 'styled-components';

const VideoPlayButton = styled.button`
	color:#7C7976;
	background-color:transparent;
	border:0;
	outline:none;
	position:absolute;
	top:50%;
	left:50%;
	transform:translate(-50%,-50%);
`;

const LandingScreenComponent = (props) => {
	return (
		<div className="ml-content">
			<div className='landing-content'>
			<Card className='card-video'>
				<CardBody className='d-flex flex-column p-0'>
					<Center className="video-thumbnail-wrapper">
					<ImageIcon
						src={imgNoTemplate}
						alt="No templates available"
						className="img-video-thumbnail"
					/>
					<VideoPlayButton>
						<Icon name="play" size={124}/>
					</VideoPlayButton>
					</Center>
				</CardBody>
			</Card>
			
				{/* <SearchBox/> */}
				
				{/* <ButtonBar/> */}
			</div>

			<div className="d-flex justify-content-center align-items-center">
				<DeckCard/>
				<DeckCard/>
			</div>
		</div>
	);
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
