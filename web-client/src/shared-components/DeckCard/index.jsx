import React from 'react';

import styled from 'styled-components';
// import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { Text, Icon } from 'shared-components';
import {
	Button, 
} from 'core-components';

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
`;
const DeckContent = styled.div`
position:relative;
background:#fff url('/images/deck-card-bg.svg')no-repeat;
> button{
	min-width:7.063rem;
	height:2.5rem;
	line-height:1.125rem;
}
`;
const DeckTitle = styled.div`
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
`;
const ActionButton = styled.button`
	position: absolute;
	top:0;
	right:4rem;
	z-index:2;
  display: block;
	outline: none;
	border: 0;
	background: transparent;
`;
 
const SemiCircleOutterBox = styled.div`
	width: 62px;
	height: 40px;
	background-color: #B3D9D0;
	border-bottom-left-radius: 88px;
	border-bottom-right-radius: 88px;
	box-shadow: 0px -3px 6px rgb(0 0 0 / 31%);
	display: flex;
	justify-content: center;
	align-items: center;
`;

const SemiCircularButton = styled.div `
	width: 44px;
	height: 44px;
	border-radius: 50%;
	border: 1px solid #F0801D;
	background-color: #F38220;
	z-index: 1;
	text-decoration: none;
	margin-top: -22px;
	display: flex;
	justify-content: center;
	align-items: center;
	color:#fff;
`;

const DeckCard = (props) => {
	return (
		<DeckCardBox className="d-flex justify-content-start align-items-center">
			<DeckTitle className="d-flex justify-content-center align-items-center"> 
				<Text Tag="label" size="20" >Deck A</Text>
			</DeckTitle>
			<DeckContent className="d-flex justify-content-between align-items-center p-4 w-100 h-100">
				<ActionButton className="d-flex justify-content-center align-items-center">
					<SemiCircleOutterBox>
							<SemiCircularButton>
								<Icon name='play' size={18} />
							</SemiCircularButton>
					</SemiCircleOutterBox>
				</ActionButton>
				
				<div className="">
					<Text Tag="h5">Recipe Name</Text>
					<Text Tag="label">Current Processes - (Process Name)</Text>
					{/* <ProgressBar variant="info" now={20} /> */}
				</div>
				<Button
					color="primary"
					className="ml-auto"
					size="sm"
				>	Login       
				</Button>
			</DeckContent>
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
