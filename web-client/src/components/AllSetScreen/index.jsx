import React from 'react';
import { Button} from 'core-components';
import { Text, Center } from 'shared-components';
import { AllSetScreen } from './AllSetScreen';

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
