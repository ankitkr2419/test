import React from 'react';

import styled from 'styled-components'
import { Button} from 'core-components';
import { Text, Center } from 'shared-components';

const AllSetScreen = styled.div`
display: flex;
min-height: 100vh;
flex-direction: column;
justify-content: space-between;
	.footer-bg{
		background-color:#F9E7D7;
		height:6.25rem;
		border-radius:2.25rem 2.25rem 0 0;
		.next-btn{
			position:absolute;
			top:-1rem;
			left: 0;
    	right: 0;
		}
	}
`;

const AllSetScreenComponent = (props) => {
	return (
		<AllSetScreen>
			<Text className="d-flex justify-content-center align-items-center flex-grow-1">Tip Control</Text>
			<Center>
				<Text className="text-primary font-weight-light">Tip All Set!</Text>
				<div className="footer-bg position-relative text-center">
					<div className="next-btn">
						<Button
						color="primary"
						>
							Next
						</Button>
					</div>
				</div>
			</Center>
		</AllSetScreen>
	);
};

AllSetScreenComponent.propTypes = {};

export default AllSetScreenComponent;
