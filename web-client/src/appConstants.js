export const API_HOST_URL = process.env.REACT_APP_API_HOST_URL;
export const WS_HOST_URL = process.env.REACT_APP_WS_HOST_URL;
export const API_HOST_VERSION = process.env.REACT_APP_API_HOST_VERSION;
// Target capacity is used to restrict selection of targets
export const TARGET_CAPACITY = process.env.REACT_APP_TARGET_CAPACITY || 6;

export const ROOT_URL_PATH = '/';

export const HTTP_METHODS = {
	GET: 'GET',
	POST: 'POST',
	PUT: 'PUT',
	DELETE: 'DELETE',
};

export const TARGET_LINE_COLORS = [
	'#F590B2',
	'#7986CB',
	'#4FC4F7',
	'#9D27B0',
	'#F3811F',
	'#EFD600',
];

export const SOCKET_MESSAGE_TYPE = {
	wellsData : 'Wells',
	graphData: 'Graph',
	success: 'Success',
	failure: 'Fail',
	ErrorPCRAborted: 'ErrorPCRAborted',
	ErrorPCRStopped: 'ErrorPCRStopped',
	ErrorPCRMonitor: 'ErrorPCRMonitor',
	ErrorPCRDead: 'ErrorPCRDead',
	temperatureData: 'Temperature',
};

export const EXPERIMENT_STATUS = {
	running: 'running',
	runFailed: 'run-failed',
	stopped: 'stopped',

	// socket
	success: 'success',
	failed: 'failed',
};

export const ROUTES = {
	landing: "landing",
	splashScreen: "splashscreen"
};

export const API_ENDPOINTS = {
	homing: "homing/",
	run: "run",
	pause: "pause",
	resume: "resume",
	abort: "abort",
	discardDeck: "discard-box/cleanup",
	restoreDeck: "restore-deck",
	recipeListing: "recipes"
};

export const MODAL_MESSAGE = {
	setPosition : "Please check the position of tip and magnet!",
	homingConfirmation: "Homing Confirmation"
};

export const MODAL_BTN = {
	okay: "Okay",
	cancel: "Cancel"
};
