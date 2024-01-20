import styled from "styled-components";


export const StyledTodo = styled.div`
margin-top: 15px;
box-sizing: border-box;
width: 100%;
height: 80px;
background-color: var(--MainGray);
display: flex;
align-items: center;
padding: 25px;
color: white;
font-size: 10px;
border-radius: 5px;
border-left: 10px solid ${props => props.completed ? 'var(--MainGreen)' : 'var(--MainRed)'};
justify-content: space-between;
p{
    font-size: 15px;
}
`

