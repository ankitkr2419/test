import React from 'react';
import styled from 'styled-components';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import imgLogo from 'assets/images/mylab-logo.png';
import imgSymbol from 'assets/images/mylab-symbol.png';

const StyledLogo = styled(Link)`
  display: flex;
  width: ${props => (props.size === 'sm' ? '58px' : '150px')};
  height: 48px;
  align-items: center;
  justify-content: center;
  margin: ${props => (props.size === 'sm' ? '0 12px' : '')};
`;

const Logo = (props) => {
	const { isUserLoggedIn } = props;
	// if user logged-in will show small icon
	const size = isUserLoggedIn ? 'sm' : 'lg';

	return (
		<StyledLogo size={size} to={isUserLoggedIn === true ? '/login' : ''} className="logo">
			{isUserLoggedIn === true ? (
				<img src={imgSymbol} alt="My Lab logo" className="h-100" />
			) : (
				<img src={imgLogo} alt="My Lab logo" className="w-100" />
			)}
		</StyledLogo>
	);
};

Logo.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

Logo.defaultProps = {
	isUserLoggedIn: false,
};

export default Logo;
