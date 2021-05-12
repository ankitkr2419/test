import styled from "styled-components";
export const StyledButtonIcon = styled.button`
    width: ${(props) => props.size + 8 + `px`};
    height: ${(props) => props.size + 8 + `px`};
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: transparent;
    color: #707070;
    font-weight: normal;
    padding: 4px;
    border: 1px solid white;
    border-radius: 50%;
    box-shadow: ${(props) => (props.isShadow ? "0 2px 6px #00000020" : "")};
    position: ${(props) => props.position};
    top: ${(props) =>
        props.position === "absolute" || props.position === "fixed"
            ? `${props.top}px`
            : ""};
    right: ${(props) =>
        (props.position === "absolute" || props.position === "fixed") &&
        props.placement === "right"
            ? `${props.right}px`
            : ""};
    left: ${(props) =>
        (props.position === "absolute" || props.position === "fixed") &&
        props.placement === "left"
            ? `${props.left}px`
            : ""};
`;
