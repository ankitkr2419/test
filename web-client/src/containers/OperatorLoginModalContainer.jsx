import React, { useReducer } from 'react';

import {reducer, initialState, authDataStateActions} from "components/modals/OperatorLoginModal/state"
import OperatorLoginModal from 'components/modals/OperatorLoginModal';

import { operatorLoginInitiated } from 'action-creators/operatorLoginModalActionCreators';
import { login } from '../action-creators/loginActionCreators'
import { useDispatch, useSelector } from 'react-redux';
import { Redirect } from 'react-router';
import { USER_ROLES } from '../appConstants'
const OperatorLoginModalContainer = (props) => {

    const {
        operatorLoginModalOpen,
        toggleOperatorLoginModal,
        deckName
    } = props;

    const {
        SET_EMAIL,
        SET_PASSWORD,
        SET_EMAIL_INVALID,
        SET_PASSWORD_INVALID,
    } = authDataStateActions;

    const dispatch = useDispatch();
    // const operatorLoginModalReducer = useSelector((state) => state.operatorLoginModalReducer);
    // const { isOperatorLoggedIn, error } = operatorLoginModalReducer.toJS();

    const loginReducer = useSelector(
        (state) => state.loginReducer
    );
    const loginReducerData = loginReducer.toJS()
    let activeDeckObj = loginReducerData && loginReducerData.decks.find(deck => deck.isActive)
    let { isLoggedIn, error }  = activeDeckObj ? activeDeckObj : {}


    const [authData, setAuthData] = useReducer(reducer, initialState);

    //change local state value of email
    const handleEmailChange = (event) => {
        const email = event.target.value;
        setAuthData({ type:SET_EMAIL, payload:{value:email} });
        setAuthData({type:SET_EMAIL_INVALID, payload:{invalid:false}});
    };
    
    //change local state value of password
    const handlePasswordChange = (event) => {
        const password = event.target.value;
        setAuthData({ type:SET_PASSWORD, payload:{value:password} });
        setAuthData({type:SET_PASSWORD_INVALID, payload:{invalid:false}});
    };
    
    //email and password validation and setting local state
    const handleLoginButtonClick = () => {
        const emailValue = authData.email.value;//emailValue example => username@role.com
        const password = authData.password.value;
        // dispatch(operatorLoginInitiated({email:email, password: password, deckName: deckName, role: "admin"}));
        let role = emailValue ? emailValue.split('@').pop().split('.')[0] : undefined 
        let email = emailValue.split('@')[0]
        
        dispatch(login({email, password, deckName, role}))
    }

    if (isLoggedIn && !error) {
        return <Redirect to="/recipe-listing"/>
    }

    return(
        <OperatorLoginModal 
            operatorLoginModalOpen={operatorLoginModalOpen}
            toggleOperatorLoginModal={toggleOperatorLoginModal}
            handleEmailChange={handleEmailChange}
            handlePasswordChange={handlePasswordChange}
            handleLoginButtonClick={handleLoginButtonClick}
            authData={authData}
        />
    )
}

OperatorLoginModalContainer.propTypes = {};

export default OperatorLoginModalContainer;

