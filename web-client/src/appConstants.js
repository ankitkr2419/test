export const API_HOST_URL = process.env.REACT_APP_API_HOST_URL;
export const WS_HOST_URL = process.env.REACT_APP_WS_HOST_URL;
export const API_HOST_VERSION = process.env.REACT_APP_API_HOST_VERSION;

// Target capacity is used to restrict selection of targets
export const TARGET_CAPACITY = process.env.REACT_APP_TARGET_CAPACITY || 6;

export const ROOT_URL_PATH = "/";

export const EMAIL_REGEX = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
export const EMAIL_REGEX_OR_EMPTY_STR = /^$|^.*@.*\..*$/;
export const NAME_REGEX = /^[^\\\/&]*$/;

export const CREDS_FOR_HOMING = {
  email: "main",
  password: "main",
  role: "admin",
  deckName: "",
  showToast: false,
};

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
  ErrorPCR: "ErrorPCR",
  temperatureData: "Temperature",
  PIDProgress: "PROGRESS_PID",
  PIDSuccess: "SUCCESS_PID",
  rtpcrProgress: "RTPCR_PROGRESS",
  rtpcrSuccess: "RTPCR_SUCCESS",
  homingProgress: "PROGRESS_HOMING",
  homingSuccess: "SUCCESS_HOMING",
  runRecipeProgress: "PROGRESS_RECIPE",
  runRecipeSuccess: "SUCCESS_RECIPE",
  uvLightProgress: "PROGRESS_UVLIGHT",
  uvLightSuccess: "SUCCESS_UVLIGHT",
  discardTipProgress: "DISCARD_TIP_PROGRESS",
  discardTipSuccess: "DISCARD_TIP_SUCCESS",
  ErrorExtractionMonitor: "ErrorExtractionMonitor",
  progressHeater: "PROGRESS_HEATER",
  progressPidTuning: "PROGRESS_SHAKERPIDTUNING",
  successPidTuning: "SUCCESS_SHAKERPIDTUNING",
  PROGRESSLidPIDTuning: "PROGRESS_LidPIDTuning",
  ErrorPIDTuning: "ErrorPIDTuning",
  SUCCESSLidPIDTuning: "SUCCESS_LidPIDTuning",
  progressShakerRun: "PROGRESS_SHAKERRUN",
  successShakerRun: "SUCCESS_SHAKERRUN",
  progressHeaterRun: "PROGRESS_HEATERRUN",
  successHeaterRun: "SUCCESS_HETERRUN",
  ErrorOperationAborted: "ErrorOperationAborted",
  progressDyeCalibration: "PROGRESS_OPTCALIB",
  completedDyeCalibration: "SUCCESS_OPTCALIB",
};

export const HOMING_STATUS = {
  progressStarted: "progressStarted",
  progressing: "progressing",
  progressComplete: "progressComplete",
  progressFailed: "progressFailed",
};

export const CLEAN_UP_STATUS = {
  aborted: "aborted",
  aborting: "aborting",
  abortFailed: "abortFailed",

  progressing: "progressing",
  progressComplete: "progressComplete",
};

export const HEATER_STATUS = {
  progressing: "progressing",
  progressComplete: "progressComplete",
};

export const PID_STATUS = {
  running: "running",
  runFailed: "run-failed",
  stopped: "stopped",
  aborted: "aborted",

  aborting: "aborting",
  abortFailed: "abortFailed",

  progressing: "progressing",
  progressComplete: "progressComplete",
};

export const SHAKER_RUN_STATUS = {
  progressing: "progressing",
  progressComplete: "progressComplete",
  progressAborted: "progressAborted",
};

export const HEATER_RUN_STATUS = {
  progressing: "progressing",
  progressComplete: "progressComplete",
  progressAborted: "progressAborted",
};

export const EXPERIMENT_STATUS = {
  running: "running",
  runFailed: "run-failed",
  stopped: "stopped",

  //while running experiment
  progressing: "progressing",
  progressComplete: "progressComplete",

  // socket
  success: "success",
  failed: "failed",
};

