import React, { useState } from 'react';
import { useHistory } from "react-router-dom";

import styled from 'styled-components'
import { Modal, ModalBody, Button} from 'core-components';
import { ImageIcon, ButtonIcon, Icon, Text } from 'shared-components';

import CirclelogoIcon from 'assets/images/mylab-logo-with-circle.png';
import thumbsUpGraphics from 'assets/images/thumbs-up-graphic.svg';
import { ROUTES } from "appConstants"

import Radio from 'core-components/Radio';

const SplashScreen = styled.div`
    background: url('/images/logo-bg.svg') left -4.875rem top -5.5rem no-repeat,
    url('/images/honey-bees-bg.svg') right -1.75rem bottom -1.5rem no-repeat;
    .circle-image{
        margin-right: 14.313rem;
        margin-left: auto;
    };
    cursor: pointer
`;

const OptionBox = styled.div`
	background-color:#F5F5F5;
	border-radius:36px 0 0 36px;
	.large-btn{
		width:15.125rem;
		height:8.5rem;
		margin-bottom: 2.125rem;
	}
`;
const CloseButton = styled.div`
		position:absolute;
		top:1rem;
		right:1rem;
`;

const SplashScreenComponent = (props) => {

	const [modal, setModal] = useState(false);
  const history = useHistory();
	const toggle = () => setModal(!modal);

  const redirectToLandingPage = () => {
    return history.push(ROUTES.landing)
  }

  return (
		<SplashScreen className='splash-screen-content h-100 py-0 bg-white d-flex justify-content-center align-items-center' onClick={redirectToLandingPage}>
			<div className="circle-image">
					<ImageIcon
					src={CirclelogoIcon}
					alt="My Lab"
					/>
			</div>
			{/* Alert pop up2 */}
    </SplashScreen>
	);
};

SplashScreenComponent.propTypes = {};

export default SplashScreenComponent;
