import {
  runCleanUpAction,
  pauseCleanUpAction,
  resumeCleanUpAction,
  abortCleanUpAction,
  cleanUpHourActions,
  cleanUpMinsActions,
  cleanUpSecsActions,
  setShowCleanUpAction,
} from "actions/cleanUpActions";
import { DECKCARD_BTN, DECKNAME } from "appConstants";
import { getUpdatedDecks } from "utils/helpers";

export const initialState = {
  decks: [
    {
      name: DECKNAME.DeckA,
      isLoading: false,
      cleanUpData: null,
      showCleanUp: false,
      serverErrors: {},
      isCleanUpActionCompleted: false,
      cleanUpApiError: null,
      cleanUpWebSocketError: null,
      leftActionBtn: DECKCARD_BTN.text.startUv,
      rightActionBtn: DECKCARD_BTN.text.cancel,
      leftActionBtnDisabled: false,
      rightActionBtnDisabled: false,
      runCleanUpError: null,
      pauseCleanUpError: null,
      resumeCleanUpError: null,
      abortCleanUpError: null,
      hours: 0,
      mins: 0,
      secs: 0,
      progress: 0,
    },
    {
      name: DECKNAME.DeckB,
      isLoading: false,
      cleanUpData: null,
      serverErrors: {},
      isCleanUpActionCompleted: false,
      cleanUpApiError: null,
      cleanUpWebSocketError: null,
      leftActionBtn: DECKCARD_BTN.text.startUv,
      rightActionBtn: DECKCARD_BTN.text.cancel,
      leftActionBtnDisabled: false,
      rightActionBtnDisabled: false,
      runCleanUpError: null,
      pauseCleanUpError: null,
      resumeCleanUpError: null,
      abortCleanUpError: null,
      hours: 0,
      mins: 0,
      secs: 0,
      progress: 0,
    },
  ],
};

