import loginActions from "actions/loginActions";
import {
  runRecipeAction,
  pauseRecipeAction,
  resumeRecipeAction,
  abortRecipeAction,
  recipeListingAction,
  saveRecipeDataAction,
  publishRecipeAction,
  deleteRecipeAction,
  actionBtnStates,
} from "actions/recipeActions";
import { DECKCARD_BTN, DECKNAME, RUN_RECIPE_TYPE } from "appConstants";
import {
  getUpdatedDecks,
  getUpdatedDecksAfterRecipeListChanged,
} from "utils/helpers";

export const initialState = {
  // recipeData: [], //all recipe data
  tempDeckName: "", //used for fetch recipe list
  tempRecipeId: "",
  tempIsPublished: "",
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
      runRecipeType: RUN_RECIPE_TYPE.CONTINUOUS_RUN,
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
      runRecipeType: RUN_RECIPE_TYPE.CONTINUOUS_RUN,
    },
  ],
};

export const recipeActionReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case saveRecipeDataAction.saveRecipeDataForDeck: //set and update: depend on deckName
      let deckNameForRecipe = action.payload.deckName;

      let isAdmin = action.payload.recipeData?.isAdmin
        ? action.payload.recipeData.isAdmin
        : false;

      const saveChangesForDeck = {
        recipeData: action.payload.recipeData,
        showProcess: true,
        runRecipeType: isAdmin
          ? action.payload.recipeData.runRecipeType
          : RUN_RECIPE_TYPE.CONTINUOUS_RUN,
      };
      const newDecksAfterRecipeDataAdded = getUpdatedDecks(
        state,
        deckNameForRecipe,
        saveChangesForDeck
      );
      return {
        ...state,
        decks: newDecksAfterRecipeDataAdded,
      };

    case saveRecipeDataAction.updateRecipeReducerDataForDeck: //update data in a deck
      let deckNameToUpdate = action.payload.deckName;
      const updateChangesForDeck = action.payload.body;
      let newDecksAfterRequiredUpdations = getUpdatedDecks(
        state,
        deckNameToUpdate,
        updateChangesForDeck
      );
      return {
        ...state,
        decks: newDecksAfterRequiredUpdations,
      };

    case saveRecipeDataAction.resetRecipeDataForDeck:
      let deckNameToResetData = action.payload.deckName;

      const resetChanges = {
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
      };
      const newDeckDataAfterReset = getUpdatedDecks(
        state,
        deckNameToResetData,
        resetChanges
      );
      return {
        ...state,
        decks: newDeckDataAfterReset,
      };

    case runRecipeAction.runRecipeInitiated:
      let deckNameToInitiateRun = action.payload.params.deckName;

      const runInitChanges = {
        runRecipeResponse: {
          recipeId: action.payload.params.recipeId,
        },
      };
      const decksAfterRunInitiated = getUpdatedDecks(
        state,
        deckNameToInitiateRun,
        runInitChanges
      );

      return {
        ...state,
        decks: decksAfterRunInitiated,
      };
    case runRecipeAction.runRecipeSuccess:
      let deckNameRunSuccess =
        action.payload.response.deck === DECKNAME.DeckAShort
          ? DECKNAME.DeckA
          : DECKNAME.DeckB;

      const deckRunSuccessChanges = {
        leftActionBtn: DECKCARD_BTN.text.pause,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };
      const decksAfterRunSuccess = getUpdatedDecks(
        state,
        deckNameRunSuccess,
        deckRunSuccessChanges
      );
      return {
        ...state,
        decks: decksAfterRunSuccess,
      };

    case runRecipeAction.runRecipeFailed:
      return { ...state };

    case runRecipeAction.runRecipeReset:
      let deckNameToReset = action.payload.deckName;

      const decksAfterRecipeReset = state.decks.map((deckObj) => {
        let recipeListOfDeckObj = deckObj.allRecipeData;
        return deckObj.name === deckNameToReset
          ? {
              ...initialState.decks.find(
                (initialDeckObj) => initialDeckObj.name === deckNameToReset
              ),
              allRecipeData: recipeListOfDeckObj,
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
        let isStepRun = deckObj.runRecipeType === RUN_RECIPE_TYPE.STEP_RUN;

        //for admin: step-run: if current_step !== old_step then activate next button
        let shouldActivateNextProcess =
          isStepRun &&
          deckObj.runRecipeInProgress?.operation_details?.current_step &&
          deckObj.runRecipeInProgress.operation_details.current_step !==
            response.operation_details.current_step;

        return deckObj.name === deckNameRunInProgress
          ? {
              ...deckObj,
              runRecipeInProgress: {
                ...response,
              },
              ...(shouldActivateNextProcess && {
                leftActionBtn: DECKCARD_BTN.text.next,
              }),
            }
          : deckObj;
      });

      return {
        ...state,
        decks: decksAfterRunInProgress,
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
      };

    case pauseRecipeAction.pauseRecipeInitiated:
      return { ...state /*...action.payload, isLoading: true */ };

    case pauseRecipeAction.pauseRecipeSuccess:
      let responsePauseSuccess = action.payload.response;
      let deckNamePauseSuccess =
        responsePauseSuccess.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const pauseRecipeChanges = {
        pauseRecipeResponse: responsePauseSuccess,
        pauseRecipeError: false,
        leftActionBtn: DECKCARD_BTN.text.resume,
      };

      const decksAfterPauseSuccess = getUpdatedDecks(
        state,
        deckNamePauseSuccess,
        pauseRecipeChanges
      );

      return {
        ...state,
        decks: decksAfterPauseSuccess,
      };

    case pauseRecipeAction.pauseRecipeFailed:
      return { ...state };

    case pauseRecipeAction.pauseRecipeReset:
      return { ...state };

    case abortRecipeAction.abortRecipeInitiated:
      return { ...state };

    case abortRecipeAction.abortRecipeSuccess:
      let responseAbortSuccess = action.payload.response;
      let deckNameAbortSuccess =
        responseAbortSuccess.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const abortRecipeChanges = {
        abortRecipeResponse: responseAbortSuccess,
        abortRecipeError: false,
        leftActionBtn: DECKCARD_BTN.text.done,
        rightActionBtn: DECKCARD_BTN.text.cancel,
        leftActionBtnDisabled: true,
        rightActionBtnDisabled: true,
      };

      const decksAfterAbortSuccess = getUpdatedDecks(
        state,
        deckNameAbortSuccess,
        abortRecipeChanges
      );

      return {
        ...state,
        decks: decksAfterAbortSuccess,
      };

    case abortRecipeAction.abortRecipeFailed:
      return { ...state };

    case abortRecipeAction.abortRecipeReset:
      return { ...state };

    case resumeRecipeAction.resumeRecipeInitiated:
      return { ...state /*...action.payload, isLoading: true*/ };

    case resumeRecipeAction.resumeRecipeSuccess:
      let deckNameResumeSuccess =
        action.payload.response.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const resumeRecipeChanges = {
        leftActionBtn: DECKCARD_BTN.text.pause,
        resumeRecipeError: false,
      };

      const decksAfterResumeSuccess = getUpdatedDecks(
        state,
        deckNameResumeSuccess,
        resumeRecipeChanges
      );

      return {
        ...state,
        decks: decksAfterResumeSuccess,
      };
    case resumeRecipeAction.resumeRecipeFailed:
      return { ...state };

    case resumeRecipeAction.resumeRecipeReset:
      return { ...state };

    case resumeRecipeAction.resumeRecipeInProgress:
      return { ...state };

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
        tempDeckName: action.payload.deckName,
      };

    case recipeListingAction.recipeListingSuccess:
      const recipeListingSuccessChanges = {
        allRecipeData: action.payload.response,
      };

      const newDecksAfterRecipeList = getUpdatedDecks(
        state,
        state.tempDeckName,
        recipeListingSuccessChanges
      );

      return {
        ...state,
        tempDeckName: "",
        decks: newDecksAfterRecipeList,
      };

    case recipeListingAction.recipeListingFailed:
      return { ...state };

    case recipeListingAction.recipeListingReset:
      return { ...state };

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

    case publishRecipeAction.publishRecipeInitiated:
      return {
        ...state,
        tempDeckName: action.payload.params.deckName,
        tempRecipeId: action.payload.params.recipeId,
        tempIsPublished: action.payload.params.isPublished,
        isLoading: true,
      };

    case publishRecipeAction.publishRecipeSuccess:
      const changesInMatchedRecipe = {
        is_published: !state.tempIsPublished,
      };

      const newDecksAfterPublishSuccess = getUpdatedDecksAfterRecipeListChanged(
        state,
        state.tempDeckName,
        state.tempRecipeId,
        changesInMatchedRecipe
      );

      return {
        ...state,
        tempDeckName: "",
        tempRecipeId: "",
        tempIsPublished: "",
        isLoading: false,
        decks: newDecksAfterPublishSuccess,
      };

    case publishRecipeAction.publishRecipeFailed:
      return {
        ...state,
        tempDeckName: "",
        tempRecipeId: "",
        tempIsPublished: "",
        isLoading: false,
      };

    case deleteRecipeAction.deleteRecipeInitiated:
      return {
        ...state,
        tempDeckName: action.payload.params.deckName,
        tempRecipeId: action.payload.params.recipeId,
      };

    case deleteRecipeAction.deleteRecipeSuccess:
      const newDecksAfterDeleteRecipe = getUpdatedDecksAfterRecipeListChanged(
        state,
        state.tempDeckName,
        state.tempRecipeId,
        {}, //changes in matched
        {}, //changes in unmatched
        true //isDeleteRecipe
      );
      return {
        ...state,
        tempDeckName: "",
        tempRecipeId: "",
        decks: newDecksAfterDeleteRecipe,
      };
    case deleteRecipeAction.deleteRecipeFailure:
      return {
        ...state,
        tempDeckName: "",
        tempRecipeId: "",
      };

    // enable disable action btns seperately
    case actionBtnStates.enableActionBtn:
      let enableBtnDeckname = action.payload.deckName;

      let enableBtnChanges = null;
      if (action.payload.isLeftBtn === true) {
        enableBtnChanges = {
          leftActionBtnDisabled: false,
        };
      } else {
        enableBtnChanges = {
          rightActionBtnDisabled: false,
        };
      }

      const newDeckDataAfterEnableBtn = getUpdatedDecks(
        state,
        enableBtnDeckname,
        enableBtnChanges
      );
      return {
        ...state,
        decks: newDeckDataAfterEnableBtn,
      };

    case actionBtnStates.disableActionBtn:
      let disableBtnDeckname = action.payload.deckName;

      let disableBtnChanges = null;
      if (action.payload.isLeftBtn === true) {
        disableBtnChanges = {
          leftActionBtnDisabled: true,
        };
      } else {
        disableBtnChanges = {
          rightActionBtnDisabled: true,
        };
      }

      const newDeckDataAfterDisableBtn = getUpdatedDecks(
        state,
        disableBtnDeckname,
        disableBtnChanges
      );
      return {
        ...state,
        decks: newDeckDataAfterDisableBtn,
      };

    default:
      return state;
  }
};
