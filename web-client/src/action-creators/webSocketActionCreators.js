import webSocketActions from 'actions/webSocketActions';

export const webSocketOpened = () => ({ type: webSocketActions.onOpen });
export const webSocketError = () => ({ type: webSocketActions.onError });
export const webSocketClosed = () => ({ type: webSocketActions.onClose });

export const webSocketMessage = data => ({
	type: webSocketActions.onMessage,
	payload: {
		data,
	},
});
