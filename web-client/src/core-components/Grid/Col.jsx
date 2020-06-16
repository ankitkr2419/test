import React from "react";
import { Col } from "reactstrap";

const CustomCol = (props) => {
	return <Col {...props}>{props.children}</Col>;
};

CustomCol.propTypes = {};

export default CustomCol;
