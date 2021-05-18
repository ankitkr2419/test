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

      let dockAfterRunInit = state.decks.map((deckObj) => {
        return deckObj.name === deckInitiateName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckInitiateName
              ),
              isLoading: true,
              isCleanUpActionInProgress: false,
              isCleanUpActionCompleted: false,
              showCleanUp: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterRunInit,
      };

    case runCleanUpAction.runCleanUpSuccess:
      const deckSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterRunSuccess = state.decks.map((deckObj) => {
        return deckObj.name === deckSuccessName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckSuccessName
              ),
              leftActionBtn: DECKCARD_BTN.text.pauseUv,
              rightActionBtn: DECKCARD_BTN.text.abort,
              isLoading: false,
              runCleanUpError: false,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterRunSuccess,
      };

    case runCleanUpAction.runCleanUpFailed:
      const deckFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterRunFailure = state.decks.map((deckObj) => {
        return deckObj.name === deckFailureName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckFailureName
              ),
              isLoading: false,
              runCleanUpError: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterRunFailure,
      };

    case runCleanUpAction.runCleanUpReset:
      const deckResetName =
        action.payload.params.deckName === DECKNAME.DeckA
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      let dockAfterRunReset = state.decks.map((deckObj) => {
        return deckObj.name === deckResetName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckResetName
              ),
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
            }
          : deckObj;
      });
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

      let dockAfterProgress = state.decks.map((deckObj) => {
        return deckObj.name === deckInProgressName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckInProgressName
              ),
              isLoading: false,
              isCleanUpActionCompleted: false,
              isCleanUpActionInProgress: true,
              cleanUpData: action.payload.cleanUpActionInProgress,
            }
          : deckObj;
      });
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

      let dockAfterProgressCompleted = state.decks.map((deckObj) => {
        return deckObj.name === deckInCompleteResponse
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckInCompleteResponse
              ),
              isLoading: false,
              isCleanUpActionCompleted: true,
              isCleanUpActionInProgress: false,
              leftActionBtn: DECKCARD_BTN.text.done,
              rightActionBtn: DECKCARD_BTN.text.cancel,
              rightActionBtnDisabled: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterProgressCompleted,
      };

    case pauseCleanUpAction.pauseCleanUpInitiated:
      const deckPauseInitiateName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      let dockAfterPauseInit = state.decks.map((deckObj) => {
        return deckObj.name === deckPauseInitiateName
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckPauseInitiateName
              ),
              isLoading: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterPauseInit,
      };

    case pauseCleanUpAction.pauseCleanUpSuccess:
      const deckPauseSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterPauseSuccess = state.decks.map((deckObj) => {
        return deckObj.name === deckPauseSuccessName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckPauseSuccessName
              ),
              isLoading: false,
              pauseCleanUpError: false,
              leftActionBtn: DECKCARD_BTN.text.resumeUv,
              rightActionBtn: DECKCARD_BTN.text.abort,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterPauseSuccess,
      };

    case pauseCleanUpAction.pauseCleanUpFailed:
      const deckPauseFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterPauseFailure = state.decks.map((deckObj) => {
        return deckObj.name === deckPauseFailureName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckPauseFailureName
              ),
              isLoading: false,
              pauseCleanUpError: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterPauseFailure,
      };

    case pauseCleanUpAction.pauseCleanUpReset:
      const deckPauseResetName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      let dockAfterPauseReset = state.decks.map((deckObj) => {
        return deckObj.name === deckPauseResetName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckPauseResetName
              ),
              pauseCleanUpError: null,
              leftActionBtn: DECKCARD_BTN.text.startUv,
              rightActionBtn: DECKCARD_BTN.text.cancel,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterPauseReset,
      };

    case resumeCleanUpAction.resumeCleanUpInitiated:
      const deckResumeInitiateName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      let dockAfterResumeInit = state.decks.map((deckObj) => {
        return deckObj.name === deckResumeInitiateName
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckResumeInitiateName
              ),
              isLoading: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterResumeInit,
      };

    case resumeCleanUpAction.resumeCleanUpSuccess:
      const deckResumeSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterResumeSuccess = state.decks.map((deckObj) => {
        return deckObj.name === deckResumeSuccessName
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckResumeSuccessName
              ),
              isLoading: false,
              resumeCleanUpError: false,
              leftActionBtn: DECKCARD_BTN.text.pauseUv,
              rightActionBtn: DECKCARD_BTN.text.abort,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterResumeSuccess,
      };

    case resumeCleanUpAction.resumeCleanUpFailed:
      const deckResumeFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterResumeFailure = state.decks.map((deckObj) => {
        return deckObj.name === deckResumeFailureName
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckResumeFailureName
              ),
              isLoading: false,
              resumeCleanUpError: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterResumeFailure,
      };

    case resumeCleanUpAction.resumeCleanUpReset:
      const deckResumeResetName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      let dockAfterResumeReset = state.decks.map((deckObj) => {
        return deckObj.name === deckResumeResetName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckResumeResetName
              ),
              resumeCleanUpError: null,
              leftActionBtn: DECKCARD_BTN.text.startUv,
              rightActionBtn: DECKCARD_BTN.text.cancel,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterResumeReset,
      };

    case abortCleanUpAction.abortCleanUpInitiated:
      const deckAbortInitiateName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      let dockAfterAbortInit = state.decks.map((deckObj) => {
        return deckObj.name === deckAbortInitiateName
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckAbortInitiateName
              ),
              isLoading: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterAbortInit,
      };

    case abortCleanUpAction.abortCleanUpSuccess:
      const deckAbortSuccessName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterAbortSuccess = state.decks.map((deckObj) => {
        return deckObj.name === deckAbortSuccessName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckAbortSuccessName
              ),

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
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterAbortSuccess,
      };

    case abortCleanUpAction.abortCleanUpFailed:
      const deckAbortFailureName =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      let dockAfterAbortFailure = state.decks.map((deckObj) => {
        return deckObj.name === deckAbortFailureName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckAbortFailureName
              ),
              isLoading: false,
              abortCleanUpError: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterAbortFailure,
      };

    case abortCleanUpAction.abortCleanUpReset:
      const deckAbortResetName =
        action.payload.params.deckName === "A"
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      let dockAfterAbortReset = state.decks.map((deckObj) => {
        return deckObj.name === deckAbortResetName
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckAbortResetName
              ),
              abortCleanUpError: null,
              leftActionBtn: DECKCARD_BTN.text.startUv,
              rightActionBtn: DECKCARD_BTN.text.cancel,
              hours: 0,
              mins: 0,
              secs: 0,
              progress: 0,
              cleanUpData: null,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterAbortReset,
      };

    case cleanUpHourActions.setHours:
      let deckNameToSetHours = action.payload.params.deckName;
      let newHours = parseInt(action.payload.params.hours);

      let dockAfterHoursSet = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToSetHours
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckNameToSetHours
              ),
              hours: newHours,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterHoursSet,
      };

    case cleanUpMinsActions.setMins:
      let deckNameToSetMins = action.payload.params.deckName;
      let newMins = parseInt(action.payload.params.mins);

      let dockAfterMinsSet = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToSetMins
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckNameToSetMins
              ),
              mins: newMins,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterMinsSet,
      };

    case cleanUpSecsActions.setSecs:

      let deckNameToSetSecs = action.payload.params.deckName;
      let newSecs = parseInt(action.payload.params.secs);

      let dockAfterSecsSet = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToSetSecs
          ? {
              ...state.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckNameToSetSecs
              ),
              secs: newSecs,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterSecsSet,
      };

    case setShowCleanUpAction.setShowCleanUp:
      let deckNameToShowCleanUp = action.payload.params.deckName;
      let dockAfterShowCleanUp = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToShowCleanUp
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckNameToShowCleanUp
              ),
              showCleanUp: true,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterShowCleanUp,
      };

    case setShowCleanUpAction.resetShowCleanUp:
      let deckNameToHideCleanUp = action.payload.params.deckName;
      let dockAfterHideCleanUp = state.decks.map((deckObj) => {
        return deckObj.name === deckNameToHideCleanUp
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckNameToHideCleanUp
              ),
              showCleanUp: false,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: dockAfterHideCleanUp,
      };

    default:
      return state;
  }
};
