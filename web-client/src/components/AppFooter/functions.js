/**
 * Returns Hours, Mins, Secs based on type.
 */

export const getTime = (type, recipeData, cleanUpData) => {
  let reducerData;
  if (recipeData?.showProcess) {
    reducerData = recipeData.runRecipeInProgress;
  } else {
    reducerData = JSON.parse(cleanUpData.cleanUpData);
  }

  const time = reducerData?.operation_details.remaining_time[type];
  return time ? time : 0;
};
