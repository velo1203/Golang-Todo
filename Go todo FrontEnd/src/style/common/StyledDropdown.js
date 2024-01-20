import styled from "styled-components";

export const StyledDropdown = styled.div `
    position: relative;
    display: inline-block;
    `;

export const StyledDropdownMenu = styled.div `
    right: 20px;
    position: absolute;
    min-width: 160px;
    font-size: 15px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
`;

export const StyledDropdownItem = styled.div `
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    cursor: pointer;
    background-color: var(--White);
    &:hover{
        opacity: 0.9;
    }
`;