export const ROUTES = {
  login: "login", //rtpcr login page
  landing: "landing", //extraction login page
  splashScreen: "splashscreen", //application homepage
  recipeListing: "recipe-listing",
  labware: "labware",
  processListing: "process-listing",
  selectProcess: "select-process",
  piercing: "piercing",
  tipPickup: "tip-pickup",
  aspireDispense: "aspire-dispense",
  shaking: "shaking",
  heating: "heating",
  magnet: "magnet",
  tipDiscard: "tip-discard",
  delay: "delay",
  templates: "templates",
  plate: "plate",
  activity: "activity",
  tipPosition: "tip-position",
  calibration: "calibration", //rtpcr flow: engineer homepage
  users: "users", //manage users
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
  cartridges: "cartridges",
  cartridge: "cartridge",
  tubes: "tube",
  tips: "tip",
  stepRun: "step-run",
  runNextStep: "run-next-step",
  tipOperation: "tip-operation",
  login: "login",
  logout: "logout",
  piercing: "piercing",
  aspireDispense: "aspire-dispense",
  shaking: "shaking",
  heating: "heating",
  recipe: "recipe",
  tipDiscard: "tip-operation",
  magnet: "attach-detach",
  duplicateProcess: "duplicate-process",
  duplicateRecipe: "duplicate-recipe",
  attachDetach: "attach-detach",
  tipDocking: "tip-docking",
  delay: "delay",
  rearrangeProcesses: "rearrange-processes",
  processes: "processes",
  appInfo: "app-info",
  experiments: "experiments",
  configs: "configs",
  pidCalibration: "pid-calibration",
  pidUpdate: "configs/extraction",
  manual: "manual",
  motor: "motor",
  senseAndHit: "calibrations",
  calibrations: "calibrations",
  startShaking: "start-shaking",
  startHeating: "start-heating",
  emailReport: "email-report",
  graphUpdate: "rtpcr/graph-update-scale",
  uploadReport: "upload-report",
  emission: "emission",
  experiments: "experiments",
  temperature: "temperature",
  setThreshold: "set-threshold",
  getBaseline: "get-baseline",
  tipTube: "tiptube",
  rtpcrConfigs: "configs/rtpcr",
  tecConfigs: "configs/tec",
  lidPidStart: "lid/pid-calibration/start",
  lidPidStop: "/lid/pid-calibration/stop",
  resetTEC: "tec/reset-device",
  autoTuneTEC: "tec/auto-tune",
  dyes: "dyes",
  consumable: "consumable-distance",
  dyeCalibration: "optical-caliberation",
  users: "users",
  whiteLight: "light",
  shutdown: "shutdown",
};

export const MODAL_MESSAGE = {
  runConfirmMsg: "Wells are not configured. Are you sure you want to proceed?",
  abortExpInfo: "Can't logout while experiment is still running.",
  abortExpWarning: "Are you sure you want to abort experiment?",
  setPosition: "Please check the position of tip and magnet!",
  homingConfirmation: "Homing Confirmation",
  experimentSuccess: "Experiment was successful",
  experimentAborted: "Experiment was aborted",
  abortConfirmation: "Are you sure you want to abort now?",
  abortCleanupConfirmation: "Are you sure you want to Abort Cleanup?",
  uvSuccess: "UV Clean Up was successful",
  publishConfirmation: "Are you sure you want to Publish this recipe?",
  unpublishConfirmation: "Are you sure you want to Unpublish this recipe?",
  finishProcessListConfirmation: "Are you sure you want to save changes to ",
  deleteProcessConfirmation: "Are you sure you want to delete this process?",
  deleteRecipeConfirmation: "Are you sure you want to delete this recipe?",
  exitConfirmation: "Are you sure you want to exit?",
  deleteTemplateConfirmation: "Are you sure you want to delete this template?",
  deleteStepConfirmation: "Are you sure you want to delete this step?",
  deleteActivityConfirmation: "Are you sure you want to delete this activity?",
  backConfirmation: "Are you sure you want to go back?",
  forgotPasswordMsg: "Contact admin to change your password",
  senseAndHitHomingMsg:
    "System will now be homed to verify that the system is calibrated",
  logoutConformation: "Are you sure you want to logout?",
};

