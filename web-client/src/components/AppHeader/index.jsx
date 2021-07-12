import React, { useState, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import PropTypes from "prop-types";
import { Logo, ButtonIcon, Text, Icon, MlModal } from "shared-components";
import {
  loginReset,
  logoutInitiated,
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
// import { getExperimentId } from 'selectors/experimentSelector';
import {
  runExperiment,
  stopExperiment,
} from "action-creators/runExperimentActionCreators";
import { getExperimentId } from "selectors/experimentSelector";
import { getRunExperimentReducer } from "selectors/runExperimentSelector";
import { getWells, getFilledWellsPosition } from "selectors/wellSelectors";
// import PrintDataModal from './PrintDataModal';
// import ExportDataModal from './ExportDataModal';
import {
  APP_TYPE,
  EXPERIMENT_STATUS,
  MODAL_BTN,
  MODAL_MESSAGE,
  ROUTES,
} from "appConstants";
import { NAV_ITEMS } from "./constants";
import { Header } from "./Header";
import { ActionBtnList, ActionBtnListItem } from "./ActionBtnList";
import { useHistory } from "react-router";

const AppHeader = (props) => {
  const {
    isUserLoggedIn,
    isPlateRoute,
    isLoginTypeAdmin,
    isLoginTypeOperator,
    isTemplateRoute,
    token,
    deckName,
    app,
    activeWidgetID,
  } = props;

  const dispatch = useDispatch();
  const history = useHistory();
  const experimentId = useSelector(getExperimentId);
  const runExperimentReducer = useSelector(getRunExperimentReducer);
  const wellListReducer = useSelector(getWells);

  const filledWellsPositions = getFilledWellsPosition(wellListReducer);
  const experimentStatus = runExperimentReducer.get("experimentStatus");
  const isExperimentRunning = experimentStatus === EXPERIMENT_STATUS.running;
  const isExperimentStopped = experimentStatus === EXPERIMENT_STATUS.stopped;
  const isRunFailed = experimentStatus === EXPERIMENT_STATUS.runFailed;
  const isExperimentSucceeded = experimentStatus === EXPERIMENT_STATUS.success;

  const [isExitModalVisible, setExitModalVisibility] = useState(false);
  const [isWarningModalVisible, setWarningModalVisibility] = useState(false);
  const [isAbortModalVisible, setAbortModalVisibility] = useState(false);
  const [userDropdownOpen, setUserDropdownOpen] = useState(false);
  const [isExpSuccessModalVisible, setExpSuccessModalVisibility] =
    useState(false);
  const [isRunConfirmModalVisible, setRunConfirmModalVisibility] =
    useState(false);

  const toggleUserDropdown = () =>
    setUserDropdownOpen((prevState) => !prevState);

  // useEffect(() => {
  // 	if (isExperimentRunning === true) {
  // 		// connectSocket(dispatch);
  // 	}
  // }, [isExperimentRunning, dispatch]);

  useEffect(() => {
    if (isExperimentSucceeded) {
      setExpSuccessModalVisibility(true);
    }
  }, [isExperimentSucceeded]);

  useEffect(() => {
    if (isExperimentStopped === true) {
      // disConnectSocket();
      dispatch(loginReset());
    }
  }, [isExperimentStopped, dispatch]);

  // logout user
  const logoutClickHandler = () => {
    dispatch(logoutInitiated({ deckName: deckName, token: token }));
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

  /** Hide plates tab if the user is admin */
  const getIsNavLinkHidden = (pathname) => {
    if (pathname === "/plate" && isLoginTypeAdmin === true) {
      return true;
    }
    return false;
  };

  const getIsNavLinkDisabled = (pathname) => {
    switch (pathname) {
      /* Disable plate navlink if user logged in is admin. Admin just has access to templates
		and activity log. Also user can't navigate to plate directly from templates route
		without selecting a template. isTemplateRoute is true untill user selects a template */
      case "/plate":
        if (isLoginTypeAdmin === true || isTemplateRoute === true) {
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

      default:
        return false;
    }
  };

  const onNavLinkClickHandler = (event, pathname) => {
    if (getIsNavLinkDisabled(pathname)) {
      event.preventDefault();
    }
  };

  const onCrossClick = () => {
    if (isExperimentRunning === true) {
      setWarningModalVisibility(true);
    } else {
      setExitModalVisibility(true);
    }
  };

  // Exit modal confirmation click handler
  const confirmationClickHandler = (isConfirmed) => {
    setExitModalVisibility(false);
    if (isConfirmed) {
      if (isExperimentRunning === true) {
        // show warning that user needs to abort first in order to log out.
        setWarningModalVisibility(true);
      } else {
        // user log out
        dispatch(logoutInitiated({ deckName: deckName, token: token }));
      }
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

  return (
    <Header>
      <Logo isUserLoggedIn={isUserLoggedIn} />
      {isUserLoggedIn && app === APP_TYPE.RTPCR && (
        <Nav className="ml-3 mr-auto">
          {NAV_ITEMS.map(
            (ele) =>
              !getIsNavLinkHidden(ele.path) && (
                <NavItem key={ele.name}>
                  <NavLink
                    onClick={(event) => {
                      onNavLinkClickHandler(event, ele.path);
                    }}
                    to={ele.path}
                    disabled={getIsNavLinkDisabled(ele.path)}
                  >
                    {ele.name}
                  </NavLink>
                </NavItem>
              )
          )}
        </Nav>
      )}
      {isUserLoggedIn && (
        <div className="header-elements d-flex align-items-center">
          {/* <PrintDataModal /> */}
          {/* <ExportDataModal /> */}
          {app === APP_TYPE.RTPCR && (
            <div className="experiment-info text-right mx-3">
              <Text
                size={12}
                className={`text-default ${isExperimentRunning ? "show" : ""}`}
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

              {isExperimentSucceeded === false && isPlateRoute === true && (
                <div className="d-flex align-items-center">
                  <Button
                    color={isExperimentRunning ? "primary" : "secondary"}
                    size="sm"
                    className="font-weight-light border-2 border-gray shadow-none mr-3"
                    outline={
                      isExperimentRunning === false &&
                      isExperimentSucceeded === false
                    }
                    onClick={startExperiment}
                    disabled={isExperimentRunning}
                  >
                    Run
                  </Button>

                  <Button
                    color={isExperimentSucceeded ? "primary" : "danger"}
                    size="sm"
                    className="font-weight-light border-2 border-gray shadow-none"
                    onClick={() => setAbortModalVisibility(true)}
                    disabled={!isExperimentRunning}
                  >
                    Abort
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
              {isExperimentSucceeded === true && (
                <Button
                  color="success"
                  size="sm"
                  className="font-weight-light border-2 border-gray shadow-none"
                  onClick={() => setExpSuccessModalVisibility(true)}
                >
                  Result - Successful
                </Button>
              )}
            </div>
          )}

          {isLoginTypeAdmin === true && (
            <Dropdown
              isOpen={userDropdownOpen}
              toggle={toggleUserDropdown}
              className="ml-2"
            >
              <DropdownToggle icon name="user" size={32} />
              <DropdownMenu right>
                <DropdownItem onClick={logoutClickHandler}>
                  Log out
                </DropdownItem>
              </DropdownMenu>
            </Dropdown>
          )}
          {isLoginTypeOperator === true && (
            <ButtonIcon
              size={34}
              name="cross"
              onClick={onCrossClick}
              className="ml-2"
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

          <ActionBtnList className="d-flex justify-content-between align-items-center list-unstyled mb-0">
            <ActionBtnListItem>
              <Icon name="setting" size={18} />
            </ActionBtnListItem>
            <ActionBtnListItem>
              <Icon name="notifications" size={18} />
            </ActionBtnListItem>
            <ActionBtnListItem>
              <Icon name="menu" size={18} />
            </ActionBtnListItem>
          </ActionBtnList>
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
