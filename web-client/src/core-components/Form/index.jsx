import React from "react";
import { Form } from "reactstrap";

const CustomForm = (props) => {
	return <Form {...props}>{props.children}</Form>;
};

CustomForm.propTypes = {};

export default CustomForm;
