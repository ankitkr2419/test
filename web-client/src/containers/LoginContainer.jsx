import React from 'react';
import { CardBody, Row, Col } from "reactstrap";
import Card from 'core-components/Card';
import ButtonGroup from 'shared-components/ButtonGroup';
import Link from 'shared-components/Link';

const LoginContainer = (props) => {
  return (
		<>
			<ButtonGroup>
				<Link to="/" className="btn-secondary mr-4">
					Admin
				</Link>
				<Link to="/" className="btn-secondary mr-4">
					Supervisor
				</Link>
			</ButtonGroup>
			<Card className="ml-card">
				<CardBody>
					<Row>
						<Col>
							<h1 className="card-title">Compact 32</h1>
							<Link to="/templates" className="btn-primary">
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
