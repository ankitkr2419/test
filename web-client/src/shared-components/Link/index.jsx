import React from "react";
import { Link } from "react-router-dom";
import styled from "styled-components";

const StyledLink = styled(Link)`
	display: inline-block;
	width: 202px;
	height: 40px;
	font-size: 16px;
	line-height: 19px;
	font-weight: bold;
	text-align: center;
	vertical-align: middle;
	padding: 10px 20px;
	border-radius: 27px;
	border-width: 1px;
	border-style: solid;
	box-shadow: 0 2px 6px #00000029;
	user-select: none;
	transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out,
		border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;

	&:focus {
		outline: none;
		box-shadow: 0 2px 6px #00000029;
	}

	&:hover {
		text-decoration: none;
	}
`;

const BLink = (props) => {
	return (
		<StyledLink to={props.to} className={props.className}>
			{props.children}
		</StyledLink>
	);
};

BLink.propTypes = {};

export default BLink;
