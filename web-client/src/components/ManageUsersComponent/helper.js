import { USER_ROLES } from "appConstants";

export const roleOptions = Object.keys(USER_ROLES).map((key) => {
  return {
    label: key,
    value: USER_ROLES[key],
  };
});
