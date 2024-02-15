import {useEffect, useState} from 'react';
import Search from '../../components/Common/Search/Search';
import Todo from '../../components/Common/Todo/Todo';
import {StyledDefaultPage} from '../../style/layout/StyledDefaultPage';
import {StyledTodoList, StyledTodoSearch} from '../../style/layout/StyledTodo';
import PopupWrapper from '../../components/Common/PopupWrapper/PopupWrapper';
import AddTodo from '../../components/AddTodo/AddTodo';
import {getTodo} from '../../service/Todo/Todo';
import {useNavigate} from 'react-router-dom';
import useAuthStore from '../../store/userStore';

function TodoPage() {
    const [addTodo, setAddTodo] = useState(false);
    const {isLoggedIn} = useAuthStore();
    const [todo, setTodo] = useState([]);
    const [reloadTodos, setReloadTodos] = useState(false);
    const [search, setSearch] = useState('');

    const navigate = useNavigate()
    useEffect(() => {
        const fetchTodos = async () => {
            try {
                const Todos = await getTodo();
                setTodo(Todos.todos)
                console.log(Todos.todos)
            } catch (error) {
                if (error.response.status === 401) {
                    navigate('/login')
                }
            }
        };

        fetchTodos();
    }, [isLoggedIn, reloadTodos]);
    const handleChangeTodo = () => {
        setReloadTodos(prev => !prev); // 투두 추가 후 목록 다시 불러오기 트리거
    }

    const filteredTodos = todo.filter((item) => {
        const filter = item.Title.toLowerCase().includes(search.toLowerCase()); 
        console.log(filter)
        return filter
    });
    return (
        <> {
            addTodo && <PopupWrapper
                    onOutsideClick={() => {
                        setAddTodo(false)
                    }}>
                    <AddTodo onTodoAdded={handleChangeTodo} handleClose={()=>{setAddTodo(false)}}/>
                </PopupWrapper>
        } < StyledDefaultPage > <StyledTodoSearch>
            <Search
            onChange={(e)=>{setSearch(e.target.value)}}
                onButtonClick={() => {
                    setAddTodo(!addTodo)
                }}/>
        </StyledTodoSearch>
        <StyledTodoList>
            {
                filteredTodos.map((item) => {
                    return <Todo item={item} id={item.ID} onTodoChange={handleChangeTodo}/>
                })
            }
        </StyledTodoList>
    </StyledDefaultPage>
</ >
    );
}

export default TodoPage;
