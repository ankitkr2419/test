import { fromJS } from 'immutable';
import webSocketActions from 'actions/webSocketActions';

const socketInitialState = fromJS({
	isOpened: false,
	isClosed: false,
	isError: false,
});

export const socketReducer = (state = socketInitialState, action) => {
	switch (action.type) {
	case webSocketActions.onOpen:
		return state.setIn(['isOpened'], true);
	case webSocketActions.onClose:
		return state.setIn(['isClosed'], true);
	case webSocketActions.onError:
		return state.setIn(['isError'], true);
	default:
		return state;
	}
};
