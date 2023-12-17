import React from "react";
import "./App.scss";
import { useFetchApi } from "./hooks/useFetchAPI";
import { SearchField } from "./components/SearchField/SearchField";
import { LoadingIcon } from "./components/LoadingIcon/LoadingIcon";
import { TickTable } from "./components/TickTable/TickTable";
import { Company } from "./types";

export const App = () => {
  const { data, error, isLoading } = useFetchApi("/company");
  const companyNames =
    data && data.map((company: Company) => company.companyName);
  return (
    <main className="App">
      <header className="App__hero-container">
        <h1>Tick Summary Tool</h1>
      </header>
      <main className="App__content-container">
        <div>Tick Summary Tool</div>
        {isLoading ? (
          <LoadingIcon size="sm" />
        ) : (
          <>
            <SearchField options={companyNames} />
          </>
        )}
        <TickTable companies={data} />
      </main>
    </main>
  );
};