export const MODAL_BTN = {
  okay: "Okay",
  cancel: "Cancel",
  next: "Next",
  yes: "Yes",
  no: "No",
  viewResults: "View Results",
  complete: "Complete",
  withHoming: "With Homing",
  withoutHoming: "Without Homing",
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
  ENGINEER: "engineer",
};
export const TOAST_MESSAGE = {
  error: "Something went wrong!",
  calRedirect: "Cannot redirect to calibration while adding/editing processes!",
  deckBlockForProcess:
    "Decks cannot be switched while adding/editing processes!",
  deckBlockForCalibration: "Decks cannot be switched in calibration!",
  sendingMailSuccess: "Mail sent successfully!",
  sendingMailFailure: "Something went wrong!",
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

export const SELECT_PROCESS_PROPS = [
  // if processType not found, use default process icon
  {
    processName: "Process Name",
    processType: "default",
    iconName: "process",
    route: "",
    apiEndpoint: null,
  },

  // process properties
  {
    iconName: "piercing",
    processType: "Piercing",
    processName: "Piercing",
    route: ROUTES.piercing,
    apiEndpoint: API_ENDPOINTS.piercing,
  },
  {
    iconName: "tip-pickup",
    processType: "TipPickup",
    processName: "Tip Pickup",
    route: ROUTES.tipPickup,
    apiEndpoint: API_ENDPOINTS.tipOperation,
  },
  {
    iconName: "aspire-dispense",
    processType: "AspireDispense",
    processName: "Aspire & Dispense",
    route: ROUTES.aspireDispense,
    apiEndpoint: API_ENDPOINTS.aspireDispense,
  },
  {
    iconName: "shaking",
    processType: "Shaking",
    processName: "Shaking",
    route: ROUTES.shaking,
    apiEndpoint: API_ENDPOINTS.shaking,
  },
  {
    iconName: "heating",
    processType: "Heating",
    processName: "Heating",
    route: ROUTES.heating,
    apiEndpoint: API_ENDPOINTS.heating,
  },
  {
    iconName: "magnet",
    processType: "AttachDetach",
    processName: "Magnet",
    route: ROUTES.magnet,
    apiEndpoint: API_ENDPOINTS.attachDetach,
  },
  {
    iconName: "tip-discard",
    processType: "TipDiscard",
    processName: "Tip Discard",
    route: ROUTES.tipDiscard,
    apiEndpoint: API_ENDPOINTS.tipOperation,
  },
  {
    iconName: "delay",
    processType: "Delay",
    processName: "Delay",
    route: ROUTES.delay,
    apiEndpoint: API_ENDPOINTS.delay,
  },
  {
    iconName: "tip-position",
    processType: "TipDocking",
    processName: "Tip Position",
    route: ROUTES.tipPosition,
    apiEndpoint: API_ENDPOINTS.tipDocking,
  },
];

export const ASPIRE_DISPENSE_SIDEBAR_LABELS = [
  "Cartridge 1",
  "Cartridge 2",
  "Shaker",
  "Deck Position",
];

// for testing, will be removed
export const TEST_TOKEN =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MjE0MjMwMTIsInN1YiI6ImFkbWluIiwicm9sZSI6ImFkbWluIiwiZGVjayI6IkEiLCJhdXRoX2lkIjoiOWFjYTYxMWMtODJkZS00MzJkLWIxNGQtMWQwZjM2MmQ3OTYyIn0.5xvpGAhljqk2cKrmfIEJvmFwHm0bVuNZUEXG2zs9nF0";

export const TEST_RECIPE_ID = "28585f66-8aa7-4e55-bff9-d0fb0240a147";

export const ASPIRE_DISPENSE_DECK_POS_OPTNS = [
  { value: "6", label: "6" },
  { value: "7", label: "7" },
  { value: "8", label: "8" },
  { value: "9", label: "9" },
  { value: "10", label: "10" },
  { value: "11", label: "11" },
];

export const TIP_PICKUP_PROCESS_OPTIONS = [
  { value: "1", label: "1" },
  { value: "2", label: "2" },
  { value: "3", label: "3" },
  { value: "4", label: "4" },
  { value: "5", label: "5" },
];

export const CATEGORY_NAME = {
  1: "well",
  2: "well",
  3: "shaker",
  4: "deck",
};

export const CATEGORY_LABEL = {
  1: "Cartridge 1",
  2: "Cartridge 2",
  3: "Shaker",
  4: "Deck Position",
};

export const APP_TYPE = {
  COMBINED: "combined",
  RTPCR: "rtpcr",
  EXTRACTION: "extraction",
  NONE: "none",
};

// common constants used for aspire-dispense process and tip-position process
export const NUMBER_OF_WELLS = 13;

// constants for aspire-dispense process
export const ASPIRE_WELLS = 0;
export const DISPENSE_WELLS = 1;

// constants for tip position process
export const TIP_HEIGHT_MAX_ALLOWED_VALUE = 25;
export const TIP_HEIGHT_MIN_ALLOWED_VALUE = 0;
export const TIP_POSTION_ERROR_MSG =
  "Invalid request sent! Please check input values.";
export const CARTRIDGE_1_WELLS = 0;
export const CARTRIDGE_2_WELLS = 1;
export const TAB_TYPE_CARTRIDGE_1 = "1";
export const TAB_TYPE_DECK = "2";
export const TAB_TYPE_CARTRIDGE_2 = "3";

//constants for shaking process
export const MAX_TEMP_ALLOWED = 120;
export const MIN_TEMP_ALLOWED = 20;
export const MAX_TIME_ALLOWED = 3660; // 1 hour 1 min
export const timeConstants = {
  SEC_IN_ONE_MIN: 60,
  SEC_IN_ONE_HOUR: 3600,
  MIN_IN_ONE_HOUR: 60,
};
export const MAX_RPM_VALUE = 1500;
export const MIN_RPM_VALUE = 800;

/**
 * Cartridge types
 */
export const CARTRIDGE_1 = "CARTRIDGE_1";
export const CARTRIDGE_2 = "CARTRIDGE_2";

/**
 * Maximum number of wells that can be present in a plate.
 * Maximum number of wells are 16.
 * 96 was the old version, this is changed in future implementations.
 */
export const MAX_NO_OF_WELLS = 16;

//constants for RTPCR - templates
export const MIN_VOLUME = 10;
export const MAX_VOLUME = 250;

export const MIN_LID_TEMP = 80;
export const MAX_LID_TEMP = 120;

//constants for motor
export const MIN_MOTOR_NUMBER = 5;
export const MAX_MOTOR_NUMBER = 10;
export const MIN_MOTOR_DISTANCE = 0;
export const MAX_MOTOR_DISTANCE = 100;
export const MIN_MOTOR_DIRECTION = 0;
export const MAX_MOTOR_DIRECTION = 1;
export const MOTOR_NUMBER_OPTIONS = [
  { value: 5, label: "Deck" },
  { value: 6, label: "Magnet Up Down" },
  { value: 7, label: "Magnet Rev For" },
  { value: 9, label: "Syringe Module" },
  { value: 10, label: "Syringe" },
];
export const MOTOR_DIRECTION_OPTIONS = [
  { value: 0, label: "Against sensor" },
  { value: 1, label: "Towards sensor" },
];

//constants for pid
export const MIN_PID_TEMP = 50;
export const MAX_PID_TEMP = 75;
export const MIN_PID_MIN = 0;
export const MAX_PID_MIN = 9999;

// constants for engineer's flow for extraction
export const MAX_ROOM_TEMPERATURE = 30;
export const MIN_ROOM_TEMPERATURE = 20;

//engineer's flow tips & tubes constants
export const MIN_TIPTUBE_ID = 0;
export const MAX_TIPTUBE_ID = 9999;
export const MIN_TIPTUBE_VOLUME = 0;
export const MAX_TIPTUBE_VOLUME = 9999;
export const MIN_TIPTUBE_HEIGHT = 0;
export const MAX_TIPTUBE_HEIGHT = 9999;
export const MIN_TIPTUBE_TTBASE = 0;
export const MAX_TIPTUBE_TTBASE = 9999;

//engineer's flow cartridges constants
export const MAX_WELLS_COUNT = 13;
export const MIN_WELLS_COUNT = 1;
export const MAX_CARTRIDGE_ID = 15;
export const MIN_CARTRIDGE_ID = 1;
export const CARTRIDGE_TYPE_OPTIONS = [
  { value: "Cartridge 1", label: "Cartridge 1" },
  { value: "Cartridge 2", label: "Cartridge 2" },
];
export const CARTRIDGE_WELLS = {
  MAX_DISTANCE: 87,
  MIN_DISTANCE: 0,
  MAX_VOLUME: 5000,
  MIN_VOLUME: 10,
  MAX_HEIGHT: 40,
  MIN_HEIGHT: 1,
};

export const MAX_TOLERANCE_ALLOWED = 100;
export const MIN_TOLERANCE_ALLOWED = 0;

export const TEMPERATURE_GRAPH_OPTIONS = {
  legend: {
    display: false,
  },
  animation: false,
  scales: {
    xAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Time (minutes)",
          fontSize: 15,
          fontStyle: "bold",
          padding: 10,
        },
        offset: true,
        type: "linear",
        ticks: {
          source: "data",
          beginAtZero: true,
          suggestedMin: 0,
          min: 0,
          fontSize: 15,
          fontStyle: "bold",
        },
      },
    ],
    yAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Temperature (??C)",
          fontSize: 15,
          fontStyle: "bold",
          padding: 10,
        },
        ticks: {
          fontSize: 15,
          fontStyle: "bold",
        },
      },
    ],
  },
};

// default min values for amplification plot
export const DEFAULT_MIN_VALUE = {
  yAxisMin: 0,
  xAxisMin: 1,
};

//analyse data graph constants
export const GRAY_COLOR = "rgba(148,147,147,1)";
export const PINK_COLOR = "rgba(245,144,178,1)";
