import styled from "styled-components";

export const ButtonBarBox = styled.div`
  width: 93.82%;
  height: 3.25rem;
  background-color: #fff;
  z-index: 2;
  border-radius: 2rem 0 0 2rem;
  padding: 0.5rem 4.938rem 0.5rem 2.375rem;
  box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.06);
  position: absolute;
  right: 0;
  bottom: 0rem;
  top: 29rem;

  &.btn-bar-adjust-tipDiscard {
    top: unset !important;
    bottom: 2rem !important;
  }

  &.btn-bar-adjust-tipPickup {
    top: unset !important;
    bottom: 2rem !important;
  }

  &.btn-bar-adjust-aspireDispense {
    top: unset !important;
    bottom: 3rem !important;
  }

  &.btn-bar-adjust-shaking {
    top: unset !important;
    bottom: 2rem !important;
  }

  &.btn-bar-adjust-heating {
    top: unset !important;
    bottom: 2rem !important;
  }

  &.btn-bar-adjust-magnet {
    top: unset !important;
    bottom: 2rem !important;
  }

  &.btn-bar-adjust-delay {
    top: unset !important;
    bottom: 2rem !important;
  }

  &.btn-bar-adjust-labware {
    top: unset !important;
    bottom: 1rem !important;
  }

  > button {
    width: 160px;
    &:hover,
    &:focus {
      color: #ffffff !important;
      > i {
        color: #ffffff !important;
      }
    }
    > i {
      color: #f38220;
    }
  }
`;
export const PrevBtn = styled.div`
  min-width: inherit;
  border: 0;
  box-shadow: none;
  color: #f38220;
`;
