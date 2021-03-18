import operatorLoginModalActions  from "actions/operatorLoginModalActions";

export const operatorLoginInitiated = (authData) => ({
    type: operatorLoginModalActions.loginInitiated,
    payload: authData,
});

export const operatorLoginSuccess = (successMsg) => ({
    type: operatorLoginModalActions.successAction,
    payload: { msg: successMsg }
});

export const operatorLoginFailed = (errorMsg) => ({
    type: operatorLoginModalActions.successAction,
    payload: { msg: errorMsg }
});