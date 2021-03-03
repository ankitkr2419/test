import React from 'react';

import styled from 'styled-components';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import './searchBox.scss';

const StyledSearchBox = styled(Link)`
  display: flex;
  width: ${props => (props.size === 'sm' ? '58px' : '150px')};
  height: 48px;
  align-items: center;
  justify-content: center;
  margin: ${props => (props.size === 'sm' ? '0 12px' : '')};
`;

const SearchBox = (props) => {
	return (
		<div className="d-flex justify-content-start align-items-center search-box-container m-3">
			<input type="search" placeholder="Search"/>
		</div>
	);
};

SearchBox.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

SearchBox.defaultProps = {
	isUserLoggedIn: false,
};

export default SearchBox;
