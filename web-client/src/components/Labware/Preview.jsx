import React from "react";
import { Text } from "shared-components";
import { LABWARE_NAME } from "appConstants";

const getSubHead = (key, recipeData) => {
  const nestedKeys = Object.keys(recipeData[key].processDetails);
  const LEN = nestedKeys.length;
  const previewInfoSubHead = [];

  nestedKeys.forEach((nestedKey) => {
    recipeData[key].processDetails[nestedKey].id &&
      previewInfoSubHead.push(
        <Text>
          {LEN > 1 && (
            <Text Tag="span" className="font-weight-bold">
              {LABWARE_NAME[nestedKey]} :{" "}
            </Text>
          )}
          <Text Tag="span" className={LEN === 1 ? "font-weight-bold" : ""}>
            {recipeData[key].processDetails[nestedKey].label}{" "}
          </Text>
        </Text>
      );
  });
  return previewInfoSubHead;
};

export const Preview = (props) => {
  const { recipeData } = props;

  Object.keys(recipeData).map((key, index) => {
    if (recipeData[key].isTicked) {
      return (
        <li key={index} className="d-flex justify-content-between">
          <Text className="d-flex w-25 font-weight-bold">
            {LABWARE_NAME[key]} :{" "}
          </Text>
          <div className="w-75">
            <div className="ml-2 setting-value">
              <Text>{getSubHead(key, recipeData)}</Text>
            </div>
          </div>
        </li>
      );
    }
  });
};
