import React from 'react';
import { Modal } from 'reactstrap';
import "./Modal.scss";

const CustomModal = props => {
  return (
    <Modal {...props}>
      {props.children}
    </Modal>
  );
};

CustomModal.propTypes = {};

export default CustomModal;