/**
 * Returns sting of totalTime or remainingTime based on type.
 * if hours field is 0, this string will not include hours.
 */
export const getTimeStr = (recipeData, cleanUpData, isTotalTime = false) => {
  let reducerData;
  if (recipeData?.showProcess) {
    reducerData = recipeData.runRecipeInProgress;
  } else {
    reducerData = JSON.parse(cleanUpData.cleanUpData);
  }

  let hours = 0;
  let mins = 0;
  let secs = 0;
  if (reducerData?.operation_details) {
    const time =
      reducerData.operation_details[
        isTotalTime ? "total_time" : "remaining_time"
      ];

    hours = time.hours;
    mins = time.minutes;
    secs = time.seconds;
  }

  let timeStr = "";
  if (hours !== 0) {
    timeStr = `${hours} hr `;
  }
  return timeStr + `${mins} min ${secs} sec`;
};
