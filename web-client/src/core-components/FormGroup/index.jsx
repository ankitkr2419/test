import React from "react";
import { FormGroup } from "reactstrap";

const CustomFormGroup = (props) => {
	return <FormGroup {...props}>{props.children}</FormGroup>;
};

CustomFormGroup.propTypes = {};

export default CustomFormGroup;
