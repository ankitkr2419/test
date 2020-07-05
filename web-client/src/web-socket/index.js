import { w3cwebsocket as W3CWebSocket } from 'websocket';
import {
	webSocketMessage,
	webSocketOpened,
	webSocketClosed,
	webSocketError,
} from 'action-creators/webSocketActionCreators';
import { WS_HOST_URL } from '../constants';

let webSocket = null;

export const connectSocket = (dispatch) => {
	webSocket = new W3CWebSocket(`${WS_HOST_URL}/monitor`);
	webSocket.onopen = (event) => {
		console.info('socket connection opened', event);
		dispatch(webSocketOpened());
	};
	webSocket.onmessage = (event) => {
		if (event.data) {
			dispatch(webSocketMessage(JSON.parse(event.data)));
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
