import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";

import {
  calibrationInitiated,
  updateCalibrationInitiated,
  commonDetailsInitiated,
  updateCommonDetailsInitiated,
} from "action-creators/calibrationActionCreators";
import CalibrationComponent from "components/Calibration";
import { formikInitialState } from "components/Calibration/helper";
import {
  deckBlockInitiated,
  logoutInitiated,
} from "action-creators/loginActionCreators";

const CalibrationRtpcrContainer = () => {
  const dispatch = useDispatch();

  //formik state for common fields
  const formik = useFormik({
    initialValues: formikInitialState,
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

  //initially populate with previous data
  useEffect(() => {
    if (token) {
      // fetch commonDetails (name, email, roomTemp) API called initially
      dispatch(commonDetailsInitiated(token));
      //another deck must be blocked
      dispatch(deckBlockInitiated({ deckName: name }));
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

  const handleSaveDetailsBtn = (data) => {
    const { name, email, roomTemperature } = data;
    const requestBody = {
      receiver_name: name.value,
      receiver_email: email.value,
      room_temperature: roomTemperature.value,
    };
    dispatch(updateCommonDetailsInitiated({ token, data: requestBody }));
  };

  /**to change formik field */
  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  return (
    <CalibrationComponent
      formik={formik}
      isAdmin={isAdmin}
      handleOnChange={handleOnChange}
      handleSaveDetailsBtn={handleSaveDetailsBtn}
    />
  );
};

export default React.memo(CalibrationRtpcrContainer);

//old code for reference
//TODO remove this if not needed for reference

// import React, { useEffect } from "react";
// import { useDispatch, useSelector } from "react-redux";
// import {
//   calibrationInitiated,
//   updateCalibrationInitiated,
//   commonDetailsInitiated,
// } from "action-creators/calibrationActionCreators";
// import CalibrationComponent from "components/Calibration";

// const CalibrationRtpcrContainer = () => {
//   const dispatch = useDispatch();

//   //get login reducer details
//   const loginReducer = useSelector((state) => state.loginReducer);
//   const loginReducerData = loginReducer.toJS();
//   let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
//   const { token } = activeDeckObj;

//   //get calibration configurations from reducer
//   const calibrationReducer = useSelector((state) => state.calibrationReducer);
//   const calibrationReducerData = calibrationReducer.toJS();
//   const { configs } = calibrationReducerData;

//   //get calibration configurations from reducer
//   const updateCalibrationReducer = useSelector(
//     (state) => state.updateCalibrationReducer
//   );
//   const updateCalibrationReducerData = updateCalibrationReducer.toJS();
//   const { isLoading, error } = updateCalibrationReducerData;

//   //initially populate with previous data
//   useEffect(() => {
//     if (token) {
//       // dispatch(calibrationInitiated(token));

//       // fetch commonDetails (name, email, roomTemp) API called initially
//       dispatch(commonDetailsInitiated(token));
//     }
//   }, [dispatch, token]);

//   //after update API call is successful populate fields with new data
//   useEffect(() => {
//     if (token && error === false && isLoading === false) {
//       dispatch(calibrationInitiated(token));
//     }
//   }, [dispatch, isLoading, error, token]);

//   const saveButtonClickHandler = (configData) => {
//     let data = {
//       receiver_name: configData.name,
//       receiver_email: configData.email,
//       room_temperature: configData.roomTemperature,
//       homing_time: configData.homingTime,
//       no_of_homing_cycles: configData.noOfHomingCycles,
//       cycle_time: configData.cycleTime,
//       pid_temperature: configData.pidTemperature,
//       pid_minutes: configData.pidMinutes,
//     };

//     dispatch(updateCalibrationInitiated({ token, data }));
//   };

//   return (
//     <CalibrationComponent
//       configs={configs}
//       saveButtonClickHandler={saveButtonClickHandler}
//     />
//   );
// };

// export default React.memo(CalibrationRtpcrContainer);
