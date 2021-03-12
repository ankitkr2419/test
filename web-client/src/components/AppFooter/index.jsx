import React from 'react';

import DeckCard from 'shared-components/DeckCard';

const AppFooter = (props) => {
	return (
        <div className="d-flex justify-content-center align-items-center">
            <DeckCard/>
            <DeckCard/>
		</div>
	);
};

AppFooter.propTypes = {};

export default AppFooter;
