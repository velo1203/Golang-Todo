import styled from "styled-components";

export const StyledSearch = styled.div`
height: 35px;
display: flex;
align-items: center;
input{
    box-sizing: border-box;
    width: 300px;
    height: 100%;
    border-radius: 5px;
    border: none;
    background-color: var(--MainGray);
    padding: 15px;
    color: white;
    transition: all 0.3s ease-in-out;
    outline: none;
    &:hover{
        background-color: var(--MainGrayHover);
    
    }
}
button{
    width: 50px;
    height: 100%;
    border-radius: 5px;
    background-color: var(--MainGreen);
    color: var(--White);
    border: none;
    margin-left: 10px;
    font-size: 25px;
    cursor: pointer;
}



`