import React from "react";
import { Row } from "reactstrap";

const CustomRow = (props) => {
	return <Row {...props}>{props.children}</Row>;
};

CustomRow.propTypes = {};

export default CustomRow;
