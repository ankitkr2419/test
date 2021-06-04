export const getRequestBody = (formik, activeTab) => {
  const formikValues = formik.values;

  const time1 =
    parseInt(formikValues.hours1) * 60 * 60 +
    parseInt(formikValues.mins1) * 60 +
    parseInt(formikValues.secs1);

  const time2 =
    parseInt(formikValues.hours2) * 60 * 60 +
    parseInt(formikValues.mins2) * 60 +
    parseInt(formikValues.secs2);

  if (time1 !== 0) {
    if (time1 > 3660) {
      return false;
    }
  }

  if (time2 !== 0) {
    if (time2 > 3660) {
      return false;
    }
  }

  const rpm1 = parseInt(formikValues.rpm1);
  if (!rpm1 || rpm1 === 0) {
    return false;
  }

  const temperature = parseInt(formikValues.temperature);
  if (temperature !== 0) {
    if (temperature < 20 || temperature > 120) {
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
    hours1 = Math.floor(time1 / 3600);
    mins1 = Math.floor((time1 % 3600) / 60);
    secs1 = Math.floor(time1 % 60);
  }

  if (editData && editData.time_2) {
    const time2 = editData.time_2;
    hours2 = Math.floor(time2 / 3600);
    mins2 = Math.floor((time2 % 3600) / 60);
    secs2 = Math.floor(time2 % 60);
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
