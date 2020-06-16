import React from 'react';
import { Form } from "core-components";
import { FormGroup } from "core-components";
import { Label } from "core-components";
import { Input } from "core-components";
import { Button } from "core-components";
import { Link } from "react-router-dom";
import Center from "shared-components/Center";

const LoginForm = () => {
  return (
		<Form className="pt-3 pb-2">
			<FormGroup>
				<Label for="username">Username</Label>
				<Input
					type="text"
					name="username"
					id="username"
					placeholder="Type here"
				/>
			</FormGroup>
			<FormGroup>
				<Label for="password">Password</Label>
				<Input
					type="password"
					name="password"
					id="password"
					placeholder="Type here"
				/>
			</FormGroup>
			<Center className="py-4 mb-3">
				<Button color="primary" className="mx-auto">
					Login
				</Button>
			</Center>
			<Center>
				<Link to="/">Forgot username or password?</Link>
			</Center>
		</Form>
	);
};

export default LoginForm;