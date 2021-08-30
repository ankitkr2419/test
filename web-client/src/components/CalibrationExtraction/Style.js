import styled from "styled-components";

export const ShakingProcessBox = styled.div`
  .process-box {
    width: 90.55%;
    .custom-checkbox {
      > label {
        font-size: 0.875rem;
        line-height: 1rem;
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
    .temperature-box {
      margin-left: 2.375rem;
    }
    .rpm-input {
      width: 6rem;
      height: 2.25rem;
    }
    .border-left-line {
      border-left: 1px solid rgba(112, 112, 112, 0.16);
    }
    .time-box {
      label {
        font-size: 0.75rem;
        line-height: 0.875rem;
      }
    }
  }
`;

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const ShakingBox = styled.div`
  .process-shaking {
    &::after {
      background: url("/images/shaking-bg.svg") no-repeat;
    }
  }
`;

export const TopContent = styled.div`
  margin-bottom: 0.75rem;
  .frame-icon {
    > button {
      > i {
        font-size: 40px;
      }
    }
  }
`;

export const HeatingBox = styled.div`
  .process-heating {
    &::after {
      background: url("/images/heating-bg.svg") no-repeat;
    }
  }
`;

export const HeatingProcessBox = styled.div`
  .process-box {
    width: 90.55%;
  }
`;
