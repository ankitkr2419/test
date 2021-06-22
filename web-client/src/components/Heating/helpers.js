export const getFormikInitialState = (editData = null) => {
  let hours, mins, secs;

  if (editData && editData.duration) {
    const duration = editData.duration;
    hours = Math.floor(duration / 3600);
    mins = Math.floor((duration % 3600) / 60);
    secs = Math.floor(duration % 60);
  }

  return {
    temperature: editData?.temperature ? editData.temperature : 0,
    followTemperature: editData?.follow_temp ? editData.follow_temp : false,
    hours: hours ? hours : 0,
    mins: mins ? mins : 0,
    secs: secs ? secs : 0,
  };
};

export const getRequestBody = (formik) => {
  const formikValues = formik.values;

  const time =
    parseInt(formikValues.hours) * 60 * 60 +
    parseInt(formikValues.mins) * 60 +
    parseInt(formikValues.secs);

  if (time !== 0) {
    if (time < 10 || time > 3660) {
      return false;
    }
  }

  const temperature = parseInt(formikValues.temperature);
  if (temperature !== 0) {
    if (temperature < 20 || temperature > 120) {
      return false;
    }
  }

  const body = {
    temperature: temperature,
    follow_temp: formikValues.followTemperature,
    duration: time,
  };

  return body;
};
