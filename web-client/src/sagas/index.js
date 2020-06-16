import { all } from 'redux-saga/effects';
import {
	createTemplateSaga,
	fetchTemplatesSaga,
	createTemplateSuccessSaga,
	deleteTemplateSaga,
} from 'sagas/templateSaga';
import {
	saveTargetSaga,
	fetchMasterTargetsSaga,
	fetchTargetsByTemplateIDSaga,
} from 'sagas/targetSaga';

const allSagas = [
	createTemplateSaga(),
	fetchTemplatesSaga(),
	createTemplateSuccessSaga(),
	saveTargetSaga(),
	fetchMasterTargetsSaga(),
	fetchTargetsByTemplateIDSaga(),
	deleteTemplateSaga(),
];

export default function* rootSaga() {
	yield all(allSagas);
}
