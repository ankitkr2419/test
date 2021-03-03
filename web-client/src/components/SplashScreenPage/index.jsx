import React from 'react';

import { ImageIcon } from 'shared-components';
import './splashscreen.scss';
import CirclelogoIcon from 'assets/images/mylab-logo-with-circle.svg';


const SplashScreenComponent = (props) => {
	return (
		<div className='splash-screen-content h-100 py-0 bg-white d-flex justify-content-center align-items-center'>
            <ImageIcon 
            src={CirclelogoIcon} 
            alt="My Lab" 
            className='mylab-logo-circle' />
        </div>
	);
};

SplashScreenComponent.propTypes = {};

export default SplashScreenComponent;
