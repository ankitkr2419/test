import React from 'react';
import PropTypes from "prop-types";
import { ButtonIcon, Text } from "shared-components";
import { Popover, PopoverBody } from 'core-components';
import styled from 'styled-components';

const StyledTemplatePopover = styled.div`
	.btn-toggle {
		width: 22px;
		height: 22px;
		border: 0 none;
    padding: 4px 0 0 0;
	}
`;

const TemplatePopover = props => {
  return (
		<StyledTemplatePopover {...props}>
			<ButtonIcon
				name="angle-down"
				size={40}
				id="PopoverTemplate"
				className="btn-toggle"
			/>
			<Popover
				trigger="legacy"
				target="PopoverTemplate"
				placement="bottom"
				popperClassName="popover-template"
			>
				<PopoverBody className="d-flex flex-column justify-content-center flex-100 scroll-y">
					<Text className="font-weight-bold">(Template name)</Text>
					<Text>Cycle Count - x</Text>
					<Text>Current temperature - x</Text>
					<Text className="mb-0">Lid temperature - x</Text>
				</PopoverBody>
			</Popover>
		</StyledTemplatePopover>
	);
};

TemplatePopover.propTypes = {
  className: PropTypes.string,
};


TemplatePopover.defaultProps = {
  className: "",
};

export default TemplatePopover;