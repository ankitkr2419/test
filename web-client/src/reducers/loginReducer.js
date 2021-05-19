import { fromJS } from "immutable";
import loginActions from "actions/loginActions";
import { DECKNAME, USER_ROLES } from "../appConstants";
import { getUpdatedDecks } from "utils/helpers";

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
      const changesInLoginInitMatchedDeck = {
        error: false,
        msg: "",
        isError: false,
        isActive: true,
      };
      const changesInLoginInitUnMatchedDeck = {
        isActive: false,
      };
      const updatedDecks = getUpdatedDecks(
        state,
        action.payload.body.deckName,
        changesInLoginInitMatchedDeck,
        changesInLoginInitUnMatchedDeck,
        true
      );

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

        const changesInLoginSuccessMatchedDeckA = {
          isLoggedIn: true,
          error: false,
          msg: "",
          isAdmin: isAdminTemp,
          isActive: true,
          token,
        };
        const changesInLoginSuccessUnMatchedDeckA = {
          isActive: false,
        };
        const newDeckA = getUpdatedDecks(
          state,
          DECKNAME.DeckA,
          changesInLoginSuccessMatchedDeckA,
          changesInLoginSuccessUnMatchedDeckA,
          true
        );

        return state.merge({
          isLoading: false,
          decks: newDeckA,
        });
      } else if (deckName && deckName === DECKNAME.DeckB) {
        //update and login deck B

        const changesInLoginSuccessMatchedDeckB = {
          isLoggedIn: true,
          error: false,
          msg: "",
          isAdmin: isAdminTemp,
          isActive: true,
          token,
        };
        const changesInLoginSuccessUnMatchedDeckB = {
          isActive: false,
        };
        const newDeckB = getUpdatedDecks(
          state,
          DECKNAME.DeckB,
          changesInLoginSuccessMatchedDeckB,
          changesInLoginSuccessUnMatchedDeckB,
          true
        );

        return state.merge({
          isLoading: false,
          decks: newDeckB,
        });
      } else {
        //if deck name dont match, then dont update state
        return state;
      }
    case loginActions.failureAction:
      let err = action.payload.serverErrors?.msg
        ? action.payload.serverErrors.msg
        : "Something went wrong!";
      let newDecks = getUpdatedDecks(
        state,
        state.toJS().deckName,
        { error: true, msg: err, token: "" },
        {},
        true
      );

      return state.merge({
        isLoading: false,
        decks: newDecks,
      });

    case loginActions.setActiveDeck:
      let deckNameToSetActive = action.payload.deckName;

      let newDecksForActive = getUpdatedDecks(
        state,
        deckNameToSetActive,
        { isActive: true },
        { isActive: false },
        true
      );

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
      return state.setIn(["isTemplateRoute"], action.payload.isTemplateRoute);

    case loginActions.loginReset:
      let deckShouldLogout = action.payload.deckName;

      let newDecksAfterLogout = getUpdatedDecks(
        state,
        deckShouldLogout,
        { isLoggedIn: false, token: "" },
        {},
        true
      );

      return state.merge({
        decks: newDecksAfterLogout,
      });

    default:
      return state;
  }
};
