import styled from "styled-components";

//For Tray Discard Modal
// Need to toggle this class for gray scale effect
export const TrayDiscardSection = styled.div`
    .status-box {
        &.gray-scale-box {
            filter: grayscale(1);
        }
    }
`;
export const DiscardTrayBox = styled.div`
    .btn-discard-tray {
        width: 10rem;
    }
`;