import styled from "styled-components";

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const TipPickupBox = styled.div`
  .-tip-pickup {
    &::after {
      background: url("/images/tip-pickup-bg.svg") no-repeat;
    }
    .process-image {
      position: absolute;
      bottom: 3.313rem;
      right: 2.313rem;
    }
    .tip-pickup-inner-box {
      padding: 5.438rem 5.188rem;
      .tip-pickup-box {
        background-color: #f4f4f4;
        border: 1px solid #e6e6e6;
        border-radius: 1.5rem;
        width: 30.313rem;
        height: 10.313rem;
      }
      .label-name {
        width: 7.688rem;
      }
      .input-field {
        width: 14.125rem;
        height: 2.25rem;
      }
    }
    .tip-pickup-note {
      font-size: 0.813rem;
      line-height: 0.938rem;
    }
  }
`;

export const TopContent = styled.div`
  margin-bottom: 0.75rem;
  .frame-icon {
    > button {
      > i {
        font-size: 2.625rem;
      }
    }
  }
`;