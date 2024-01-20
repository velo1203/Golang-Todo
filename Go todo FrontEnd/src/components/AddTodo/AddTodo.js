import React, { useState } from "react";
import { StyledInput } from "../../style/common/StyledInput";
import { StyledButton } from "../../style/common/StyledButton";
import { StyledAddTodo } from "../../style/layout/StyledAddTodo";
import { postTodo } from "../../service/Todo/Todo";



function AddTodo({onTodoAdded}) {
  const [todo, setTodo] = useState('');

  const TodoPost = async () => {
    try {
        const Todo = await postTodo(todo);
        console.log(Todo)
        onTodoAdded();
    } catch (error) {
        console.error('POST Request Error:', error);
            alert('투두가 생성되지 않았습니다.')
    }
}

  return (
    <StyledAddTodo>
      <StyledInput placeholder="할 일을 입력하세요" onChange={(e)=>{setTodo(e.target.value)}}/>
      <StyledButton onClick={()=>{TodoPost()}}>추가</StyledButton>
    </StyledAddTodo>
  );
}
export default AddTodo;