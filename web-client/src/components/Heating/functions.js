export const getFormikInitialState = () => {
  return {
    temperature: null,
    followTemperature: null,
    hours: 0,
    mins: 0,
    secs: 0,
  };
};

export const getRequestBody = (formik) => {
  const formikValues = formik.values;

  const time =
    formikValues.hours * 60 * 60 + formikValues.mins * 60 + formikValues.secs;

  if (time > 3660) {
    return false;
  }

  const body = {
    temperature: formikValues.temperature,
    follow_temp: formikValues.followTemperature,
    duration: time,
  };

  return body;
};
