import { w3cwebsocket as W3CWebSocket } from 'websocket';
import {
	webSocketOpened,
	webSocketClosed,
	webSocketError,
} from 'action-creators/webSocketActionCreators';
import { WS_HOST_URL, SOCKET_MESSAGE_TYPE } from 'appConstants';
import { updateWellThroughSocket } from 'action-creators/wellActionCreators';
import { wellGraphSucceeded } from 'action-creators/wellGraphActionCreators';
<<<<<<< HEAD
import { experimentedCompleted, experimentedFailed } from 'action-creators/runExperimentActionCreators';
import { showErrorModal } from 'action-creators/modalActionCreators';
=======
import { experimentedCompleted } from 'action-creators/runExperimentActionCreators';
import { temperatureDataSucceeded } from 'action-creators/temperatureGraphActionCreators';
import { tempData } from 'utils/temperatureData';
>>>>>>> Actions, Action creators and reducer added for temperature graph

let webSocket = null;
export const connectSocket = (dispatch) => {
	webSocket = new W3CWebSocket(`${WS_HOST_URL}/monitor`);
	webSocket.onopen = (event) => {
		console.info('socket connection opened', event);
		dispatch(webSocketOpened());
	};
	webSocket.onmessage = (event) => {
		if (event.data) {
			const { type, data } = JSON.parse(event.data);
			switch (type) {
			case SOCKET_MESSAGE_TYPE.graphData:
				dispatch(wellGraphSucceeded(data));
				break;
			case SOCKET_MESSAGE_TYPE.wellsData:
				dispatch(updateWellThroughSocket(data));
				break;
			case SOCKET_MESSAGE_TYPE.temperatureData:
				dispatch(temperatureDataSucceeded(data));
				break;
			case SOCKET_MESSAGE_TYPE.success:
				dispatch(experimentedCompleted(data));
				break;
			case SOCKET_MESSAGE_TYPE.failure:
				dispatch(experimentedFailed(data));
				break;
			// TODO after discussion with shailesh
			// case SOCKET_MESSAGE_TYPE.ErrorPCRAborted:
			// case SOCKET_MESSAGE_TYPE.ErrorPCRStopped:
			case SOCKET_MESSAGE_TYPE.ErrorPCRMonitor:
			case SOCKET_MESSAGE_TYPE.ErrorPCRDead:
				dispatch(showErrorModal(data));
				break;
			default:
				break;
			}
		}
	};

	webSocket.onclose = (event) => {
		console.info('socket connection closed');
		dispatch(webSocketClosed());
	};

	webSocket.onError = (event) => {
		console.error('socket error : ', event);
		dispatch(webSocketError());
	};
};
export const disConnectSocket = () => {
	if (webSocket !== null) {
		webSocket.close();
	}
};
