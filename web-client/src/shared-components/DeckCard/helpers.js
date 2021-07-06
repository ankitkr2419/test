const { SELECT_PROCESS_PROPS } = require("appConstants");

/* get icon by process type
 * if process type not found use 'default'
 * if process type found but icon not found, use 'default'
 */
export const getIconName = (processType) => {
  let processTypeText = processType ? processType : "default";

  let obj = SELECT_PROCESS_PROPS.find(
    (obj) => obj.processType === processTypeText
  );

  let iconName = obj?.iconName
    ? obj.iconName
    : SELECT_PROCESS_PROPS.find((obj) => obj.processType === "default")
        .iconName;
  return iconName;
};
