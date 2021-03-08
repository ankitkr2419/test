import React, {useState} from 'react';

import { Card, CardBody, Button, Row, Col } from 'core-components';
import {
	Icon,
    Text
} from 'shared-components';


import styled from 'styled-components';
import AppFooter from 'components/AppFooter';
import RecipeFlowModal from 'components/modals/RecipeFlowModal';
import ConfirmationModal from 'components/modals/ConfirmationModal';

const TopContent = styled.div`
	margin-bottom:2.25rem;
`;

const HeadingTitle = styled.label`
    font-size: 1.25rem;
    line-height: 1.438rem;
`;
const RecipeCard = styled.div`
    padding: 1rem 0.5rem;
    border: 1px solid #E3E3E3;
    border-radius: 0.5rem;
    margin-bottom:0.688rem;
    // width:27.5rem;
    // height: 5.563rem;
    .recipe-heading{
        padding-bottom:0.781rem;
    }
    .recipe-card-body{
        padding-top:0.25rem;
        border-top: 1px solid #d9d9d9;
        
        .recipe-name{
            font-size:0.875rem;
            line-height:1rem;
        }
        .recipe-value{
            font-size:1.125rem;
            line-height:1.313rem;
        }
    }
`;


const RecipeListingComponent = (props) => {
	return (
		<div className="ml-content">
			<div className='landing-content'>
            <ConfirmationModal
                isOpen={true}
            />
                <RecipeFlowModal />
                <TopContent className="d-flex justify-content-between align-items-center">
                    <div className="d-flex align-items-center">
                        <Icon name="angle-left" size={32} className="text-white"/>
                        <HeadingTitle Tag="h5" className="text-white font-weight-bold ml-3 mb-0">Select a Recipe for Deck B</HeadingTitle>
                    </div>
                    <div className="">
                     <Icon name="download" size={19} className="text-white mr-3"/>
                     <Button
                        color="secondary"
                        className="ml-auto"
                        size="sm"
                    >	Clean Up       
				</Button>
                    </div>
                </TopContent>
                <Card>
                    <CardBody className="p-5">
                        <Row>
                            <Col>
                                <RecipeCard>
                                    <div className="font-weight-bold recipe-heading">Name Name Name Name Name Name Name</div>
                                    <div className="recipe-card-body">
                                        <Text Tag="span" className="recipe-name">Total Processes -</Text>
                                        <Text Tag="span" className="text-primary font-weight-bold recipe-value ml-2">347 </Text>
                                    </div>
                                </RecipeCard>
                            </Col>
                            <Col>
                                <RecipeCard>
                                    <div className="font-weight-bold recipe-heading">Name Name Name Name Name Name Name</div>
                                    <div className="recipe-card-body">
                                        <Text Tag="span" className="recipe-name">Total Processes -</Text>
                                        <Text Tag="span" className="text-primary font-weight-bold recipe-value ml-2">347 </Text>
                                    </div>
                                </RecipeCard>
                            </Col>
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
