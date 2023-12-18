import React from "react";
import { useParams } from "react-router-dom";
import { Company } from "../../types";
import { LoadingIcon } from "../LoadingIcon/LoadingIcon";
import { useFetchApi } from "../../hooks/useFetchAPI";

export const TickTable = () => {
  let { companyName } = useParams();
  const { data, error, isLoading } = useFetchApi("/company");
  let formatedData: any[] = [];

  if (data && data.length > 0) {
    const combinedTotals = formatData(data);
    for (const companySummary in combinedTotals) {
      formatedData.push(combinedTotals[companySummary]);
    }
  }
  let singleCompany;
  if (companyName && formatedData.length > 0) {
    singleCompany = formatedData.find(
      (company) => company.companyName.toLowerCase() === companyName
    );
  }
  console.log({ companyName, singleCompany });
  return (
    <>
      {companyName ? (
        <div style={{ display: "flex", flexDirection: "column" }}>
          <div>
            <b>Name:</b>
            {singleCompany.companyName}
          </div>
          <div>
            <b>Total bank-auth:</b> {singleCompany.totalBankAuth}
          </div>
          <div>
            <b>Total bank-sign: </b>
            {singleCompany.totalBankSign}
          </div>
          <div>
            <b>Total SPAR Lookups: </b>
            {singleCompany.totalSpar}
          </div>
          <div>
            <b>Total SMS: </b>
            {singleCompany.totalSms}
          </div>
          <br />
        </div>
      ) : (
        formatedData.map((company) => (
          <div style={{ display: "flex", flexDirection: "column" }}>
            <div>
              <b>Name:</b>
              {company.companyName}
            </div>
            <div>
              <b>Total bank-auth:</b> {company.totalBankAuth}
            </div>
            <div>
              <b>Total bank-sign: </b>
              {company.totalBankSign}
            </div>
            <div>
              <b>Total SPAR Lookups: </b>
              {company.totalSpar}
            </div>
            <div>
              <b>Total SMS: </b>
              {company.totalSms}
            </div>
            <br />
          </div>
        ))
      )}
    </>
  );
};

const formatData = (data: any[]) => {
  const test = data.reduce((next, company) => {
    const { companyId } = company;
    return {
      ...next,
      [companyId]: {
        ...company,
        totalBankAuth: 0,
        totalBankSign: 0,
        totalSpar: 0,
        totalSms: 0,
      },
    };
  }, {});

  data.forEach((company) => {
    const { typeTick, companyId } = company;

    if (typeTick === "bankid-auth") {
      test[companyId].totalBankAuth = test[companyId].totalBankAuth + 1;
    }
    if (typeTick === "bankid-sign") {
      test[companyId].totalBankSign = test[companyId].totalBankSign + 1;
    }
    if (typeTick === "spar") {
      test[companyId].totalSpar = test[companyId].totalSpar + 1;
    }
    if (typeTick === "sms") {
      test[companyId].totalSms = test[companyId].totalSms + 1;
    }
  });
  return test;
};
