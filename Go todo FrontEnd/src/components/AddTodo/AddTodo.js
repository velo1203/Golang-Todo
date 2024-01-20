import React from "react";
import { StyledInput } from "../../style/common/StyledInput";
import { StyledButton } from "../../style/common/StyledButton";
import { StyledAddTodo } from "../../style/layout/StyledAddTodo";


function AddTodo() {
  return (
    <StyledAddTodo>
      <StyledInput placeholder="할 일을 입력하세요" />
      <StyledButton>추가</StyledButton>
    </StyledAddTodo>
  );
}
export default AddTodo;