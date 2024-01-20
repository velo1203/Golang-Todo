import React, { useState, useEffect,useRef} from 'react';
import { StyledIcon } from '../../../style/common/StyledIcon';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEllipsisVertical } from '@fortawesome/free-solid-svg-icons';
import { StyledDropdown, StyledDropdownItem, StyledDropdownMenu } from '../../../style/common/StyledDropdown';
import { deleteTodo, putTodo } from '../../../service/Todo/Todo';

function Dropdown({todo,onTodoChange}) {
    const [isOpen, setIsOpen] = useState(false);
    const dropdownRef = useRef(null);

    const toggleDropdown = () => setIsOpen(!isOpen);

    const handleClickOutside = (event) => {
        if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
            setIsOpen(false);
        }
    };

    useEffect(() => {
        document.addEventListener('mousedown', handleClickOutside);
        return () => {
            document.removeEventListener('mousedown', handleClickOutside);
        };
    }, []);

    const handleTodoDelete = async () => {
        try {
            const DeleteTodo = await deleteTodo(todo.ID);
            toggleDropdown()
            onTodoChange()
        } catch (error) {
                alert('투두가 삭제되지 않았습니다.')
        }
    };  
    const handleTodoStatusChange = async () => {
        try {
            const ChnagedTodo = await putTodo(todo.ID,!todo.Completed);
            toggleDropdown()
            onTodoChange()
        } catch (error) {
                alert('투두가 삭제되지 않았습니다.')
        }
    };  
    return (
        <StyledDropdown ref={dropdownRef}>
            <StyledIcon onClick={toggleDropdown}>
            <FontAwesomeIcon icon={faEllipsisVertical}/>
            </StyledIcon>

            {isOpen && (
                <StyledDropdownMenu>
                    <StyledDropdownItem onClick={()=>{handleTodoDelete()}}>Delete</StyledDropdownItem>
                    <StyledDropdownItem onClick={()=>{handleTodoStatusChange()}}>Change Status</StyledDropdownItem>
                </StyledDropdownMenu>
            )}
        </StyledDropdown>
    );
}

export default Dropdown;
