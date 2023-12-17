import React from "react";
import { useParams } from "react-router-dom";
import { Company } from "../../types";
import { LoadingIcon } from "../LoadingIcon/LoadingIcon";

export const TickTable = ({ companies }: { companies: Company[] }) => {
  let { companyName } = useParams();

  const company =
    companies &&
    companies.find(
      (company) => companyName === company.companyName.toLowerCase()
    );
  return (
    <>
      {!companies && <LoadingIcon />}
      {company && (
        <>
          Name: {company?.companyName} <br />
          Address: {company?.companyAddress}
        </>
      )}
    </>
  );
};
