import React, { useState } from 'react';

import DeckCard from 'shared-components/DeckCard';

const DeckCardContainer = (props) => {
    const [operatorLoginModalOpen, setOperatorLoginModalOpen] = useState(false);
    const toggleOperatorLoginModal = () => setOperatorLoginModalOpen(!operatorLoginModalOpen);
    
    return(
        <DeckCard 
            operatorLoginModalOpen={operatorLoginModalOpen}
            toggleOperatorLoginModal={toggleOperatorLoginModal}
        />
    )
}

DeckCardContainer.propTypes = {};

export default DeckCardContainer;

