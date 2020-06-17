import React from "react";
import styled from "styled-components";
import PropTypes from "prop-types";
import { ButtonIcon, Text } from "shared-components";

const StyledTemplate = styled.div`
	position: relative;
	width: ${(props) => (props.isActive ? "315px" : "220px")};
	height: ${(props) => (props.isActive ? "60px" : "44px")};
	background: ${(props) => (props.isActive ? "#aedbd5" : "#ffffff")};
	display: flex;
	justify-content: center;
	align-items: center;
	font-size: 18px;
	line-height: 22px;
	color: ${(props) => (props.isActive ? "#ffffff" : "#707070")};
	font-weight: ${(props) => (props.isActive ? "bold" : "")};
	box-shadow: 0px 3px 16px #0000000b;
	border: 1px solid #e5e5e5;
	border-radius: 8px;
	padding: ${(props) => (props.isActive ? "8px 74px" : "8px 16px")};
	overflow: hidden;
`;

export const Template = (props) => {
	return (
		<StyledTemplate {...props}>
			{props.isEditable && props.isActive ? 		
				<ButtonIcon
					position="absolute"
					placement="left"
					left={16}
					size={28}
					name="pencil"
					isShadow
					className="text-reset" 
				/> : "" }
			<Text tag="span" className="text-truncate">
				{props.title}
			</Text>
			{props.isDeletable && props.isActive ? 	
				<ButtonIcon
					position="absolute"
					placement="right"
					right={16}
					isShadow
					size={28}
					name="trash"
					className="text-reset" 
				/> : "" }
		</StyledTemplate>
	);
};

Template.propTypes = {
	isActive: PropTypes.bool,
	isEditable: PropTypes.bool,
	isDeletable: PropTypes.bool,
};

Template.defaultProps = {
	isActive: false,
	isEditable: false,
	isDeletable: false,
};