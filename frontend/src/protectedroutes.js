import { Navigate } from "react-router-dom";
import { accountService } from "./_services/account.service";

export const ProtectedRoute = ({ children }) => {
 
  if(accountService.isLogged()) {

  
    return children
  }else {
    return <Navigate to="/login" />;


  }
}
   
   
