import styled from "styled-components";
import { Card } from "reactstrap";

const StyledCard = styled(Card)`
	height: 528px;
	background: #fafafa 0% 0% no-repeat padding-box;
	box-shadow: 0px 3px 16px #0000000f;
	border: 1px solid #e5e5e5;
  border-radius: 36px;
  
  .card-body {
    padding: 24px 48px 24px 72px;
    overflow-x: hidden;
    overflow-y: auto;
  }
`;

export default StyledCard;