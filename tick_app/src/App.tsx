import React from "react";
import "./App.scss";
import { useFetchApi } from "./hooks/useFetchAPI";
import { SearchField } from "./components/SearchField/SearchField";
import { LoadingIcon } from "./components/LoadingIcon/LoadingIcon";

export const App = () => {
  const { data, error, isLoading } = useFetchApi("/company");
  console.log({ data });
  return (
    <main className="App">
      <header className="App__hero-container">
        <h1>Tick Summary Tool</h1>
      </header>
      <main className="App__content-container">
        <SearchField />
        <div>Tick Summary Tool</div>
        {
          /*     <LoadingIcon size="sm" /> */
          //Use for latter
        }
      </main>
    </main>
  );
};
