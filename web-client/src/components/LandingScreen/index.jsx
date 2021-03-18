import React, { useState } from 'react';
import { useDispatch } from "react-redux";
import {
	ImageIcon,
	Center,
	Icon,
	ButtonIcon,
	Text
} from 'shared-components';
import Radio from 'core-components/Radio';
import ConfirmationModal from 'components/modals/ConfirmationModal';
import thumbsUpGraphics from 'assets/images/thumbs-up-graphic.svg';

// import SearchBox from 'shared-components/SearchBox';
// import ButtonBar from 'shared-components/ButtonBar';
import imgNoTemplate from 'assets/images/video-thumbnail-poster.jpg';
import styled from 'styled-components';
import AppFooter from 'components/AppFooter';
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants"
import { Modal, ModalBody, Button, CardBody, Card } from 'core-components';
import { homingActionInitiated } from "action-creators/homingActionCreators";

const VideoPlayButton = styled.button`
	color:#7C7976;
	background-color:transparent;
	border:0;
	outline:none;
	position:absolute;
	top:50%;
	left:50%;
	transform:translate(-50%,-50%);
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
const LandingScreenComponent = (props) => {
	const [ homingStatus, setHomingStatus] = useState(true);
	const [modal, setModal] = useState(false);
	const dispatch = useDispatch();
	const toggle = () => setModal(!modal);
	const homingConfirmation = (isConfirmed) => {
		setHomingStatus(isConfirmed)
		if(isConfirmed){
			setModal(true)
			dispatch(homingActionInitiated())
		}
	}
	return (
		<div className="ml-content">
			<div className='landing-content'>
			<Card className='card-video'>
				<CardBody className='d-flex flex-column p-0'>
					<Center className="video-thumbnail-wrapper">
					<ImageIcon
						src={imgNoTemplate}
						alt="No templates available"
						className="img-video-thumbnail"
					/>
					<VideoPlayButton>
						<Icon name="play" size={124}/>
					</VideoPlayButton>
					</Center>
				</CardBody>
			</Card>
			<ConfirmationModal
				isOpen={homingStatus}
				toggleModal={homingConfirmation}
				confirmationText={MODAL_MESSAGE.homingConfirmation}
				confirmationClickHandler={homingConfirmation}
				successBtn={MODAL_BTN.okay}
				cancelBtn={MODAL_BTN.cancel}
			/>
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

				{/* <SearchBox/> */}

				{/* <ButtonBar/> */}
			</div>
			<AppFooter />
		</div>
	);
};

LandingScreenComponent.propTypes = {};

export default LandingScreenComponent;
