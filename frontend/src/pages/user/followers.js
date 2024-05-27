import { useParams, useNavigate } from 'react-router-dom';
import React, { useState, useEffect } from 'react';

//api
import { followService } from '../../_services/follow';

//component
import MyImageComponent from '../../components/profilepicture';

//style
import '../../css/Follow.css'; 



const UserFollowers = () => {
    //local state
    const [response, setResponse] = useState([]);
    const [error, setError] = useState(null); 
    const navigate = useNavigate();
    
    //utils
    let userNickname = useParams().userId;

    useEffect(() => {
        const fetchFollowers = async () => {
            try {
                const res = await followService.followersApi(userNickname);
                setResponse(res.data);
            } catch (error) {
                console.error('Erreur lors de la récupération des followers :', error.response.data.error);
                setError(error.response.data.error)

            }
        };

        fetchFollowers();
    }, [userNickname]); 

 

    return (
        <div className='followers-list'>
            {response && response.length > 0 ? (
                <ul className='listeusers'>
                    {response.map(follower => (
                        <li className='avatarname' onClick={() => navigate(`/user/${follower.nickname}`)} key={follower.id}>
                         <MyImageComponent avatarPath={follower.avatar} avatarSize={50} route={"avatar"}/>
                         <div className='nickname'>{follower.nickname}</div>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>{error}</p>
            )}
        </div>
    );
}

export default UserFollowers;
