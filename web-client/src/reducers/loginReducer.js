import { fromJS } from "immutable";
import loginActions from "actions/loginActions";
import { DECKNAME } from "appConstants";

// const loginInitialState = fromJS({
//   isLoading: true,
//   isUserLoggedIn: false,
//   isLoginTypeAdmin: false,
//   isLoginTypeOperator: false,
//   isError: false,
//   isPlateRoute: false,
//   isTemplateRoute: false,
// });

const loginInitialState = fromJS({
  isLoading: false,
  isPlateRoute: false,
  isTemplateRoute: false,
  deckName: "",
  isAdmin: false,
  decks: [
    {
      name: DECKNAME.DeckA,
      isLoggedIn: false,
      error: "",
      msg: "",
      isAdmin: false,
      isActive: false,
    },
    {
      name: DECKNAME.DeckB,
      isLoggedIn: false,
      error: "",
      msg: "",
      isAdmin: false,
      isActive: false,
    },
  ],
});

export const loginReducer = (state = loginInitialState, action) => {
  switch (action.type) {
    case loginActions.loginInitiated:
      return state.merge({
        isLoading: true,
        isUserLoggedIn: false,
        isLoginTypeAdmin: false,
        isError: false,
      });
    case loginActions.successAction:
      return state.merge({
        isLoading: false,
        isUserLoggedIn: true,
        isLoginTypeAdmin: true,
        isError: false,
      });
    case loginActions.failureAction:
      return state.merge({
        isLoading: true,
        isUserLoggedIn: false,
        isLoginTypeAdmin: false,
        isError: true,
      });
    case loginActions.setLoginTypeAsOperator:
      return state.merge({
        isLoginTypeOperator: true,
        isUserLoggedIn: true,
        isLoginTypeAdmin: false,
      });
    case loginActions.setIsPlateRoute:
      return state.setIn(["isPlateRoute"], action.payload.isPlateRoute);
    case loginActions.setIsTemplateRoute:
      return state.setIn(["isTemplateRoute"], action.payload.isTemplateRoute);
    case loginActions.loginReset:
      return loginInitialState;
    default:
      return state;
  }
};
