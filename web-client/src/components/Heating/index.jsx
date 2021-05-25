import React from 'react';

import { 
	Card, 
	CardBody
} from 'core-components';
import {
	ButtonIcon,
  ButtonBar
} from 'shared-components';

import styled from 'styled-components';
import AppFooter from 'components/AppFooter';
import HeatingProcess from './HeatingProcess';
import TopHeading from 'shared-components/TopHeading';

const PageBody = styled.div`
background-color:#f5f5f5;
`;
const HeatingBox = styled.div`
.process-heating {
	&::after {
		background:url('/images/heating-bg.svg')no-repeat;
	}
}
`;

const TopContent = styled.div`
	margin-bottom:0.75rem;
	.frame-icon{
		> button {
			> i{
				font-size:40px;
			}
		}
	}
`;

const HeatingComponent = (props) => {
	return (
		<>
			<PageBody>
				<HeatingBox>
					<div className="process-content process-heating px-2">
					<TopContent className="d-flex justify-content-between align-items-center mx-5">
						<div className="d-flex flex-column">
							<div className="d-flex align-items-center frame-icon">
								<ButtonIcon
									size={60}
									name='heating'
									className='text-primary bg-white border-gray'
								/>
								<TopHeading titleHeading="Heating"/>
							</div>
						</div>
					</TopContent>
					<Card>
						<CardBody className="p-0 overflow-hidden">
							<HeatingProcess />
						</CardBody>
					</Card>
					<ButtonBar />
					</div>
				</HeatingBox>
				<AppFooter />
			</PageBody>
		</>
	);
};

HeatingComponent.propTypes = {};

export default HeatingComponent;
