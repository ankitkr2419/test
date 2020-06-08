import styled from "styled-components";
import { Nav, NavItem } from "reactstrap";
import { NavLink } from "react-router-dom";

export const Step = styled(Nav)`
  margin: 0 0 40px;
  padding: 0 56px;
`;

export const StepItem = styled(NavItem)`
  width: 200px;
  height: 40px;
  margin: 0 8px 0 0;
  text-align: center;
`;

export const StepLink = styled(NavLink).attrs({
  className: "nav-link",
})`
  background-color: white;
  font-size: 16px;
  line-height: 24px;
  color: #707070;
  border-radius: 4px;
  opacity: ${(props) => (props.isDisabled ? "0.53" : "")};
  pointer-events: ${(props) => (props.isDisabled ? "none" : "")};

  &:hover,
  &:focus {
    color: #707070;
  }
`;