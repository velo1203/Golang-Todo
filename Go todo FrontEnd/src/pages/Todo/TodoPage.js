import {useState} from 'react';
import Search from '../../components/Common/Search/Search';
import Todo from '../../components/Common/Todo/Todo';
import {StyledDefaultPage} from '../../style/layout/StyledDefaultPage';
import {StyledTodoList, StyledTodoSearch} from '../../style/layout/StyledTodo';
import PopupWrapper from '../../components/Common/PopupWrapper/PopupWrapper';
import AddTodo from '../../components/AddTodo/AddTodo';

function TodoPage() {
    const [addTodo, setAddTodo] = useState(false);

    return (
        <> 
        {addTodo && <PopupWrapper onOutsideClick={()=>{setAddTodo(false)}}>
          <AddTodo/>
          </PopupWrapper>}
      < StyledDefaultPage >
          <StyledTodoSearch>
              <Search onButtonClick={() => {setAddTodo(!addTodo)} }/>
          </StyledTodoSearch>

          <StyledTodoList>
              <Todo/>
          </StyledTodoList>
      </StyledDefaultPage>
  </ >
    );
}

export default TodoPage;
