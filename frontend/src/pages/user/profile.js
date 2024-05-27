import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

//components
import MyImageComponent from '../../components/profilepicture';
import FollowButonn from '../../components/follow/followbutton';

//api
import { usersService } from '../../_services/users.service';

//style
import '../../css/UserLayout.css';





const UserProfile = (props) => {

    const { userId } = useParams(); 
    const { onProfileInfo } = props;
    
    //local state
    const [response, setResponse] = useState([]);
    const [error, setError] = useState(null); 


    useEffect(() => {
        setError(null); 
    }, [userId]); 

    useEffect(() => {
        const fetchUserInfo = async () => {
            try {
                const res = await usersService.UserInformation(userId);
                setResponse(res.data[0]);
                onProfileInfo(res.data[0].public_private, res.data[0].follow_status);
            } catch (error) {
                setError("L'utilisateur n'existe pas.");

            }
        };

        fetchUserInfo();
    }, [userId]); 

    if (error) {
        return <div className='profile-head'>Erreur : {error}</div>;
    }
    return (
        <div className='profile-head'>
        {!response.public_private || response.follow_status == 1 || response.follow_status == 3 ?(
            <div className='profile'>
                <div>
                <ul className='profile-name'>{response.first_name} {response.last_name} ({response.nickname})</ul>
                <ul className='profile-dob'>Date of birth: {response.date_of_birth}</ul>
                <ul className='profile-aboutme'>{response.about_me}</ul>
                <div className='profile-butonn'>
                <FollowButonn className followStatus={response.follow_status} isprivate={response.public_private} nickname={response.nickname} />
                </div>
                </div>
                <div className='avatar-in-profile'><MyImageComponent avatarPath={response.avatar} avatarSize={150} route={"avatar"}/></div>
            </div>
        ) : (
            <div>
            <ul>Nickname: {response.nickname}</ul>
            <ul>Profile :  private</ul>
            <FollowButonn followStatus={response.follow_status} isprivate={response.public_private} nickname={response.nickname} />
        </div>
        )}
    </div>
       
    )
}

export default UserProfile