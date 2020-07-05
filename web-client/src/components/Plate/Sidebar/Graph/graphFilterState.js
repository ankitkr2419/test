import { fromJS, List } from "immutable";

const targetColors = ["#9eff48", "#889accf7", "#fd5d55", "yellow", "#666666", "#f3811f"];

// const action types
export const graphFilterActions = {
  UPDATE_GRAPH_FILTER_STATE: "UPDATE_GRAPH_FILTER_VALUES",
  SET_GRAPH_FILTER_VALUES: "SET_GRAPH_FILTER_VALUES",
  RESET_GRAPH_FILTER_STATE: "RESET_GRAPH_FILTER_STATE",
  UPDATE_TARGET_LIST: "UPDATE_TARGET_LIST",
  UPDATE_GRAPH_FILTER_ACTIVE_STATE: "UPDATE_GRAPH_FILTER_ACTIVE_STATE",
};

export const graphFilterInitialState = fromJS({
  isSidebarOpen: false,
  targets: List([]),
});

const graphFilterState = (state, action) => {
  switch (action.type) {
    case graphFilterActions.SET_GRAPH_FILTER_VALUES:
      if (action.key === "targets") {
        return state
          .setIn([action.key], action.value)
          .updateIn(["targets"], (myDefaultList) =>
            myDefaultList.map((ele, index) => ele.merge({ isActive: true, lineColor: targetColors[index] }))
          );
      }
      return state.setIn([action.key], action.value);
    case graphFilterActions.UPDATE_GRAPH_FILTER_STATE:
      return state.merge(action.value);
    case graphFilterActions.UPDATE_TARGET_LIST:
      return state.setIn(
        ["targets", action.index, "threshold"],
        action.threshold
      );
    case graphFilterActions.UPDATE_GRAPH_FILTER_ACTIVE_STATE:
      return state.setIn(
        ["targets", action.index, "isActive"],
        action.isActive
      );
    case graphFilterActions.RESET_GRAPH_FILTER_STATE:
      return graphFilterInitialState;
    default:
      throw new Error("Invalid action type");
  }
};

export default graphFilterState;
