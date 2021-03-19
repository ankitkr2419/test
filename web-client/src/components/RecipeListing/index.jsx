import React, { useState } from 'react';

import { Card, CardBody, Button, Row, Col } from 'core-components';
import {
	Icon,
    Text,
    ButtonIcon
} from 'shared-components';


import styled from 'styled-components';
import AppFooter from 'components/AppFooter';
import RecipeFlowModal from 'components/modals/RecipeFlowModal';
import ConfirmationModal from 'components/modals/ConfirmationModal';
import { Fade } from 'reactstrap';
import TrayDiscardModal from 'components/modals/TrayDiscardModal';
import RecipeCard from 'components/RecipeListing/RecipeCard';

const TopContent = styled.div`
	margin-bottom:2.25rem;
`;

const HeadingTitle = styled.label`
    font-size: 1.25rem;
    line-height: 1.438rem;
`;

const RecipeListingComponent = (props) => {
    const [fadeIn, setFadeIn] = useState(true);
    const toggle = () => setFadeIn(!fadeIn);

	return (
		<div className="ml-content">
			<div className='landing-content px-2'>
				<ConfirmationModal
						isOpen={false}
				/>
				<RecipeFlowModal/>
			

				<TopContent className="d-flex justify-content-between align-items-center mx-5">
						<div className="d-flex align-items-center">
								<Icon name="angle-left" size={32} className="text-white"/>
								<HeadingTitle Tag="h5" className="text-white font-weight-bold ml-3 mb-0">Select a Recipe for Deck B</HeadingTitle>
						</div>
						<div className="">
							<Icon name="download" size={19} className="text-white mr-3"/>
							<Button
								color="secondary"
								className="ml-auto"
						>	Clean Up       
						</Button>
						<TrayDiscardModal />
					</div>
				</TopContent>
				<Card>
						<CardBody className="p-5">
								<Row>
										<Col> <RecipeCard /> </Col>
										<Col> <RecipeCard /> </Col>
								</Row>
						</CardBody>
				</Card>
			</div>
      <AppFooter />
		</div>
	);
};

RecipeListingComponent.propTypes = {};

export default RecipeListingComponent;
