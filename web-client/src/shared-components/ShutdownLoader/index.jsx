import React from "react";
import { Spinner } from "reactstrap";

const ShutdownLoader = () => {
  return (
    <div
      className="position-absolute py-5 d-flex flex-column align-items-center justify-content-center overlay"
      style={{ position: "fixed !important" }}
    >
      <Spinner color="primary" style={{ width: "5rem", height: "5rem" }} />
      <div style={{ color: "white", fontWeight: "bold" }}>Shutdown....</div>
    </div>
  );
};

export default ShutdownLoader;

