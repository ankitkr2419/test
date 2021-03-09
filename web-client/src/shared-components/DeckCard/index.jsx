import React from 'react';

import styled from 'styled-components';
// import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { Text } from 'shared-components';
import {
	Button, 
} from 'core-components';
import ActionButton from "./ActionButton";
// import { ProgressBar } from './reactstrap';

const DeckCardBox = styled.div`
  width: 32rem;
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
		box-shadow: 0 -3px 6px rgba(0,0,0,0.16);
		> label{
			transform:rotate(-90deg);
			white-space:nowrap;
			margin-bottom:0;
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
	}
`;

const CardOverlay= styled.div`
	position: absolute;
	display: none;
	width: 32rem;
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
	return (
		<DeckCardBox className="d-flex justify-content-start align-items-center">
			<CardOverlay />
			<div className="d-flex justify-content-center align-items-center deck-title"> 
				<Text Tag="label" size="20" >Deck A</Text>
			</div>
			<div className="d-flex justify-content-between align-items-center p-4 w-100 h-100 deck-content">
				<div className="d-none1">
				<ActionButton/>
				
				<div className="d-none">
					<Text Tag="h5">Recipe Name</Text>
					<Text Tag="label">Current Processes - (Process Name)</Text>
					{/* <ProgressBar variant="info" now={20} /> */}
				</div>
				</div>
				<Button
					color="primary"
					className="ml-auto"
					size="sm"
				>	Login       
				</Button>
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
