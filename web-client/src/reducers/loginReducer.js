import { fromJS } from "immutable";
import loginActions from "actions/loginActions";
import { DECKNAME, USER_ROLES } from "../appConstants";

// const loginInitialState = fromJS({
//   isLoading: true,
//   isUserLoggedIn: false,
//   isLoginTypeAdmin: false,
//   isLoginTypeOperator: false,
//   isError: false,
//   isPlateRoute: false,
//   isTemplateRoute: false,
// });

// const loginInitialState = fromJS({
// 	isLoading: true,
// 	isUserLoggedIn: false,
// 	isLoginTypeAdmin: false,
// 	isLoginTypeOperator: false,
// 	isError: false,
// 	isPlateRoute: false,
// 	isTemplateRoute: false,
// });
const initialStateOfDecks = () => {
    return [
        {
            name: DECKNAME.DeckA,
            isLoggedIn: false,
            error: false,
            msg: "",
            isAdmin: false,
            isActive: true,
            token: "",
        },
        {
            name: DECKNAME.DeckB,
            isLoggedIn: false,
            error: false,
            msg: "",
            isAdmin: false,
            isActive: false,
            token: "",
        },
    ];
};

const loginInitialState = fromJS({
    isLoading: false,
    isPlateRoute: false,
    isTemplateRoute: false,
    deckName: "", //initiated deckName while login process only (temp)
    isAdmin: false, //initiated isAdmin while login process only (temp)
    decks: initialStateOfDecks(),
});

export const loginReducer = (state = loginInitialState, action) => {
    switch (action.type) {
        case loginActions.loginInitiated:
            let updatedDecks = state.toJS().decks.map((deckObj) => {
                return deckObj.name === action.payload.body.deckName
                    ? {
                          ...deckObj,
                          error: false,
                          msg: "",
                          isError: false,
                          isActive: true,
                      }
                    : {
                          ...deckObj,
                          isActive: false,
                      };
            });

            return state.merge({
                isLoading: true,
                deckName: action.payload.body.deckName,
                isAdmin: action.payload.body.role === USER_ROLES.ADMIN,
                decks: updatedDecks,
            });
        case loginActions.successAction:
            const token = action.payload.response.token;
            let deckName = state.toJS().deckName;
            let isAdminTemp = state.toJS().isAdmin;
            if (deckName && deckName === DECKNAME.DeckA) {
                //update and login deck A
                let newDecks = state.toJS().decks.map((deckObj) => {
                    let newDeckObj =
                        deckObj.name === DECKNAME.DeckA
                            ? {
                                  ...deckObj,
                                  isLoggedIn: true,
                                  error: false,
                                  msg: "",
                                  isAdmin: isAdminTemp,
                                  isActive: true,
                                  token
                              }
                            : {
                                  ...deckObj,
                                  isActive: false,
                              };
                    return newDeckObj;
                });

                return state.merge({
                    isLoading: false,
                    decks: newDecks,
                });
            } else if (deckName && deckName === DECKNAME.DeckB) {
                //update and login deck B
                let newDecks = state.toJS().decks.map((deckObj) => {
                    let newDeckObj =
                        deckObj.name === DECKNAME.DeckB
                            ? {
                                  ...deckObj,
                                  isLoggedIn: true,
                                  error: false,
                                  msg: "",
                                  isAdmin: isAdminTemp,
                                  isActive: true,
                                  token
                              }
                            : {
                                  ...deckObj,
                                  isActive: false,
                              };
                    return newDeckObj;
                });

                return state.merge({
                    isLoading: false,
                    decks: newDecks,
                });
            } else {
                //if deck name dont match, then dont update state
                return state;
            }
        case loginActions.failureAction:
            // console.log("action: ", action);
            let err =
                action.payload.serverErrors && action.payload.serverErrors.msg
                    ? action.payload.serverErrors.msg
                    : "Something went wrong!";
            let newDecks = state.toJS().decks.map((deckObj) => {
                return deckObj.name === state.toJS().deckName
                    ? {
                          ...deckObj,
                          error: true,
                          msg: err,
                          token: "",
                      }
                    : deckObj;
            });

            return state.merge({
                isLoading: false,
                decks: newDecks,
            });

        case loginActions.setActiveDeck:
            let deckNameToSetActive = action.payload.deckName;
            let newDecksForActive = state.toJS().decks.map((deckObj) => {
                return deckObj.name === deckNameToSetActive
                    ? {
                          ...deckObj,
                          isActive: true,
                      }
                    : {
                          ...deckObj,
                          isActive: false,
                      };
            });
            return state.merge({
                decks: newDecksForActive,
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
            return state.setIn(
                ["isTemplateRoute"],
                action.payload.isTemplateRoute
            );
        case loginActions.loginReset:
            let deckShouldLogout = action.payload.deckName;
            let newDecksAfterLogout = state.toJS().decks.map((deckObj) => {
                return deckObj.name === deckShouldLogout
                    ? {
                          ...deckObj,
                          isLoggedIn: false,
                          token: "",
                      }
                    : {
                          ...deckObj,
                      };
            });
            return state.merge({
                decks: newDecksAfterLogout,
            });
        default:
            return state;
    }
};

/**
 * //old code for reference
 * 
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
			return state.setIn(['isPlateRoute'], action.payload.isPlateRoute);
		case loginActions.setIsTemplateRoute:
			return state.setIn(['isTemplateRoute'], action.payload.isTemplateRoute);
		case loginActions.loginReset:
			return loginInitialState;
		default:
			return state;
	}
};
 */
