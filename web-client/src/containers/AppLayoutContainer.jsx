import React, { useEffect } from 'react';
import { HashRouter as Router, Switch, Redirect } from 'react-router-dom';
import RouteWithSubRoutes from 'RouteHelper';
import { useSelector, useDispatch } from 'react-redux';
import AppHeader from 'components/AppHeader';

import '../assets/scss/default.scss';
import { fetchActiveWells } from 'action-creators/activeWellActionCreators';
import { getActiveLoadedWellFlag } from 'selectors/activeWellSelector';
import { connectSocket } from 'web-socket';
import ModalContainer from './ModalContainer';

/**
 * AppLayoutContainer Will contain routes(content), headers, sub-headers, notification etc.
 * @param {*} props
 */
const AppLayoutContainer = (props) => {
	const { routes } = props;
	const dispatch = useDispatch();
	const loginReducer = useSelector(state => state.loginReducer);
	const isActiveWellDataLoaded = useSelector(getActiveLoadedWellFlag);
	const socketReducer = useSelector(state => state.socketReducer);
	const isOpened = socketReducer.get('isOpened');

	// connect to websocket on mount
	useEffect(() => {
		// if connection is already opened
		if (isOpened === false) {
			connectSocket(dispatch);
		}
	}, [isOpened, dispatch]);

	useEffect(() => {
		if (loginReducer.get('isLoginTypeOperator') === true
		&& isActiveWellDataLoaded === false
		) {
			dispatch(fetchActiveWells());
		}
	}, [loginReducer, isActiveWellDataLoaded, dispatch]);

	return (
		<Router>
			<AppHeader
				isPlateRoute={loginReducer.get('isPlateRoute')}
				isUserLoggedIn={loginReducer.get('isUserLoggedIn')}
				isLoginTypeAdmin={loginReducer.get('isLoginTypeAdmin')}
				isLoginTypeOperator={loginReducer.get('isLoginTypeOperator')}
			/>
			{/* Modal container will helps in displaying error/info/warning through modal */}
			<ModalContainer />
			<section className="ml-content">
				<Switch>
					<Redirect exact from="/" to="/login" />
					{routes.map(route => (
						<RouteWithSubRoutes key={route.key} {...route} />
					))}
				</Switch>
			</section>
		</Router>
	);
};

AppLayoutContainer.propTypes = {};

export default AppLayoutContainer;
