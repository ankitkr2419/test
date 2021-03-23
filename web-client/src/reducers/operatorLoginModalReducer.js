import { fromJS } from 'immutable';

import operatorLoginModalActions from 'actions/operatorLoginModalActions';

const operatorLoginModalInitialState = fromJS({
    isOperatorLoggedIn: false,
    error: false,
    msg: '',
    deckName: ''
});

export const operatorLoginModalReducer = (state=operatorLoginModalInitialState, action) => {

    switch (action.type) {
        case operatorLoginModalActions.loginInitiated: 
            return state.merge({ deckName: action.payload.deckName });

        case operatorLoginModalActions.successAction:
            return state.merge({ isOperatorLoggedIn: true, error: false, msg: action.payload.response.msg });

        case operatorLoginModalActions.failureAction:
            return state.merge({ isOperatorLoggedIn: false, error: true, msg: action.payload.response.msg });
            
        default: return state;
    }
}


