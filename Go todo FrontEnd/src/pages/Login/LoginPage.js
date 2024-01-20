import React from "react";
import { StyledDefaultPage } from "../../style/layout/StyledDefaultPage";
import { StyledForm, StyledFormContainer } from "../../style/layout/StyledForm";
import { StyledInput } from "../../style/common/StyledInput";
import { StyledButton } from "../../style/common/StyledButton";


function LoginPage() {
    return (
        <StyledDefaultPage>
            <StyledFormContainer>
                <h1>Login</h1>
                <StyledForm>
                    <StyledInput type="text" placeholder="아이디"/>
                    <StyledInput type="password" placeholder="비밀번호"/>
                    <StyledButton>Login</StyledButton>
                </StyledForm>
            </StyledFormContainer>
        </StyledDefaultPage>
    );
    }

export default LoginPage;