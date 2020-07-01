import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import styled from 'styled-components';
import PropTypes from 'prop-types';
import { Logo, ButtonIcon, Text } from 'shared-components';
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
import PrintDataModal from './PrintDataModal';
import ExportDataModal from './ExportDataModal';

const Header = styled.header`
	position: relative;
	display: flex;
	align-items: center;
	height: 80px;
	background: white 0% 0% no-repeat padding-box;
	padding: 16px 24px 16px 48px;
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
				<Nav className='ml-3 mr-auto'>
					{NAV_ITEMS.map((ele) => (
						<NavItem key={ele.name}>
							<NavLink to={ele.path}>{ele.name}</NavLink>
						</NavItem>
					))}
				</Nav>
			)}
			{isUserLoggedIn && (
				<div className='d-flex align-items-center'>
					<PrintDataModal />
					<ExportDataModal />
					<div className='experiment-info text-right mx-3'>
						{/* TODO: Add "show" class to <Text> component when experiment starts and remove it when experiment ends */}
						<Text size={12} className='text-default mb-1'>
							Experiment started at 12:39 PM
						</Text>
						{/* TODO: When user clicks on Run button remove outline, disabled props and change value of color prop to "primary" */}
						<Button
							color='secondary'
							size='sm'
							className='font-weight-light border-2 border-gray shadow-none'
							outline
							disabled
						>
							Run
						</Button>
						{/* TODO: Show this button after experiment ends, depending on result change value of color prop to "success" or "failure"  */}
						{/* <Button
							color='success'
							size='sm'
							className='font-weight-light border-2 border-gray shadow-none'
						>
							Result - Successful
						</Button> */}
					</div>
					<Dropdown
						isOpen={userDropdownOpen}
						toggle={toggleUserDropdown}
						className='ml-2'
					>
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
				</div>
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
