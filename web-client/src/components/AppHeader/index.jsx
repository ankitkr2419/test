import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';
import { Logo } from 'shared-components';
import {
	Button, Nav, NavItem, NavLink,
} from 'core-components';
import { NAV_ITEMS } from './constants';

const Header = styled.header`
  position: relative;
  display: flex;
  align-items: center;
  height: 80px;
  background: white 0% 0% no-repeat padding-box;
  padding: 16px 48px;
  box-shadow: 0 4px 16px #00000029;
  z-index: 1;
`;

const AppHeader = (props) => {
	const { isUserLoggedIn } = props;

	return (
		<Header>
			<Logo isUserLoggedIn={isUserLoggedIn} />
			{isUserLoggedIn && (
				<Nav className="mx-3">
					{NAV_ITEMS.map(ele => (
						<NavItem key={ele.name}>
							<NavLink to={ele.path}>{ele.name}</NavLink>
						</NavItem>
					))}
				</Nav>
			)}
			{isUserLoggedIn && (
				<>
					<Button
						color="secondary"
						size="sm"
						className="ml-auto mr-5"
						outline
						disabled
					>
            Run
					</Button>
					{/* TODO change it to button */}
					{/* <Link to="/" className="ml-2" icon>
						<Icon size={32} name="cross" />
					</Link> */}
				</>
			)}
		</Header>
	);
};

AppHeader.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

AppHeader.defaultProps = {
	isUserLoggedIn: false,
};

export default AppHeader;
