import styled from "styled-components";

export const StyledFormContainer = styled.div`
box-sizing: border-box;
padding: 50px;
border: 2px solid var(--MainGray);
height: 450px;
width: 600px;
margin: 0 auto;
margin-top: 100px;
border-radius: 5px;
text-align: center;

`;

export const StyledForm = styled.form`
margin-top: 50px;
display: flex;
flex-direction: column;
justify-content: center;
align-items: center;
gap: 10px;
button{
    margin-top: 40px;
}

`;