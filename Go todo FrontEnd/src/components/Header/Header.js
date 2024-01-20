import { StyledHeader, StyledHeaderContent, StyledHeaderOptions, StyledHeaderTitle } from "../../style/layout/StyledHeader";
import { useNavigate } from "react-router-dom";
import useAuthStore from "../../store/userStore";

function Header() {
const navigate = useNavigate();
const {isLoggedIn,logout} = useAuthStore();

  const navigatePage = (route) => {
    navigate(route);
  }
  const handleLogout = () => {
    logout(); // 로그아웃 함수 호출
    navigate('/'); // 홈페이지나 로그인 페이지로 리디렉션
  }
  return (
    <StyledHeader>
      <StyledHeaderContent>
        <StyledHeaderTitle onClick={()=>{navigatePage('/')}}>Go todo</StyledHeaderTitle>
        <StyledHeaderOptions>
          {!isLoggedIn && <li onClick={()=>{navigatePage("login")}}>login</li>}
          {isLoggedIn && <li onClick={()=>{handleLogout()}}>logout</li>}
          {!isLoggedIn && <li onClick={()=>{navigate("register")}}>register</li>}
        </StyledHeaderOptions>
      </StyledHeaderContent>
    </StyledHeader>
  );
}

export default Header;
