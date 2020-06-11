import React from 'react';
import { Table } from 'reactstrap';
import "./Table.scss";

const StyledTable = props => {
  return (
    <Table {...props} />
  );
};

StyledTable.propTypes = {};

export default StyledTable;