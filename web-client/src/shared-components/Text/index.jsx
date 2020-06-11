import React from "react";
import PropTypes from "prop-types";

const Text = props => {
  return <props.tag className={props.className}>{props.children}</props.tag>;
};

export default Text;

Text.propTypes = {
	tag: PropTypes.string,
};

Text.defaultProps = {
  tag: "p"
};