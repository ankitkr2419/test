import {
  MAX_TEMP_ALLOWED,
  MIN_TEMP_ALLOWED,
  MAX_TIME_ALLOWED,
  timeConstants,
} from "appConstants";

const { SEC_IN_ONE_HOUR, SEC_IN_ONE_MIN, MIN_IN_ONE_HOUR } = timeConstants;

/** This function checks for validity of input data and
 *  returns the request body.
 */
export const getRequestBody = (formik) => {
  const formikValues = formik.values;

  const time1 =
    parseInt(formikValues.hours1) * MIN_IN_ONE_HOUR * SEC_IN_ONE_MIN +
    parseInt(formikValues.mins1) * SEC_IN_ONE_MIN +
    parseInt(formikValues.secs1);

  const time2 =
    parseInt(formikValues.hours2) * MIN_IN_ONE_HOUR * SEC_IN_ONE_MIN +
    parseInt(formikValues.mins2) * SEC_IN_ONE_MIN +
    parseInt(formikValues.secs2);

  if (time1 !== 0) {
    if (time1 > MAX_TIME_ALLOWED) {
      return false;
    }
  }

  if (time2 !== 0) {
    if (time2 > MAX_TIME_ALLOWED) {
      return false;
    }
  }

  const rpm1 = parseInt(formikValues.rpm1);
  if (!rpm1 || rpm1 === 0) {
    return false;
  }

  const temperature = parseInt(formikValues.temperature);
  if (temperature !== 0) {
    if (temperature < MIN_TEMP_ALLOWED || temperature > MAX_TEMP_ALLOWED) {
      return false;
    }
  }

  const body = {
    with_temp: activeTab === "2",
    temperature: temperature ? temperature : 0,
    follow_temp: formikValues.followTemperature,
    rpm_1: parseInt(formikValues.rpm1) ? parseInt(formikValues.rpm1) : 0,
    rpm_2: parseInt(formikValues.rpm2) ? parseInt(formikValues.rpm2) : 0,
    time_1: time1 ? time1 : 0,
    time_2: time2 ? time2 : 0,
  };

  return body;
};

export const getFormikInitialState = (editData = null) => {
  let hours1, mins1, secs1, hours2, mins2, secs2;

  if (editData && editData.rpm_1 !== 0) {
    isDisabled.rpm2 = false;
  }

  if (editData && editData.time_1) {
    const time1 = editData.time_1;
    hours1 = Math.floor(time1 / SEC_IN_ONE_HOUR);
    mins1 = Math.floor((time1 % SEC_IN_ONE_HOUR) / MIN_IN_ONE_HOUR);
    secs1 = Math.floor(time1 % MIN_IN_ONE_HOUR);
  }

  if (editData && editData.time_2) {
    const time2 = editData.time_2;
    hours2 = Math.floor(time2 / SEC_IN_ONE_HOUR);
    mins2 = Math.floor((time2 % SEC_IN_ONE_HOUR) / MIN_IN_ONE_HOUR);
    secs2 = Math.floor(time2 % MIN_IN_ONE_HOUR);
  }

  return {
    withHeating: editData?.with_temp ? editData.with_temp : null,
    temperature: editData?.temperature ? editData.temperature : null,
    followTemperature: editData?.follow_temp ? editData.follow_temp : false,
    rpm1: editData?.rpm_1 ? editData.rpm_1 : 0,
    rpm2: editData?.rpm_2 ? editData.rpm_2 : 0,
    hours1: hours1 ? hours1 : 0,
    mins1: mins1 ? mins1 : 0,
    secs1: secs1 ? secs1 : 0,
    hours2: hours2 ? hours2 : 0,
    mins2: mins2 ? mins2 : 0,
    secs2: secs2 ? secs2 : 0,
    rpm2IsDisabled: editData && editData.rpm_1 !== 0 ? false : true,
  };
};

export const isDisabled = {
  withHeating: false,
  withoutHeating: false,
};

/** This function is used  */
export const setRpmFormikField = (formik, activeTab, fieldName, fieldValue) => {
  formik.setFieldValue(fieldName, fieldValue);

  if (fieldName === "rpm2") {
    return;
  }

  const currentTab = activeTab === "1" ? "withHeating" : "withoutHeating";
  const rpm1Value = parseInt(fieldValue);

  //toggle disabled state of rpm2 based on rpm1Value
  const rpm2IsDisabled = rpm1Value && rpm1Value !== 0 ? false : true;
  formik.setFieldValue("rpm2IsDisabled", rpm2IsDisabled);

  //toggle disabled state of tabs based on rpm1Value
  isDisabled[currentTab] = rpm1Value ? true : false;
};
