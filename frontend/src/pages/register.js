import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";

//api
import { accountService } from "../_services/account.service";

//style
import "../css/register.css"


export const Registerpage = () => {
  //utils
  let navigate = useNavigate()
  function handleChange(event) {
    setAvatar(event.target.files[0])
  }

  //local state
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [firstName, setFirstName] = useState("")
  const [lastName, setLastName] = useState("")
  const [dateOfBirth, setDateOfBirth] = useState("");
  const [avatar, setAvatar] = useState("")
  const [nickname, setNickname] = useState("")
  const [aboutMe, setAboutme] = useState("")

  const [errors, setErrors] = useState({});


  //register operation
  const handleregister = async (e) => {
     e.preventDefault();
    
      let payload = {
        email:email,
        password:password,
        firstName:firstName,
        lastName:lastName,
        dateOfBirth:dateOfBirth,
        nickname:nickname,
        avatar:avatar,
        aboutMe:aboutMe
}
  //call api
    accountService.register(payload)   
    .then((r) => {
        navigate('/login', {replace: true})
    })
    .catch((e) => {
      if (!e.response.data) {
        navigate( '/error', {replace: true}) // à changer plus tard avec un vraie affichage quand l'api est down
        return
      }
        console.error("Error while register", e.response.data);
        setErrors(e.response.data);
 });

    
};
  return (
    <div className="register-page">
      <div className="register-container">
      {errors.general && (
          <small>{errors.general}</small>
        )}
      <form onSubmit={handleregister} className="register-form">
        <div className="register-form-group">
          <label htmlFor="email">Email:</label>
          <input
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="register-form-control"
          />
           {errors.email && (
                <small>* {errors.email}</small>
              )}
        </div>
        <div className="reigster-form-group">
          <label htmlFor="password">Password:</label>
          <input
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="register-form-control"
          />
          {errors.password && (
                <small>* {errors.password}</small>
              )}
        </div>
        <div className="reigster-form-group">
          <label htmlFor="firstname">Firstname:</label>
          <input
            id="firstname"
            type="text"
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
            className="register-form-control"
          />
          {errors.firstname && (
                <small>* {errors.firstname}</small>
              )}
        </div>
        <div className="reigster-form-group">
          <label htmlFor="lastname">Lastname:</label>
          <input
            id="lastname"
            type="text"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
            className="register-form-control"
          />
          {errors.lastname && (
                <small>* {errors.lastname}</small>
              )}
        </div>
        <div className="reigster-form-group">
          <label htmlFor="dateOfBirth">DateOfBirth:</label>
          <input
            id="dateOfBirth"
            type="date"
            value={dateOfBirth}
            onChange={(e) => setDateOfBirth(e.target.value)}
          />
          {errors.dateofbirth && (
                <small>* {errors.dateofbirth}</small>
              )}
        </div>
        <div className="reigster-form-group">
          <label htmlFor="nickname">Nickname:</label>
          <input
            id="nickname"
            type="text"
            value={nickname}
            onChange={(e) => setNickname(e.target.value)}
            className="register-form-control"
          />
          {errors.nickname && (
                <small>* {errors.nickname}</small>
              )}
        </div>
        <div className="reigster-form-group">
          <label htmlFor="aboutme">Aboutme:</label>
          <input
            id="aboutme"
            type="text"
            value={aboutMe}
            onChange={(e) => setAboutme(e.target.value)}
            className="register-form-control"
          />
          {errors.avoutme && (
                <small>* {errors.aboutme}</small>
              )}
        </div>
        <div >
          <label htmlFor="avatar">Avatar:</label>
          <input
            id="avatar"
            type="file"
            onChange={handleChange}
          />
          {errors.avatar && (
                <small>* {errors.avatar}</small>
              )}
        </div>
      
        <button type="submit" className="register-btn btn-primary">Register</button>
      </form>
      <Link to="/login" className="reigster-login-link">Login</Link>

    </div>
    </div>

  
  );
};

