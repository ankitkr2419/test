import React from "react";
import PropTypes from "prop-types";
import { Link } from "react-router-dom";
import styled, { css } from "styled-components";

const StyledLink = styled(Link)`
	display: ${(props) => (props.isIcon ? "flex" : "inline-block")};
	width: ${(props) => (props.isIcon ? "40px" : "202px")};
	height: 40px;
	font-size: 16px;
	line-height: 19px;
	font-weight: bold;
	text-align: center;
	vertical-align: middle;
	padding: ${(props) => (props.isIcon ? "4px" : "10px 20px")};
	border-radius: ${(props) => (props.isIcon ? "50%" : "27px")};
	border-width: ${(props) => (props.isIcon ? "" : "1px")};
	border-style: ${(props) => (props.isIcon ? "" : "solid")};
	box-shadow: ${(props) => (props.isIcon ? "" : "0 2px 6px #00000029")};
	user-select: none;
	transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out,
		border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;

	&:focus {
		outline: none;
		box-shadow: ${(props) => (props.isIcon ? "" : "0 2px 6px #00000029")};
	}

	&:hover {
		text-decoration: none;
	}

	${(props) =>
		props.isIcon &&
		css`
			align-items: center;
			justify-content: center;

			i {
				color: #707070;
			}
		`};
`;

const CustomLink = (props) => {
	return (
		<StyledLink {...props}>
			{props.children}
		</StyledLink>
	);
};

CustomLink.propTypes = {
	isIcon: PropTypes.bool,
};

CustomLink.defaultProps = {
	isIcon: false,
};

export default CustomLink;
