import React, { useReducer } from "react";
import PropTypes from "prop-types";
import { useHistory } from "react-router";

import { Button, Card, CardBody } from "core-components";
import { Text, Icon } from "shared-components";
import CreateUserModal from "components/modals/ManageUserModals/CreateUserModal";
import { roleOptions } from "./helper";
import DeleteUserModal from "components/modals/ManageUserModals/DeleteUserModal";
import UpdateUserModal from "components/modals/ManageUserModals/UpdateUserModal";

const ManageUsersComponent = (props) => {
  const { handleCreateUser, handleDeleteUser, handleUpdateUser } = props;
  const history = useHistory();
  const [showCreateUserModal, toggleCreateUserModal] = useReducer(
    (showCreateUserModal) => !showCreateUserModal,
    false
  );

  const [showUpdateUserModal, toggleUpdateUserModal] = useReducer(
    (showUpdateUserModal) => !showUpdateUserModal,
    false
  );

  const [showDeleteUserModal, toggleDeleteUserModal] = useReducer(
    (showDeleteUserModal) => !showDeleteUserModal,
    false
  );

  const onHandleCreateUser = (userData) => {
    handleCreateUser(userData);
    toggleCreateUserModal();
  };

  const onHandleDeleteUser = (username) => {
    handleDeleteUser(username);
    toggleDeleteUserModal();
  };

  const onHandleUpdateUser = (userData) => {
    handleUpdateUser(userData);
    toggleUpdateUserModal();
  };

  const handleBack = () => {
    history.goBack();
  };

  return (
    <div className="manageUsers-content px-5">
      <div className="d-flex">
        <div
          className="d-flex"
          style={{ cursor: "pointer" }}
          onClick={handleBack}
        >
          <Icon name="angle-left" size={32} className="text-white" />
          <Text
            Tag="h5"
            size={24}
            className="text-white font-weight-bold ml-1 mt-1 mb-0"
          >
            Back
          </Text>
        </div>
      </div>

      {showCreateUserModal && (
        <CreateUserModal
          isOpen={showCreateUserModal}
          toggleModal={toggleCreateUserModal}
          roleOptions={roleOptions}
          handleCreateUser={onHandleCreateUser}
        />
      )}
      {showUpdateUserModal && (
        <UpdateUserModal
          isOpen={showUpdateUserModal}
          toggleModal={toggleUpdateUserModal}
          roleOptions={roleOptions}
          handleUpdateUser={onHandleUpdateUser}
        />
      )}
      {showDeleteUserModal && (
        <DeleteUserModal
          isOpen={showDeleteUserModal}
          toggleModal={toggleDeleteUserModal}
          handleDeleteUser={onHandleDeleteUser}
        />
      )}

      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          <Text
            Tag="h4"
            size={24}
            className="text-center text-primary mt-2 mb-4"
          >
            User Manager
          </Text>
          <div className="d-flex justify-content-center mb-4">
            <Button
              onClick={toggleCreateUserModal}
              color="secondary"
              className={"mr-3"}
              size="md"
            >
              Create User
            </Button>
            <Button
              onClick={toggleUpdateUserModal}
              color="secondary"
              className={"mr-3"}
              size="md"
            >
              Update User
            </Button>
            <Button
              onClick={toggleDeleteUserModal}
              color="secondary"
              className={"mr-3"}
              size="md"
            >
              Delete User
            </Button>
          </div>
        </CardBody>
      </Card>
    </div>
  );
};

ManageUsersComponent.propTypes = {
  handleCreateUser: PropTypes.func.isRequired,
  handleDeleteUser: PropTypes.func.isRequired,
  handleUpdateUser: PropTypes.func.isRequired,
};

export default React.memo(ManageUsersComponent);
