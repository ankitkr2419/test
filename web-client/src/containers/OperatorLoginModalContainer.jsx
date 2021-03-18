import React, { useReducer } from 'react';

import {reducer, initialState, authDataStateActions} from "components/modals/OperatorLoginModal/state"
import {EMAIL_REGEX, PASSWORD_REGEX} from "components/modals/OperatorLoginModal/constants"
import OperatorLoginModal from 'components/modals/OperatorLoginModal';

import { operatorLoginInitiated } from 'action-creators/operatorLoginModalActionCreators';

import { useDispatch, useSelector } from 'react-redux';
import { Redirect } from 'react-router';

const OperatorLoginModalContainer = (props) => {

    const {
        operatorLoginModalOpen,
        toggleOperatorLoginModal
    } = props;

    const {
        SET_EMAIL,
        SET_PASSWORD,
        SET_EMAIL_INVALID,
        SET_PASSWORD_INVALID,
    } = authDataStateActions;

    const dispatch = useDispatch();
    const operatorLoginModalReducer = useSelector((state) => state.operatorLoginModalReducer);
    const { isOperatorLoggedIn, error } = operatorLoginModalReducer.toJS();

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
        const email = authData.email.value;
        const password = authData.password.value;

        const invalidEmail = !EMAIL_REGEX.test(email);
        const invalidPassword = !PASSWORD_REGEX.test(password);
        
        setAuthData({type:SET_EMAIL_INVALID, payload:{invalid:invalidEmail}});
        setAuthData({type:SET_PASSWORD_INVALID, payload:{invalid:invalidPassword}});

        if(!invalidEmail && !invalidPassword){
            dispatch(operatorLoginInitiated({email:email, password: password, role: "admin"}));
        }
    }   

    if (isOperatorLoggedIn && !error) {
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

