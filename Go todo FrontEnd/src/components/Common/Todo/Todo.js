import React from "react";
import { StyledTodo } from "../../../style/common/StyledTodo";
import Dropdown from "../Dropdown/Dropdown";

function Todo() {
  return (
    <StyledTodo>
      <h1>디미고 숙제하기</h1>
      <p>2021-09-29</p>
      <Dropdown/>
    </StyledTodo>
  );
}

export default Todo;