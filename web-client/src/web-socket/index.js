import { w3cwebsocket as W3CWebSocket } from "websocket";
import {
  webSocketOpened,
  webSocketClosed,
  webSocketError,
} from "action-creators/webSocketActionCreators";
import { WS_HOST_URL, SOCKET_MESSAGE_TYPE } from "appConstants";
import { updateWellThroughSocket } from "action-creators/wellActionCreators";
import { wellGraphSucceeded } from "action-creators/wellGraphActionCreators";
import { experimentedCompleted } from "action-creators/runExperimentActionCreators";
import { showErrorModal } from "action-creators/modalActionCreators";
import { temperatureDataSucceeded } from "action-creators/temperatureGraphActionCreators";
import {
  homingActionInProgress,
  homingActionInCompleted,
} from "action-creators/homingActionCreators";

import {
  runRecipeInProgress,
  runRecipeInCompleted,
} from "action-creators/recipeActionCreators";

import {
  cleanUpActionInProgress,
  cleanUpActionInCompleted,
} from "action-creators/cleanUpActionCreators";

import {
  discardTipInProgress,
  discardTipInCompleted,
} from "action-creators/discardDeckActionCreators";

import { toast } from "react-toastify";

let webSocket = null;
export const connectSocket = (dispatch) => {
  webSocket = new W3CWebSocket(`${WS_HOST_URL}/monitor`);
  webSocket.onopen = (event) => {
    console.info("socket connection opened", event);
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
        case SOCKET_MESSAGE_TYPE.homingProgress:
          dispatch(homingActionInProgress(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.homingSuccess:
          toast.success("Homing Successfull");
          dispatch(homingActionInCompleted(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.runRecipeProgress:
          dispatch(runRecipeInProgress(data));
          break;
        case SOCKET_MESSAGE_TYPE.runRecipeSuccess:
          dispatch(runRecipeInCompleted(data));
          break;
        case SOCKET_MESSAGE_TYPE.uvLightProgress:
          dispatch(cleanUpActionInProgress(data));
          break;
        case SOCKET_MESSAGE_TYPE.uvLightSuccess:
          dispatch(cleanUpActionInCompleted(data));
          break;
        case SOCKET_MESSAGE_TYPE.discardTipProgress:
          dispatch(discardTipInProgress(data));
          break;
        case SOCKET_MESSAGE_TYPE.discardTipSuccess:
          dispatch(discardTipInCompleted(data));
          break;
        // case SOCKET_MESSAGE_TYPE.failure:
        // 	dispatch(experimentedFailed(data));
        // 	break;
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
    console.info("socket connection closed");
    dispatch(webSocketClosed());
  };

  webSocket.onError = (event) => {
    console.error("socket error : ", event);
    dispatch(webSocketError());
  };
};
export const disConnectSocket = () => {
  if (webSocket !== null) {
    webSocket.close();
  }
};
