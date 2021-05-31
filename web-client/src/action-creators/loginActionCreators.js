import loginActions, {
  deckBlockActions,
  logoutActions,
} from "actions/loginActions";

export const login = (body) => ({
  type: loginActions.loginInitiated,
  payload: {
    body,
  },
});

export const loginFailed = (errorResponse) => ({
  type: loginActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const setActiveDeck = (deckName) => ({
  type: loginActions.setActiveDeck,
  payload: {
    deckName,
  },
});

export const loginReset = (deckName) => ({
  type: loginActions.loginReset,
  payload: {
    deckName,
  },
});

export const loginAsOperator = () => ({
  type: loginActions.setLoginTypeAsOperator,
});

export const setIsPlateRoute = (isPlateRoute) => ({
  type: loginActions.setIsPlateRoute,
  payload: {
    isPlateRoute,
  },
});

export const setIsTemplateRoute = (isTemplateRoute) => ({
  type: loginActions.setIsTemplateRoute,
  payload: {
    isTemplateRoute,
  },
});

//log out action-creators
export const logoutInitiated = (params) => ({
  type: logoutActions.logoutActionInitiated,
  payload: params,
});

export const logoutSuccess = (response) => ({
  type: logoutActions.logoutActionSuccess,
  payload: response,
});

export const logoutFailure = (error) => ({
  type: logoutActions.logoutActionFailure,
  payload: error,
});

//log out action-creators
export const deckBlockInitiated = () => ({
  type: deckBlockActions.deckBlockInitiated,
});

export const deckBlockReset = () => ({
  type: deckBlockActions.deckBlockReset,
});
