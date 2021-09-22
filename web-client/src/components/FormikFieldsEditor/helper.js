import { EMAIL_REGEX_OR_EMPTY_STR, NAME_REGEX } from "appConstants";

export const validateAllFields = (state) => {
  for (const key in state) {
    const { name, value, isInvalid } = state[key];
    if (
      value === null ||
      value === "" ||
      isInvalid === true ||
      isValueValid(state, name, value) === false
    ) {
      return false;
    }
  }
  return true;
};

/**
 * generalized value validator
 */
export const isValueValid = (formik, name, value) => {
  const element = formik[name];
  const { type, min, max } = element;

  if (type === "number" && (value < min || value > max)) {
    return false;
  } else if (
    type === "email" &&
    value.match(EMAIL_REGEX_OR_EMPTY_STR) === null
  ) {
    return false;
  }

  return true;
};

export const getRequestBody = (state) => {
  const body = {};
  for (const key in state) {
    const element = state[key];
    const { type, apiKey, value } = element;
    body[apiKey] = type === "number" ? parseInt(value) : value;
  }
  return body;
};

/**
 * populate values from api into formik state
 */
export const populateFormikStateFromApi = (formik, apiData) => {
  if (apiData) {
    Object.keys(formik.values).map((element) => {
      const { apiKey, name } = formik.values[element];

      const newValue = apiData[apiKey] ? apiData[apiKey] : "";
      const isValid = isValueValid(formik.values, name, newValue);

      // set formik fields
      formik.setFieldValue(`${name}.isInvalid`, !isValid);
      formik.setFieldValue(`${name}.value`, newValue);
    });
  }
};
