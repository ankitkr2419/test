import { addWellActions, listWellActions } from "actions/wellActions";

export const addWell = (experimentId, body, token) => ({
  type: addWellActions.addAction,
  payload: {
    body,
    experimentId,
    token,
  },
});

export const addWellFailed = (errorResponse) => ({
  type: addWellActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const addWellReset = () => ({
  type: addWellActions.addWellReset,
});

export const fetchWells = (experimentId, token) => ({
  type: listWellActions.listAction,
  payload: {
    experimentId,
    token,
  },
});

export const fetchWellsFailed = (errorResponse) => ({
  type: listWellActions.failureAction,
  payload: {
    ...errorResponse,
    error: true,
  },
});

export const updateWellThroughSocket = (response) => ({
  type: listWellActions.updateWellThroughSocket,
  payload: {
    response,
  },
});

export const setSelectedWell = (index, isSelected) => ({
  type: listWellActions.setSelectedWell,
  payload: {
    isSelected,
    index,
  },
});

export const resetSelectedWells = () => ({
  type: listWellActions.resetSelectedWell,
});

export const setMultiSelectedWell = (index, isMultiSelected) => ({
  type: listWellActions.setMultiSelectedWell,
  payload: {
    isMultiSelected,
    index,
  },
});

export const resetMultiSelectedWells = () => ({
  type: listWellActions.resetMultiSelectedWell,
});

export const toggleMultiSelectOption = () => ({
  type: listWellActions.toggleMultiSelectOption,
});

export const selectAllWellsOption = (isAllWellsSelected) => ({
  type: listWellActions.selectAllWellsOption,
  payload: {
    isAllWellsSelected,
  },
});
