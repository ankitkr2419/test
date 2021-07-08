import { fromJS } from "immutable";
import { wizardList } from "./templateConstant";

// const action types
export const templateLayoutActions = {
  SET_ACTIVE_WIDGET: "SET_ACTIVE_WIDGET",
  SET_TEMPLATE_ID: "SET_TEMPLATE_ID",
  SET_STAGE_ID: "SET_STAGE_ID",
  RESET_WIZARD_LIST: "RESET_WIZARD_LIST",
};

// Initial state wrap with fromJS for immutability
export const templateInitialState = fromJS({
  // active wizard id
  activeWidgetID: "template",
  // Pre-filled template initial list with saved wizard list
  wizardList,
  // templateID: '948297f7-4166-49af-bad0-4a1742d9ca8a',
  // stageId: 'd57185b8-952b-4369-a4a3-6317de7f35dd',
  templateID: null,
});

// getUpdateList will update all disabled to true and set false to selected wizard
const getUpdatedList = (state, selectedID) => {
  let isDisable = false;
  let updatedState = state.update("wizardList", (item) =>
    item.map((keyValue) => {
      if (keyValue.get("id") === selectedID) {
        isDisable = true;
        return keyValue.set("isDisable", false);
      } else {
        return keyValue.set("isDisable", isDisable);
      }
    })
  );
  updatedState = updatedState.setIn(["activeWidgetID"], selectedID);
  return updatedState;
};

export const getWizardListByLoginType = (list, isAdmin) => {
  // return all options for admin
  if (isAdmin === true) {
    return list.filter((ele) => ele.get("isAdmin") === true);
  }
  // return option visible to operator
  return list.filter((ele) => ele.get("isOperatorVisible") === !isAdmin);
};

const templateLayoutReducer = (state, action) => {
  switch (action.type) {
    case templateLayoutActions.SET_ACTIVE_WIDGET:
      return getUpdatedList(state, action.value);
    case templateLayoutActions.RESET_WIZARD_LIST:
      return state.set("wizardList", fromJS(wizardList));
    case templateLayoutActions.SET_TEMPLATE_ID:
      return state.setIn(["templateID"], action.value);
    default:
      throw new Error("Invalid action type");
  }
};

export default templateLayoutReducer;
