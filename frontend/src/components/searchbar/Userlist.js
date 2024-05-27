import "../../css/leftmenu.css"

import { Link , useNavigate } from 'react-router-dom'; 

function UserList({ users, onItemClick }) {
  // Si aucun résultat, affichez rien
  const navigate = useNavigate()
  if (users == null) {
    return ( 
      <div className="user-list"></div>
    )
  }

  // Affichage de la liste des utilisateurs
  return (  
    <div className="user-list">
    {users.map((user) => (
        <div className="user-item" key={user.id} onClick={() => { onItemClick(); navigate(`/user/${user.nickname}`); }}>     
               <li>{user.nickname}</li>
        </div>
    ))}
</div>

  );
}

export default UserList;
