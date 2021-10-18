import React, { useState, useEffect, useReducer } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useHistory, useLocation } from "react-router";
import { toast } from "react-toastify";
import PropTypes from "prop-types";

import { Logo, ButtonIcon, Text, MlModal, WhiteLight } from "shared-components";
import {
  logoutInitiated,
  loginReset,
} from "action-creators/loginActionCreators";
import {
  Button,
  Dropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
  Nav,
  NavItem,
  NavLink,
} from "core-components";
import {
  runExperiment,
  stopExperiment,
} from "action-creators/runExperimentActionCreators";
import { getExperimentId } from "selectors/experimentSelector";
import {
  getRunExperimentReducer,
  getTimeNow,
} from "selectors/runExperimentSelector";
import { getWells, getFilledWellsPosition } from "selectors/wellSelectors";
import {
  APP_TYPE,
  EXPERIMENT_STATUS,
  MODAL_BTN,
  MODAL_MESSAGE,
  ROUTES,
  TOAST_MESSAGE,
  USER_ROLES,
} from "appConstants";
import {
  NAV_ITEMS,
  getBtnPropObj,
  PATH_TO_SHOW_CROSS_BTN,
  getRedirectObj,
} from "./constants";
import { Header } from "./Header";
import { ActionBtnList, ActionBtnListItem } from "./ActionBtnList";
import { whiteLightInitiated } from "action-creators/whiteLightActionCreators";

