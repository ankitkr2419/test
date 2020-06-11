import React from 'react';
import { CardBody } from "reactstrap";
import Card from 'core-components/Card';
import ButtonGroup from 'shared-components/ButtonGroup';
import Link from 'shared-components/Link';

const LoginContainer = (props) => {
  return (
		<div className="login-content">
			<ButtonGroup>
				<Link to="/" className="btn-secondary mr-4">
					Admin
				</Link>
				<Link to="/" className="btn-secondary mr-4">
					Supervisor
				</Link>
			</ButtonGroup>
			<Card className="card-login">
				<CardBody className="d-flex scroll-y">
					<div className="flex-100">
						<h1 className="card-title">Compact 32</h1>
						<Link to="/templates" className="btn-primary">
							Login as Operator
						</Link>
					</div>
					<div className="flex-100" />
				</CardBody>
			</Card>
		</div>
	);
};

LoginContainer.propTypes = {};

export default LoginContainer;
