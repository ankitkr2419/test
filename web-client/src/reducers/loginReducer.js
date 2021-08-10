import { fromJS } from "immutable";
import loginActions, {
  deckBlockActions,
  logoutActions,
} from "actions/loginActions";
import { DECKNAME, USER_ROLES } from "../appConstants";
import { getUpdatedDecks } from "utils/helpers";

const initialStateOfDecks = () => {
  return [
    {
      name: DECKNAME.DeckA,
      isLoggedIn: false,
      error: false,
      msg: "",
      role: "",
      isAdmin: false,
      isEngineer: false,
      isOperator: false,
      isActive: true,
      isDeckBlocked: false,
      token: "",
    },
    {
      name: DECKNAME.DeckB,
      isLoggedIn: false,
      error: false,
      msg: "",
      role: "",
      isAdmin: false,
      isEngineer: false,
      isOperator: false,
      isActive: false,
      isDeckBlocked: false,
      token: "",
    },
  ];
};

const loginInitialState = fromJS({
  tempDeckName: "",
  tokenForHoming: "",
  isLoggedInForHoming: false,
  isLoading: false,
  isPlateRoute: false,
  isTemplateRoute: false,
  deckName: "", //initiated deckName while login process only (temp)
  isAdmin: false, //initiated isAdmin while login process only (temp)
  isEngineer: false, //initiated isEngineer while login process only (temp)
  isOperator: false, //initiated isOperator while login process only (temp)
  decks: initialStateOfDecks(),
});

export const loginReducer = (state = loginInitialState, action) => {
  switch (action.type) {
    case loginActions.loginInitiated:
      let updatedDecks = state.toJS().decks;

      // only if deckName is provided i.e. NOT for homing
      // if deckName is not for homing i.e. it will contain some deckName
      if (action.payload.body.deckName !== "") {
        const changesInLoginInitMatchedDeck = {
          error: false,
          msg: "",
          isActive: true,
        };
        const changesInLoginInitUnMatchedDeck = {
          isActive: false,
        };
        updatedDecks = getUpdatedDecks(
          state,
          action.payload.body.deckName,
          changesInLoginInitMatchedDeck,
          changesInLoginInitUnMatchedDeck,
          true
        );
      }

      return state.merge({
        isLoading: true,
        isLoggedInForHoming: false,
        deckName: action.payload.body.deckName,
        // isAdmin: action.payload.body.role === USER_ROLES.ADMIN,//TODO remove after tested
        // isEngineer: action.payload.body.role === USER_ROLES.ENGINEER,
        decks: updatedDecks,
      });

    case loginActions.successAction:
      const role = action.payload.response?.role;
      const token = action.payload.response.token;
      let deckName = state.toJS().deckName; //current deckname
      let isAdminTemp = role === USER_ROLES.ADMIN;
      let isEngineerTemp = role === USER_ROLES.ENGINEER;
      let isOperatorTemp = role === USER_ROLES.OPERATOR;
      if (deckName && deckName === DECKNAME.DeckA) {
        //update and login deck A

        const changesInLoginSuccessMatchedDeckA = {
          isLoggedIn: true,
          error: false,
          msg: "",
          role,
          isAdmin: isAdminTemp,
          isEngineer: isEngineerTemp,
          isOperator: isOperatorTemp,
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
          role,
          isAdmin: isAdminTemp,
          isEngineer: isEngineerTemp,
          isOperator: isOperatorTemp,
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
      }
      // only if deckName is provided i.e. NOT for homing
      // if no deckname is provided that means we login for homing
      else if (deckName === "") {
        return state.merge({
          tokenForHoming: action.payload.response.token,
          isLoading: false,
          isLoggedInForHoming: true,
          decks: state.toJS().decks,
        });
      }

    case loginActions.failureAction:
      let err = action.payload.serverErrors?.msg
        ? action.payload.serverErrors.msg
        : "Something went wrong!";

      let newDecks = state.toJS().decks; // current deck state

      // only if deckName is provided i.e. NOT for homing
      if (state.toJS().deckName !== "") {
        newDecks = getUpdatedDecks(
          state,
          state.toJS().deckName,
          { error: true, msg: err, token: "" },
          {},
          true
        );
      }

      return state.merge({
        isLoggedInForHoming: false,
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

    // login reset
    case loginActions.loginReset:
      let deckShouldLogout = action.payload.deckName;

      let newDecksAfterLogout = getUpdatedDecks(
        state,
        deckShouldLogout,
        {
          error: null,
          isLoggedIn: false,
          isAdmin: false,
          isEngineer: false,
          isOperator: false,
          token: "",
          isDeckBlocked: false,
        },
        {},
        true
      );

      return state.merge({
        decks: newDecksAfterLogout,
      });

    //logout init
    case logoutActions.logoutActionInitiated:
      return state.merge({
        isLoading: true,
        deckName: action.payload.deckName,
        decks: state.toJS().decks,
      });

    //logout success
    case logoutActions.logoutActionSuccess:
      let newdeckStateAferLogoutSuccess = state.toJS().decks; // current decks

      // only if deckName is provided i.e. NOT for homing
      if (state.toJS().deckName !== "") {
        newdeckStateAferLogoutSuccess = getUpdatedDecks(
          state,
          state.toJS().deckName,
          {
            error: false,
            isLoggedIn: false,
            isAdmin: false,
            isEngineer: false,
            isOperator: false,
            token: "",
            isDeckBlocked: false,
          },
          {},
          true
        );
      }

      return state.merge({
        tempDeckName: "",
        isLoading: false,
        isLoggedInForHoming: false,
        decks: newdeckStateAferLogoutSuccess,
      });

    //logout fail
    case logoutActions.logoutActionFailure:
      let errorMsg = action.payload.error?.msg
        ? action.payload.error.msg
        : "Something went wrong!";

      let newDeckStateAfterLogoutFail = state.toJS().decks; // current decks

      // only if deckName is provided i.e. NOT for homing
      if (state.toJS().deckName !== "") {
        newDeckStateAfterLogoutFail = getUpdatedDecks(
          state,
          state.toJS().deckName,
          { error: true, msg: errorMsg, token: "" },
          {},
          true
        );
      }

      return state.merge({
        tempDeckName: "",
        isLoading: false,
        decks: newDeckStateAfterLogoutFail,
      });

    //block current deck : in case of recipe edit/add process
    case deckBlockActions.deckBlockInitiated:
      let newDeckStateAfterDeckBlocked = getUpdatedDecks(
        state,
        action.payload.deckName,
        { isDeckBlocked: true },
        {},
        true
      );

      return state.merge({
        decks: newDeckStateAfterDeckBlocked,
      });

    case deckBlockActions.deckBlockReset:
      let newDeckStateAfterDeckReset = getUpdatedDecks(
        state,
        action.payload.deckName,
        { isDeckBlocked: false },
        {},
        true
      );

      return state.merge({
        decks: newDeckStateAfterDeckReset,
      });

    default:
      return state;
  }
};
