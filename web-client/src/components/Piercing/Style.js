import styled from "styled-components";

export const ExtractionBox = styled.div`
  .process-box {
    width: 90.55%;
    .custom-checkbox {
      > label {
        font-size: 1.063rem;
        line-height: 1.25rem;
        color: #666666;
      }
      .custom-control-label::after {
        left: -2.25rem;
      }
    }
    label {
      font-size: 1rem;
      line-height: 1.125rem;
      color: #666666;
    }
    .well-box {
      // counter-reset: section;
      .well {
        &:active {
          background-color: #ffffff;
        }
      }
      .well-no {
        .coordinate-item {
          color: #999999;
          font-size: 18px;
          line-height: 21px;
        }
      }
      .well{
      	position:relative;
      // 	&::before{
      // 	counter-increment: section;
      // 	content: counter(section);
      // 	position: absolute;
      // 	top: -28px;
      // 	left: 0;
      // 	right: 0;
      // 	display: flex;
      // 	justify-content: center;
      // 	align-items: center;
      // 	color:#999999;
      // 	}
      // }
      .selected {
        border: 3px solid #abd5ce;
      }
    }
  }
`;

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const PiercingBox = styled.div`
  .process-piercing {
    &::after {
      background: url("/images/piercing-bg.svg") no-repeat;
    }
  }
`;

export const TopContent = styled.div`
  margin-bottom: 0.75rem;
  .frame-icon {
    > button {
      > i {
        font-size: 53px;
      }
    }
  }
`;
