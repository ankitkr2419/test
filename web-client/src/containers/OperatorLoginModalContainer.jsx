import React, {useReducer, useCallback} from 'react';
// import validate from 'yup';
import produce from 'immer';

import OperatorLoginModal from 'components/modals/OperatorLoginModal';

const OperatorLoginModalContainer = (props) => {

    const {
        operatorLoginModalOpen,
        toggleOperatorLoginModal
    } = props;

    const reducer = (state, action) => {
        switch (action.type) {
            case 'email':
            return produce(state, (draft) => {
                draft.email.value = action.payload.value;
                draft.email.state = action.payload.state;
            });

            case 'password':
            return produce(state, (draft) => {
                draft.password.value = action.payload.value;
                draft.password.state = action.payload.state;
            });

            case 'emailInvalid':
            return produce(state, (draft) => { draft.email.state = action.payload; });

            case 'passwordInvalid':
            return produce(state, (draft) => { draft.password.state = action.payload; });

            default: return state;
        }
    };

    const initialState = {
        email: { value: '', state: { valid: true, message: '' } },
        password: { value: '', state: { valid: true, message: '' } },
      };
    
    const [authData, setAuthData] = useReducer(reducer, initialState);

    const handleEmailChange = useCallback((event) => {
        const email = event.target.value;
        setAuthData({ type:"email", payload:{value:email} });
    }, []);
    
    const handlePasswordChange = useCallback((event) => {
        const password = event.target.value;
        setAuthData({ type:"password", payload:{value:password} });
    }, []);
    
    const handleLoginButtonClick = useCallback((event) => {
        console.log(event.target.value);
    }, []);

    console.log(authData);

    return(
        <OperatorLoginModal 
            operatorLoginModalOpen={operatorLoginModalOpen}
            toggleOperatorLoginModal={toggleOperatorLoginModal}
            handleEmailChange={handleEmailChange}
            handlePasswordChange={handlePasswordChange}
            handleLoginButtonClick={handleLoginButtonClick}
        />
    )
}

OperatorLoginModalContainer.propTypes = {};

export default OperatorLoginModalContainer;

