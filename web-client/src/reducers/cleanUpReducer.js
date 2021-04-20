import {
  runCleanUpAction,
  pauseCleanUpAction,
  resumeCleanUpAction,
  abortCleanUpAction,
} from "actions/cleanUpActions";
import { DECKCARD_BTN } from "appConstants";

export const initialState = {
  isLoading: false,
  cleanUpData: {},
  serverErrors: {},
  isCleanUpActionCompleted: false,
  cleanUpApiError: null,
  cleanUpWebSocketError: null,
  leftActionBtn: DECKCARD_BTN.text.startUv,
  rightActionBtn: DECKCARD_BTN.text.cancel,
  runCleanUpError: null,
  pauseCleanUpError: null,
  resumeCleanUpError: null,
  abortCleanUpError: null,
};

export const cleanUpReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case runCleanUpAction.runCleanUpInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
        isCleanUpActionInProgress: false,
        isCleanUpActionCompleted: false,
        cleanUpActionInProgress: {},
        cleanUpActionInCompleted: {},
      };
    case runCleanUpAction.runCleanUpSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        runCleanUpError: false,
        leftActionBtn: DECKCARD_BTN.text.pauseUv,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };
    case runCleanUpAction.runCleanUpFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        runCleanUpError: true,
      };
    case runCleanUpAction.runCleanUpReset:
      return {
        ...state,
        runCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };
    case runCleanUpAction.runCleanUpInProgress:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isCleanUpActionInProgress: true,
        isCleanUpActionCompleted: false,
      };
    case runCleanUpAction.runCleanUpInCompleted:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        isCleanUpActionInProgress: false,
        isCleanUpActionCompleted: true,
      };

    case pauseCleanUpAction.pauseCleanUpInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
      };
    case pauseCleanUpAction.pauseCleanUpSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        pauseCleanUpError: false,
        leftActionBtn: DECKCARD_BTN.text.resumeUv,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };
    case pauseCleanUpAction.pauseCleanUpFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        pauseCleanUpError: true,
      };
    case pauseCleanUpAction.pauseCleanUpReset:
      return {
        ...state,
        ...action.payload,
        pauseCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };

    case resumeCleanUpAction.resumeCleanUpInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
      };
    case resumeCleanUpAction.resumeCleanUpSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        resumeCleanUpError: false,
        leftActionBtn: DECKCARD_BTN.text.pause,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };
    case resumeCleanUpAction.resumeCleanUpFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        resumeCleanUpError: true,
      };
    case resumeCleanUpAction.resumeCleanUpReset:
      return {
        ...state,
        ...action.payload,
        resumeCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };

    case abortCleanUpAction.abortCleanUpInitiated:
      return {
        ...state,
        ...action.payload,
        isLoading: true,
      };
    case abortCleanUpAction.abortCleanUpSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        abortCleanUpError: false,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };
    case abortCleanUpAction.abortCleanUpFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        abortCleanUpError: true,
      };
    case abortCleanUpAction.abortCleanUpReset:
      return {
        ...state,
        ...action.payload,
        abortCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };

    default:
      return state;
  }
};
