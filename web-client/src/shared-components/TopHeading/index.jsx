import React from 'react';

import styled from 'styled-components';
// import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';

const HeadingTitle = styled.label`
    font-size: 1.25rem;
    line-height: 1.438rem;
`;

const TopHeading = ({titleHeading}) => {
	return (
		<HeadingTitle Tag="h5" className="text-primary font-weight-bold ml-4 mb-0">{titleHeading}</HeadingTitle>
	);
};

TopHeading.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

TopHeading.defaultProps = {
	isUserLoggedIn: false,
};

export default TopHeading;
