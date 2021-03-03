import React from 'react';

import styled from 'styled-components';
// import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';

const StyledSearchBox = styled.div`
input {
    outline: none;

    &:-moz-placeholder {
        color: rgba(102, 102, 102, 0.26);
    }

    &::-webkit-input-placeholder {
        color: rgba(102, 102, 102, 0.26);
    }
}

input[type=search] {
    //-webkit-appearance: textfield;
    //-webkit-box-sizing: content-box;
    font-family: inherit;
    font-size: 100%;
    background: #fafafa url('images/search-icon.png') no-repeat center right 12px;
    border: 1px solid #717171;
    padding: 0.5rem 0.563rem 0.5rem 1.938rem;
    width: 0px;
    color: transparent;
    cursor: pointer;
    border-radius:10rem;
    -webkit-transition: all .5s;
    -moz-transition: all .5s;
    transition: all .5s;
    &:hover{
        background-color: #fff;
    }

    &:focus {
        width: 37.25rem;
        padding-right: 2.5rem;
        color: #000;
        background-color: #fff;
        cursor: auto;
    }
}

input::-webkit-search-decoration,
input::-webkit-search-cancel-button {
    display: none;
}
`;

const SearchBox = (props) => {
	return (
		<StyledSearchBox className="d-flex justify-content-start align-items-center m-3">
			<input type="search" placeholder="Search"/>
		</StyledSearchBox>
	);
};

SearchBox.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

SearchBox.defaultProps = {
	isUserLoggedIn: false,
};

export default SearchBox;
