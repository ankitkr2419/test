import React from 'react';
import PropTypes from 'prop-types'
import styled from 'styled-components';
import { Text } from 'shared-components';
import TemplatePopover from 'components/Plate/Popover';

const StyledSubHeader = styled.div`
	display: flex;
	align-items: center;
	height: 40px;
	padding: 8px 16px 8px 88px;
	color: #707070;

	h6 {
		font-size: 14px;
		line-height: 1.25;
	}
`;

const SubHeader = ({ experimentTemplate }) => (
	<StyledSubHeader className="plate-subheader">
		<Text Tag="h6" className="mb-0">
			{experimentTemplate.templateId}
		</Text>
		<Text Tag="h6" className="mb-0 mx-5">
			{experimentTemplate.templateName}
		</Text>
		{/* <Text Tag="h6" className="mb-0 ml-5">
				22/06/2020
		</Text>
		<Text Tag="h6" className="mb-0 ml-3">
				23:50 PM to 01:21 AM
		</Text>
		<Text Tag="h6" className="mb-0 ml-5">
				No. of wells - 5
		</Text> */}
		<TemplatePopover className="ml-auto" />
	</StyledSubHeader>
);

SubHeader.propTypes = {
	experimentTemplate: PropTypes.shape({
		templateId: PropTypes.string,
		templateName: PropTypes.string,
	}).isRequired,
};

export default SubHeader;
