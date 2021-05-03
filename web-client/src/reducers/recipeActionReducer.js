import loginActions from "actions/loginActions";
import {
  runRecipeAction,
  pauseRecipeAction,
  resumeRecipeAction,
  abortRecipeAction,
  recipeListingAction,
  saveRecipeDataAction,
} from "actions/recipeActions";
import { DECKCARD_BTN, DECKNAME } from "appConstants";

export const initialState = {
  // recipeData: [], //all recipe data
  tempDeckName: "", //used for fetch recipe list
  decks: [
    {
      name: DECKNAME.DeckA,
      allRecipeData: [], //all recipe list
      recipeData: null, //current recipe
      showProcess: false,
      isLoading: false,
      serverErrors: {},
      runRecipeError: null,
      abortRecipeError: null,
      pauseRecipeError: null,
      resumeRecipeError: null,
      recipeListingError: null,
      runRecipeResponse: {},
      abortRecipeResponse: {},
      pauseRecipeResponse: {},
      resumeRecipeResponse: {},
      leftActionBtn: DECKCARD_BTN.text.run,
      rightActionBtn: DECKCARD_BTN.text.cancel,
      isRunRecipeCompleted: null,
      isResumeRecipeCompleted: null,
      runRecipeInCompleted: {},
      runRecipeInProgress: null,
      hours: 0,
      mins: 0,
      secs: 0,
      showCleanUp: false,
      leftActionBtnDisabled: false,
      rightActionBtnDisabled: false,
    },
    {
      name: DECKNAME.DeckB,
      allRecipeData: [], //all recipe list
      recipeData: null, //current recipe
      showProcess: false,
      isLoading: false,
      serverErrors: {},
      runRecipeError: null,
      abortRecipeError: null,
      pauseRecipeError: null,
      resumeRecipeError: null,
      recipeListingError: null,
      runRecipeResponse: {},
      abortRecipeResponse: {},
      pauseRecipeResponse: {},
      resumeRecipeResponse: {},
      leftActionBtn: DECKCARD_BTN.text.run,
      rightActionBtn: DECKCARD_BTN.text.cancel,
      isRunRecipeCompleted: null,
      isResumeRecipeCompleted: null,
      runRecipeInCompleted: {},
      runRecipeInProgress: null,
      hours: 0,
      mins: 0,
      secs: 0,
      showCleanUp: false,
      leftActionBtnDisabled: false,
      rightActionBtnDisabled: false,
    },
  ],
};

