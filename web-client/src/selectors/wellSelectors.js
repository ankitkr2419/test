import { MAX_NO_OF_WELLS } from "appConstants";
import { createSelector } from "reselect";

const getWellListSelector = (state) => state.wellListReducer;
const getAddWellsSelector = (state) => state.addWellsReducer;

export const getWells = createSelector(
  getWellListSelector,
  (wellListReducer) => wellListReducer
);

/**
 * return array of selected position
 * array wells that have isSelected or isMultiSelected field set to true
 */
export const getWellsPosition = createSelector(
  (wellListReducer) => wellListReducer,
  (wellListReducer) =>
    wellListReducer
      .get("defaultList")
      .map((ele, index) => {
        if (
          ele !== null &&
          (ele.get("isSelected") === true ||
            ele.get("isMultiSelected") === true)
        ) {
          return index;
        }
        return null;
      })
      .filter((ele) => ele !== null)
);

/**
 * returns array of indexes of filled wells
 * Here filled wells means that the wells that are selected 
 * and filled with information for - targets, samples and tasks
*/
export const getFilledWellsPosition = createSelector(
  (wellListReducer) => wellListReducer,
  (wellListReducer) =>
    wellListReducer
      .get("defaultList")
      .map((ele, index) => {
        if (ele !== null && ele.get("isWellFilled") === true) {
          return index;
        }
        return null;
      })
      .filter((ele) => ele !== null)
);
// returns array of targets_ids configured for a filled well
export const getFilledWellTargets = createSelector(
  (wellListReducer) => wellListReducer,
  (_, wellPosition) => wellPosition,
  (wellListReducer, wellPosition) => {
    const temp = wellListReducer
      .getIn(["defaultList", wellPosition, "targets"])
      .map((ele) => ele.target_id);
    return temp;
  }
);
// set isSelected flag to true for given index
export const setSelectedToList = (state, { isSelected, index }) =>
  state.setIn(["defaultList", index, "isSelected"], isSelected);
// set isMultiSelected flag to true for given index
export const setMultiSelectedToList = (state, { isMultiSelected, index }) =>
  state.setIn(["defaultList", index, "isMultiSelected"], isMultiSelected);
// makes all wells isSelected flag to  false
export const resetWellDefaultList = (state) =>
  state.updateIn(["defaultList"], (myDefaultList) =>
    myDefaultList.map((ele) => ele && ele.setIn(["isSelected"], false))
  );
// makes all wells isMultiSelected flag to  false
export const resetMultiWellDefaultList = (state) =>
  state.updateIn(["defaultList"], (myDefaultList) =>
    myDefaultList.map((ele) => ele && ele.setIn(["isMultiSelected"], false))
  );
// makes all wells isSelected flag to  true
export const setAllWellsSelected = (state, { isAllWellsSelected }) =>
  state.updateIn(["defaultList"], (myDefaultList) =>
    myDefaultList.map(
      (ele) => ele && ele.setIn(["isSelected"], !isAllWellsSelected)
    )
  );

/**
 * getDefaultPlatesList return wells default data w.r.t MAX_NO_OF_WELLS.
 */
export const getDefaultWellsList = createSelector(() => {
  const arr = [];
  // insert a null value at beginning of list so that wells data will start from index 1
  // we have the wells numbered from from 1 and we use the index of each element in array
  // as it's well number so to maintain the number ordering we have added a null at start
  arr.push(null);
  const initialPlateState = {
    isSelected: false,
    isWellFilled: false,
    isRunning: false,
    isMultiSelected: false,
    status: "", // red, green, orange
    initial: "",
    id: null,
    isWellActive: false
  };
  // Initial plate state added for wells in array from index 1
  for (let i = 0; i !== MAX_NO_OF_WELLS; i += 1) {
    arr.push(initialPlateState);
  }
  return arr;
});

/**
 * Return isWellSaved flag from addWellReducer
 */
export const getWellsSavedStatus = createSelector(
  getAddWellsSelector,
  (addWellReducer) => addWellReducer.get("isWellSaved")
);
/**
 * Get well data with w.r.t position
 * @param {*} wells
 * @param {*} position
 */
const getSelectedWell = (wells, position) =>
  wells.filter((ele) => ele.position === position)[0];

const getSelectedTargets = (selectedWell) =>
  selectedWell.targets.filter((ele) => ele.selected);
/**
 * updateWellListSelector accepts current state and action
 * It will update default list with updated action response
 * It will fill the popover data and well config data
 */
export const updateWellListSelector = createSelector(
  (state) => state,
  (state, action) => action,
  (state, action) => {
    // if no wells selected
    if (action.payload.response !== null) {
      const {
        payload: { response }
      } = action;
      // get position of selected wells from response
      const positions = response.map((ele) => ele.position);
      // iteration over default map
      const tempState = state.updateIn(["defaultList"], (myDefaultList) =>
        myDefaultList.map((ele, index) => {
          // find the index present in response data
          if (positions.includes(index)) {
            // get selected well data by index
            const selectedWell = getSelectedWell(response, index);
            // filter the targets to get only the selected targets with selected property true
            selectedWell.targets = getSelectedTargets(selectedWell);
            // merge selected well data and modify local fields.
            return ele.merge({
              isWellFilled: true,
              ...selectedWell,
              initial: selectedWell.task[0],
              status: selectedWell.color_code || "green",
              sample: selectedWell.sample_name
            });
          }
          return ele;
        })
      );
      return tempState.merge({ isLoading: false, isWellFilled: true });
    }
    return state.merge({ isLoading: false, isWellFilled: false });
  }
);

/**
 * configuring active wells with existing wells data
 */
export const setActiveWells = createSelector(
  (state) => state,
  (state, action) => action,
  (state, action) => {
    if (action.payload.response !== null) {
      const activeWellsPositions = action.payload.response;
      const tempState = state.updateIn(["defaultList"], (myDefaultList) =>
        myDefaultList.map((ele, index) => {
          // find the index present in response data
          if (activeWellsPositions.includes(index) && ele !== null) {
            return ele.merge({
              isWellActive: true
            });
          }
          return ele;
        })
      );
      return tempState;
    }
    return state;
  }
);
