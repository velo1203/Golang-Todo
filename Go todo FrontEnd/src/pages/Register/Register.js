import React, { useState } from "react";
import { StyledDefaultPage } from "../../style/layout/StyledDefaultPage";
import { StyledForm, StyledFormContainer } from "../../style/layout/StyledForm";
import { StyledInput } from "../../style/common/StyledInput";
import { StyledButton } from "../../style/common/StyledButton";
import { AuthRegister } from "../../service/Auth/Auth";
import useAuthStore from "../../store/userStore";
import { useNavigate } from "react-router-dom";

function RegisterPage() {
    const {setUserToken} = useAuthStore();
    const [usename, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [passwordCheck, setPasswordCheck] = useState('')
    const navigate = useNavigate()

    const fetchRegister = async () => {
        try {
            if (password !== passwordCheck){
                alert('비밀번호가 일치하지 않습니다.')
                return
            }
            const User = await AuthRegister(usename,password);
            setUserToken(User.token);
            navigate('/')
        } catch (error) {
                alert('유저가 이미 있거나 오류가 발생했습니다.')
        }
    }

    return (
        <StyledDefaultPage>
            <StyledFormContainer>
                <h1>Register</h1>
                <StyledForm>
                    <StyledInput type="text" placeholder="유저 이름" onChange={(e)=>{setUsername(e.target.value)}}/>
                    <StyledInput type="password" placeholder="비밀번호" onChange={(e)=>{setPassword(e.target.value)}}/>
                    <StyledInput type="password" placeholder="비밀번호 확인 " onChange={(e)=>{setPasswordCheck(e.target.value)}}/>
                    <StyledButton onClick={()=>{fetchRegister()}}>Register</StyledButton>
                </StyledForm>
            </StyledFormContainer>
        </StyledDefaultPage>
    );
    }

export default RegisterPage;