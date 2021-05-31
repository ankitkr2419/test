import { createSelector } from "reselect";

/**
 * Purpose of this function is to convert array of elements as required by react-select component
 * It requires array of objects as [{label : <label>, value: <value>}]
 * @param {*} list => immutable list
 * @param {*} labelKey => stting
 * @param {*} valueKey =>string
 */
export const covertToSelectOption = (list, labelKey, valueKey) => {
  const arr = [];
  list.map((ele) =>
    arr.push({ label: ele.get(labelKey), value: ele.get(valueKey) })
  );
  return arr;
};

export const convertStringToSeconds = (timeString) => {
  if (timeString.indexOf(":") !== -1) {
    const a = timeString.split(":"); // split it at the colons

    // minutes are worth 60 seconds.
    return parseInt(+a[0] * 60 + +a[1], 10);
  }
  return 0;
};

export const convertSecondsToString = (seconds) => {
  let min = Math.floor(seconds / 60);
  let sec = seconds - min * 60;
  min = min < 10 ? `0${min}` : min;
  sec = sec < 10 ? `0${sec}` : sec;
  return `${min}:${sec}`;
};

export const formatDate = createSelector(
  (inputDate) => inputDate,
  (inputDate) => {
    const dt = new Date(inputDate);
    let dd = dt.getDate();

    let mm = dt.getMonth() + 1;
    const yyyy = dt.getFullYear();
    if (dd < 10) {
      dd = `0${dd}`;
    }

    if (mm < 10) {
      mm = `0${mm}`;
    }
    return `${mm}/${dd}/${yyyy}`;
  }
);

const checkTime = (i) => (i < 10 ? `0${i}` : i);

export const formatTime = createSelector(
  (inputDate) => inputDate,
  (inputDate) => {
    const dt = new Date(inputDate);
    const h = dt.getHours();
    let m = dt.getMinutes();
    let s = dt.getSeconds();

    m = checkTime(m);
    s = checkTime(s);

    return `${h}:${m}:${s}`;
  }
);

//  To avoid parseFloat returning NaN return the value as it is or else return parsed output
export const parseFloatWrapper = (value) => {
  if (value === undefined || value === "" || value === null) {
    return value;
  }
  return parseFloat(value);
};

// Get the time differnce in current and start time in minutes
export const getTimeDiff = (startTime, currTime) => {
  const hour_diff = currTime.getHours() - startTime.getHours();
  const min_diff = currTime.getMinutes() - startTime.getMinutes();
  const sec_diff = currTime.getSeconds() - startTime.getSeconds();
  const time_diff = min_diff + hour_diff * 60 + sec_diff / 60;
  return time_diff;
};

// This function is used by multiple reducers in order to
// get the updated state structure.
export const getUpdatedDecks = (
  state,
  deckName,
  changesInMatchedDeck,
  changesInUnMatchedDeck = {},
  isLoginReducer = false
) => {
  const arrayOfDecks = isLoginReducer ? state.toJS().decks : state.decks;
  const array = arrayOfDecks.map((deckObj) => {
    return deckObj.name === deckName
      ? {
          ...deckObj,
          ...changesInMatchedDeck,
        }
      : {
          ...deckObj,
          ...changesInUnMatchedDeck,
        };
  });
  return array;
};

/**returns deckArray
 * if we want recipeList changes, pass changes required and recipeId and this method will return new deckArray
 */
export const getUpdatedDecksAfterRecipeListChanged = (
  state,
  deckName,
  recipeId,
  changesInMatchedRecipe,
  changesInUnMatchedRecipe
) => {
  const arrayOfDecks = state.decks;

  const array = arrayOfDecks.map((deckObj) => {
    return deckObj.name === deckName
      ? {
          ...deckObj,
          allRecipeData: deckObj.allRecipeData.map((recipeObj) => {
            return recipeObj.id === recipeId 
              ? {
                ...recipeObj,
                ...changesInMatchedRecipe
              } : {
                ...recipeObj,
                ...changesInUnMatchedRecipe
              }
          })
          
        }
      : deckObj
  });
  return array;
}
