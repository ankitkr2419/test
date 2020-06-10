import React from 'react';
import { Link } from "react-router-dom";
import styled from "styled-components";
import { CardBody, Row, Col } from "reactstrap";
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
				<Link to="/" className="btn btn-secondary mr-4">
					Admin
				</Link>
				<Link to="/" className="btn btn-secondary mr-4">Supervisor</Link>
			</ButtonGroup>
			<Card className="ml-card">
				<CardBody>
					<Row>
						<Col>
							<h1 className="card-title">Compact 32</h1>
							<Link to="/templates" className="btn btn-primary">
								Login as Operator
							</Link>
						</Col>
						<Col />
					</Row>
				</CardBody>
			</Card>
		</>
	);
};

LoginContainer.propTypes = {};

export default LoginContainer;
