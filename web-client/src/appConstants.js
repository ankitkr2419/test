export const API_HOST_URL = process.env.REACT_APP_API_HOST_URL;
export const WS_HOST_URL = process.env.REACT_APP_WS_HOST_URL;
export const API_HOST_VERSION = process.env.REACT_APP_API_HOST_VERSION;
// Target capacity is used to restrict selection of targets
export const TARGET_CAPACITY = process.env.REACT_APP_TARGET_CAPACITY || 6;

export const ROOT_URL_PATH = "/";

export const HTTP_METHODS = {
  GET: "GET",
  POST: "POST",
  PUT: "PUT",
  DELETE: "DELETE",
};

export const TARGET_LINE_COLORS = [
  "#F590B2",
  "#7986CB",
  "#4FC4F7",
  "#9D27B0",
  "#F3811F",
  "#EFD600",
];

export const SOCKET_MESSAGE_TYPE = {
  wellsData: "Wells",
  graphData: "Graph",
  success: "Success",
  failure: "Fail",
  ErrorPCRAborted: "ErrorPCRAborted",
  ErrorPCRStopped: "ErrorPCRStopped",
  ErrorPCRMonitor: "ErrorPCRMonitor",
  ErrorPCRDead: "ErrorPCRDead",
  temperatureData: "Temperature",
  homingProgress: "PROGRESS_HOMING",
  homingSuccess: "SUCCESS_HOMING",
  runRecipeProgress: "PROGRESS_RECIPE",
  runRecipeSuccess: "SUCCESS_RECIPE",
  uvLightProgress: "PROGRESS_UVLIGHT",
  uvLightSuccess: "SUCCESS_UVLIGHT",
  discardTipProgress: "DISCARD_TIP_PROGRESS",
  discardTipSuccess: "DISCARD_TIP_SUCCESS",
  ErrorExtractionMonitor: "ErrorExtractionMonitor",
};

export const EXPERIMENT_STATUS = {
  running: "running",
  runFailed: "run-failed",
  stopped: "stopped",

  // socket
  success: "success",
  failed: "failed",
};

export const ROUTES = {
  landing: "landing",
  splashScreen: "splashscreen",
  recipeListing: "recipe-listing",
  labware: "labware",
  selectProcess: "select-process",
};

export const API_ENDPOINTS = {
  homing: "homing/",
  run: "run",
  pause: "pause",
  resume: "resume",
  abort: "abort",
  discardDeck: "discard-box/cleanup",
  restoreDeck: "restore-deck",
  recipeListing: "recipes",
  saveAndUpdateRecipes: "recipes",
  discardTipAndHoming: "discard-tip-and-home",
  cleanUp: "uv",
  tipsTubes: "tips-tubes",
  cartridge: "cartridges",
  tubes: "tube",
  tips: "tip",
  stepRun: "step-run",
  runNextStep: "run-next-step",
};

export const MODAL_MESSAGE = {
  setPosition: "Please check the position of tip and magnet!",
  homingConfirmation: "Homing Confirmation",
  experimentSuccess: "Experiment was successful",
  abortConfirmation: "Are you sure you want to abort now?",
  abortCleanupConfirmation: "Are you sure you want to Abort Cleanup?",
  uvSuccess: "UV Clean Up was successful",
  publishConfirmation: "Are you sure you want to Publish this recipe?",
};

export const MODAL_BTN = {
  okay: "Okay",
  cancel: "Cancel",
  next: "Next",
  yes: "Yes",
  no: "No",
};

export const DECKCARD_BTN = {
  text: {
    run: "RUN",
    abort: "ABORT",
    cancel: "CANCEL",
    pause: "PAUSE",
    done: "DONE",
    resume: "RESUME",
    startUv: "START UV",
    pauseUv: "PAUSE UV",
    resumeUv: "RESUME UV",
    next: "NEXT",
  },
  icon: {
    run: "play",
    startUv: "play",
    abort: "abort",
    cancel: "cancel",
    pause: "pause",
    done: "done",
    resume: "resume",
    next: "next",
  },
};

export const DECKNAME = {
  DeckA: "Deck A",
  DeckB: "Deck B",
  DeckAShort: "A",
  DeckBShort: "B",
};

export const USER_ROLES = {
  ADMIN: "admin",
  OPERATOR: "operator",
};
export const TOAST_MESSAGE = {
  error: "Something went wrong!",
};

export const LABWARE_ITEMS_NAME = [
  "Tips",
  "Tip Piercing",
  "Deck Position 1",
  "Deck Position 2 ",
  "Cartidge 1",
  "Deck Position 3",
  "Cartidge 2",
  "Deck Position 4",
];

// do not change the order!
export const LABWARE_INITIAL_STATE = {
  tips: {
    isTicked: false,
    processDetails: {
      tipPosition1: { id: null, label: null },
      tipPosition2: { id: null, label: null },
      tipPosition3: { id: null, label: null },
    },
  },
  tipPiercing: {
    isTicked: false,
    processDetails: { position1: { id: null }, position2: { id: null } },
  },
  deckPosition1: {
    isTicked: false,
    processDetails: { tubeType: { id: null, label: null } },
  },
  deckPosition2: {
    isTicked: false,
    processDetails: { tubeType: { id: null, label: null } },
  },
  cartridge1: {
    isTicked: false,
    processDetails: { cartridgeType: { id: null, label: null } },
  },
  deckPosition3: {
    isTicked: false,
    processDetails: { tubeType: { id: null, label: null } },
  },
  cartridge2: {
    isTicked: false,
    processDetails: { cartridgeType: { id: null, label: null } },
  },
  deckPosition4: {
    isTicked: false,
    processDetails: { tubeType: { id: null, label: null } },
  },
};

export const RUN_RECIPE_TYPE = {
  CONTINUOUS_RUN: 0,
  STEP_RUN: 1,
};

export const LABWARE_NAME = {
  tips: "Tips",
  tipPosition1: "Tip Position 1",
  tipPosition2: "Tip Position 2",
  tipPosition3: "Tip Position 3",
  tipPiercing: "Tip Piercing",
  position1: "Position 1",
  position2: "Position 2",
  deckPosition1: "Deck Position 1",
  deckPosition2: "Deck Position 2",
  deckPosition3: "Deck Position 3",
  deckPosition4: "Deck Position 4",
  cartridge1: "Cartridge 1",
  cartridge2: "Cartridge 2",
  cartridgeType: "Cartridge Type",
};

/**
 * get process icon_name associated with processType
 * */
export const PROCESS_ICON_CONSTANTS = [
  //if processType not found, use default process icon
  {
    processType: "default",
    iconName: "process",
  },

  //other icons:
  {
    processType: "Piercing",
    iconName: "piercing",
  },
  {
    processType: "AspireDispense",
    iconName: "aspire-dispense",
  },
  {
    processType: "Heating",
    iconName: "heating",
  },
  {
    processType: "Shaking",
    iconName: "shaking",
  },
  {
    processType: "Delay",
    iconName: "delay",
  },
  /**TODO:
   * following icons not available in fonts.scss
   */
  {
    processType: "TipOperation",
    iconName: "",
  },
  {
    processType: "TipDocking",
    iconName: "",
  },
  {
    processType: "AttachDetach",
    iconName: "",
  },
];
