import styled from "styled-components";

export const RecipeFlowSlider = styled.div`
    .slides {
        .slides-inner-box {
            width: 43.5rem;
            height: 100%;
            // height:25rem;
            margin: 0 auto;
            overflow: hidden !important;
            img {
                border-radius: 1.5rem !important;
                box-shadow: 0px 3px 6px rgba(0, 0, 0, 0.16) !important;
            }
        }
    }
    .slick-dots {
        bottom: -2.5rem !important;
        li button:before {
            font-size: 0.75rem;
        }
        li.slick-active button:before {
            transform: scale(1.5);
            color: #9ad0c8;
        }
    }
    .center {
        .slick-list {
            padding-top: 1.875rem !important;
            padding-bottom: 1.875rem !important;
        }
        .slick-center .slides-inner-box {
            transform: scale(1.12);
            overflow: hidden;
            border-radius: 1.5rem;
        }
        .slides {
            -webkit-transition: all 0.3s ease-out;
            transition: all 0.3s ease-out;
        }
        .slick-next,
        .slick-prev {
            background-color: #9ad0c8;
            z-index: 1;
            width: 3rem;
            height: 6.063rem;
            box-shadow: 0px 3px 6px rgba(0, 0, 0, 0.16);
        }
        .slick-next {
            right: -1px;
            border-radius: 3.125rem 0 0 3.125rem;
            &::before {
                background: url("/images/next-arrow.svg") no-repeat;
                background-position: top center;
                position: relative;
                left: 5px;
                background-size: contain;
                color: transparent;
            }
        }
        .slick-prev {
            left: -1px;
            border-radius: 0 3.125rem 3.125rem 0;
            &::before {
                background: url("/images/prev-arrow.svg") no-repeat;
                background-position: top center;
                position: relative;
                right: 5px;
                background-size: contain;
                color: transparent;
            }
        }
    }
`;
export const NextButton = styled.div`
    position: absolute;
    bottom: 1.5rem;
    right: 6rem;
`;
