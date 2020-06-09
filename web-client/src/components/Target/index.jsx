import styled from "styled-components";

export const TargetList = styled.ul.attrs({ className: "list-target" })`
	padding: 0 32px;
	list-style: none;
	overflow-x: hidden;
  overflow-y: auto;
  margin: 0
`;

export const TargetListItem = styled.li.attrs({className: "list-target-item"})`
  display: flex;
  align-items: center;

  + .list-target-item {
    margin-top: 16px;
  }

  &:nth-child(2) {
    margin-top: 0;
  }

  .custom-checkbox {
    margin: 0 8px 0 0;
  }

  .ml-select {
    padding: 0 8px;
  }

  > p {
    font-size: 14px;
    line-height: 16px;
    color: #999999;
    min-width: 24px;
    margin: 0 0 4px;

    &:first-child {
      margin-right: 8px;
    }

    &:nth-child(2),
    &:nth-child(3) {
      padding: 0 24px;
    }
  }

  .-target {
    flex: 1;
  }

  .-threshold {
    flex: 0 0 40%;
    max-width: 40%;
    padding-right: 0;
  }
`;