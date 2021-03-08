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
	right:4rem;
	z-index:2;
    display: block;
	outline: none;
	border: 0;
	background: transparent;
	.semi-circle-outter-box{
		width: 62px;
		height: 40px;
		background-color: #B3D9D0;
		border-bottom-left-radius: 88px;
		border-bottom-right-radius: 88px;
		box-shadow: 0px -3px 6px rgb(0 0 0 / 31%);
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.semi-circular-button{
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
	}
`;
 
const ActionButton = (props) => {
	return (
		<ActionBtn className="d-flex justify-content-center align-items-center">
            <div className="semi-circle-outter-box">
                <div className="semi-circular-button">
                    <Icon name='play' size={18} />
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
