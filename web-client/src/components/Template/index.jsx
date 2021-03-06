import React, { useState, useEffect } from "react";
import { Button } from "core-components";
import PropTypes from "prop-types";
import {
  StyledUl,
  StyledLi,
  CustomButton,
  Center,
  Text,
  ImageIcon,
} from "shared-components";
import imgNoTemplate from "assets/images/no-template-available.svg";
import MlModal from "shared-components/MlModal";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";

const TemplateComponent = (props) => {
  const {
    templates,
    deleteTemplate,
    updateSelectedWizard,
    updateTemplateID,
    isLoginTypeOperator,
    isLoginTypeAdmin,
    createExperiment,
    toggleTemplateModal,
    isTemplatesLoading,
    isCreateTemplateModalVisible,
  } = props;

  // Local state to store template name
  const [selectedTemplateId, setSelectedTemplateId] = useState(null);
  const [showDeleteTemplateModal, setShowDeleteTemplateModal] = useState(false);

  const deleteClickHandler = (e) => {
    e.stopPropagation();
    toggleDeleteTemplateModal();
  };

  const toggleDeleteTemplateModal = () => {
    //if we are hiding delete modal, then clear selected template data to hide edit/delete buttons
    if (showDeleteTemplateModal) {
      setSelectedTemplateId(null);
    }
    //update state to toggle delete confirmation modal
    setShowDeleteTemplateModal(!showDeleteTemplateModal);
  };

  const onConfirmedDeleteTemplate = () => {
    // Delete api call
    deleteTemplate(selectedTemplateId);
    toggleDeleteTemplateModal();
  };

  const editClickHandler = () => {
    // Updates template id to templateState maintain over templateLayout
    updateTemplateID(selectedTemplateId);
    // navigate to next wizard
    updateSelectedWizard("target");
  };

  const onTemplateButtonClickHandler = ({ id: templateId, description }) => {
    // if its admin save template id and show edit & delete options on button
    // if its operator save template id an navigate to target wizard
    // set selected template id to local state for maintaining active state of button

    // toggle button to show edit/delete buttons
    let toggleTemplateButton = null;
    if (selectedTemplateId === null) {
      toggleTemplateButton = templateId;
    }
    setSelectedTemplateId(toggleTemplateButton);

    if (isLoginTypeOperator === true) {
      // make api call to save experiments
      createExperiment({
        template_id: templateId,
        description,
      });
      // Updates template id to templateState maintain over templateLayout
      updateTemplateID(templateId);
      // navigate to next wizard
      // updateSelectedWizard('target');
    }
  };

  useEffect(() => {
    // If admin login make creat modal open if no data is available
    // isTemplatesLoading will tell us weather api calling is finish or not
    // templates.size = 0  will tell us there is no records present
    // isCreateTemplateModalVisible is check as we have to make modal visible only once
    if (
      isLoginTypeAdmin === true &&
      isTemplatesLoading === false &&
      templates.size === 0 &&
      isCreateTemplateModalVisible === false
    ) {
      toggleTemplateModal();
    }
    // isCreateStageModalVisible skipped in dependency because its causing issue with modal state
    // eslint-disable-next-line
  }, [isTemplatesLoading, templates]);

  return (
    <div className="d-flex flex-100 flex-column p-4 mt-3">
      {templates.size === 0 && (
        <Center className="no-template-wrap">
          <ImageIcon
            src={imgNoTemplate}
            alt="No templates available"
            className="img-no-template"
          />
          <Text className="d-flex justify-content-center" Tag="p">
            No templates available
          </Text>
        </Center>
      )}

      {/**Delete template confirmation modal */}
      {showDeleteTemplateModal && (
        <MlModal
          isOpen={showDeleteTemplateModal}
          textHead={""}
          textBody={MODAL_MESSAGE.deleteTemplateConfirmation}
          handleSuccessBtn={onConfirmedDeleteTemplate}
          handleCrossBtn={toggleDeleteTemplateModal}
          successBtn={MODAL_BTN.yes}
          failureBtn={MODAL_BTN.no}
        />
      )}

      {/* templates size check before iteration */}
      {templates.size !== 0 && (
        <StyledUl>
          {templates.map((template) => (
            <StyledLi key={template.get("id")}>
              <CustomButton
                title={template.get("name")}
                isActive={template.get("id") === selectedTemplateId}
                isEditable={
                  isLoginTypeAdmin === true && isLoginTypeOperator === false
                }
                onButtonClickHandler={() => {
                  onTemplateButtonClickHandler(template.toJS());
                }}
                onEditClickHandler={editClickHandler}
                isDeletable={
                  isLoginTypeAdmin === true && isLoginTypeOperator === false
                }
                onDeleteClickHandler={deleteClickHandler}
              />
            </StyledLi>
          ))}
        </StyledUl>
      )}
      <Center className="mb-5">
        {isLoginTypeAdmin === true && (
          <Button color="primary" onClick={toggleTemplateModal}>
            Create New
          </Button>
        )}
      </Center>
    </div>
  );
};

TemplateComponent.propTypes = {
  templates: PropTypes.shape({}).isRequired,
  deleteTemplate: PropTypes.func.isRequired,
  updateSelectedWizard: PropTypes.func.isRequired,
  updateTemplateID: PropTypes.func.isRequired,
  isLoginTypeOperator: PropTypes.bool.isRequired,
  isLoginTypeAdmin: PropTypes.bool.isRequired,
};

export default React.memo(TemplateComponent);
