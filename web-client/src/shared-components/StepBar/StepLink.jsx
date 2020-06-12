import React from "react";
import PropTypes from "prop-types";
import classNames from 'classnames';
import { NavLink } from "react-router-dom";
import "./StepLink.scss";

export const StepLink = (props) => {
	
	const stepLinkClass = classNames("step-link", props.className);

	if(props.tag === "a") {
		return <NavLink {...props} className={stepLinkClass} />;
	}
	return <props.tag {...props} className={stepLinkClass} />;
};

StepLink.propTypes = {
	tag: PropTypes.oneOf(["a", "button"]),
};

StepLink.defaultProps = {
	tag: "button",
};
