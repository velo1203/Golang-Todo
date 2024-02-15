import React from "react";
import { StyledSearch } from "../../../style/common/StyledSearch";

function Search({onButtonClick,onChange}) {
  return (
    <StyledSearch>
    <input type="text" placeholder="Search Todo" onChange={onChange}/>
    <button onClick={onButtonClick}>+</button>
    </StyledSearch>
  );
}

export default Search