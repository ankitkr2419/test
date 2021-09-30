import {
  saveNewRecipeAction,
  getTipsAndTubesAction,
  getCartridgeAction,
  getTubesAction,
  getTipsAction,
  storeExistingRecipeAction,
  getCartridgeByIdAction,
  getCartridge1Action,
  getCartridge2Action,
} from "actions/saveNewRecipeActions";

const cartridgeDetailsInitialState = {
  isLoading: false,
  error: null,
  cartridgeDetails: null,
};

export const cartridge1DetailsReducer = (
  state = cartridgeDetailsInitialState,
  actions
) => {
  switch (actions.type) {
    case getCartridge1Action.getCartridge1Initiated:
      return {
        ...state,
        isLoading: true,
        ...actions.payload,
      };

    case getCartridge1Action.getCartridge1Success:
      return {
        ...state,
        isLoading: false,
        error: false,
        cartridgeDetails: actions.payload.response,
      };

    case getCartridge1Action.getCartridge1Failure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getCartridge1Action.getCartridge1Reset:
      return cartridgeDetailsInitialState;

    default:
      return state;
  }
};

export const cartridge2DetailsReducer = (
  state = cartridgeDetailsInitialState,
  actions
) => {
  switch (actions.type) {
    case getCartridge2Action.getCartridge2Initiated:
      return {
        ...state,
        isLoading: true,
      };

    case getCartridge2Action.getCartridge2Success:
      return {
        ...state,
        isLoading: false,
        error: false,
        cartridgeDetails: actions.payload.response,
      };

    case getCartridge2Action.getCartridge2Failure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getCartridge2Action.getCartridge2Reset:
      return cartridgeDetailsInitialState;

    default:
      return state;
  }
};

const initialState = {
  tempDeckName: "",
  isLoading: null,
  error: null,
  isSuccess: null,
  recipeDetails: {
    id: "",
    name: "",
    deckName: "",
  },
  tipsAndTubesOptions: null,
  tubesOptions: null,
  tipsOptions: null,
  cartridgeOptions: null,
  isSaved: false,
  errorInSaving: false,
  token: "",
};

export const updateRecipeDetailsReducer = (state = initialState, actions) => {
  switch (actions.type) {
    //saving recipe name
    case saveNewRecipeAction.saveRecipeName:
      return {
        ...state,
        ...actions.payload,
      };

    //updating new recipe : init
    case saveNewRecipeAction.updateRecipeInitiated:
      const token = actions.payload.token;

      return {
        ...state,
        isLoading: true,
        token: token,
      };

    case saveNewRecipeAction.updateRecipeSuccess:
      return {
        ...state,
        recipeDetails: {
          ...state.recipeDetails,
          id: actions.payload.response.id,
        },
        isLoading: false,
        error: false,
        isSuccess: true,
      };

    case saveNewRecipeAction.updateRecipeFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
        isSuccess: false,
      };

    case saveNewRecipeAction.updateRecipeReset:
      return {
        ...state,
        isLoading: null,
        error: null,
        isSuccess: null,
      };

    //get tips and tubes options : init
    case getTipsAndTubesAction.getTipsAndTubesInitiated:
      return {
        ...state,
        isLoading: true,
      };

    case getTipsAndTubesAction.getTipsAndTubesSuccess:
      return {
        ...state,
        tipsAndTubesOptions: actions.payload.response,
        isLoading: false,
        error: false,
      };

    case getTipsAndTubesAction.getTipsAndTubesFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getTipsAndTubesAction.getTipsAndTubesReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    //get cartridge options : init
    case getCartridgeAction.getCartridgeInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getCartridgeAction.getCartridgeSuccess:
      return {
        ...state,
        cartridgeOptions: actions.payload.response,
        isLoading: false,
        error: false,
      };

    case getCartridgeAction.getCartridgeFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getCartridgeAction.getCartridgeReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    // Tips
    case getTipsAction.getTipsInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getTipsAction.getTipsSuccess:
      return {
        ...state,
        tipsOptions: actions.payload.response,
        isLoading: false,
        error: false,
      };

    case getTipsAction.getTipsFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getTipsAction.getTipsReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    //tubes options
    case getTubesAction.getTubesInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getTubesAction.getTubesSuccess:
      return {
        ...state,
        tubesOptions: actions.payload.response,
        isLoading: false,
        error: false,
      };

    case getTubesAction.getTubesFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getTubesAction.getTubesReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    default:
      return state;
  }
};
