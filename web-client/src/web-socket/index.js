import { w3cwebsocket as W3CWebSocket } from "websocket";
import {
  webSocketOpened,
  webSocketClosed,
  webSocketError,
} from "action-creators/webSocketActionCreators";
import { WS_HOST_URL, SOCKET_MESSAGE_TYPE, DECKNAME } from "appConstants";
import { updateWellThroughSocket } from "action-creators/wellActionCreators";
import { wellGraphSucceeded } from "action-creators/wellGraphActionCreators";
import {
  experimentedCompleted,
  runExperimentInProgress,
  runExperimentSuccess,
} from "action-creators/runExperimentActionCreators";
import { showErrorModal } from "action-creators/modalActionCreators";
import { temperatureDataSucceeded } from "action-creators/temperatureGraphActionCreators";
import {
  homingActionInProgress,
  homingActionInCompleted,
  homingActionInProgressDeckA,
  homingActionInProgressDeckB,
  homingActionInCompletedDeckB,
  homingActionInCompletedDeckA,
} from "action-creators/homingActionCreators";

import {
  runRecipeInProgress,
  runRecipeInCompleted,
} from "action-creators/recipeActionCreators";

import {
  runCleanUpActionInProgress,
  runCleanUpActionInCompleted,
} from "action-creators/cleanUpActionCreators";

import {
  discardTipInProgress,
  discardTipInCompleted,
} from "action-creators/discardDeckActionCreators";

import { toast } from "react-toastify";
import {
  heaterProgress,
  heaterRunInProgress,
  heaterRunInSuccess,
  heaterRunInAborted,
  runPidInProgress,
  runPidInSuccess,
  progressLidPid,
  successLidPid,
  shakerRunInProgress,
  shakerRunInSuccess,
  shakerRunInAborted,
  progressDyeCalibration,
  completedDyeCalibration,
  pidInProgress,
  pidInSuccess,
  pidInAborted,
} from "action-creators/calibrationActionCreators";

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
        case SOCKET_MESSAGE_TYPE.PIDProgress:
          dispatch(runPidInProgress(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.PIDSuccess:
          dispatch(runPidInSuccess(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.rtpcrProgress:
          dispatch(runExperimentInProgress(data));
          break;
        case SOCKET_MESSAGE_TYPE.rtpcrSuccess:
          dispatch(runExperimentSuccess(data));
          break;
        case SOCKET_MESSAGE_TYPE.success:
          dispatch(experimentedCompleted(data));
          break;
        case SOCKET_MESSAGE_TYPE.homingProgress:
          let homingParsedData = JSON.parse(data);
          const { deck } = homingParsedData;
          dispatch(homingActionInProgress(homingParsedData));
          if (deck === DECKNAME.DeckAShort) {
            dispatch(homingActionInProgressDeckA(homingParsedData));
          } else if (deck === DECKNAME.DeckBShort) {
            dispatch(homingActionInProgressDeckB(homingParsedData));
          }
          break;
        case SOCKET_MESSAGE_TYPE.homingSuccess:
          let parsedData = JSON.parse(data);
          let homingSuccessMsg =
            parsedData &&
            parsedData.operation_details &&
            parsedData.operation_details.message
              ? parsedData.operation_details.message
              : "Homing Successfull";
          toast.success(homingSuccessMsg);
          dispatch(homingActionInCompleted(data));
          if (parsedData?.deck) {
            if (parsedData.deck === DECKNAME.DeckAShort) {
              dispatch(homingActionInCompletedDeckA(parsedData));
            } else if (parsedData.deck === DECKNAME.DeckBShort) {
              dispatch(homingActionInCompletedDeckB(parsedData));
            }
          }
          break;
        case SOCKET_MESSAGE_TYPE.runRecipeProgress:
          dispatch(runRecipeInProgress(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.runRecipeSuccess:
          dispatch(runRecipeInCompleted(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.uvLightProgress:
          dispatch(runCleanUpActionInProgress(data));
          break;
        case SOCKET_MESSAGE_TYPE.uvLightSuccess:
          dispatch(runCleanUpActionInCompleted(data));
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
        case SOCKET_MESSAGE_TYPE.ErrorPCR:
          dispatch(showErrorModal(data));
          break;
        case SOCKET_MESSAGE_TYPE.ErrorExtractionMonitor:
          let parsedErrorData = JSON.parse(data);
          let errorMessage = parsedErrorData.message;
          if (errorMessage) {
            toast.error(errorMessage, { autoClose: false });
          }
          break;
        case SOCKET_MESSAGE_TYPE.progressHeater:
          dispatch(heaterProgress(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.progressPidTuning:
          dispatch(pidInProgress());
          break;
        case SOCKET_MESSAGE_TYPE.successPidTuning:
          dispatch(pidInSuccess());
          break;
        // global abort for PID, Shaker and heater
        case SOCKET_MESSAGE_TYPE.ErrorOperationAborted:
          dispatch(pidInAborted());
          dispatch(shakerRunInAborted());
          dispatch(heaterRunInAborted());
          break;
        // case SOCKET_MESSAGE_TYPE.abortShakerRun:
        //   dispatch(shakerRunInAborted());
        //   break;
        // case SOCKET_MESSAGE_TYPE.abortHeaterRun:
        //   dispatch(heaterRunInAborted());
        //   break;
        case SOCKET_MESSAGE_TYPE.PROGRESSLidPIDTuning:
          dispatch(progressLidPid(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.SUCCESSLidPIDTuning:
          dispatch(successLidPid(JSON.parse(data)));
          break;
        case SOCKET_MESSAGE_TYPE.progressShakerRun:
          dispatch(shakerRunInProgress());
          break;
        case SOCKET_MESSAGE_TYPE.successShakerRun:
          dispatch(shakerRunInSuccess());
          break;
        case SOCKET_MESSAGE_TYPE.progressHeaterRun:
          dispatch(heaterRunInProgress());
          break;
        case SOCKET_MESSAGE_TYPE.successHeaterRun:
          dispatch(heaterRunInSuccess());
          break;
        case SOCKET_MESSAGE_TYPE.progressDyeCalibration:
          dispatch(progressDyeCalibration());
          break;
        case SOCKET_MESSAGE_TYPE.completedDyeCalibration:
          dispatch(completedDyeCalibration());
          toast.success(data);
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
