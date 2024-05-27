import { useParams, useNavigate } from 'react-router-dom';
import React, { useState, useEffect } from 'react';

//api
import { followService } from '../../_services/follow';

//component
import MyImageComponent from '../../components/profilepicture';

const UserFollowing = () => {

    //local state
    const [response, setResponse] = useState([]);
    const [error, setError] = useState(null); 
    const navigate = useNavigate()
    
    //utils
    let userNickname = useParams().userId;

    useEffect(() => {
        const fetchFollowing = async () => {
            try {
                const res = await followService.followingApi(userNickname);
                setResponse(res.data);
            } catch (error) {
                console.error('Erreur lors de la récupération des following :', error.response.data.error);
                setError(error.response.data.error)
            }
        };

        fetchFollowing();
    }, [userNickname]); 

   

    return (
        <div className='followers-list'>
            {response && response.length > 0 ? (
                <ul className='listeusers'>
                    {response.map(following => (
                        <li className='avatarname' onClick={() => navigate(`/user/${following.nickname}`)} key={following.id}>
                         <MyImageComponent avatarPath={following.avatar} avatarSize={50} route={"avatar"}/>
                         <div className='nickname'>{following.nickname}</div>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>{error}</p>
            )}
        </div>
    );
}

export default UserFollowing;


