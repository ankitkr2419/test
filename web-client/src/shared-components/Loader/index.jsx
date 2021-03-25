import React from "react";
import { Spinner } from "reactstrap";

const Loader = () => {
  return (
    <div className="position-absolute py-5 d-flex flex-column align-items-center justify-content-center overlay">
      <Spinner color="primary" style={{ width: "5rem", height: "5rem" }} />
      <div>Loading....</div>
    </div>
  );
};

export default Loader;