const AppHeader = (props) => {
  const {
    role,
    isDeckBlocked,
    isUserLoggedIn,
    isPlateRoute,
    isLoginTypeAdmin,
    isLoginTypeOperator,
    isLoginTypeEngineer,
    isTemplateRoute,
    token,
    deckName,
    app,
    activeWidgetID,
  } = props;

  const location = useLocation();
  const dispatch = useDispatch();
  const history = useHistory();
  const experimentId = useSelector(getExperimentId);
  const runExperimentReducer = useSelector(getRunExperimentReducer);
  const wellListReducer = useSelector(getWells);
  const createExperimentReducer = useSelector(
    (state) => state.createExperimentReducer
  );
  const whiteLightReducer = useSelector((state) => state.whiteLightReducer);
  const recipeActionReducer = useSelector((state) => state.recipeActionReducer);
  const recipeReducerDataOfActiveDeck = recipeActionReducer.decks.find(
    (deck) => deck.name === deckName
  );

  const filledWellsPositions = getFilledWellsPosition(wellListReducer);
  const experimentStatus = runExperimentReducer.get("experimentStatus");
  const isExperimentRunning = experimentStatus === EXPERIMENT_STATUS.running;
  const isExperimentStopped = experimentStatus === EXPERIMENT_STATUS.stopped;
  const isRunFailed = experimentStatus === EXPERIMENT_STATUS.runFailed;
  const isExperimentSucceeded = experimentStatus === EXPERIMENT_STATUS.success;
  const isExpanded = createExperimentReducer.get("isExpanded");
  const result = createExperimentReducer.get("result");
  const btnProps = getBtnPropObj(result);
  const { isLightOn } = whiteLightReducer.toJS();

  const [isExitModalVisible, setExitModalVisibility] = useState(false);
  const [isWarningModalVisible, setWarningModalVisibility] = useState(false);
  const [isAbortModalVisible, setAbortModalVisibility] = useState(false);
  const [userDropdownOpen, setUserDropdownOpen] = useState(false);
  const [menuDropdownOpen, setMenuDropdownOpen] = useState(false);

  const [isExpSuccessModalVisible, setExpSuccessModalVisibility] =
    useState(false);
  const [isRunConfirmModalVisible, setRunConfirmModalVisibility] =
    useState(false);

  const [showConfirmBackModal, toggleConfirmBackModal] = useReducer(
    (showConfirmBackModal) => !showConfirmBackModal,
    false
  );

  //local state to handle extraction flow cross button confirmation modal
  const [showRedirectionModal, toggleRedirectionModal] = useReducer(
    (showRedirectionModal) => !showRedirectionModal,
    false
  );

  const toggleUserDropdown = () => {
    setUserDropdownOpen((prevState) => !prevState);
  };

  const toggleMenuDropdown = () => {
    setMenuDropdownOpen((prevState) => !prevState);
  };

  useEffect(() => {
    if (isExperimentSucceeded) {
      setExpSuccessModalVisibility(true);
    }
  }, [isExperimentSucceeded]);

  // logout user
  const logoutClickHandler = () => {
    setExitModalVisibility(true);
  };

  const startExperiment = () => {
    // if no well is filled then show confirmation modal.
    if (filledWellsPositions.toJS().length === 0) {
      //showModal = True
      setRunConfirmModalVisibility(true);
    } else if (
      isExperimentRunning === false &&
      isExperimentSucceeded === false
    ) {
      dispatch(runExperiment(experimentId, token));
    }
  };

  const handleAbortConrfirmation = () => {
    dispatch(stopExperiment(experimentId, token));
    setAbortModalVisibility(false);
  };

  /** Hide plates tab accordingly */
  const getIsNavLinkHidden = (pathname) => {
    // type : rtpcr
    if (app === APP_TYPE.RTPCR) {
      if (isLoginTypeEngineer === true && pathname !== ROUTES.calibration) {
        // hide all for engineer except for calibrations
        return true;
      } else if (
        isLoginTypeOperator === true &&
        pathname === ROUTES.calibration
      ) {
        // hide only calibrations for operator
        return true;
      } else if (isLoginTypeAdmin === true && pathname === ROUTES.plate) {
        // hide only plates for operator
        return true;
      }
    }
    // type extraction
    else {
      if (isLoginTypeEngineer === true && pathname !== ROUTES.calibration) {
        // hide all for engineer except for calibrations
        return true;
      } else if (isLoginTypeOperator === true) {
        // hide everything for operator in extraction
        return true;
      } else if (
        isLoginTypeAdmin === true &&
        (pathname !== ROUTES.calibration ||
          recipeReducerDataOfActiveDeck.showProcess)
      ) {
        // hide everything except calibrations for admin in extractions
        // AND hide calibrations when process is running
        return true;
      }
    }
    return false;
  };

  const getIsNavLinkDisabled = (pathname) => {
    switch (pathname) {
      /* Disable plate navlink if user logged in is admin. Admin just has access to templates
		and activity log. Also user can't navigate to plate directly from templates route
		without selecting a template. isTemplateRoute is true untill user selects a template */
      case "/plate":
        if (isLoginTypeAdmin === false && isExpanded === true) {
          return false;
        } else if (isLoginTypeAdmin === true || isTemplateRoute === true) {
          return true;
        }
        return false;

      /* Disable Templates navlink when isPlateRoute is true or
		after experiment is started */
      case "/templates":
        if (isPlateRoute === true || experimentStatus !== null) {
          return true;
        }
        return false;

      /* Disable Activity log navlink while experiment is running */
      case "/activity":
        if (isExperimentRunning === true) {
          return true;
        }
        return false;

      case "/calibration":
        if (
          app === APP_TYPE.EXTRACTION &&
          isLoginTypeAdmin === true &&
          isDeckBlocked === true
        ) {
          return true;
        }
        return false;

      default:
        return false;
    }
  };

  const onNavLinkClickHandler = (event, pathname) => {
    if (getIsNavLinkDisabled(pathname)) {
      if (pathname === `/${ROUTES.calibration}`) {
        toast.warning(TOAST_MESSAGE.calRedirect, { autoClose: false });
      }
      event.preventDefault();
    }
  };

  // Exit modal confirmation click handler
  const confirmationClickHandler = (isConfirmed) => {
    setExitModalVisibility(false);
    if (isExperimentRunning === true) {
      // show warning that user needs to abort first in order to log out.
      setWarningModalVisibility(true);
    } else {
      // user log out
      dispatch(logoutInitiated({ deckName: deckName, token: token }));
    }
  };

  // view results modal button click
  const handleSuccessModalConfirmation = () => {
    setExpSuccessModalVisibility(false);
    // redirect to activity
    history.push("activity");
  };

  // if user selects 'yes', run experiment
  const handleRunModalConfirmation = () => {
    setRunConfirmModalVisibility(false);
    dispatch(runExperiment(experimentId, token));
  };

  const handleBackBtn = () => {
    toggleConfirmBackModal();
    history.push("templates");
  };

  const handleRedirectionButton = () => {
    toggleRedirectionModal();
    const currentPathname = location.pathname;
    const redirectPath = getRedirectObj(currentPathname).redirectPath;
    history.push(redirectPath);
  };

  const handleManageUsersClick = () => {
    history.push(ROUTES.users);
  };

  const handleWhiteLightClick = () => {
    if (isLightOn === true) {
      // dispatch(whiteLightInitiated())
    } else {
      // dispatch(whiteLightInitiated())
    }
  };

  const handleHelpSupportBtn = () => {
    toggleMenuDropdown();
    // model => open
  };

  return (
    <Header>
      <Logo isUserLoggedIn={isUserLoggedIn} app={app} />
      {isUserLoggedIn && (
        <Nav className="ml-3 mr-auto">
          {NAV_ITEMS.map(
            (navItem) =>
              !getIsNavLinkHidden(navItem.path) && (
                <NavItem key={navItem.name}>
                  <NavLink
                    onClick={(event) => {
                      onNavLinkClickHandler(event, `/${navItem.path}`);
                    }}
                    to={navItem.path}
                    disabled={getIsNavLinkDisabled(`/${navItem.path}`)}
                  >
                    {navItem.name}
                  </NavLink>
                </NavItem>
              )
          )}
        </Nav>
      )}

      <>
        {/**extraction flow cross button used in process creation/edition flow */}
        {PATH_TO_SHOW_CROSS_BTN.includes(location.pathname) && (
          <ButtonIcon
            size={34}
            name="cross"
            onClick={toggleRedirectionModal}
            className="ml-auto"
          />
        )}

        {/* For new menu items, we need to toggle it manually. */}
        {app === APP_TYPE.EXTRACTION && (
          <Dropdown
            className="ml-2"
            isOpen={menuDropdownOpen}
            toggle={toggleMenuDropdown}
          >
            <DropdownToggle icon name="menu" size={16} />
            <DropdownMenu right onClick={toggleMenuDropdown}>
              <DropdownItem>
                <WhiteLight
                  isLightOn={isLightOn}
                  handleWhiteLightClick={handleWhiteLightClick}
                />
              </DropdownItem>
              {/* <DropdownItem onClick={handleHelpSupportBtn}>
              {"Help & Support"}
            </DropdownItem> */}
            </DropdownMenu>
          </Dropdown>
        )}
      </>

      {isUserLoggedIn && (
        <div className="header-elements d-flex align-items-center ml-auto">
          {/* <PrintDataModal /> */}
          {/* <ExportDataModal /> */}
          {app === APP_TYPE.RTPCR && (
            <>
              <div className="experiment-info text-right mx-3">
                <Text
                  size={12}
                  className={`text-default ${
                    isExperimentRunning ? "show" : ""
                  }`}
                >
                  {`Experiment started at ${runExperimentReducer.get(
                    "experimentStartedTime"
                  )}`}
                </Text>
                <Text
                  size={12}
                  className={`text-error ${isRunFailed ? "show" : ""}`}
                >
                  Experiment failed to run.
                </Text>

                <Text
                  size={12}
                  className={`text-error ${isExperimentStopped ? "show" : ""}`}
                >
                  {`Experiment aborted at ${runExperimentReducer.get(
                    "experimentStoppedTime"
                  )}.`}
                </Text>

                {isPlateRoute === true && (
                  <div className="d-flex align-items-center">
                    <Button
                      color="secondary"
                      size="sm"
                      className={`font-weight-light border-2 border-gray shadow-none mr-3`}
                      outline={true}
                      onClick={toggleConfirmBackModal}
                      disabled={isExperimentRunning}
                    >
                      Back
                    </Button>

                    <Button
                      color={isExperimentSucceeded ? "primary" : "secondary"}
                      size="sm"
                      className={`font-weight-light border-2 border-gray shadow-none  mr-3 
                      ${isExperimentSucceeded || isExpanded ? "d-none" : ""}`}
                      onClick={() => setAbortModalVisibility(true)}
                      disabled={!isExperimentRunning}
                    >
                      Abort
                    </Button>
                    <Button
                      color={isExperimentRunning ? "primary" : "secondary"}
                      size="sm"
                      className={`font-weight-light border-2 border-gray shadow-none ${
                        isExperimentSucceeded || isExpanded ? "d-none" : ""
                      }`}
                      outline={
                        isExperimentRunning === false &&
                        isExperimentSucceeded === false
                      }
                      onClick={startExperiment}
                      disabled={
                        isExperimentRunning ||
                        isExperimentSucceeded ||
                        isExperimentStopped
                      }
                    >
                      Run
                    </Button>
                    <Button
                      color={btnProps.color}
                      size="sm"
                      className={`font-weight-light border-2 border-gray shadow-none
                       ${
                         isExperimentSucceeded && isExpanded === false
                           ? ""
                           : "d-none"
                       }`}
                      onClick={() => setExpSuccessModalVisibility(true)}
                    >
                      {btnProps.msg}
                    </Button>
                  </div>
                )}

                {isLoginTypeAdmin && activeWidgetID === "step" && (
                  <Button
                    color="primary"
                    size="sm"
                    className="font-weight-light border-2 border-gray shadow-none"
                    // onClick={startExperiment}
                  >
                    Save
                  </Button>
                )}
              </div>
              <div className="user-dropdown-wrapper position-relative ml-2">
                <Text
                  size={10}
                  className="user position-absolute font-weight-bold text-capitalize my-auto"
                >
                  {role || ""}
                </Text>

                <Dropdown isOpen={userDropdownOpen} toggle={toggleUserDropdown}>
                  <DropdownToggle icon name="user" size={32} />
                  <DropdownMenu right>
                    {/**manage users accessible only for admin */}
                    {role === USER_ROLES.ADMIN && (
                      <DropdownItem onClick={handleManageUsersClick}>
                        Manage Users
                      </DropdownItem>
                    )}
                    <DropdownItem
                      onClick={logoutClickHandler}
                      disabled={isExperimentRunning}
                    >
                      Log out
                    </DropdownItem>
                  </DropdownMenu>
                </Dropdown>
              </div>
            </>
          )}

          {showRedirectionModal && (
            <MlModal
              isOpen={showRedirectionModal}
              successBtn={MODAL_BTN.yes}
              failureBtn={MODAL_BTN.no}
              handleSuccessBtn={handleRedirectionButton}
              handleCrossBtn={toggleRedirectionModal}
              textHead={deckName}
              textBody={getRedirectObj(location.pathname).msg}
            />
          )}

          {/* MODALS */}

          {isExitModalVisible && (
            <MlModal
              isOpen={isExitModalVisible}
              textBody={MODAL_MESSAGE.exitConfirmation}
              successBtn={MODAL_BTN.yes}
              failureBtn={MODAL_BTN.no}
              handleSuccessBtn={confirmationClickHandler}
              handleCrossBtn={() => setExitModalVisibility(false)}
            />
          )}
          {isWarningModalVisible && (
            <MlModal
              isOpen={isWarningModalVisible}
              textBody={MODAL_MESSAGE.abortExpInfo}
              successBtn={MODAL_BTN.okay}
              handleSuccessBtn={() => setWarningModalVisibility(false)}
              handleCrossBtn={() => setWarningModalVisibility(false)}
            />
          )}
          {isAbortModalVisible && (
            <MlModal
              isOpen={isAbortModalVisible}
              textBody={MODAL_MESSAGE.abortExpWarning}
              successBtn={MODAL_BTN.yes}
              failureBtn={MODAL_BTN.no}
              handleSuccessBtn={handleAbortConrfirmation}
              handleCrossBtn={() => setAbortModalVisibility(false)}
            />
          )}
          {isExpSuccessModalVisible && (
            <MlModal
              isOpen={isExpSuccessModalVisible}
              textBody={MODAL_MESSAGE.experimentSuccess}
              successBtn={MODAL_BTN.viewResults}
              failureBtn={MODAL_BTN.cancel}
              handleSuccessBtn={handleSuccessModalConfirmation}
              handleCrossBtn={() => setExpSuccessModalVisibility(false)}
            />
          )}

          {isRunConfirmModalVisible && (
            <MlModal
              isOpen={isRunConfirmModalVisible}
              textBody={MODAL_MESSAGE.runConfirmMsg}
              successBtn={MODAL_BTN.yes}
              failureBtn={MODAL_BTN.cancel}
              handleSuccessBtn={handleRunModalConfirmation}
              handleCrossBtn={() => setRunConfirmModalVisibility(false)}
            />
          )}

          {showConfirmBackModal && (
            <MlModal
              isOpen={showConfirmBackModal}
              textBody={MODAL_MESSAGE.backConfirmation}
              successBtn={MODAL_BTN.yes}
              failureBtn={MODAL_BTN.cancel}
              handleSuccessBtn={handleBackBtn}
              handleCrossBtn={toggleConfirmBackModal}
            />
          )}
        </div>
      )}
    </Header>
  );
};

AppHeader.propTypes = {
  isUserLoggedIn: PropTypes.bool,
};

AppHeader.defaultProps = {
  isUserLoggedIn: false,
};

export default React.memo(AppHeader);
