import styled from "styled-components";

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const HeatingBox = styled.div`
  .process-heating {
    &::after {
      background: url("/images/heating-bg.svg") no-repeat;
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

export const HeatingProcessBox = styled.div`
  .process-box {
    width: 90.55%;
  }
`;
