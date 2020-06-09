import React from "react";
import styled from "styled-components";
import ButtonIcon from "shared-components/ButtonIcon";
import Icon from "shared-components/Icon";
import Text from "shared-components/Text";

const BtnEdit = (props) => {
	return (
		<ButtonIcon position="absolute" placement="left" left="16">
			<Icon name="pencil" />
		</ButtonIcon>
	);
};

const BtnDelete = (props) => {
	return (
		<ButtonIcon position="absolute" placement="right" right="16">
			<Icon name="trash" />
		</ButtonIcon>
	);
};

export const Template = (props) => {
	return (
		<StyledTemplate {...props}>
			{props.isEditable && props.isActive ? <BtnEdit /> : ""}
			<Text tag="span" className="text-truncate">{props.title}</Text>
			{props.isDeletable && props.isActive ? <BtnDelete /> : ""}
		</StyledTemplate>
	);
};

export const TemplateList = styled.ul.attrs({className: "list-template"})`
	display: flex;
	flex-wrap: wrap;
	padding: 0 60px;
	list-style: none;
	margin: 0 0 56px;
	overflow-x: hidden;
	overflow-y: auto
`;

export const TemplateListItem = styled.li.attrs({className: "list-template-item"})`
	display: flex;
	justify-content: center;
	align-items: center;
	flex: 0 0 50%;
	max-width: 50%;
	min-height: 60px;
	text-align: center;
`;

const StyledTemplate = styled.div.attrs({ className: "template" })`
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