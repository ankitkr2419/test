import React from 'react';
import { Table } from 'core-components';
import { ButtonIcon } from 'shared-components';
import SearchBox from 'shared-components/SearchBox';
// import './activity.scss';

const LandingScreenComponent = (props) => {
	return (
		<div className='landing-content h-100 py-0'>
            <SearchBox/>
		</div>
	);
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
