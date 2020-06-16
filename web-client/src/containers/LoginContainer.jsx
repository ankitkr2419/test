import React from "react";
import { Card, CardBody } from "core-components";
import { ButtonGroup, Link } from "shared-components";

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
