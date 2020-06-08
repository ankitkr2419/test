import React from 'react';
import styled from "styled-components";
import { CardBody, Row, Col } from "reactstrap";
import Button from 'core-components/Button';
import Card from 'core-components/Card';

const ButtonGroup = styled.div`
	margin: 0 0 40px 0;
	text-align: right;
	padding: 0 24px 0 24px;
`;

const LoginContainer = (props) => {
  return (
		<>
			<ButtonGroup>
				<Button className="mr-4">Admin</Button>
				<Button className="mr-4">Supervisor</Button>
			</ButtonGroup>
			<Card className="ml-card">
				<CardBody>
					<Row>
						<Col>
							<h1 className="card-title">Compact 32</h1>
							<Button color="primary">Login as Operator</Button>
						</Col>
						<Col/>
					</Row>
				</CardBody>
			</Card>
		</>
	);
};

LoginContainer.propTypes = {};

export default LoginContainer;
