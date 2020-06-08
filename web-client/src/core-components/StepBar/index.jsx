import styled from "styled-components";
import { Nav, NavItem } from "reactstrap";
import { NavLink } from "react-router-dom";
import bgStep from "../../assets/images/steps.svg";

export const Step = styled(Nav).attrs({className: "nav-steps"})`
  margin: 0 0 40px;
  padding: 0 56px;
`;

export const StepItem = styled(NavItem)`
  width: 185px;
  height: 40px;
  background: transparent url(${bgStep}) no-repeat;
  background-size: auto 84px;
  background-position: -22px -19px;
  margin: 0;
  padding: 0 23px;
  text-align: center;
  opacity: ${(props) => (props.isDisable ? "0.53" : "")};
  pointer-events: ${(props) => (props.isDisable ? "none" : "")};

  + .nav-item {
    margin: 0 0 0 -16px;
  }
`;

export const StepLink = styled(NavLink).attrs({className: "nav-link"})`
  font-size: 16px;
  line-height: 24px;
  color: #707070;
  padding: 8px 0;
  border-radius: 4px;

  &:hover,
  &:focus {
    color: #707070;
  }
`;