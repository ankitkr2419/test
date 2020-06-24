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

import {
	fetchStagesSaga,
	addStageSaga,
	deleteStageSaga,
	updateStageSaga,
} from 'sagas/stageSaga';

import {
	fetchStepsSaga,
	addStepSaga,
	deleteStepSaga,
	updateStepSaga,
} from 'sagas/stepSaga';

import {
	loginSaga,
} from 'sagas/loginSaga';

const allSagas = [
	createTemplateSaga(),
	fetchTemplatesSaga(),
	createTemplateSuccessSaga(),
	saveTargetSaga(),
	fetchMasterTargetsSaga(),
	fetchTargetsByTemplateIDSaga(),
	deleteTemplateSaga(),
	fetchStagesSaga(),
	addStageSaga(),
	deleteStageSaga(),
	updateStageSaga(),
	fetchStepsSaga(),
	addStepSaga(),
	deleteStepSaga(),
	updateStepSaga(),
	loginSaga(),
];

export default function* rootSaga() {
	yield all(allSagas);
}
