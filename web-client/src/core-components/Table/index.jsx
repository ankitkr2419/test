import React from "react";
import { Table } from "reactstrap";
import styled from "styled-components";

const StyledTable = styled(Table)`
	margin: 0;
`;

const CustomTable = (props) => {
	return <StyledTable {...props} />;
};

CustomTable.propTypes = {};

export default CustomTable;
