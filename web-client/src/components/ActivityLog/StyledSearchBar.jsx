import styled from 'styled-components';
import { FormGroup } from 'core-components';
export const StyledSearchBar = styled(FormGroup)`
	position: relative;
	margin-bottom: -1px;

	.form-control {
		padding: 4px 51px 4px 16px;
		border-color: #cbcbcb;
		border-radius: 0;

		&:focus {
			box-shadow: none;
		}
	}

	i {
		position: absolute;
		top: 3px;
		right: 10px;
		color: #666666;
	}
`;