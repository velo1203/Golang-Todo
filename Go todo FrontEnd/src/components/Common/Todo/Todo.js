import React from "react";
import { StyledTodo } from "../../../style/common/StyledTodo";
import Dropdown from "../Dropdown/Dropdown";

function Todo({item,onTodoChange}) {
  return (
    <StyledTodo completed={item.Completed}>
      <h1>{item.Title}</h1>
      <p>ID: {item.ID}</p>
      <Dropdown todo={item} onTodoChange={onTodoChange}/>
    </StyledTodo>
  );
}

export default Todo;