import React from 'react';

import styled from 'styled-components'
import { ImageIcon } from 'shared-components';
import CirclelogoIcon from 'assets/images/mylab-logo-with-circle.svg';

const SplashScreen = styled.div`
    background: url('/images/logo-bg.svg') left -4.875rem top -5.5rem no-repeat,
    url('/images/honey-bees-bg.svg') right -1.75rem bottom -1.5rem no-repeat;
`;
const CircleImg = styled.div`
    margin-right: 14.313rem;
    margin-left: auto;
`;

const SplashScreenComponent = (props) => {
	return (
		<SplashScreen className='splash-screen-content h-100 py-0 bg-white d-flex justify-content-center align-items-center'>
            <CircleImg>
                <ImageIcon 
                src={CirclelogoIcon} 
                alt="My Lab" 
                />
            </CircleImg>
        </SplashScreen>
	);
};

SplashScreenComponent.propTypes = {};

export default SplashScreenComponent;
