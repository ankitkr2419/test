import styled from "styled-components";

export const StyledProcessListing = styled.div`
    .landing-content {
        padding: 1.25rem 4.5rem 0.875rem 4.5rem;
        &::after {
            height: 9.125rem;
        }
        .recipe-listing-cards {
            height: 30.75rem;
        }
        .selected-button {
            color: #fff !important;
            background: #aedbd5;
        }
    }
`;

export const InnerBox = styled.div`
    height: 50px;
`;
