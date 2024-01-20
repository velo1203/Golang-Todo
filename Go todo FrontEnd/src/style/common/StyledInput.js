import styled from "styled-components";

export const StyledInput = styled.input`
    box-sizing: border-box;
    width: 400px;
    height: 50px;
    border-radius: 5px;
    border: none;
    background-color: var(--MainGray);
    padding: 15px;
    color: white;
    transition: all 0.3s ease-in-out;
    &:hover{
        background-color: var(--MainGrayHover);
    
    }
`