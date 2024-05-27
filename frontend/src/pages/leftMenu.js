import { Link, useNavigate } from "react-router-dom";
import { useState, useEffect, useContext , useRef} from "react";


//api
import { accountService } from "../_services/account.service";
import { searchService } from "../_services/searchbar";


//components
import SearchBar from "../components/searchbar/SearchBar";
import UserList from "../components/searchbar/Userlist";
import MyImageComponent from "../components/profilepicture";

import NotificationsCounter from "../components/notificationsCounter"
import { WebsocketContext } from "../_services/websocketProvider"

//css
import "../css/leftmenu.css"

const LeftMenu = () => {
  //utils
  let navigate = useNavigate()

  //local state
  const [searchResults, setSearchResults] = useState([]);
  const [nikcnanme, setNickname] = useState([]);
  const [avatar, setAvatar] = useState([])
  const userListRef = useRef(null);

  // searchbar
  const handleSearch = (query) => {
    console.log(query)
    searchService.searchbarleftmenu(query)
      .then((data) => setSearchResults(data.data))
  } 
 // useEffect pour recuperer le pseudo sur le leftmenu
  useEffect(() => {
    const fetchNickname = async () => {
      try {
        const response = await accountService.nickname();
        setNickname(response.data.nickname);
        setAvatar(response.data.avatar)
      } catch (error) {
        console.error("Erreur lors de la récupération du nickname :", error);
        accountService.logout()
        navigate('/login')
      }
    };

    fetchNickname(); 
  }, []); 


  const [notifCounter, setNotifCounter] = useState(0)
  const [ready, value, send, connect] = useContext(WebsocketContext);

  useEffect(() => {
      if (!ready) {
          setNotifCounter(0)
          return
      }

      if (value === null) {
          return
      }

      const data = JSON.parse(value?.data)

      if (value?.type === "notif") {
          setNotifCounter(data.count)
      }

      if (value?.type === "newnotif") {
          setNotifCounter(notifCounter + 1)
      }
  }, [ready, value])

  const resetCounter = () => {
    setNotifCounter(0)
  }

// deconnexion
  const handleLogout = () => {
    accountService.logout()
    navigate('/login')

  };

  const handleItemClick = () => {
    console.log("oui je ferme la page")
    setSearchResults([]); 
  };

  useEffect(() => {
    function handleClickOutside(event) {
      if (userListRef.current && !userListRef.current.contains(event.target)) {
        setSearchResults([]);
      }
    }

    document.addEventListener("mousedown", handleClickOutside);
    
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
      <div className="leftmenu">
        <AvatarName nickname = {nikcnanme} avatar = {avatar}/>
       
        <div className="searchbarandlist"> 
      <SearchBar onSearch={handleSearch} />
      <div className="listusers" ref={userListRef}>
        <UserList users={searchResults} onItemClick={handleItemClick} />
      </div>   
       </div>
        <nav><ul>
          <li><Link to="/home">Home</Link></li>
          <li><Link to={`/user/${nikcnanme}/posts`}>Mes posts</Link></li>
          <li><Link to={`/user/${nikcnanme}/followers`}>Mes followers</Link></li>
          <li><Link to={`/user/${nikcnanme}/following`}>Mes followings</Link></li>
          <li><Link to={`/user/${nikcnanme}/groups`}>Mes groupes</Link></li>
          <li><Link to={"/groupcreation"}>Create a group</Link></li>
          <li><Link to={"/notifications"}>Notifications <NotificationsCounter/></Link></li>
        </ul></nav>
        <button onClick={handleLogout}>Logout</button>
        
       
      </div>
  )
};
  
export default LeftMenu;
  

 export const AvatarName = ({nickname,avatar}) => {
  let navigate = useNavigate()
  const handleAvatarNameClick = () => {
    navigate(`/user/${nickname}/posts`);
  };

  return (
    <div onClick={handleAvatarNameClick} className="leftmenu-avatarname">
      <MyImageComponent avatarPath={avatar} avatarSize={50} route={"avatar"}/>
      <div className="leftmenu-nickname">{nickname}</div>
  </div>

  )


}