import React from 'react';
import { Link } from "react-router-dom";
import { Button, Form, FormGroup, FormError, Input, Label } from "core-components";
import { Center } from "shared-components";

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
					invalid
				/>
				<FormError>Oh noes! that name is already taken</FormError>
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