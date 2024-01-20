import styled from "styled-components";


export const StyledButton = styled.button`

    width: 130px;
    height: 55px;
    border-radius: 5px;
    border: 2px solid var(--MainGray);
    background-color: rgba(217, 217, 217, 0.00);
    color: var(--White);
    font-size: 15px;

    cursor: pointer;
    transition: all 0.3s ease-in-out;

    &:hover{
        background-color: var(--MainGray);
        color: var(--White);
    }

    &:focus{
        outline: none;
    }
`