export const recipeActionReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case saveRecipeDataAction.saveRecipeDataForDeck: //set and update: depend on deckName
      let deckNameForRecipe = action.payload.deckName;
      let newDecksAfterRecipeDataAdded = state.decks.map((deckObj) => {
        return deckObj.name === deckNameForRecipe
          ? {
              ...deckObj,
              recipeData: action.payload.recipeData,
              showProcess: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: newDecksAfterRecipeDataAdded,
      };
    case saveRecipeDataAction.updateRecipeReducerDataForDeck: //update data in a deck
      let deckNameToUpdate = action.payload.deckName;
      let newDecksAfterRequiredUpdations = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToUpdate
          ? {
              ...deckObj,
              ...action.payload.body,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: newDecksAfterRequiredUpdations,
      };

    case saveRecipeDataAction.resetRecipeDataForDeck:
      let deckNameToResetData = action.payload.deckName;
      let newDeckDataAfterReset = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToResetData
          ? {
              ...deckObj,
              ...action.payload.body,
              showProcess: false,
              showCleanUp: false,
              recipeData: null,
              runRecipeInCompleted: {},
              runRecipeInProgress: null,
              leftActionBtn: DECKCARD_BTN.text.run,
              rightActionBtn: DECKCARD_BTN.text.cancel,
              leftActionBtnDisabled: false,
              rightActionBtnDisabled: false,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: newDeckDataAfterReset,
      };

    case runRecipeAction.runRecipeInitiated:
      let deckNameToInitiateRun = action.payload.params.deckName;
      let decksAfterRunInitiated = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToInitiateRun
          ? {
              ...deckObj,
              runRecipeResponse: {
                recipeId: action.payload.params.recipeId,
              },
              leftActionBtn: DECKCARD_BTN.text.pause,
              rightActionBtn: DECKCARD_BTN.text.abort,
            }
          : deckObj;
      });

      return {
        ...state,
        decks: decksAfterRunInitiated,
      };
    case runRecipeAction.runRecipeSuccess:
      return state;
    // return {
    //     ...state,
    //     // runRecipeResponse: action.payload.response,
    //     // isLoading: false,
    //     // runRecipeError: false,
    // };
    case runRecipeAction.runRecipeFailed:
      return {
        ...state,
        // serverErrors: action.payload.serverErrors,
        // isLoading: false,
        // runRecipeError: true,
      };
    case runRecipeAction.runRecipeReset:
      let deckNameToReset = action.payload.deckName;
      const decksAfterRecipeReset = state.decks.map((deckObj) => {
        const prevRecipeData = deckObj.recipeData;
        return deckObj.name === deckNameToReset
          ? {
              ...initialState.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckNameToReset
              ),
              recipeData: null,
              allRecipeData: prevRecipeData,
              showProcess: false,
            }
          : deckObj;
      });
      return {
        ...state,
        // runRecipeError: null,
        decks: decksAfterRecipeReset,
      };
    case runRecipeAction.runRecipeInProgress:
      let response = action.payload.runRecipeInProgress;
      let deckNameRunInProgress =
        response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let decksAfterRunInProgress = state.decks.map((deckObj) => {
        return deckObj.name === deckNameRunInProgress
          ? {
              ...deckObj,
              runRecipeInProgress: {
                ...response,
              },
            }
          : deckObj;
      });

      return {
        ...state,
        decks: decksAfterRunInProgress,
        // ...action.payload,
        // isLoading: false,
        // isRunRecipeCompleted: false,
        // leftActionBtn: DECKCARD_BTN.text.pause,
        // rightActionBtn: DECKCARD_BTN.text.abort,
      };
    case runRecipeAction.runRecipeInCompleted:
      let responseRunRecipeInCompleted = action.payload.runRecipeInCompleted;
      let deckNameOfRunRecipeInCompleted =
        responseRunRecipeInCompleted.deck === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;
      let decksAfterRunRecipeInCompleted = state.decks.map((deckObj) => {
        return deckObj.name === deckNameOfRunRecipeInCompleted
          ? {
              ...deckObj,
              runRecipeInCompleted: {
                ...responseRunRecipeInCompleted,
              },
              leftActionBtn: DECKCARD_BTN.text.done,
              rightActionBtn: DECKCARD_BTN.text.cancel,
              rightActionBtnDisabled: true,
              isRunRecipeCompleted: true,
            }
          : deckObj;
      });

      return {
        ...state,
        decks: decksAfterRunRecipeInCompleted,
        // ...action.payload,
        // isRunRecipeCompleted: true,
        // leftActionBtn: DECKCARD_BTN.text.done,
        // rightActionBtn: DECKCARD_BTN.text.cancel,
      };

    case pauseRecipeAction.pauseRecipeInitiated:
      return { ...state /*...action.payload, isLoading: true */ };
    case pauseRecipeAction.pauseRecipeSuccess:
      let responsePauseSuccess = action.payload.response;
      let deckNamePauseSuccess =
        responsePauseSuccess.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;
      let decksAfterPauseSuccess = state.decks.map((deckObj) => {
        return deckObj.name === deckNamePauseSuccess
          ? {
              ...deckObj,
              pauseRecipeResponse: responsePauseSuccess,
              pauseRecipeError: false,
              leftActionBtn: DECKCARD_BTN.text.resume,
            }
          : deckObj;
      });

      return {
        ...state,
        decks: decksAfterPauseSuccess,
      };
    case pauseRecipeAction.pauseRecipeFailed:
      return {
        ...state,
        // ...action.payload,
        // isLoading: false,
        // pauseRecipeError: true,
      };
    case pauseRecipeAction.pauseRecipeReset:
      return {
        ...state,
        // pauseRecipeError: null,
      };

    case abortRecipeAction.abortRecipeInitiated:
      return {
        ...state,
        // ...action.payload,
        // isLoading: true,
        // abortRecipeError: null,
      };
    case abortRecipeAction.abortRecipeSuccess:
      return {
        ...state,
        // ...action.payload,
        // isLoading: false,
        // abortRecipeError: false,
        // leftActionBtn: DECKCARD_BTN.text.run,
        // rightActionBtn: DECKCARD_BTN.text.cancel,
      };
    case abortRecipeAction.abortRecipeFailed:
      return {
        ...state,
        // ...action.payload,
        // isLoading: false,
        // abortRecipeError: true,
      };
    case abortRecipeAction.abortRecipeReset:
      return {
        ...state,
        // abortRecipeError: null,
      };

    case resumeRecipeAction.resumeRecipeInitiated:
      return { ...state /*...action.payload, isLoading: true*/ };
    case resumeRecipeAction.resumeRecipeSuccess:
      let deckNameResumeSuccess =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;
      let decksAfterResumeSuccess = state.decks.map((deckObj) => {
        return deckObj.name === deckNameResumeSuccess
          ? {
              ...deckObj,
              leftActionBtn: DECKCARD_BTN.text.pause,
              resumeRecipeError: false,
            }
          : deckObj;
      });

      return {
        ...state,
        decks: decksAfterResumeSuccess,
        // ...action.payload,
        // isLoading: false,
        // resumeRecipeError: false,
        // leftActionBtn: DECKCARD_BTN.text.pause,
      };
    case resumeRecipeAction.resumeRecipeFailed:
      return {
        ...state,
        // ...action.payload,
        // isLoading: false,
        // resumeRecipeError: true,
      };

    case resumeRecipeAction.resumeRecipeReset:
      return {
        ...state,
        // resumeRecipeError: null,
      };

    case resumeRecipeAction.resumeRecipeInProgress:
      return {
        ...state,
        // ...action.payload,
        // isLoading: false,
        // isResumeRecipeCompleted: false,
      };
    case resumeRecipeAction.resumeRecipeInCompleted:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isResumeRecipeCompleted: false,
      };

    case recipeListingAction.recipeListingInitiated:
      return {
        ...state,
        // ...action.payload,
        // isLoading: true,
        tempDeckName: action.payload.deckName,
      };
    case recipeListingAction.recipeListingSuccess:
      const newDecksAfterRecipeList = state.decks.map((deckObj) => {
        return deckObj.name === state.tempDeckName
          ? {
              ...deckObj,
              allRecipeData: action.payload.response,
            }
          : deckObj;
      });
      return {
        ...state,
        tempDeckName: "",
        decks: newDecksAfterRecipeList,
      };
    case recipeListingAction.recipeListingFailed:
      return {
        ...state,
        // serverErrors: action.payload.serverErrors,
        // recipeListingError: true,
        // isLoading: false,
      };
    case recipeListingAction.recipeListingReset:
      return {
        ...state,
        // recipeListingError: null,
      };

    case loginActions.loginReset:
      let deckToResetOnLogout = action.payload.deckName;
      let newDecksAfterLoggedOut = state.decks.map((deckObj) => {
        return deckObj.name === deckToResetOnLogout
          ? {
              ...initialState.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckToResetOnLogout
              ),
            }
          : deckObj;
      });
      return {
        ...state,
        decks: newDecksAfterLoggedOut,
      };

    default:
      return state;
  }
};

