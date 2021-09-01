import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";

import {
  commonDetailsInitiated,
  updateCommonDetailsInitiated,
  fetchRtpcrConfigsInitiated,
  updateRtpcrConfigsInitiated,
} from "action-creators/calibrationActionCreators";
import CalibrationComponent from "components/Calibration";
import {
  formikInitialState,
  formikInitialStateRtpcrVars,
} from "components/Calibration/helper";
import { deckBlockInitiated } from "action-creators/loginActionCreators";
import { populateFormikStateFromApi } from "components/FormikFieldsEditor/helper";

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

  //initially populate with previous data
  useEffect(() => {
    if (token) {
      // fetch commonDetails (name, email, roomTemp) API called initially
      dispatch(commonDetailsInitiated(token));
      //another deck must be blocked
      dispatch(deckBlockInitiated({ deckName: name }));
      //fetch rtpcr variables from api
      dispatch(fetchRtpcrConfigsInitiated(token));
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

  /**to change formik field */
  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  return (
    <CalibrationComponent
      formik={formik}
      isAdmin={isAdmin}
      handleSaveDetailsBtn={handleSaveDetailsBtn}
      formikRtpcrVars={formikRtpcrVars}
      handleRtpcrConfigSubmitButton={handleRtpcrConfigSubmitButton}
    />
  );
};

export default React.memo(CalibrationRtpcrContainer);
