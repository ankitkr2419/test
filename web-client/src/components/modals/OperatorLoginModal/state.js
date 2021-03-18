import produce from "immer";

export const initialState = {
    email: { value: '', invalid: false },
    password: { value: '', invalid: false },
};

export const authDataStateActions = {
	SET_EMAIL: 'SET_EMAIL',
    SET_PASSWORD: 'SET_PASSWORD',
    SET_EMAIL_INVALID: 'SET_EMAIL_INVALID',
    SET_PASSWORD_INVALID: 'SET_PASSWORD_INVALID',
};

export const reducer = (state, action) => {
    switch (action.type) {
        case authDataStateActions.SET_EMAIL:
        return produce(state, (draft) => { draft.email.value = action.payload.value; });

        case authDataStateActions.SET_PASSWORD:
        return produce(state, (draft) => { draft.password.value = action.payload.value; });

        case authDataStateActions.SET_EMAIL_INVALID:
        return produce(state, (draft) => { draft.email.invalid = action.payload.invalid; });

        case authDataStateActions.SET_PASSWORD_INVALID:
        return produce(state, (draft) => { draft.password.invalid = action.payload.invalid; });

        default: return state;
    }
};
