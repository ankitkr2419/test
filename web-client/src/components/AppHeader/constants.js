import { ROUTES } from "appConstants";

export const NAV_ITEMS = [
  {
    name: "Template",
    path: ROUTES.templates,
  },
  {
    name: "Plate",
    path: ROUTES.plate,
  },
  {
    name: "Activity Log",
    path: ROUTES.activity,
  },
  {
    name: "Calibration",
    path: ROUTES.calibration,
  },
];

export const PATH_TO_SHOW_CROSS_BTN = [
  "/process-listing",
  "/select-process",
  "/piercing",
  "/tip-pickup",
  "/aspire-dispense",
  "/shaking",
  "/heating",
  "/magnet",
  "/tip-discard",
  "/delay",
  "/tip-poistion",
];

export const getRedirectObj = (currentPathname) => {
  switch (currentPathname) {
    case "/select-process":
      return {
        redirectPath: ROUTES.recipeListing,
        msg: "Are you sure you want to exit adding process?",
      };
    case "/process-listing":
      return {
        redirectPath: ROUTES.recipeListing,
        msg: "Are you sure you want to exit editing?",
      };
    default:
      return {
        redirectPath: ROUTES.processListing,
        msg: "Are you sure you want to exit?",
      };
  }
};

export const getBtnPropObj = (result) => {
  switch (result) {
    case "success":
      return { msg: "Result - Successful", color: "success" };
    case "aborted":
      return { msg: "Result - Aborted", color: "danger" };
    case "running":
      return { msg: "Running", color: "info" };
    default:
      return { msg: "Result - NA", color: "info" };
  }
};
