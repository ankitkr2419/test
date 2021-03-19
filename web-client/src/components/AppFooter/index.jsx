import React from 'react';
import DeckCard from 'shared-components/DeckCard';

const AppFooter = (props) => {

	return (
        <div className="d-flex justify-content-center align-items-center">
            <DeckCard cardName={"Deck A"}/>
            <DeckCard cardName={"Deck B"}/>
		</div>
	);
};

AppFooter.propTypes = {};

export default AppFooter;
