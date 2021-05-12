import React from "react";
import PropTypes from "prop-types";
import { ButtonIcon, Text } from "shared-components";
import { StyledTemplate } from './StyledTemplate'

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