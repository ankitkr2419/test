import styled from "styled-components";

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;

export const TipDiscardBox = styled.div`
  .-tip-discard {
    &::after {
      background: url("/images/tip-discard-bg.svg") no-repeat;
    }
    .magnet-large-btn {
      width: 15.188rem;
      height: 20.875rem;
      margin-bottom: 2.125rem;
      border: 1px solid #e6e6e6;
      border-radius: 1.5rem;
      padding: 61px 55px;
      &.btn-bg {
        background-color: #f4f4f4;
      }
      &.selected {
        border: 2px solid #b2dad1;
      }
    }
    .process-image {
      position: absolute;
      bottom: 53px;
      right: 37px;
    }
    .process-box {
      width: 90.55%;
    }
    .pickup-point {
      width: 4px;
      height: 4px;
      background-color: #f38220;
      border: 2px solid #f38220;
      position: absolute;
      top: 12%;
      left: 48%;
      animation: at-pickup 2s 2s linear infinite;
      @keyframes at-pickup {
        0% {
          top: 12%;
        }
        35% {
          transform: scale(4);
        }
        50% {
          transform: scale(4);
        }
        85% {
          transform: scale(1);
        }
        100% {
          transform: scale(1);
        }
      }
    }

    .discard-point {
      width: 12px;
      height: 12px;
      background-color: #f38220;
      border: 2px solid #f38220;
      position: absolute;
      top: 5%;
      left: 50%;
      transform: translate(-50%, -5%);
      z-index: 1;
      animation: at-discard 2.5s ease-out 3s infinite;
      @keyframes at-discard {
        0% {
          opacity: 0;
        }
        5% {
          opacity: 1;
        }
        50% {
          opacity: 1;
        }
        60% {
          top: 65%;
          opacity: 0;
        }
        100% {
          top: 65%;
          opacity: 0;
        }
      }
    }
    .long-down-arrow {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      margin-top: -10%;
      height: 0;
      opacity: 0;
      animation: long-arrow 2.5s ease-out 3s infinite;
      @keyframes long-arrow {
        0% {
          opacity: 0;
        }
        10% {
          opacity: 0;
          height: 0;
        }
        50% {
          opacity: 1;
          height: 41px;
        }
        60% {
          opacity: 0;
        }
        100% {
          opacity: 0;
        }
      }
    }
  }
`;

export const TopContent = styled.div`
  margin-bottom: 0.75rem;
  .frame-icon {
    > button {
      > i {
        font-size: 26px;
      }
    }
  }
`;
