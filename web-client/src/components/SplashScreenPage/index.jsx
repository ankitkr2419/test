import React, { useState } from 'react';

import styled from 'styled-components'
import { Modal, ModalBody, Button} from 'core-components';
import { ImageIcon, ButtonIcon, Icon, Text } from 'shared-components';

import CirclelogoIcon from 'assets/images/mylab-logo-with-circle.png';
import thumbsUpGraphics from 'assets/images/thumbs-up-graphic.svg';

import Radio from 'core-components/Radio';

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
	}
`;
const CloseButton = styled.div`
		position:absolute;
		top:1rem;
		right:1rem;
`;

const SplashScreenComponent = (props) => {
       
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);

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
							</OptionBox>
						</div>
						<div className="w-50 border-left">
								<div className="d-flex justify-content-center align-items-center flex-column p-5">
									<CloseButton>
										<ButtonIcon
										size={34}
										name="cross"
										onClick={toggle}
										className="ml-auto border-0"
										/>
									</CloseButton>
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
								</div>
							</div>
						</div>
				</ModalBody>
			</Modal>
    </SplashScreen>
	);
};

SplashScreenComponent.propTypes = {};

export default SplashScreenComponent;
