import React, { useState } from 'react';

import styled from 'styled-components'
import { Modal, ModalBody, Button, Row, Col } from 'core-components';
import { ImageIcon, ButtonIcon, Icon } from 'shared-components';

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
            <Button color="danger" onClick={toggle}>Show Modal</Button>
            <Modal isOpen={modal} toggle={toggle} centered size="lg">
                <ModalBody className="p-0">
                    <Row>
                        <Col>
                            <div className="option-box p-5">
                                <Radio
                                    id='radio1'
                                    name='radio1'
                                    label='I see a problem with the position of the tip and/or magnet!'
                                    className='mb-3 mr-4'
                                />
                                <Button
								color="default"
								size="sm"
								className="font-weight-light border-2 border-gray shadow-none">
                                <Icon
                                    size={34}
                                    name="tip-pickup"
                                    onClick={toggle}
                                    className="ml-auto"
                                />Fix Tip Control
							</Button>
                                 
                            <Button
								color="default"
								size="sm"
								className="font-weight-light border-2 border-gray shadow-none">
                                <Icon
                                    size={34}
                                    name="magnet"
                                    onClick={toggle}
                                    className="ml-auto"
                                />Fix Magnet Control
							</Button>
                            </div>
                        </Col>
                        <Col className="border-left">
                            <div className="d-flex justify-content-center align-items-center flex-column px-3 py-3">
                                <ButtonIcon
                                size={34}
                                name="cross"
                                onClick={toggle}
                                className="ml-auto border-0"
                            />
                                <Radio
                                id='radio2'
                                name='radio1'
                                label='I declare the position of tip and magnet is Okay'
                                className='mb-3 mr-4'
                                />
                                <ImageIcon
                                src={thumbsUpGraphics}
                                alt="No templates available"
                                className="img-video-thumbnail"
                                />
                                <Button
								color="primary"
								className="font-weight-light border-2 border-gray shadow-none"
							>
                                 Next
							</Button>
                            </div>
                        </Col>
                    </Row>
                </ModalBody>
            </Modal>
        </SplashScreen>
	);
};

SplashScreenComponent.propTypes = {};

export default SplashScreenComponent;
