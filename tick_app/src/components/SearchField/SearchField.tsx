import React from "react";

export const SearchField = () => {
  // ? Would be nice to have auto complete, or suggestions in a dropdown
  const [searchInput, setSearchInput] = React.useState("");

  const handleOnChange = (
    e: React.FormEvent<HTMLInputElement> & { target: { value: string } }
  ) => {
    e.preventDefault();
    setSearchInput(e.target.value);
  };
  return (
    <div className="search-field">
      <i className="search-field__icon fa-solid fa-magnifying-glass" />
      <input type="text" placeholder="Company Name" onChange={handleOnChange} />
      <button onClick={() => console.log({ submitvalue: searchInput })}>
        Submit
      </button>
    </div>
  );
};
