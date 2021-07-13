import React, { useCallback, useMemo, useEffect } from "react";
import PropTypes from "prop-types";
import { useSelector, useDispatch } from "react-redux";
import {
  createTemplate as createTemplateAction,
  updateTemplate as updateTemplateAction,
  updateTemplateReset,
  fetchTemplates,
} from "action-creators/templateActionCreators";
import TemplateModal from "components/TemplateModal/TemplateModal";
import { templateModalActions } from "components/TemplateModal/templateModalState";

const TemplateModalContainer = (props) => {
  // constants
  const {
    templateModalState,
    templateModalDispatch,
    templateID,
    toggleTemplateModal,
  } = props;

  const {
    templateDescription,
    templateName,
    volume,
    lidTemperature,
    isCreateTemplateModalVisible,
    isTemplateEdited,
  } = templateModalState.toJS();

  const dispatch = useDispatch();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  // useSelector section
  const { isTemplateUpdated, errorUpdatingTemplate } = useSelector(
    (state) => state.updateTemplateReducer
  );
  // reading templates from redux
  const templates = useSelector((state) => state.listTemplatesReducer).get(
    "list"
  );

  // helper method section
  // helper method to set template name
  const setTemplateName = useCallback(
    (name) => {
      templateModalDispatch({
        type: templateModalActions.SET_TEMPLATE_NAME,
        value: name,
      });
    },
    [templateModalDispatch]
  );

  // helper method to set template description
  const setTemplateDescription = useCallback(
    (description) => {
      templateModalDispatch({
        type: templateModalActions.SET_TEMPLATE_DESCRIPTION,
        value: description,
      });
    },
    [templateModalDispatch]
  );

  // helper method to set volume
  const setVolume = useCallback(
    (volume) => {
      templateModalDispatch({
        type: templateModalActions.SET_VOLUME,
        value: volume,
      });
    },
    [templateModalDispatch]
  );

  // helper method to set lid temperature
  const setLidTemperature = useCallback(
    (lidTemperature) => {
      templateModalDispatch({
        type: templateModalActions.SET_LID_TEMPERATURE,
        value: lidTemperature,
      });
    },
    [templateModalDispatch]
  );

  // fetch old template values for comparing with the edited values
  const prevTemplate = useMemo(
    () => templates.find((ele) => ele.get("id") === templateID),
    [templates, templateID]
  );

  // Auto fill the template name and description with old values
  const autofillNameDescription = useCallback(() => {
    setTemplateName(prevTemplate.get("name"));
    setTemplateDescription(prevTemplate.get("description"));
    setVolume(prevTemplate.get("volume"));
    setLidTemperature(prevTemplate.get("lidTemperature"));
  }, [prevTemplate, setTemplateName, setTemplateDescription]);

  // check if changes are persent by comparing previous values
  const checkForChanges = () => {
    if (
      templateDescription !== prevTemplate.get("description") ||
      templateName !== prevTemplate.get("name") ||
      volume !== prevTemplate.get("volume") ||
      lidTemperature !== prevTemplate.get("lidTemperature")
    ) {
      return true;
    }
    return false;
  };

  // Validate create template form
  const validateTemplateForm = () => {
    if (
      templateDescription !== "" &&
      templateName !== "" &&
      volume &&
      lidTemperature
    ) {
      return true;
    }
    return false;
  };

  const createTemplate = (template) => {
    // creating template though api
    dispatch(createTemplateAction(template, token));
  };

  const updateTemplate = (template) => {
    // updating template though api
    dispatch(updateTemplateAction(templateID, template, token));
  };

  // helper method to reset the local modal state
  const resetModalState = () =>
    templateModalDispatch({
      type: templateModalActions.RESET_TEMPLATE_MODAL,
    });

  // reset form input values
  const resetFormValues = () => {
    setTemplateDescription("");
    setTemplateName("");
    setVolume(null);
    setLidTemperature(null);
  };

  // save/edit click handler
  const addClickHandler = () => {
    if (validateTemplateForm()) {
      if (isTemplateEdited) {
        // check if the templateDescriptions and templateName values
        // are changed from previous values before update api call
        if (checkForChanges()) {
          // Update template rest api call.
          updateTemplate({
            description: templateDescription,
            name: templateName,
            volume: volume,
            lid_temp: lidTemperature,
          });
        }
      } else {
        // Create new template rest api call.
        createTemplate({
          description: templateDescription,
          name: templateName,
          volume: volume,
          lid_temp: lidTemperature,
        });
      }
      // reset modal state to initial values
      resetModalState();
    }
    // TODO show error notification
  };

  // useEffect Section
  useEffect(() => {
    if (isTemplateEdited === true) {
      autofillNameDescription();
      toggleTemplateModal();
    }
  }, [isTemplateEdited, autofillNameDescription, toggleTemplateModal]);

  // after template is updated we reset updateTemplateReducer and
  // re-fetch templates
  useEffect(() => {
    if (isTemplateUpdated === true && errorUpdatingTemplate === false) {
      dispatch(updateTemplateReset());
      dispatch(fetchTemplates({ token }));
    }
  }, [isTemplateUpdated, dispatch]);

  return (
    <TemplateModal
      isCreateTemplateModalVisible={isCreateTemplateModalVisible}
      toggleTemplateModal={toggleTemplateModal}
      templateDescription={templateDescription}
      setTemplateDescription={setTemplateDescription}
      templateName={templateName}
      setTemplateName={setTemplateName}
      volume={volume}
      setVolume={setVolume}
      lidTemperature={lidTemperature}
      setLidTemperature={setLidTemperature}
      addClickHandler={addClickHandler}
      isFormValid={validateTemplateForm()}
      resetFormValues={resetFormValues}
      isTemplateEdited={isTemplateEdited}
      resetModalState={resetModalState}
    />
  );
};

TemplateModalContainer.propTypes = {
  templateModalState: PropTypes.object.isRequired,
  toggleTemplateModal: PropTypes.func.isRequired,
  templateModalDispatch: PropTypes.func.isRequired,
  templateID: PropTypes.string,
};

export default TemplateModalContainer;
