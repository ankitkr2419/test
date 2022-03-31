import { pageBtnAction } from "actions/PageActions";

export const nextRecipePageBtn = (payload) => ({
  type: pageBtnAction.nextRecipePageBtn,
  payload: payload,
});

export const prevRecipePageBtn = (payload) => ({
  type: pageBtnAction.previousRecipePageBtn,
  payload: payload,
});

export const nextProcessPageBtn = (payload) => ({
  type: pageBtnAction.nextProcessPageBtn,
  payload: payload,
});

export const prevProcessPageBtn = (payload) => ({
  type: pageBtnAction.previousProcessPageBtn,
  payload: payload,
});

export const ResetRecipePageState = (payload) => ({
  type: pageBtnAction.ResetRecipePageState,
  payload: payload,
});

export const ResetProcessPageState = (payload) => ({
  type: pageBtnAction.ResetProcessPageState,
  payload: payload,
});
