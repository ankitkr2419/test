import React, { useState } from 'react';

import styled from 'styled-components';
import { 
	Modal, 
	ModalBody, 
	Button,
	CheckBox
} from 'core-components';
import { ImageIcon, ButtonIcon, Icon, Text} from 'shared-components';

import CirclelogoIcon from 'assets/images/mylab-logo-with-circle.png';
import OperatorLoginModal from 'components/modals/OperatorLoginModal';
import TimeModal from 'components/modals/TimeModal';

const SplashScreen = styled.div`
    background: url('/images/logo-bg.svg') left -4.875rem top -5.5rem no-repeat,
    url('/images/honey-bees-bg.svg') right -1.75rem bottom -1.5rem no-repeat;
    .circle-image{
        margin-right: 14.313rem;
        margin-left: auto;
    }
`;

const OptionBox = styled.div`
	.large-btn{
		width:16.75rem;
		height:6.25rem;
		background-color:#F5F5F5;
		margin-bottom: 2.125rem;
		border:1px solid #DBDBDB;
		border-radius:1rem;
	}
	.title-heading{
		color:
	}
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
			{/* Homing Confirmation Popup */}
      <Button color="danger" onClick={toggle}>Alert pop up2</Button>
				<Modal isOpen={modal} toggle={toggle} centered size="sm">
					<ModalBody className="p-0">
						<OptionBox className="p-5">
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
							
							<div className="d-flex justify-content-center align-items-center flex-column mt-5">
								<Text Tag="label" size="20" className="mb-4 title-heading">Homing Confirmation</Text>
									<div
										className="d-flex justify-content-center align-items-center font-weight-light border-1 border-gray shadow-none bg-gray large-btn">
											<div className="d-flex justify-content-center align-items-center flex-column">
												<Icon
													size={21}
													name="tip-pickup"
													onClick={toggle}
													className="text-primary mt-3 mb-3"
												/>
												<CheckBox
													id='tip-discard'
													name='tip-discard'
													label='Tip Discard'
													className='mb-3'
												/>
										</div>
									</div>
									{/* <Button
										color="default"
										className="font-weight-light border-1 border-gray shadow-none bg-white large-btn">
											<div className="d-flex justify-content-center align-items-center flex-column">
												<Text Tag="span">Tip Discard</Text>
												<Icon
														size={21}
														name="tip-pickup"
														onClick={toggle}
														className="text-primary mt-3"
												/>
										</div>
									</Button> */}
									<Button
										color="primary"
									>
										Yes
									</Button>
							</div>
						</OptionBox>
				</ModalBody>
			</Modal>

			{/* Operator Login Modal */}

			<OperatorLoginModal />
			<TimeModal />
    </SplashScreen>
	);
};

SplashScreenComponent.propTypes = {};

export default SplashScreenComponent;
