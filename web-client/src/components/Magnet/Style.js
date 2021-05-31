import styled from "styled-components";

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const MagnetBox = styled.div`
  .process-magnet {
    &::after {
      background: url("/images/magnet-bg.svg") no-repeat;
    }
    .magnet-large-btn {
      width: 15.188rem;
      height: 20.875rem;
      margin-bottom: 2.125rem;
      border: 1px solid #e6e6e6;
      border-radius: 1.5rem;
      padding: 61px 79px;
      &.btn-bg {
        background-color: #f4f4f4;
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
    .animated-attach-icon {
      position: relative;
      .icon-upward-magnet {
        position: absolute;
        top: -75px;
        left: auto;
        right: auto;
        animation: attach-upward-magnet 1.75s 3s linear infinite;
        opacity: 1;
        @keyframes attach-upward-magnet {
          0% {
            top: -75px;
          }
          70% {
            top: -35px;
            opacity: 1;
            animation-delay: 10s;
          }
          100% {
            top: -35px;
            opacity: 0;
            animation-delay: 10s;
          }
        }
      }
      .icon-downward-magnet {
        position: absolute;
        top: 0;
        left: auto;
        right: auto;
        animation: attach-downward-magnet 1.75s 3s linear infinite;
        @keyframes attach-downward-magnet {
          0% {
            top: 35px;
          }
          70% {
            top: 0;
            opacity: 1;
            animation-delay: 10s;
          }
          100% {
            top: 0;
            opacity: 0;
            animation-delay: 10s;
          }
        }
      }
    }

    .animated-detach-icon {
      position: relative;
      .icon-upward-magnet {
        position: absolute;
        top: -35px;
        left: auto;
        right: auto;
        animation: detach-upward-magnet 1.75s 2s linear infinite;
        opacity: 1;
        color: #717171;
        @keyframes detach-upward-magnet {
          0% {
            top: -35px;
          }
          70% {
            top: -75px;
            opacity: 1;
          }
          100% {
            top: -75px;
            opacity: 0;
          }
        }
      }
      .icon-downward-magnet {
        position: absolute;
        top: 0;
        left: auto;
        right: auto;
        color: #717171;
        animation: detach-downward-magnet 1.75s 2s linear infinite;
        @keyframes detach-downward-magnet {
          0% {
            top: 0px;
          }
          70% {
            top: 35px;
            opacity: 1;
          }
          100% {
            top: 35px;
            opacity: 0;
          }
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
