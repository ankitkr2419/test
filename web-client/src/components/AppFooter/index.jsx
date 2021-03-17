import React from 'react';

import DeckCardContainer from 'containers/DeckCardContainer';

const AppFooter = (props) => {
	return (
        <div className="d-flex justify-content-center align-items-center">
            <DeckCardContainer/>
            <DeckCardContainer/>
		</div>
	);
};

AppFooter.propTypes = {};

export default AppFooter;
