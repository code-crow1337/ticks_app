import React from "react";
import { render, screen } from "@testing-library/react";
import { App } from "./App";

test("renders learn react link", () => {
  render(<App />);
  const linkElement = screen.getByText(/learn react/i); //TODO add some test after I have some text to test for
  expect(linkElement).toBeInTheDocument();
});
