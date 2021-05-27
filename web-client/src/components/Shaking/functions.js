export const getRequestBody = (formik) => {
  const formikValues = formik.values;

  const time1 =
    formikValues.hours1 * 60 * 60 +
    formikValues.mins1 * 60 +
    formikValues.secs1;

  const time2 =
    formikValues.hours2 * 60 * 60 +
    formikValues.mins2 * 60 +
    formikValues.secs2;

  if (time1 > 3660 || time2 > 3660) {
    return false;
  }

  const body = {
    with_temp: formikValues.withHeating,
    temperature: formikValues.temperature,
    follow_temp: formikValues.followTemperature,
    rpm_1: formikValues.rpm1,
    rpm_2: formikValues.rpm2,
    time_1: time1,
    time_2: time2,
  };

  return body;
};

export const getFormikInitialState = () => {
  return {
    withHeating: null,
    temperature: null,
    followTemperature: false,
    rpm1: null,
    rpm2: null,
    hours1: 0,
    mins1: 0,
    secs1: 0,
    hours2: 0,
    mins2: 0,
    secs2: 0,
  };
};

export const isDisabled = {
  withHeating: false,
  withoutHeating: false,
  rpm2: true, //initially it is disabled
};

export const setRpmFormikField = (formik, activeTab, fieldName, fieldValue) => {
  formik.setFieldValue(fieldName, fieldValue);

  if (fieldName === "rpm2") {
    return;
  }

  const currentTab = activeTab === "1" ? "withHeating" : "withoutHeating";
  const rpm1Value = fieldValue;

  //set with heatingField in formik
  let isWithHeating = null;
  if (rpm1Value) {
    isWithHeating = activeTab === "2";
  }
  formik.setFieldValue("withHeating", isWithHeating);

  //toggle disabled state of rpm2 based on rpm1Value
  isDisabled["rpm2"] = rpm1Value ? false : true;

  //toggle disabled state of tabs based on rpm1Value
  isDisabled[currentTab] = rpm1Value ? true : false;
};
