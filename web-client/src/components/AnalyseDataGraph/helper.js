/**
 *  Transform targets into (label, value) from (target_name, target_id,...)
 * to use in dropdown of analyse data filter
 */
export const generateTargetOptions = (targetList) => {
  return targetList?.map((target) => {
    return {
      label: target.target_name,
      value: target.target_id,
    };
  });
};
