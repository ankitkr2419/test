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
  const experimentId = useSelector(getExperimentId);
  const runExperimentReducer = useSelector(getRunExperimentReducer);
  const experimentStatus = runExperimentReducer.get("experimentStatus");
  const isExperimentRunning = experimentStatus === EXPERIMENT_STATUS.running;
  const isExperimentStopped = experimentStatus === EXPERIMENT_STATUS.stopped;
  const isRunFailed = experimentStatus === EXPERIMENT_STATUS.runFailed;
  const isExperimentSucceeded = experimentStatus === EXPERIMENT_STATUS.success;

  const [isExitModalVisible, setExitModalVisibility] = useState(false);
  const [userDropdownOpen, setUserDropdownOpen] = useState(false);
  const toggleUserDropdown = () =>
    setUserDropdownOpen((prevState) => !prevState);

  useEffect(() => {
    if (isExperimentStopped === true) {
      dispatch(loginReset());
    }
  }, [isExperimentStopped, dispatch]);

  // logout user
  const logoutClickHandler = () => {
    dispatch(logoutInitiated({ deckName: deckName, token: token }));
  };

  const startExperiment = () => {
    if (isExperimentRunning === false && isExperimentSucceeded === false) {
      dispatch(runExperiment(experimentId, token));
    }
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
    setExitModalVisibility(true);
  };

  // Exit modal confirmation click handler
  const confirmationClickHandler = (isConfirmed) => {
    setExitModalVisibility(false);
    if (isConfirmed === true) {
      if (isExperimentRunning === true) {
        // user aborted experiment
        dispatch(stopExperiment(experimentId, token));
      } else {
        dispatch(logoutInitiated({ deckName: deckName, token: token }));
      }
    }
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
                className={`text-default mb-1 ${
                  isExperimentRunning ? "show" : ""
                }`}
              >
                {`Experiment started at ${runExperimentReducer.get(
                  "experimentStartedTime"
                )}`}
              </Text>
              <Text
                size={12}
                className={`text-error mb-1 ${isRunFailed ? "show" : ""}`}
              >
                Experiment failed to run.
              </Text>
              
              {isExperimentSucceeded === false && isPlateRoute === true && (
                <Button
                  color={isExperimentRunning ? "primary" : "secondary"}
                  size="sm"
                  className="font-weight-light border-2 border-gray shadow-none"
                  outline={
                    isExperimentRunning === false &&
                    isExperimentSucceeded === false
                  }
                  onClick={startExperiment}
                >
                  Run
                </Button>
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
