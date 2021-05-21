import React from 'react';

import { 
	Row,
	Card, 
	CardBody
} from 'core-components';
import {
	ButtonBar
} from 'shared-components';

import styled from 'styled-components';
import AppFooter from 'components/AppFooter';
import Process from './Process';
const PageBody = styled.div`
background-color:#f5f5f5;
`;
const ProcessBox = styled.div`
.select-process-bg {
	padding:16px 94px;
	&::after {
		background:url('/images/process-bg.svg')no-repeat;
		background-position: top right;
		margin-top:20px;
	}
	.process-content-box{
		background:none;
		border:none;
		box-shadow:none;
	}
	.process-card{
		border-radius:9px;
		border:1px solid #CCCCCC;
		padding:29px 34px;
		height:108px;
		.process-name{
			font-size:18px;
			line-height:21px;
			color: #666666;
		}
	}
	// .frame-icon{
	// 	> button {
	// 		> i{
	// 			font-size:28px;
	// 		}
	// 	}
	// }
	.row-small-gutter {
			margin-left: -7px !important;
			margin-right: -7px !important;
	}

	.row-small-gutter > * {
			padding-left: 7px !important;
			padding-right: 7px !important;
	}
	.icon-piercing{
		font-size:28px;
	}
	.icon-tip-pickup{
		font-size:21px;
	}
	.icon-aspire-dispense{
		font-size:17px;
	}
	.icon-shaking{
		font-size:26px;
	}
	.icon-heating{
		font-size:22px;
	}
	.icon-magnet{
		font-size:17px;
	}
	.icon-tip-discard{
		font-size:24px;
	}
	.icon-delay{
		font-size:19px;
	}
	.icon-tip-position{
		font-size:24px;
	}
`;

const TopContent = styled.div`
	margin-bottom:0.75rem;
}
`;

const HeadingTitle = styled.label`
    font-size: 1.25rem;
    line-height: 1.438rem;
`;

const SelectProcessPageComponent = () => {
	return (
		<>
			<PageBody>
				<ProcessBox>
					<div className="process-content select-process-bg">
					<TopContent className="d-flex justify-content-between align-items-center my-3 py-4">
						<div className="d-flex flex-column py-1">
							<HeadingTitle Tag="h5" className="text-primary font-weight-bold mb-0">Select a process</HeadingTitle>
						</div>
					</TopContent>
					<Card className="process-content-box">
						<CardBody className="p-0">
							<Row className="row-small-gutter">
								<Process iconName="piercing" processName="Piercing" />
								<Process iconName="tip-pickup" processName="Tip Pickup" />
								<Process iconName="aspire-dispense" processName="Aspire & Dispense" />
								<Process iconName="shaking" processName="Shaking" />
								<Process iconName="heating" processName="Heating" />
								<Process iconName="magnet" processName="Magnet" />
								<Process iconName="tip-discard" processName="Tip Discard" />
								<Process iconName="delay" processName="Delay" />
								<Process iconName="tip-position" processName="Tip Position" />
								{/* <Col md={4}>
									<div className="process-card bg-white d-flex align-items-center frame-icon">
										<ButtonIcon
												size={51}
												name='piercing'
												className="border-gray text-primary"
												//onClick={toggleExportDataModal}
										/>
										<Text Tag="span" className="ml-2 process-name">
											Piercing
										</Text>
									</div>
								</Col> */}
								
							</Row>
						</CardBody>
					</Card>
					<ButtonBar />
					</div>
				</ProcessBox>
				<AppFooter />
			</PageBody>
		</>
	);
};

SelectProcessPageComponent.propTypes = {};

export default SelectProcessPageComponent;
