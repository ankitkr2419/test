import React, {useState} from 'react';

import styled from 'styled-components';
// import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { Text, Icon } from 'shared-components';
import {
	Button, 
} from 'core-components';
// import ActionButton from "./ActionButton";
// import { Progress } from 'reactstrap';
import OperatorLoginModalContainer from 'containers/OperatorLoginModalContainer';

const DeckCardBox = styled.div`
	width:50%;  
// width: 32rem;
  height: 6.625rem;
	position:relative;
	box-shadow:0px -3px 6px rgba(0,0,0,0.16);
	&::before{
		content: "";
		position:absolute;
		background-image: linear-gradient(to right, #aedbd5, #a9dac5, #afd7b0, #bed29a, #d3ca87, #dcc278, #e7b96c, #f2ae64, #f2a453, #f29942, #f38d31, #f3811f);
		width: 100%;
		height: 2px;
		top: 0;
		left: 0;
		z-index:1;
	}
	.deck-title{
		width:2.563rem;
		height:100%;
		font-size:1.25rem;
		line-height:1.688rem;
		font-weight:bold;
		color:#51575A;
		border:1px solid transparent;
		box-shadow: 0 -3px 6px rgba(0,0,0,0.16);
		> label{
			transform:rotate(-90deg);
			white-space:nowrap;
			margin-bottom:0;
		}
		&.active{
			background-color:#B2DAD1;
			border:1px solid #ffffff;
		}
	}
	.deck-content{
		position:relative;
		background:#fff url('/images/deck-card-bg.svg')no-repeat;
		> button{
			min-width:7.063rem;
			height:2.5rem;
			line-height:1.125rem;
		}
		.custom-progress-bar{
			border-radius:7px;
			background-color:#B2DAD131;
			border: 2px solid #B2DAD131;
			.progress-bar{
				//background-color:#10907A;
				border-radius:7px 0px 0px 7px;
				background-color:#72B5E6;
				animation: blink 1s linear infinite;
			}
		}
		// .uv-light-button{
		// 	position:absolute;
		// 	right:244px;
		// 	top:0;
		// }
		.resume-button{
			position:absolute;
			right:123px;
			top:0;
		}
		.abort-button{
			position:absolute;
			right:21px;
			top:0;
		}
		.hour-label{
			background-color:#F5E3D3;
			border-radius:4px 0 0 4px;
			border-right:2px solid #F38220;
			padding:3px 4px;
			font-size:0.875rem;
			line-height:1rem;
		}
		.min-label{
			font-size:0.875rem;
			line-height:1rem;
		}
		.process-count-label{
			background-color:#F5E3D3;
			border-radius:4px;
			padding:3px 4px;
			font-size:1.125rem;
			line-height:1rem;
		}
		.process-total-count{
			font-size:0.875rem;
			line-height:1rem;
		}
		.process-remaining{
			font-size:10px;
			line-height:11px;
		}
		// add this class while login 
		&.logged-in{
			background:#ffffff;
		}
	}
	@keyframes blink{
		0%{
			background-color:#9D9D9D;
		}
		50%{
			background-color:#72B5E6;
		}
		100%{
			background-color:#9D9D9D;
		}
	}
`;

const CardOverlay= styled.div`
	position: absolute;
	display: none;
	width: 50%;
	height: 6.625rem;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	background-color: rgba(0,0,0,0.28);
	z-index: 3;
	cursor: pointer;
`;
const DeckCard = (props) => {

	const { cardName } = props;
	
	const [operatorLoginModalOpen, setOperatorLoginModalOpen] = useState(false);
    const toggleOperatorLoginModal = () => setOperatorLoginModalOpen(!operatorLoginModalOpen);

	return (
		<DeckCardBox className="d-flex justify-content-start align-items-center">
			<CardOverlay />
			<div className="d-flex justify-content-center align-items-center deck-title"> 
				<Text Tag="label" size="20" >{cardName}</Text>
			</div>
			<div className="p-4 w-100 h-100 deck-content logged-in1">
			<div className="d-flex justify-content-between align-items-center">
				<div className="d-none1">
					{/* <div className="uv-light-button">
						<ActionButton/>
					</div>  */}
					{/* <div className="resume-button">
						<ActionButton/>
					</div>
					<div className="abort-button">
						<ActionButton/>
					</div> */}
					
					<div className="d-none1">
							<Text Tag="h5" size="18" className="mb-2 font-weight-bold recipe-name">Recipe Name</Text>
							{/* <Text Tag="label" className="mb-1">Current Processes - (Process Name)</Text> */}
							{/* <Text Tag="label" className="mb-1 d-flex align-items-center">
								<Icon name='timer' size={19} className="text-primary"/>
								<Text Tag="span" className="hour-label font-weight-bold ml-2"> 1 Hr </Text>
								<Text Tag="span" className="min-label ml-2 font-weight-bold">8 min</Text>
								<Text Tag="span" className="ml-1">remaining</Text>
							</Text> */}

							<Text Tag="label" className="mb-1 d-flex align-items-center">
								<Icon name='process' size={19} className="text-primary"/>
								<Text Tag="span" className="process-count-label font-weight-bold ml-2"> 4 
								<Text Tag="span" className="process-total-count font-weight-bold">/10 </Text> </Text>
								<Text Tag="span" className="ml-1 process-remaining">Processes remaining</Text>
							</Text>
						</div>
				</div>
				<Button
					color="primary"
					className="ml-auto d-flex"
					size="sm"
					onClick={toggleOperatorLoginModal}
				>	Login
				</Button>

				<OperatorLoginModalContainer 
					operatorLoginModalOpen={operatorLoginModalOpen}
					toggleOperatorLoginModal={toggleOperatorLoginModal}
				/>
				
				</div>
				{/* <Progress value="2" className="custom-progress-bar"/> */}
			</div>
		</DeckCardBox>
	);
};

DeckCard.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

DeckCard.defaultProps = {
	isUserLoggedIn: false,
};

export default DeckCard;
