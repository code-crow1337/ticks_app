import React from "react";
import { generateURLFriendlyName } from "../../utils/utils";
import { Link, useParams, useNavigate } from "react-router-dom";

export const SearchField = ({ options }: { options: string[] }) => {
  // ? Would be nice to have auto complete, or suggestions in a dropdown
  let { companyName } = useParams();
  const navigate = useNavigate();
  const [searchInput, setSearchInput] = React.useState(
    companyName ? companyName : ""
  );
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
      navigate(`/company/${generateURLFriendlyName(companyName)}`);
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
        {options
          .filter((option) => {
            const regex = new RegExp(`${searchInput}`, "gi");
            return !!option.match(regex);
          })
          .map((option) => (
            <Link
              to={`/company/${generateURLFriendlyName(option)}`}
              key={option}
              onClick={(e) => handleOnChange(e, option)}
            >
              {option}
            </Link>
          ))}
      </div>
    </div>
  );
};
