import React from "react";
import styled from "styled-components";
import { Text } from "shared-components";
import TemplatePopover from "./TemplatePopover";

const StyledPlateSubHeader = styled.div`
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

const PlateSubHeader = props => {
	return (
		<StyledPlateSubHeader className="plate-subheader">
			<Text Tag="h6" className="mb-0">
				ID002
			</Text>
			<Text Tag="h6" className="mb-0 mx-5">
				Template Name
			</Text>
			<Text Tag="h6" className="mb-0 ml-5">
				22/06/2020
			</Text>
			<Text Tag="h6" className="mb-0 ml-3">
				23:50 PM to 01:21 AM
			</Text>
			<Text Tag="h6" className="mb-0 ml-5">
				No. of wells - 5
			</Text>
			<TemplatePopover className="ml-auto" />
		</StyledPlateSubHeader>
	);
}

export default PlateSubHeader;