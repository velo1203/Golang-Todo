import React, { useState } from "react";
import { StyledDefaultPage } from "../../style/layout/StyledDefaultPage";
import { StyledForm, StyledFormContainer } from "../../style/layout/StyledForm";
import { StyledInput } from "../../style/common/StyledInput";
import { StyledButton } from "../../style/common/StyledButton";
import { AuthLogin } from "../../service/Auth/Auth";
import useAuthStore from "../../store/userStore"; 
import { Navigate, useNavigate } from "react-router-dom";


function LoginPage() {
    const {setUserToken} = useAuthStore();
    const [usename, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const navigate = useNavigate()

    const fetchLogiin = async () => {
        try {
            const User = await AuthLogin(usename,password);
            setUserToken(User.token);
            navigate('/')

        } catch (error) {
                alert('아이디 또는 비밀번호가 일치하지 않습니다.')
        }
    };    

    return (
        <StyledDefaultPage>
            <StyledFormContainer>
                <h1>Login</h1>
                <StyledForm>
                    <StyledInput type="text" placeholder="유저 이름" onChange={(e)=>{setUsername(e.target.value)}}/>
                    <StyledInput type="password" placeholder="비밀번호" onChange={(e)=>{setPassword(e.target.value)}}/>
                    <StyledButton onClick={()=>{fetchLogiin()}}>Login</StyledButton>
                </StyledForm>
            </StyledFormContainer>
        </StyledDefaultPage>
    );
    }

export default LoginPage;