/**old code for reference */
/*
export const initialState = {
  isLoading: false,
  serverErrors: {},
  runRecipeError: null,
  abortRecipeError: null,
  pauseRecipeError: null,
  resumeRecipeError: null,
  recipeListingError: null,
  runRecipeResponse: {},
  abortRecipeResponse: {},
  pauseRecipeResponse: {},
  resumeRecipeResponse: {},
  recipeData: [],
  leftActionBtn: DECKCARD_BTN.text.run,
  rightActionBtn: DECKCARD_BTN.text.cancel,
  isRunRecipeCompleted: null,
  isResumeRecipeCompleted: null,
  runRecipeInCompleted: {},
  runRecipeInProgress: null,
};

export const recipeActionReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case runRecipeAction.runRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case runRecipeAction.runRecipeSuccess:
      return {
        ...state,
        runRecipeResponse: action.payload.response,
        isLoading: false,
        runRecipeError: false,
      };
    case runRecipeAction.runRecipeFailed:
      return {
        ...state,
        serverErrors: action.payload.serverErrors,
        isLoading: false,
        runRecipeError: true,
      };
    case runRecipeAction.runRecipeReset:
      return {
        ...state,
        runRecipeError: null,
      };
    case runRecipeAction.runRecipeInProgress:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isRunRecipeCompleted: false,
        leftActionBtn: DECKCARD_BTN.text.pause,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };
    case runRecipeAction.runRecipeInCompleted:
      return {
        ...state,
        ...action.payload,
        isRunRecipeCompleted: true,
        leftActionBtn: DECKCARD_BTN.text.done,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };

    case pauseRecipeAction.pauseRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case pauseRecipeAction.pauseRecipeSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        pauseRecipeError: false,
        leftActionBtn: DECKCARD_BTN.text.resume,
      };
    case pauseRecipeAction.pauseRecipeFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        pauseRecipeError: true,
      };
    case pauseRecipeAction.pauseRecipeReset:
      return {
        ...state,
        pauseRecipeError: null,
      };

    case abortRecipeAction.abortRecipeInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
        abortRecipeError: null,
      };
    case abortRecipeAction.abortRecipeSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        abortRecipeError: false,
        leftActionBtn: DECKCARD_BTN.text.run,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };
    case abortRecipeAction.abortRecipeFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        abortRecipeError: true,
      };
    case abortRecipeAction.abortRecipeReset:
      return {
        ...state,
        abortRecipeError: null,
      };

    case resumeRecipeAction.resumeRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case resumeRecipeAction.resumeRecipeSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        resumeRecipeError: false,
        leftActionBtn: DECKCARD_BTN.text.pause,
      };
    case resumeRecipeAction.resumeRecipeFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        resumeRecipeError: true,
      };

    case resumeRecipeAction.resumeRecipeReset:
      return {
        ...state,
        resumeRecipeError: null,
      };

    case resumeRecipeAction.resumeRecipeInProgress:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isResumeRecipeCompleted: false,
      };
    case resumeRecipeAction.resumeRecipeInCompleted:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isResumeRecipeCompleted: false,
      };

    case recipeListingAction.recipeListingInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case recipeListingAction.recipeListingSuccess:
      return {
        ...state,
        recipeData: action.payload.response,
        isLoading: false,
        recipeListingError: false,
      };
    case recipeListingAction.recipeListingFailed:
      return {
        ...state,
        serverErrors: action.payload.serverErrors,
        recipeListingError: true,
        isLoading: false,
      };
    case recipeListingAction.recipeListingReset:
      return {
        ...state,
        recipeListingError: null,
      };

    default:
      return state;
  }
};
*/
