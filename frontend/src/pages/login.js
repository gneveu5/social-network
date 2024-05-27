import { useContext, useState } from "react";
import { useNavigate , Link} from "react-router-dom";

//api 
import { accountService } from "../_services/account.service";

//context
import UserContext from "../context/UserContext";
//style
import "../css/login.css"



export const LoginPage = () => {
  //utils
  let navigate = useNavigate()
  let userToken = "";

  //global  state
  const { setUserData } = useContext(UserContext);
  

  // local state
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errors, setErrors] = useState({});

  //login operation
  const handleLogin = async (e) => {
    e.preventDefault();

    // call api
    accountService.login({email , password})
     .then((r) => {
        userToken = r.data.token
        accountService.saveToken(userToken)
        //JE DOIS FINIR LA PARTIE GO POOUR LE CACHE ///////////////////////////
        // if (userToken) {
        //   accountService.getAuthenticatedUser(userToken)
        //     .then((res) => {
        //       setUserData({
        //         token: userToken,
        //         user: res.data,
        //         isAuth: true,
        //       });
        //       window.sessionStorage.setItem(
        //         "CacheUserData",
        //         JSON.stringify({
        //           token: userToken,
        //           isAuth: true,
        //           user: res.data,
        //         })
        //       );
        //     })
        //   }
       
        
     })
     .then(() => {
        navigate('/home', {replace: true})})
     .catch((e) => {
      if (!e.response.data) {
        navigate( '/error', {replace: true}) // à changer plus tard avec un vraie affichage quand l'api est down
        return
      }
        console.error("Error while login", e.response.data);
        setErrors(e.response.data);
        setPassword("")
 });

    
};
return (
  <div className="login-page">
  <div className="login-container">
        {errors.general && (
          <small>{errors.general}</small>
        )}
    <form onSubmit={handleLogin} className="login-form">
      <div className="login-form-group">
        <label htmlFor="email">Email:</label>
        <input
          id="email"
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="login-form-control"
          required
        />
           {errors.email && (
                <small>* {errors.email}</small>
              )}
      </div>
      <div className="login-form-group">
        <label htmlFor="password">Password:</label>
        <input
          id="password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="login-form-control"
          required
        />
           {errors.password && (
                <small>* {errors.password}</small>
              )}
      </div>
      <button type="submit" className="login-btn btn-primary">Login</button>
    </form>
    <Link to="/register" className="login-register-link">Register</Link>

  </div>
  </div>
);
};