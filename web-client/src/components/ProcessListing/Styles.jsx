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
        
        //TODO: remove if not needed 
        //for 2nd solution
        .box {//added with dhanraj
            max-height: 400px;
            flex: 1;
            margin: 0 -15px;
        }
        .box > div {//added with dhanraj
            margin: 0 7.5px;
            margin-bottom: 15px;
        }

        //TODO: remove if not needed
        //3rd solution
        .card-columns>.card {
            height: 50px;
            padding: 1rem;
            text-align:center;
        }
    }
`;

export const InnerBox = styled.div`
    height: 50px;
`;