export const cleanUpReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case runCleanUpAction.runCleanUpInitiated:
      const deckInitiateName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const changesForCleanUpMatched = {
        isLoading: true,
        isCleanUpActionInProgress: false,
        isCleanUpActionCompleted: false,
        showCleanUp: true,
      };

      const dockAfterRunInit = getUpdatedDecks(
        state,
        deckInitiateName,
        changesForCleanUpMatched
      );

      return {
        ...state,
        decks: dockAfterRunInit,
      };

    case runCleanUpAction.runCleanUpSuccess:
      const deckSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const changesForCleanUpSuccessMatched = {
        leftActionBtn: DECKCARD_BTN.text.pauseUv,
        rightActionBtn: DECKCARD_BTN.text.abort,
        isLoading: false,
        runCleanUpError: false,
      };

      const dockAfterRunSuccess = getUpdatedDecks(
        state,
        deckSuccessName,
        changesForCleanUpSuccessMatched
      );

      return {
        ...state,
        decks: dockAfterRunSuccess,
      };

    case runCleanUpAction.runCleanUpFailed:
      const deckFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const changesForCleanUpFailureMatched = {
        isLoading: false,
        runCleanUpError: true,
      };

      const dockAfterRunFailure = getUpdatedDecks(
        state,
        deckFailureName,
        changesForCleanUpFailureMatched
      );

      return {
        ...state,
        decks: dockAfterRunFailure,
      };

    case runCleanUpAction.runCleanUpReset:
      const deckResetName =
        action.payload.params.deckName === DECKNAME.DeckA
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const changesForCleanUpResetMatched = {
        cleanUpData: null,
        hours: 0,
        mins: 0,
        secs: 0,
        progress: 0,
        runCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
        leftActionBtnDisabled: false,
        rightActionBtnDisabled: false,
        showCleanUp: false,
      };

      const dockAfterRunReset = getUpdatedDecks(
        state,
        deckResetName,
        changesForCleanUpResetMatched
      );

      return {
        ...state,
        decks: dockAfterRunReset,
      };

    case runCleanUpAction.runCleanUpInProgress:
      const deckNameCheck =
        action.payload.cleanUpActionInProgress &&
        JSON.parse(action.payload.cleanUpActionInProgress).deck;

      const deckInProgressName =
        deckNameCheck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const changesForRunCleanUpProgressMatched = {
        isLoading: false,
        isCleanUpActionCompleted: false,
        isCleanUpActionInProgress: true,
        cleanUpData: action.payload.cleanUpActionInProgress,
      };

      const dockAfterProgress = getUpdatedDecks(
        state,
        deckInProgressName,
        changesForRunCleanUpProgressMatched
      );

      return {
        ...state,
        decks: dockAfterProgress,
      };

    case runCleanUpAction.runCleanUpInCompleted:
      const progressEndDeck =
        action.payload.cleanUpActionInCompleted &&
        JSON.parse(action.payload.cleanUpActionInCompleted).deck;

      const deckInCompleteResponse =
        progressEndDeck === DECKNAME.DeckAShort
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const changesForRunCleanUpCompletedMatched = {
        isLoading: false,
        isCleanUpActionCompleted: true,
        isCleanUpActionInProgress: false,
        leftActionBtn: DECKCARD_BTN.text.done,
        rightActionBtn: DECKCARD_BTN.text.cancel,
        rightActionBtnDisabled: true,
      };

      const dockAfterProgressCompleted = getUpdatedDecks(
        state,
        progressEndDeck,
        changesForRunCleanUpCompletedMatched
      );

      return {
        ...state,
        decks: dockAfterProgressCompleted,
      };

    case pauseCleanUpAction.pauseCleanUpInitiated:
      const deckPauseInitiateName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const cleanUpPauseInitMatchedChanges = {
        isLoading: true,
      };
      const dockAfterPauseInit = getUpdatedDecks(
        state,
        deckPauseInitiateName,
        cleanUpPauseInitMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterPauseInit,
      };

    case pauseCleanUpAction.pauseCleanUpSuccess:
      const deckPauseSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const cleanUpPauseSuccessMatchedChanges = {
        isLoading: false,
        pauseCleanUpError: false,
        leftActionBtn: DECKCARD_BTN.text.resumeUv,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };
      const dockAfterPauseSuccess = getUpdatedDecks(
        state,
        deckPauseSuccessName,
        cleanUpPauseSuccessMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterPauseSuccess,
      };

    case pauseCleanUpAction.pauseCleanUpFailed:
      const deckPauseFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const cleanUpPauseFailureMatchedChanges = {
        isLoading: false,
        pauseCleanUpError: true,
      };

      const dockAfterPauseFailure = getUpdatedDecks(
        state,
        deckPauseFailureName,
        cleanUpPauseFailureMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterPauseFailure,
      };

    case pauseCleanUpAction.pauseCleanUpReset:
      const deckPauseResetName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const cleanUpPauseResetMatchedChanges = {
        pauseCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };

      const dockAfterPauseReset = getUpdatedDecks(
        state,
        deckPauseResetName,
        cleanUpPauseResetMatchedChanges
      );
      return {
        ...state,
        decks: dockAfterPauseReset,
      };

    case resumeCleanUpAction.resumeCleanUpInitiated:
      const deckResumeInitiateName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const cleanUpResumeInitMatchedChanges = {
        isLoading: true,
      };

      const dockAfterResumeInit = getUpdatedDecks(
        state,
        deckResumeInitiateName,
        cleanUpResumeInitMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterResumeInit,
      };

    case resumeCleanUpAction.resumeCleanUpSuccess:
      const deckResumeSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const cleanUpResumeSuccessMatchedChanges = {
        isLoading: false,
        resumeCleanUpError: false,
        leftActionBtn: DECKCARD_BTN.text.pauseUv,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };

      const dockAfterResumeSuccess = getUpdatedDecks(
        state,
        deckResumeSuccessName,
        cleanUpResumeSuccessMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterResumeSuccess,
      };

    case resumeCleanUpAction.resumeCleanUpFailed:
      const deckResumeFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const cleanUpResumeFailureMatchedChanges = {
        isLoading: false,
        resumeCleanUpError: true,
      };

      const dockAfterResumeFailure = getUpdatedDecks(
        state,
        deckResumeFailureName,
        cleanUpResumeFailureMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterResumeFailure,
      };

    case resumeCleanUpAction.resumeCleanUpReset:
      const deckResumeResetName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const cleanUpResumeResetMatchedChanges = {
        resumeCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };

      const dockAfterResumeReset = getUpdatedDecks(
        state,
        deckResumeResetName,
        cleanUpResumeResetMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterResumeReset,
      };

    case abortCleanUpAction.abortCleanUpInitiated:
      const deckAbortInitiateName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const cleanUpAbortInitMatchedChanges = {
        isLoading: true,
      };

      const dockAfterAbortInit = getUpdatedDecks(
        state,
        deckAbortInitiateName,
        cleanUpAbortInitMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterAbortInit,
      };

    case abortCleanUpAction.abortCleanUpSuccess:
      const deckAbortSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const cleanUpAbortSuccessMatchedChanges = {
        showCleanUp: false,
        isLoading: false,
        abortCleanUpError: false,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
        hours: 0,
        mins: 0,
        secs: 0,
        progress: 0,
        cleanUpData: null,
      };

      const dockAfterAbortSuccess = getUpdatedDecks(
        state,
        deckAbortSuccessName,
        cleanUpAbortSuccessMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterAbortSuccess,
      };

    case abortCleanUpAction.abortCleanUpFailed:
      const deckAbortFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const cleanUpAbortFailureMatchedChanges = {
        isLoading: false,
        abortCleanUpError: true,
      };

      const dockAfterAbortFailure = getUpdatedDecks(
        state,
        deckAbortFailureName,
        cleanUpAbortFailureMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterAbortFailure,
      };

    case abortCleanUpAction.abortCleanUpReset:
      const deckAbortResetName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const cleanUpAbortResetMatchedChanges = {
        abortCleanUpError: null,
        leftActionBtn: DECKCARD_BTN.text.startUv,
        rightActionBtn: DECKCARD_BTN.text.cancel,
        hours: 0,
        mins: 0,
        secs: 0,
        progress: 0,
        cleanUpData: null,
      };

      const dockAfterAbortReset = getUpdatedDecks(
        state,
        deckAbortResetName,
        cleanUpAbortResetMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterAbortReset,
      };

    case cleanUpHourActions.setHours:
      let deckNameToSetHours = action.payload.params.deckName;
      let newHours = parseInt(action.payload.params.hours);

      const cleanUpHoursMatchedChanges = {
        hours: newHours,
      };

      const dockAfterHoursSet = getUpdatedDecks(
        state,
        deckNameToSetHours,
        cleanUpHoursMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterHoursSet,
      };

    case cleanUpMinsActions.setMins:
      let deckNameToSetMins = action.payload.params.deckName;
      let newMins = parseInt(action.payload.params.mins);

      const cleanUpMinsMatchedChanges = {
        mins: newMins,
      };

      const dockAfterMinsSet = getUpdatedDecks(
        state,
        deckNameToSetMins,
        cleanUpMinsMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterMinsSet,
      };

    case cleanUpSecsActions.setSecs:
      let deckNameToSetSecs = action.payload.params.deckName;
      let newSecs = parseInt(action.payload.params.secs);

      const cleanUpSecsMatchedChanges = {
        secs: newSecs,
      };

      const dockAfterSecsSet = getUpdatedDecks(
        state,
        deckNameToSetSecs,
        cleanUpSecsMatchedChanges
      );

      return {
        ...state,
        decks: dockAfterSecsSet,
      };

    case setShowCleanUpAction.setShowCleanUp:
      let deckNameToShowCleanUp = action.payload.params.deckName;

      const cleanUpSetShowChanges = {
        showCleanUp: true,
      };

      const dockAfterShowCleanUp = getUpdatedDecks(
        state,
        deckNameToShowCleanUp,
        cleanUpSetShowChanges
      );

      return {
        ...state,
        decks: dockAfterShowCleanUp,
      };

    case setShowCleanUpAction.resetShowCleanUp:
      let deckNameToHideCleanUp = action.payload.params.deckName;

      const cleanUpSetHideChanges = {
        showCleanUp: false,
      };

      const dockAfterHideCleanUp = getUpdatedDecks(
        state,
        deckNameToHideCleanUp,
        cleanUpSetHideChanges
      );

      return {
        ...state,
        decks: dockAfterHideCleanUp,
      };

    default:
      return state;
  }
};
