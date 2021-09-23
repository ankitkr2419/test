import { takeEvery, put } from "redux-saga/effects";

import fetchAction from "actions/fetchActions";
import { API_HOST_URL, API_HOST_VERSION } from "appConstants";
import { toast } from "react-toastify";

/**
 * getQueryString with form query string from given object
 * e.g {a : 'A', b: 'B'} => '?a=A&b=B'
 */
const getQueryString = (params) => {
  const queryString = [];
  // Purposefully disabled
  // eslint-disable-next-line no-restricted-syntax
  for (const key in params) {
    if (params[key]) {
      queryString.push(`${key}=${params[key]}`);
    }
  }
  return `?${queryString.join("&")}`;
};

// getRequestUrl with form request url based on reqPath and params
const getRequestUrl = (reqPath, params) => {
  let queryParams = "";
  // params may be present if want to add inputs over query string
  if (params !== null) {
    queryParams = getQueryString(params);
  }
  // If API_HOST_URL is present return prefix with it( API_HOST_URL can be configured from .env)
  if (API_HOST_URL) {
    return `${API_HOST_URL}/${reqPath}${queryParams}`;
  }

  // If API_HOST_URL is not present that means client and server running at same domain
  return `/${reqPath}${queryParams}`;
};

const defaultHeaders = (token) => (token ? 
  {
    "Content-Type": "application/json",
    // API_HOST_VERSION is api version is also configured from .env(constants.js)
    Accept: `application/${API_HOST_VERSION}`,
    Authorization: `Bearer ${token}`
  }: {
    "Content-Type": "application/json",
    // API_HOST_VERSION is api version is also configured from .env(constants.js)
    Accept: `application/${API_HOST_VERSION}`,
});

// Rest success status check
export const isSuccessResponse = (status) => status >= 200 && status < 300;

// dispatcherHelper is a method which requires actions and payload that is to be dispatched
const dispatcherHelper = (action, parsedResponse, error = false) => {
  return error
    ? {
        type: action,
        payload: {
          serverErrors: parsedResponse,
          error,
        },
      }
    : {
        type: action,
        payload: {
          response: parsedResponse,
          error,
        },
      };
};

/**
 * callApi is generic methods which is called on <FETCH_RESPONSE> action(saga)
 * Is used to call restApi
 * @param {*} actions
 * Action object must have
 * reqPath(required) : endpoint
 * successAction(required) : Success action to dispatch
 * failureAction(required) : Failure action to dispatch
 * method : http method default =>"GET"
 * params : query parameters,
 */
export function* callApi(actions) {
  const {
    payload: {
      body,
      reqPath,
      successAction,
      failureAction,
      method = "GET",
      params = null,
      showPopupSuccessMessage = false,
      showPopupFailureMessage = false,
      token,
      isMultipartFormData
    },
  } = actions;
  
  // request url formation
  const fetchUrl = getRequestUrl(reqPath, params);
  try {
    const fetchOptions = {
      method,
      headers: defaultHeaders(token),
      body: body && (isMultipartFormData ? body : JSON.stringify(body)),
    };

    /**
     * if want to send multipart form data, we need to delete content type explicitly 
     * and browser will set it automatically for us
     */
    if(isMultipartFormData) { 
      delete fetchOptions.headers['Content-Type']; 
    }

    const response = yield fetch(fetchUrl, fetchOptions);

    let parsedResponse;
    try {
      parsedResponse = yield response.json();
    } catch (e) {
      parsedResponse = response.statusText;
    }

    if (response.ok) {
      yield put(dispatcherHelper(successAction, parsedResponse));
      if(showPopupSuccessMessage && parsedResponse?.msg){
        toast.success(parsedResponse.msg)
      }
    } else {
      yield put(dispatcherHelper(failureAction, parsedResponse, true));
      if(showPopupFailureMessage && parsedResponse?.err){
        toast.error(parsedResponse.err, { autoClose: false })
      }
    }
  } catch (error) {
    throw error;
  }
}
/**
 * fetchResponseSaga is listening <FETCH_RESPONSE> action
 */
export function* fetchResponseSaga() {
  yield takeEvery(fetchAction.fetchResponse, callApi);
}
