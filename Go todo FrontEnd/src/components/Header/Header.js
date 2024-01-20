import { StyledHeader, StyledHeaderContent, StyledHeaderOptions, StyledHeaderTitle } from "../../style/layout/StyledHeader";
import { useNavigate } from "react-router-dom";


function Header() {
const navigate = useNavigate();

  const navigatePage = (route) => {
    navigate(route);
  }

  return (
    <StyledHeader>
      <StyledHeaderContent>
        <StyledHeaderTitle onClick={()=>{navigatePage('/')}}>Go todo</StyledHeaderTitle>
        <StyledHeaderOptions>
          <li onClick={()=>{navigatePage("login")}}>login</li>
          <li onClick={()=>{navigate("register")}}>register</li>
        </StyledHeaderOptions>
      </StyledHeaderContent>
    </StyledHeader>
  );
}

export default Header;
