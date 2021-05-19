import styled from "styled-components";

export const ProcessCardBox = styled.div`
    width: 100%;
    min-height: 3.125rem;
    transition: min-height 0.2s ease-in-out;
    border-radius: 0.5rem;
    padding: 0.313rem 1rem;
    box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.06);
    position: absolute;
    bottom: 10px;
    // &:hover {
    //     min-height: 107px;
    //     // height:auto;
    //     .hidden-box {
    //         display: flex !important;
    //     }
    // }
    &.selected-box {
        min-height: 107px;
        border: 1px solid #b2dad1;
        .hidden-box {
          display: block !important;
        }
      }
    .process-title {
        > button {
            width: 40px !important;
            height: 40px !important;
            border: 1px solid #696969 !important;
        }
        > label {
            margin-left: 0.75rem;
        }
    }
    .more-action {
        > button {
            width: 30px !important;
            height: 30px !important;
            border: 1px solid #696969 !important;
        }
    }
    .hidden-box {
        display: none;
    }
    .drop-badge {
        background-color: #b2dad1;
        border-radius: 0.875rem;
        width: 9rem;
        height: 1.688rem;
        line-height: 1.688rem;
    }
`;