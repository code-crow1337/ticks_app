import React from "react";

export const SearchField = () => {
  // ? Would be nice to have auto complete, or suggestions in a dropdown
  const [searchInput, setSearchInput] = React.useState("");
  const [searchFieldIsFocused, setSearchFieldIsFocused] = React.useState(false);

  const handleOnChange = (
    e:
      | React.ChangeEvent<HTMLInputElement>
      | React.MouseEvent<HTMLAnchorElement>,
    companyName?: string | undefined
  ) => {
    e.preventDefault();
    if ("value" in e.target) {
      setSearchInput(e.target.value);
    }
    if (companyName) {
      setSearchFieldIsFocused(false);
      setSearchInput(companyName);
    }
  };

  return (
    <div className="search-field">
      <div className="search-field__search">
        <i className="search-field__search__icon fa-solid fa-magnifying-glass" />
        <input
          type="text"
          placeholder="Company Name..."
          value={searchInput}
          onChange={handleOnChange}
          onFocus={() => setSearchFieldIsFocused(true)}
        />
        <button onClick={() => console.log({ submitvalue: searchInput })}>
          Submit
        </button>
      </div>
      <div
        className={`search-field__dropdown--${
          searchFieldIsFocused ? "show" : "hide"
        }`}
      >
        {mockCompanies
          .filter((company) => {
            const regex = new RegExp(`${searchInput}`, "gi");
            return !!company.name.match(regex);
          })
          .map((company) => (
            <a
              href={company.query} //Would be nice to use the url without actually loading the page. But then I need react-router
              key={company.query}
              onClick={(e) => handleOnChange(e, company.name)}
            >
              {company.name}
            </a>
          ))}
      </div>
    </div>
  );
};
const mockCompanies = [
  { name: "Kanelbullens konditori", query: "kanelbullens-konditori" },
  { name: "Lucifiers lusse cafe", query: "lucifiers-lusse-cafe" },
  { name: "Linnes konditori", query: "linnes-konditori" },
];
