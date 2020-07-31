import { takeEvery, put, call } from 'redux-saga/effects';
import { callApi } from 'apis/apiHelper';
import {
	addStepActions,
	updateStepActions,
	deleteStepActions,
	listHoldStepActions,
	listCycleStepActions,
} from 'actions/stepActions';
import {
	addStepFailed,
	fetchHoldStepsFailed,
	fetchCycleStepsFailed,
	updateStepFailed,
	deleteStepFailed,
} from 'action-creators/stepActionCreators';
import { HTTP_METHODS } from 'appConstants';

export function* addStep(actions) {
	const {
		payload: { body },
	} = actions;

	const { successAction, failureAction } = addStepActions;

	try {
		yield call(callApi, {
			payload: {
				method: HTTP_METHODS.POST,
				body,
				reqPath: 'steps',
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error in adding step ', error);
		yield put(addStepFailed(error));
	}
}

export function* fetchHoldSteps(actions) {
	const {
		payload: { stageId },
	} = actions;
	const { successAction, failureAction } = listHoldStepActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `stages/${stageId}/steps`,
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error in fetch steps ', error);
		yield put(fetchHoldStepsFailed(error));
	}
}

export function* fetchCycleSteps(actions) {
	const {
		payload: { stageId },
	} = actions;
	const { successAction, failureAction } = listCycleStepActions;
	try {
		yield call(callApi, {
			payload: {
				body: null,
				reqPath: `stages/${stageId}/steps`,
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error in fetch steps ', error);
		yield put(fetchCycleStepsFailed(error));
	}
}

export function* updateStep(actions) {
	const {
		payload: { stepId, body },
	} = actions;

	const { successAction, failureAction } = updateStepActions;

	try {
		yield call(callApi, {
			payload: {
				method: HTTP_METHODS.PUT,
				body,
				reqPath: `steps/${stepId}`,
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error while updating step ', error);
		yield put(updateStepFailed(error));
	}
}

export function* deleteStep(actions) {
	const {
		payload: { stepId },
	} = actions;

	const { successAction, failureAction } = deleteStepActions;

	try {
		yield call(callApi, {
			payload: {
				method: HTTP_METHODS.DELETE,
				reqPath: `steps/${stepId}`,
				successAction,
				failureAction,
			},
		});
	} catch (error) {
		console.error('error while deleting step ', error);
		yield put(deleteStepFailed(error));
	}
}

export function* addStepSaga() {
	yield takeEvery(addStepActions.addAction, addStep);
}

export function* fetchHoldStepsSaga() {
	yield takeEvery(listHoldStepActions.listAction, fetchHoldSteps);
}

export function* fetchCycleStepsSaga() {
	yield takeEvery(listCycleStepActions.listAction, fetchCycleSteps);
}

export function* updateStepSaga() {
	yield takeEvery(updateStepActions.updateAction, updateStep);
}

export function* deleteStepSaga() {
	yield takeEvery(deleteStepActions.deleteAction, deleteStep);
}
