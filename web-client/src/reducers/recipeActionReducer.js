import {
    runRecipeAction,
    pauseRecipeAction,
    resumeRecipeAction,
    abortRecipeAction,
    recipeListingAction,
} from "actions/recipeActions";
import { DECKCARD_BTN, DECKNAME } from "appConstants";

export const initialState = {
    recipeData: [],//all recipe data
    decks: [
        {
            name: DECKNAME.DeckA,
            recipeData: null, //current recipe 
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
        },
        {
            name: DECKNAME.DeckB,
            recipeData: null,//current recipe
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
        },
    ],
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
