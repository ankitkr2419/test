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
import SearchBox from 'shared-components/SearchBox';
import PaginationBox from 'shared-components/PaginationBox';

const TopContent = styled.div`
	margin-bottom:2.25rem;
`;

const HeadingTitle = styled.label`
    font-size: 1.25rem;
    line-height: 1.438rem;
`;
const RecipeCard = styled.div`
    padding: 0.8rem 0.5rem;
    border: 1px solid #E3E3E3;
    border-radius: 0.5rem;
    margin-bottom:0.688rem;
    box-shadow: 0px 3px 16px rgba(0,0,0,0.04);
    // width:27.5rem;
    // height: 5.563rem;
    .recipe-heading{
        padding-bottom:0.5rem;
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
        .recipe-action{
            button {
                width:33px !important;
                height:33px !important;
                border:1px solid #696969 !important;
                &:not(:first-child){
                    margin-left:12px;
                }
            }
        }
    }
    &:focus{
        background-color:rgba(243, 130, 32, 0.30);
    }
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
						<div className="d-flex justify-content-center align-items-center">
						<ButtonIcon
							size={19}
							name="download"
							//onClick={toggleOperatorLoginModal}
							className="border-0 text-white"
						/>
							{/* <Icon name="download" size={19} className="text-white mr-3"/> */}
							<Button
								color="secondary"
								className="ml-2"
						>	Clean Up       
						</Button>
						<TrayDiscardModal />
						<Button
								color="secondary"
								className="ml-2"
						>	<Icon name="plus-1" size={14} className="text-primary mr-2"/> Add New Recipe       
						</Button>
						<ButtonIcon
							size={36}
							name="cross"
							//onClick={toggleOperatorLoginModal}
							className="border-0 text-white mr-2"
						/>
					</div>
				</TopContent>
				<Card>
					<CardBody className="p-5">
						<div className="d-flex justify-content-between align-items-center">
							<SearchBox />
							<PaginationBox />
						</div>
						<Row>
								<Col>
										<RecipeCard onClick={toggle}>
												<div className="font-weight-bold recipe-heading">Name Name Name Name Name Name Name</div>
												<div className="recipe-card-body">
														<Text Tag="span" className="recipe-name">Total Processes -</Text>
														<Text Tag="span" className="text-primary font-weight-bold recipe-value ml-2">347 </Text>
														<Fade in={fadeIn} tag="h5" className="m-0 d-none">
														<div className="recipe-action d-flex justify-content-between align-items-center">
																<div className="d-flex justify-content-between align-items-center">
																		<ButtonIcon
																				size={14}
																				name='play'
																				className="border-gray text-primary"
																				//onClick={toggleExportDataModal}
																		/>
																		<ButtonIcon
																				size={14}
																				name='edit-pencil'
																				className="border-gray text-primary"
																				//onClick={toggleExportDataModal}
																		/>
																		<ButtonIcon
																				size={14}
																				name='publish'
																				className="border-gray text-primary"
																				//onClick={toggleExportDataModal}
																		/>
																</div>
																<ButtonIcon
																		size={20}
																		name='minus-1'
																		className="border-gray text-primary"
																		//onClick={toggleExportDataModal}
																/>
														</div>
												</Fade>
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
