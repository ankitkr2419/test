import React from "react";
import PropTypes from "prop-types";
import { Icon } from "shared-components";
import { Button } from "core-components";
import { ButtonBarBox, PrevBtn } from './Styles';

const ButtonBar = (props) => {
  const { handleTemp } = props;
  return (
    <ButtonBarBox className="d-flex justify-content-start align-items-center mt-5">
      <PrevBtn>
        <Icon name="angle-left" size={30} />
      </PrevBtn>
      <Button color="outline-primary" className="ml-auto text-dark" size="md">
        {" "}
        <Icon size={20} name="plus-2" className="mb-0 p-0" />
        Add Process
      </Button>
      <Button onClick={handleTemp} color="primary" className="ml-4" size="md">
        {" "}
        Finish
      </Button>
    </ButtonBarBox>
  );
};

ButtonBar.propTypes = {
  isUserLoggedIn: PropTypes.bool,
};

ButtonBar.defaultProps = {
  isUserLoggedIn: false,
};

export default ButtonBar;
