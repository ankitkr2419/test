import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";

import {
  commonDetailsInitiated,
  updateCommonDetailsInitiated,
  fetchRtpcrConfigsInitiated,
  updateRtpcrConfigsInitiated,
  fetchTECConfigsInitiated,
  updateTECConfigsInitiated,
  startLidPid,
  abortLidPid,
  resetTECInitiated,
  autoTuneTECInitiated,
  updateToleranceInitiated,
  fetchToleranceInitiated,
  runDyeCalibration,
} from "action-creators/calibrationActionCreators";
import CalibrationComponent from "components/Calibration";
import {
  formikInitialState,
  formikInitialStateRtpcrVars,
  formikInitialStateTECVars,
  formikInitialStateDyeCalibration,
  createDyeOptions,
} from "components/Calibration/helper";
import { deckBlockInitiated } from "action-creators/loginActionCreators";
import { populateFormikStateFromApi } from "components/FormikFieldsEditor/helper";
import { PID_STATUS } from "appConstants";

const CalibrationRtpcrContainer = () => {
  const dispatch = useDispatch();

  //formik state for common fields
  const formik = useFormik({
    initialValues: formikInitialState,
    enableReinitialize: true,
  });

  //formik state for rtpcr variables
  const formikRtpcrVars = useFormik({
    initialValues: formikInitialStateRtpcrVars,
    enableReinitialize: true,
  });

  //formik state for TEC variables
  const formikTECVars = useFormik({
    initialValues: formikInitialStateTECVars,
    enableReinitialize: true,
  });

  //formik state for dye calibration
  const formikDyeCalibration = useFormik({
    initialValues: formikInitialStateDyeCalibration,
    enableReinitialize: true,
  });

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { isAdmin, name, token } = activeDeckObj;

  const commonDetailsReducer = useSelector(
    (state) => state.commonDetailsReducer
  );
  const commonDetailsReducerData = commonDetailsReducer.toJS();
  const { isUpdateApi, details } = commonDetailsReducerData;

  //rtpcr configs
  const rtpcrConfigsReducer = useSelector((state) => state.rtpcrConfigsReducer);
  const rtpcrConfigsReducerData = rtpcrConfigsReducer.toJS();
  let isRtpcrConfigUpdateApi = rtpcrConfigsReducerData?.isUpdateApi;
  let rtpcrConfigDetails = rtpcrConfigsReducerData?.details;

  //TEC variables
  const tecConfigsReducer = useSelector((state) => state.tecConfigsReducer);
  const tecConfigsReducerData = tecConfigsReducer.toJS();
  let isTECConfigUpdateApi = tecConfigsReducerData?.isUpdateApi;
  let tecConfigDetails = tecConfigsReducerData?.details;

  //lid pid reducer
  const lidPidReducer = useSelector((state) => state.lidPidReducer);
  const lidPidReducerData = lidPidReducer.toJS();
  const { lidPidStatus } = lidPidReducerData;

  //Tolerance Variables
  const toleranceReducer = useSelector((state) => state.toleranceReducer);
  const toleranceReducerData = toleranceReducer.toJS();

  //dye calibration
  const dyeCalibrationReducer = useSelector(
    (state) => state.dyeCalibrationReducer
  );
  const dyeCalibrationReducerData = dyeCalibrationReducer.toJS();
  const { dyeCalibrationStatus } = dyeCalibrationReducerData;

  //initially populate with previous data
  useEffect(() => {
    if (token) {
      // fetch commonDetails (name, email, roomTemp) API called initially
      dispatch(commonDetailsInitiated(token));
      //another deck must be blocked
      dispatch(deckBlockInitiated({ deckName: name }));
      //fetch rtpcr variables from api
      dispatch(fetchRtpcrConfigsInitiated(token));
      //fetch TEC variables from api
      dispatch(fetchTECConfigsInitiated(token));
      //fetch Tolerance variables from api
      dispatch(fetchToleranceInitiated(token));
    }
  }, [dispatch, token]);

  useEffect(() => {
    if (
      commonDetailsReducerData.error === false &&
      commonDetailsReducerData.isLoading === false
    ) {
      // populate formik data with fetched values
      if (isUpdateApi === false) {
        const { receiver_name, receiver_email, room_temperature } = details;
        formik.setFieldValue("name.value", receiver_name);
        formik.setFieldValue("email.value", receiver_email);
        formik.setFieldValue("roomTemperature.value", room_temperature);
      } else {
        // fetch updated data after updation
        dispatch(commonDetailsInitiated(token));
      }
    }
  }, [
    commonDetailsReducerData.error,
    commonDetailsReducerData.isLoading,
    isUpdateApi,
  ]);

  useEffect(() => {
    if (
      rtpcrConfigsReducerData.error === false &&
      rtpcrConfigsReducerData.isLoading === false
    ) {
      if (isRtpcrConfigUpdateApi === false) {
        //populate formik data with fetched values
        populateFormikStateFromApi(formikRtpcrVars, rtpcrConfigDetails);
      } else {
        //fetch updated data after updation
        dispatch(fetchRtpcrConfigsInitiated(token));
      }
    }
  }, [
    rtpcrConfigsReducerData.error,
    rtpcrConfigsReducerData.isLoading,
    isRtpcrConfigUpdateApi,
  ]);

  useEffect(() => {
    if (
      tecConfigsReducerData.error === false &&
      tecConfigsReducerData.isLoading === false
    ) {
      if (isTECConfigUpdateApi === false) {
        //populate formik data with fetched values
        populateFormikStateFromApi(formikTECVars, tecConfigDetails);
      } else {
        //fetch updated data after updation
        dispatch(fetchTECConfigsInitiated(token));
      }
    }
  }, [
    tecConfigsReducerData.error,
    tecConfigsReducerData.isLoading,
    isTECConfigUpdateApi,
  ]);

  useEffect(() => {
    if (
      toleranceReducerData.error === false &&
      toleranceReducerData.isLoading === false &&
      toleranceReducerData.isUpdateApi === true
    ) {
      //fetch updated data after updation
      dispatch(fetchToleranceInitiated(token));
    }
  }, [
    toleranceReducerData.error,
    toleranceReducerData.isLoading,
    toleranceReducerData.isUpdateApi,
  ]);

  const handleSaveDetailsBtn = (data) => {
    const { name, email, roomTemperature } = data;
    const requestBody = {
      receiver_name: name.value,
      receiver_email: email.value,
      room_temperature: roomTemperature.value,
    };
    dispatch(updateCommonDetailsInitiated({ token, data: requestBody }));
  };

  const handleRtpcrConfigSubmitButton = (requestBody) => {
    dispatch(updateRtpcrConfigsInitiated(token, requestBody));
  };

  const handleTECConfigSubmitButton = (requestBody) => {
    dispatch(updateTECConfigsInitiated(token, requestBody));
  };

  const handleLidPidButton = () => {
    if (
      lidPidStatus === PID_STATUS.running ||
      lidPidStatus === PID_STATUS.progressing
    ) {
      dispatch(abortLidPid(token));
    } else {
      dispatch(startLidPid(token));
    }
  };

  const handleResetTEC = () => {
    dispatch(resetTECInitiated(token));
  };

  const handleAutoTuneTEC = () => {
    dispatch(autoTuneTECInitiated(token));
  };

  const handleDyeCalibrationButton = (requestBody) => {
    dispatch(runDyeCalibration(token, requestBody));
  };

  const handleSaveToleranceBtn = (requestBody) => {
    dispatch(updateToleranceInitiated({ token, requestBody }));
  };

  //dye options
  let dyeOptions = [];
  dyeOptions =
    toleranceReducerData?.data &&
    dyeOptions.length === 0 &&
    createDyeOptions(toleranceReducerData?.data);
  //set first dye as default selected for first time
  if (
    dyeOptions?.length &&
    formikDyeCalibration.values.selectedDye.value === null
  ) {
    formikDyeCalibration.setFieldValue("selectedDye.value", dyeOptions[0]);
  }

  return (
    <CalibrationComponent
      formik={formik}
      isAdmin={isAdmin}
      handleSaveDetailsBtn={handleSaveDetailsBtn}
      formikRtpcrVars={formikRtpcrVars}
      handleRtpcrConfigSubmitButton={handleRtpcrConfigSubmitButton}
      formikTECVars={formikTECVars}
      handleTECConfigSubmitButton={handleTECConfigSubmitButton}
      lidPidStatus={lidPidStatus}
      handleLidPidButton={handleLidPidButton}
      handleResetTEC={handleResetTEC}
      handleAutoTuneTEC={handleAutoTuneTEC}
      dyeOptions={dyeOptions}
      formikDyeCalibration={formikDyeCalibration}
      handleDyeCalibrationButton={handleDyeCalibrationButton}
      dyeCalibrationStatus={dyeCalibrationStatus}
      handleSaveToleranceBtn={handleSaveToleranceBtn}
      toleranceData={toleranceReducerData.data}
    />
  );
};

export default React.memo(CalibrationRtpcrContainer);
