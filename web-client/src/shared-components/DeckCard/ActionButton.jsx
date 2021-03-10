import React from 'react';

import styled from 'styled-components';
// import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { Icon } from 'shared-components';
import { 
} from 'core-components';

const ActionBtn = styled.button`
	position: absolute;
	top:0;
	right:0;
	z-index:2;
    display: block;
	outline: none;
	border: 0;
	background: transparent;
	.semi-circle-outter-box{
		width: 3.875rem;
		height: 2.5rem;
		background-color: #B3D9D0;
		border-bottom-left-radius: 5.5rem;
		border-bottom-right-radius: 5.5rem;
		box-shadow: 0px -3px 6px rgb(0 0 0 / 31%);
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.semi-circular-button{
		width: 2.75rem;
		height: 2.75rem;
		border-radius: 50%;
		border: 1px solid #F0801D;
		background-color: #F38220;
		z-index: 1;
		text-decoration: none;
		margin-top: -1.375rem;
		display: flex;
		justify-content: center;
		align-items: center;
		color:#fff;	
		position:relative;
		.btn-label{
			position:absolute;
			bottom: -2rem;
			left: 0;
			right: 0;
			font-size:0.75rem;
			line-height:0.875rem;
			color:#3C3C3C;
		}
	}
`;
 
const ActionButton = (props) => {
	return (
		<ActionBtn className="d-flex justify-content-center align-items-center">
            <div className="semi-circle-outter-box">
                <div className="semi-circular-button">
                    <Icon name='play' size={18} className="ml-2"/>
					<div className="btn-label font-weight-bold">RUN</div>
                </div>
            </div>
        </ActionBtn>
	);
};

ActionButton.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

ActionButton.defaultProps = {
	isUserLoggedIn: false,
};

export default ActionButton;
