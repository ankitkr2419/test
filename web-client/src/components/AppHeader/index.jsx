import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
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
// import { getExperimentId } from 'selectors/experimentSelector';
import {
	runExperiment,
	stopExperiment,
} from 'action-creators/runExperimentActionCreators';
import { getExperimentId } from 'selectors/experimentSelector';
import { getRunExperimentReducer } from 'selectors/runExperimentSelector';
// import PrintDataModal from './PrintDataModal';
// import ExportDataModal from './ExportDataModal';
import ConfirmationModal from 'components/modals/ConfirmationModal';
import { EXPERIMENT_STATUS } from 'appConstants';
import { NAV_ITEMS } from './constants';

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
	const {
		isUserLoggedIn,
		isPlateRoute,
		isLoginTypeAdmin,
		isLoginTypeOperator,
	} = props;

	const dispatch = useDispatch();
	const experimentId = useSelector(getExperimentId);
	const runExperimentReducer = useSelector(getRunExperimentReducer);
	const experimentStatus = runExperimentReducer.get('experimentStatus');
	const isExperimentRunning = experimentStatus === EXPERIMENT_STATUS.running;
	const isExperimentStopped = experimentStatus === EXPERIMENT_STATUS.stopped;
	const isRunFailed = experimentStatus === EXPERIMENT_STATUS.runFailed;
	const isExperimentSucceeded = experimentStatus === EXPERIMENT_STATUS.success;

	const [isExitModalVisible, setExitModalVisibility] = useState(false);
	const [userDropdownOpen, setUserDropdownOpen] = useState(false);
	const toggleUserDropdown = () => setUserDropdownOpen(prevState => !prevState);

	// useEffect(() => {
	// 	if (isExperimentRunning === true) {
	// 		// connectSocket(dispatch);
	// 	}
	// }, [isExperimentRunning, dispatch]);

	useEffect(() => {
		if (isExperimentStopped === true) {
			// disConnectSocket();
			dispatch(loginReset());
		}
	}, [isExperimentStopped, dispatch]);

	const logoutClickHandler = () => {
		dispatch(loginReset());
	};

	const startExperiment = () => {
		if (isExperimentRunning === false && isExperimentSucceeded === false) {
			dispatch(runExperiment(experimentId));
		}
	};

	const onNavLinkClickHandler = (event, pathname) => {
		if (pathname === '/plate' && isLoginTypeAdmin === true) {
			event.preventDefault();
		}
	};

	const onCrossClick = () => {
		setExitModalVisibility(true);
	};

	// Exit modal confirmation click handler
	const confirmationClickHandler = (isConfirmed) => {
		setExitModalVisibility(false);
		if (isConfirmed) {
			if (isExperimentRunning === true) {
				// user aborted experiment
				dispatch(stopExperiment(experimentId));
			} else {
				dispatch(loginReset());
			}
		}
	};

	return (
		<Header>
			<Logo isUserLoggedIn={isUserLoggedIn} />
			{isUserLoggedIn && (
				<Nav className="ml-3 mr-auto">
					{NAV_ITEMS.map(ele => (
						<NavItem key={ele.name}>
							<NavLink
								onClick={(event) => {
									onNavLinkClickHandler(event, ele.path);
								}}
								to={ele.path}
								disabled={
									(ele.path === '/plate' && isLoginTypeAdmin === true)
                  || (isPlateRoute === true && ele.path === '/templates')
								}
							>
								{ele.name}
							</NavLink>
						</NavItem>
					))}
				</Nav>
			)}
			{isUserLoggedIn && (
				<div className="d-flex align-items-center">
					{/* <PrintDataModal /> */}
					{/* <ExportDataModal /> */}
					<div className="experiment-info text-right mx-3">
						<Text
							size={12}
							className={`text-default mb-1 ${
								isExperimentRunning ? 'show' : ''
							}`}
						>
							{`Experiment started at ${runExperimentReducer.get(
								'experimentStartedTime',
							)}`}
						</Text>
						<Text
							size={12}
							className={`text-error mb-1 ${isRunFailed ? 'show' : ''}`}
						>
              Experiment failed to run.
						</Text>
						{isExperimentSucceeded === false && isPlateRoute === true && (
							<Button
								color={isExperimentRunning ? 'primary' : 'secondary'}
								size="sm"
								className="font-weight-light border-2 border-gray shadow-none"
								outline={
									isExperimentRunning === false
                  && isExperimentSucceeded === false
								}
								onClick={startExperiment}
							>
                Run
							</Button>
						)}
						{isExperimentSucceeded === true && (
							<Button
								color="success"
								size="sm"
								className="font-weight-light border-2 border-gray shadow-none"
							>
                Result - Successful
							</Button>
						)}
					</div>
					{isLoginTypeAdmin === true && (
						<Dropdown
							isOpen={userDropdownOpen}
							toggle={toggleUserDropdown}
							className="ml-2"
						>
							<DropdownToggle icon name="user" size={32} />
							<DropdownMenu right>
								<DropdownItem onClick={logoutClickHandler}>
                  Log out
								</DropdownItem>
							</DropdownMenu>
						</Dropdown>
					)}
					{isLoginTypeOperator === true && (
						<ButtonIcon
							size={34}
							name="cross"
							onClick={onCrossClick}
							className="ml-2"
						/>
					)}
					{isExitModalVisible === true && (
						<ConfirmationModal
							isOpen={isExitModalVisible}
							confirmationClickHandler={confirmationClickHandler}
						/>
					)}
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
