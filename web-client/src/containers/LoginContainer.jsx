import React from 'react';
import MlButton from 'components/shared/Button/Button';
import styled from 'styled-components';
import MlCard from 'components/shared/Card/Card';
import { CardBody, Row, Col } from 'reactstrap';

const MlButtonGroup = styled.div`
	margin: 0 0 40px 0;
	text-align: right;
	padding: 0 24px 0 24px;
`;

const LoginContainer = (props) => {
  return (
		<>
			<MlButtonGroup>
				<MlButton className="mr-4">Admin</MlButton>
				<MlButton className="mr-4">Supervisor</MlButton>
			</MlButtonGroup>
			<MlCard className="ml-card">
				<CardBody>
					<Row>
						<Col>
							<h1 className="card-title">Compact 32</h1>
							<MlButton color="primary">Login as Operator</MlButton>
						</Col>
						<Col/>
					</Row>
				</CardBody>
			</MlCard>
		</>
	);
};

LoginContainer.propTypes = {};

export default LoginContainer;
