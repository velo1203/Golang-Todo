import React from "react";
import { StyledSearch } from "../../../style/common/StyledSearch";

function Search({onButtonClick}) {
  return (
    <StyledSearch>
    <input type="text" placeholder="Search Todo"/>
    <button onClick={onButtonClick}>+</button>
    </StyledSearch>
  );
}

export default Search