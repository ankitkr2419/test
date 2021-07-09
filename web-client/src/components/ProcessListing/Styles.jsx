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

        //TopContent reduce space for ProcessListing page
        .process-listing-changes {
            margin-bottom: 1.75rem;
        }

        //TODO: remove if not needed
        //for 2nd solution
        .box {
            max-height: 400px;
            flex: 1;
            margin: 0 -8px;
        }
        .box > div {
            width: 50%;
            padding: 0 8px;
            margin-bottom: 16px;
        }

        //TODO: remove if not needed
        //3rd solution
        .card-columns > .card {
            // height: 50px;
            // padding: 1rem;
            // text-align:center;
            background-color: transparent;
            // border:none;
        }
    }
`;

export const InnerBox = styled.div`
    height: 50px;
`;
