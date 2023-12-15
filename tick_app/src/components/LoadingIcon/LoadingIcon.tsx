import React from "react";
import "./LoadingIcon.scss";
export const LoadingIcon = ({ size = "md" }: { size?: "sm" | "md" }) => {
  return (
    <div className="lds-ring">
      <div className={`size-${size}`}></div>
      <div className={`size-${size}`}></div>
      <div className={`size-${size}`}></div>
      <div className={`size-${size}`}></div>
    </div>
  );
};
