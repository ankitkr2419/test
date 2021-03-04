import React from 'react';

import SearchBox from 'shared-components/SearchBox';
import DeckCard from 'shared-components/DeckCard';
import ButtonBar from 'shared-components/ButtonBar';

const LandingScreenComponent = (props) => {
	return (
		<div className='landing-content h-100 py-0'>
            <SearchBox/>
			<div className="d-flex justify-content-center align-items-center">
				<DeckCard/>
				<DeckCard/>
			</div>
			<ButtonBar/>
		</div>
	);
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
