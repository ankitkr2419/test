import React, { useState } from 'react';

import styled from 'styled-components';
import { 
	Modal, 
	ModalBody, 
	Button
} from 'core-components';
import { ImageIcon, ButtonIcon, Icon, Text} from 'shared-components';

import CirclelogoIcon from 'assets/images/mylab-logo-with-circle.png';
import thumbsUpGraphics from 'assets/images/thumbs-up-graphic.svg';

import Radio from 'core-components/Radio';
import OperatorLoginModal from 'components/modals/OperatorLoginModal';

const SplashScreen = styled.div`
    background: url('/images/logo-bg.svg') left -4.875rem top -5.5rem no-repeat,
    url('/images/honey-bees-bg.svg') right -1.75rem bottom -1.5rem no-repeat;
    .circle-image{
        margin-right: 14.313rem;
        margin-left: auto;
    }
`;

const OptionBox = styled.div`
	background-color:#F5F5F5;
	border-radius:36px 0 0 36px;
	.large-btn{
		width:15.125rem;
		height:8.5rem;
		margin-bottom: 2.125rem;
		border:1px solid #DBDBDB;
	}
	.tick-icon-box{
		position:absolute;
		left:50%;
		transform:translateX(-50%);
		top:-1rem;
		> button{
			width:2.625rem;
			height:2.625rem;
		}
	}
`;
const ThumbBox = styled.div`
	background-color: #F5F5F5;
    border-radius: 0 2.25rem 2.25rem 0;
`;
//For Operator Login Form
const OperatorLoginForm = styled.div`
.has-border-left{
	.form-control{
		border-left:7px solid #F38220;
	}
}
.link{
	color:#3C70FF;
}
`;

const SplashScreenComponent = (props) => {
       
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);

	// Operator Login Modal
	const [operatorLoginModal, setOperatorLoginModal] = useState(false);
	const toggleOperatorLoginModal = () => setOperatorLoginModal(!operatorLoginModal);

  return (
		<SplashScreen className='splash-screen-content h-100 py-0 bg-white d-flex justify-content-center align-items-center'>
			<div className="circle-image">
					<ImageIcon 
					src={CirclelogoIcon} 
					alt="My Lab" 
					/>
			</div>
			{/* Alert pop up2 */}
      <Button color="danger" onClick={toggle}>Alert pop up2</Button>
			<Modal isOpen={modal} toggle={toggle} centered size="lg">
				<ModalBody className="p-0">
					<div className="d-flex w-100">
						<div className="w-50">
							<OptionBox className="option-box p-5 h-100 d-flex flex-column">
									<Radio
											id='radio1'
											name='radio1'
											label='I see a problem with the position of the tip and/or magnet!'
											className='mb-3'
									/>
									<div className="d-flex justify-content-center align-items-center flex-column mt-5">
									<div className="position-relative">
											<div className="tick-icon-box">
												<ButtonIcon size={16} name='tick' className="border-success font-weight-bold text-success bg-white"/>
											</div>
										<Button
											color="default"
											className="font-weight-light border-1 border-gray shadow-none bg-white large-btn">
												<div className="d-flex justify-content-center align-items-center flex-column">
													<Text Tag="span">Fix Tip Control</Text>
													<Icon
															size={21}
															name="tip-pickup"
															onClick={toggle}
															className="text-primary mt-3"
													/>
											</div>
										</Button>
										</div>
										<div className="position-relative">
											<div className="tick-icon-box">
												<ButtonIcon 
													size={16} 
													name='tick' 
													className="border-success font-weight-bold text-success bg-white"/>
											</div>
											<Button
												color="default"
												className="font-weight-light border-1 border-gray shadow-none bg-white large-btn">
													<div className="d-flex justify-content-center align-items-center flex-column">
														<Text Tag="span">Fix Magnet Control</Text>
														<Icon
																size={21}
																name="magnet"
																onClick={toggle}
																className="text-primary mt-3"
														/>
													</div>
											</Button>
										</div>
									</div>
							</OptionBox>
						</div>
						<div className="w-50 border-left">
								<ThumbBox className="d-flex justify-content-center align-items-center flex-column p-5">
									<ButtonIcon
										position="absolute"
										placement="right"
										top={16}
										right={16}
										size={36}
										name="cross"
										onClick={toggle}
										className="ml-auto border-0"
										/>
									<Radio
									id='radio2'
									name='radio1'
									label='I declare the position of tip and magnet is Okay'
									className='mb-5'
									/>
									<ImageIcon
									src={thumbsUpGraphics}
									alt="No templates available"
									className="img-video-thumbnail"
									/>
									<Button
										color="primary"
									>
										Next
									</Button>
								</ThumbBox>
							</div>
						</div>
				</ModalBody>
			</Modal>

			{/* Operator Login Modal */}

		<OperatorLoginModal />
    </SplashScreen>
	);
};

SplashScreenComponent.propTypes = {};

export default SplashScreenComponent;
