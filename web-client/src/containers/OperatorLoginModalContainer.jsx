import React, {useReducer, useCallback} from 'react';

import {reducer, initialState, authDataStateActions} from "components/modals/OperatorLoginModal/state"
import {EMAIL_REGEX, PASSWORD_REGEX} from "components/modals/OperatorLoginModal/constants"

import OperatorLoginModal from 'components/modals/OperatorLoginModal';

const OperatorLoginModalContainer = (props) => {

    const {
        SET_EMAIL,
        SET_PASSWORD,
        SET_EMAIL_INVALID,
        SET_PASSWORD_INVALID,
    } = authDataStateActions;

    const {
        operatorLoginModalOpen,
        toggleOperatorLoginModal
    } = props;
    
    const [authData, setAuthData] = useReducer(reducer, initialState);

    //change local state value of email
    const handleEmailChange = useCallback((event) => {
        const email = event.target.value;
        setAuthData({ type:SET_EMAIL, payload:{value:email} });
    }, [SET_EMAIL]);
    
    //change local state value of password
    const handlePasswordChange = useCallback((event) => {
        const password = event.target.value;
        setAuthData({ type:SET_PASSWORD, payload:{value:password} });
    }, [SET_PASSWORD]);
    
    //email and password validation and setting local state
    const handleLoginButtonClick = useCallback(() => {
        const email = authData.email.value;
        const password = authData.password.value;

        const invalidEmail = !EMAIL_REGEX.test(email);
        const invalidPassword = !PASSWORD_REGEX.test(password);
        
        setAuthData({type:SET_EMAIL_INVALID, payload:{invalid:invalidEmail}});
        setAuthData({type:SET_PASSWORD_INVALID, payload:{invalid:invalidPassword}});
    }, [authData, SET_EMAIL_INVALID, SET_PASSWORD_INVALID]);

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

