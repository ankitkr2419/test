import { createStore, applyMiddleware } from 'redux';
import createSagaMiddleware from 'redux-saga';

import rootReducer from './reducers';
import rootSaga from './sagas';


const bindMiddleware = (middleware) => {
	if (process.env.NODE_ENV !== 'production') {
		// eslint-disable-next-line global-require
		const { composeWithDevTools } = require('redux-devtools-extension');
		return composeWithDevTools(applyMiddleware(...middleware));
	}
	return applyMiddleware(...middleware);
};

const sagaMiddleware = createSagaMiddleware();
const store = createStore(
	rootReducer,
	bindMiddleware([sagaMiddleware]),
);

sagaMiddleware.run(rootSaga);

export default store;
