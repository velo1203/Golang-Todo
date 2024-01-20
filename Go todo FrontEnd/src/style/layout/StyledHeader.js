import styled from "styled-components";

export const StyledHeader = styled.header`
background-color: var(--DarkBlue);
height: 60px;
width: 100%;
color: var(--White);
`

export const StyledHeaderContent = styled.div`
display: flex;
justify-content: space-between;
align-items: center;
height: 100%;
max-width: 1100px;
margin: 0 auto;

`

export const StyledHeaderTitle = styled.h1`
font-size: 25px;
cursor: pointer;

`

export const StyledHeaderOptions = styled.ul`
display: flex;
font-size: 15px;
list-style: none;
li{
    margin-left: 20px;
    cursor: pointer;
}

`