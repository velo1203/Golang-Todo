import React from "react";
import { StyledDefaultPage } from "../../style/layout/StyledDefaultPage";
import { StyledForm, StyledFormContainer } from "../../style/layout/StyledForm";
import { StyledInput } from "../../style/common/StyledInput";
import { StyledButton } from "../../style/common/StyledButton";


function RegisterPage() {
    return (
        <StyledDefaultPage>
            <StyledFormContainer>
                <h1>Register</h1>
                <StyledForm>
                    <StyledInput type="text" placeholder="아이디"/>
                    <StyledInput type="password" placeholder="비밀번호"/>
                    <StyledInput type="password" placeholder="비밀번호 확인 "/>
                    <StyledButton>Register</StyledButton>
                </StyledForm>
            </StyledFormContainer>
        </StyledDefaultPage>
    );
    }

export default RegisterPage;