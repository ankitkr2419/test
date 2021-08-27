import React from "react";

const ColoredCircle = (props) => {
  const { isOnline } = props;
  const color = isOnline ? "#05b714" : "red";
  const styles = { backgroundColor: color };

  return <div className="colored-circle" style={styles} />;
};

export default ColoredCircle;
