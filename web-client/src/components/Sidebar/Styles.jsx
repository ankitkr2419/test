import styled from "styled-components";

export const StyledSidebarBody = styled.div`
    background: #fafafa 0% 0% no-repeat padding-box;
    padding: 48px;
    text-align: center;
    border: 1px solid #e5e5e5;
    border-left: 0 none;
    border-radius: 0px 24px 24px 0px;
`;

export const StyledSidebarContent = styled.div`
    display: flex;
    flex-direction: column;
    position: relative;
    background: #aedbd5 0% 0% no-repeat padding-box;
    padding: 20px 20px 20px 0;
    box-shadow: 0 2px 6px #00000029;
    overflow: hidden;
`;

export const Shadow = styled.div`
    position: absolute;
    top: 0;
    left: 0;
    height: 172px;
    width: 60px;
    border-radius: 0 40px 40px 0;
    box-shadow: 0 2px 6px #00000029;
    z-index: 1;

    &::after {
        content: "";
        position: absolute;
        top: 50%;
        transform: translate(0%, -50%);
        width: 20px;
        height: 184px;
        background-color: #aedbd5;
        left: -8px;
        z-index: 2;
    }
`;

export const StyledSidebarHandle = styled.button`
    position: absolute;
    top: 50%;
    right: -48px;
    transform: translate(0%, -50%);
    background-color: #aedbd5;
    border: 0 none;
    height: 172px;
    width: 60px;
    border-radius: 0 40px 40px 0;
    color: #ffffff;
    padding: 4px;
    z-index: 1;

    i {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
`;
