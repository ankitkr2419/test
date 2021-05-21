export const updateAllTicks = (formik) => {
  const currentState = formik.values;

  Object.keys(currentState).forEach((key) => {
    const processDetails = currentState[key].processDetails;
    let tick = false;

    for (const key in processDetails) {
      if (processDetails[key].id) {
        tick = true;
        break;
      }
    }
    formik.setFieldValue(`${key}.isTicked`, tick);
  });
};
