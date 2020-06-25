import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import styled from 'styled-components';
import PropTypes from 'prop-types';
import { Logo, ButtonIcon } from 'shared-components';
import { loginReset } from 'action-creators/loginActionCreators';
import {
	Button,
	Dropdown,
	DropdownToggle,
	DropdownMenu,
	DropdownItem,
	Nav,
	NavItem,
	NavLink,
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
	const dispatch = useDispatch();
	const [userDropdownOpen, setUserDropdownOpen] = useState(false);
	const toggleUserDropdown = () =>
		setUserDropdownOpen((prevState) => !prevState);

	const logoutClickHandler = () => {
		dispatch(loginReset());
	};

	return (
		<Header>
			<Logo isUserLoggedIn={isUserLoggedIn} />
			{isUserLoggedIn && (
				<Nav className='mx-3'>
					{NAV_ITEMS.map((ele) => (
						<NavItem key={ele.name}>
							<NavLink to={ele.path}>{ele.name}</NavLink>
						</NavItem>
					))}
				</Nav>
			)}
			{isUserLoggedIn && (
				<>
					<Button
						color='secondary'
						size='sm'
						className='ml-auto mr-5'
						outline
						disabled
					>
						Run
					</Button>
					<Dropdown isOpen={userDropdownOpen} toggle={toggleUserDropdown}>
						<DropdownToggle icon name='user' size={32} />
						<DropdownMenu right>
							<DropdownItem onClick={logoutClickHandler}>Log out</DropdownItem>
						</DropdownMenu>
					</Dropdown>
					<ButtonIcon
						size={34}
						name='cross'
						className='ml-2'
						onClick={logoutClickHandler}
					/>
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
